package _2lighting

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gorustyt/fyne/v2/canvas3d"
	"github.com/gorustyt/fyne/v2/canvas3d/canvas3d_render"
	"github.com/gorustyt/fyne/v2/canvas3d/context"
	"github.com/gorustyt/learn-opengl-go/ui/base_ui"
	"math"
	"time"
)

var vert1 = []float32{
	-0.5, -0.5, -0.5, 0.0, 0.0, -1.0,
	0.5, -0.5, -0.5, 0.0, 0.0, -1.0,
	0.5, 0.5, -0.5, 0.0, 0.0, -1.0,
	0.5, 0.5, -0.5, 0.0, 0.0, -1.0,
	-0.5, 0.5, -0.5, 0.0, 0.0, -1.0,
	-0.5, -0.5, -0.5, 0.0, 0.0, -1.0,

	-0.5, -0.5, 0.5, 0.0, 0.0, 1.0,
	0.5, -0.5, 0.5, 0.0, 0.0, 1.0,
	0.5, 0.5, 0.5, 0.0, 0.0, 1.0,
	0.5, 0.5, 0.5, 0.0, 0.0, 1.0,
	-0.5, 0.5, 0.5, 0.0, 0.0, 1.0,
	-0.5, -0.5, 0.5, 0.0, 0.0, 1.0,

	-0.5, 0.5, 0.5, -1.0, 0.0, 0.0,
	-0.5, 0.5, -0.5, -1.0, 0.0, 0.0,
	-0.5, -0.5, -0.5, -1.0, 0.0, 0.0,
	-0.5, -0.5, -0.5, -1.0, 0.0, 0.0,
	-0.5, -0.5, 0.5, -1.0, 0.0, 0.0,
	-0.5, 0.5, 0.5, -1.0, 0.0, 0.0,

	0.5, 0.5, 0.5, 1.0, 0.0, 0.0,
	0.5, 0.5, -0.5, 1.0, 0.0, 0.0,
	0.5, -0.5, -0.5, 1.0, 0.0, 0.0,
	0.5, -0.5, -0.5, 1.0, 0.0, 0.0,
	0.5, -0.5, 0.5, 1.0, 0.0, 0.0,
	0.5, 0.5, 0.5, 1.0, 0.0, 0.0,

	-0.5, -0.5, -0.5, 0.0, -1.0, 0.0,
	0.5, -0.5, -0.5, 0.0, -1.0, 0.0,
	0.5, -0.5, 0.5, 0.0, -1.0, 0.0,
	0.5, -0.5, 0.5, 0.0, -1.0, 0.0,
	-0.5, -0.5, 0.5, 0.0, -1.0, 0.0,
	-0.5, -0.5, -0.5, 0.0, -1.0, 0.0,

	-0.5, 0.5, -0.5, 0.0, 1.0, 0.0,
	0.5, 0.5, -0.5, 0.0, 1.0, 0.0,
	0.5, 0.5, 0.5, 0.0, 1.0, 0.0,
	0.5, 0.5, 0.5, 0.0, 1.0, 0.0,
	-0.5, 0.5, 0.5, 0.0, 1.0, 0.0,
	-0.5, 0.5, -0.5, 0.0, 1.0, 0.0,
}
var (
	basic_lighting_vs1 = `#version 330 core
layout (location = 0) in vec3 position;
layout (location = 1) in vec3 normal;

out vec3 FragPos;
out vec3 Normal;

uniform mat4 model;
uniform mat4 view;
uniform mat4 projection;

void main()
{
    FragPos = vec3(model * vec4(position, 1.0));
    Normal = mat3(transpose(inverse(model))) * normal;  
    
    gl_Position = projection * view * vec4(FragPos, 1.0);
}`
	basic_lighting_fs1 = `#version 330 core
out vec4 FragColor;

in vec3 Normal;  
in vec3 FragPos;  
  
uniform vec3 lightPos; 
uniform vec3 viewPos; 
uniform vec3 lightColor;
uniform vec3 objectColor;

void main()
{
    // ambient
    float ambientStrength = 0.1;
    vec3 ambient = ambientStrength * lightColor;
  	
    // diffuse 
    vec3 norm = normalize(Normal);
    vec3 lightDir = normalize(lightPos - FragPos);
    float diff = max(dot(norm, lightDir), 0.0);
    vec3 diffuse = diff * lightColor;
    
    // specular
    float specularStrength = 0.5;
    vec3 viewDir = normalize(viewPos - FragPos);
    vec3 reflectDir = reflect(-lightDir, norm);  
    float spec = pow(max(dot(viewDir, reflectDir), 0.0), 32);
    vec3 specular = specularStrength * spec * lightColor;  
        
    vec3 result = (ambient + diffuse + specular) * objectColor;
    FragColor = vec4(result, 1.0);
} `
)

type Light2 struct {
	*Lighting
}

func (l *Light2) InitChapterContent(c *base_ui.ChapterContent) {
	c.Canvas3d().SetShaderConfig(0, lightCube_vs, lightCube_fs)
	c.Canvas3d().AppendObj(0, l.lightCoord)
	c.Canvas3d().AppendObj(0, l.lightVert)

	c.Canvas3d().SetShaderConfig(1, basic_lighting_vs1, basic_lighting_fs1)
	c.Canvas3d().AppendObj(1, l.cubeVert)
	c.Canvas3d().AppendObj(1, l.cubeCoord)
	// change the light's position values over time (can be done anywhere in the render loop actually, but try to do it at least before using the light source positions)
	x := lightPos[0] + float32(math.Sin(canvas3d_render.GetGlfwTime())/2.0)
	y := lightPos[1] + float32(math.Sin(canvas3d_render.GetGlfwTime()/2.0))
	newLightPos := mgl32.Vec3{x, y, lightPos[2]}
	l.lightCoord.TranslateVec3(newLightPos)
	c.Canvas3d().AppendRenderFunc(1, func(painter context.Painter) {
		painter.UniformVec3("lightColor", mgl32.Vec3{1, 1, 1})
		painter.UniformVec3("objectColor", mgl32.Vec3{1.0, 0.5, 0.31})
		painter.UniformVec3("lightPos", newLightPos)
		painter.UniformVec3("viewPos", l.cubeCoord.Position)
	})
}
func (l *Light2) RefreshInterVal() time.Duration {
	return 1 * time.Second
}
func (l *Light2) InitParamsContent(c *base_ui.ParamsContent) {

}

func NewLight2() base_ui.IChapter {
	l1 := &Lighting{
		lightCoord: canvas3d.NewCoordinate(),
		lightVert:  canvas3d.NewVertexFloat32Array(),
		cubeCoord:  canvas3d.NewCoordinate(),
		cubeVert:   canvas3d.NewVertexFloat32Array()}
	l1.cubeVert.Arr = vert
	l1.cubeVert.PositionSize = []int{3, 0}

	l1.lightVert.Arr = vert
	l1.lightVert.PositionSize = []int{3, 0}

	l := &Light2{Lighting: l1}
	l.cubeVert.Arr = vert1
	l.cubeVert.NormalSize = []int{3, 3}
	l.cubeCoord.Rotate(45, mgl32.Vec3{1.0, 0.3, 0.5}.Normalize())
	l.cubeCoord.Scale(0.5, 0.5, 0.5)

	l.lightCoord.TranslateVec3(lightPos)
	l.lightCoord.Scale(0.1, 0.1, 0.1)

	return l
}
