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

var (
	mvs = `#version 330 core
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
	mfs = `#version 330 core
out vec4 FragColor;

struct Material {
    vec3 ambient;
    vec3 diffuse;
    vec3 specular;    
    float shininess;
}; 

struct Light {
    vec3 position;

    vec3 ambient;
    vec3 diffuse;
    vec3 specular;
};

in vec3 FragPos;  
in vec3 Normal;  
  
uniform vec3 viewPos;
uniform Material material;
uniform Light light;

void main()
{
    // ambient
    vec3 ambient = light.ambient * material.ambient;
  	
    // diffuse 
    vec3 norm = normalize(Normal);
    vec3 lightDir = normalize(light.position - FragPos);
    float diff = max(dot(norm, lightDir), 0.0);
    vec3 diffuse = light.diffuse * (diff * material.diffuse);
    
    // specular
    vec3 viewDir = normalize(viewPos - FragPos);
    vec3 reflectDir = reflect(-lightDir, norm);  
    float spec = pow(max(dot(viewDir, reflectDir), 0.0), material.shininess);
    vec3 specular = light.specular * (spec * material.specular);  
        
    vec3 result = ambient + diffuse + specular;
    FragColor = vec4(result, 1.0);
} `
)

type Material struct {
	*Lighting
	l *canvas3d.Light
	m *canvas3d.Material
}

func (l *Material) InitChapterContent(c *base_ui.ChapterContent) {
	c.Canvas3d().SetShaderConfig(0, lightCube_vs, lightCube_fs)
	c.Canvas3d().AppendObj(0, l.lightCoord)
	c.Canvas3d().AppendObj(0, l.lightVert)

	c.Canvas3d().SetShaderConfig(1, mvs, mfs)
	c.Canvas3d().AppendObj(1, l.cubeVert)
	c.Canvas3d().AppendObj(1, l.cubeCoord)
	// change the light's position values over time (can be done anywhere in the render loop actually, but try to do it at least before using the light source positions)

	x := float32(math.Sin(canvas3d_render.GetGlfwTime() * 2.0))
	y := float32(math.Sin(canvas3d_render.GetGlfwTime() * 0.7))
	z := float32(math.Sin(canvas3d_render.GetGlfwTime() * 1.3))
	lightColor := mgl32.Vec3{x, y, z}
	diffuseColor := lightColor.Mul(0.5)   // decrease the influence
	ambientColor := diffuseColor.Mul(0.2) // low influence
	l.l.Diffuse = diffuseColor
	l.l.Ambient = ambientColor
	c.Canvas3d().AppendRenderFunc(1, func(painter context.Painter) {
		painter.UniformVec3("viewPos", l.cubeCoord.Position)
	})
	c.Canvas3d().AppendObj(1, l.l)
	c.Canvas3d().AppendObj(1, l.m)
}
func (l *Material) RefreshInterVal() time.Duration {
	return 100 * time.Millisecond
}
func (l *Material) InitParamsContent(c *base_ui.ParamsContent) {

}

func NewMaterial() base_ui.IChapter {
	l1 := &Lighting{
		lightCoord: canvas3d.NewCoordinate(),
		lightVert:  canvas3d.NewVertexFloat32Array(),
		cubeCoord:  canvas3d.NewCoordinate(),
		cubeVert:   canvas3d.NewVertexFloat32Array()}
	l1.cubeVert.Arr = vert
	l1.cubeVert.PositionSize = []int{3, 0}

	l1.lightVert.Arr = vert
	l1.lightVert.PositionSize = []int{3, 0}

	l := &Material{Lighting: l1, l: canvas3d.NewLight(), m: canvas3d.NewMaterial()}
	l.cubeVert.Arr = vert1
	l.cubeVert.NormalSize = []int{3, 3}
	l.cubeCoord.Rotate(45, mgl32.Vec3{1.0, 0.3, 0.5}.Normalize())
	l.cubeCoord.Scale(0.5, 0.5, 0.5)

	l.lightCoord.TranslateVec3(lightPos)
	l.lightCoord.Scale(0.1, 0.1, 0.1)

	l.m.Shininess = 32.0
	l.m.Specular = mgl32.Vec3{0.5, 0.5, 0.5}
	l.m.Diffuse = mgl32.Vec3{1, 0.5, 0.31}
	l.m.Ambient = mgl32.Vec3{1, 0.5, 0.31}

	l.l.Position = lightPos
	l.l.Specular = mgl32.Vec3{1, 1, 1}
	return l
}
