package _2lighting

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gorustyt/fyne/v2"
	"github.com/gorustyt/fyne/v2/canvas3d"
	"github.com/gorustyt/fyne/v2/canvas3d/context"
	"github.com/gorustyt/fyne/v2/container"
	"github.com/gorustyt/learn-opengl-go/ui/base_ui"
)

var (
	cubePositions = []mgl32.Vec3{
		{0.0, 0.0, 0.0},
		{2.0, 5.0, -15.0},
		{-1.5, -2.2, -2.5},
		{-3.8, -2.0, -12.3},
		{2.4, -0.4, -3.5},
		{-1.7, 3.0, -7.5},
		{1.3, -2.0, -2.5},
		{1.5, 2.0, -2.5},
		{1.5, 0.2, -1.5},
		{-1.3, 1.0, -1.5},
	}

	vert2 = []float32{
		// positions          // normals           // texture coords
		-0.5, -0.5, -0.5, 0.0, 0.0, -1.0, 0.0, 0.0,
		0.5, -0.5, -0.5, 0.0, 0.0, -1.0, 1.0, 0.0,
		0.5, 0.5, -0.5, 0.0, 0.0, -1.0, 1.0, 1.0,
		0.5, 0.5, -0.5, 0.0, 0.0, -1.0, 1.0, 1.0,
		-0.5, 0.5, -0.5, 0.0, 0.0, -1.0, 0.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 0.0, -1.0, 0.0, 0.0,

		-0.5, -0.5, 0.5, 0.0, 0.0, 1.0, 0.0, 0.0,
		0.5, -0.5, 0.5, 0.0, 0.0, 1.0, 1.0, 0.0,
		0.5, 0.5, 0.5, 0.0, 0.0, 1.0, 1.0, 1.0,
		0.5, 0.5, 0.5, 0.0, 0.0, 1.0, 1.0, 1.0,
		-0.5, 0.5, 0.5, 0.0, 0.0, 1.0, 0.0, 1.0,
		-0.5, -0.5, 0.5, 0.0, 0.0, 1.0, 0.0, 0.0,

		-0.5, 0.5, 0.5, -1.0, 0.0, 0.0, 1.0, 0.0,
		-0.5, 0.5, -0.5, -1.0, 0.0, 0.0, 1.0, 1.0,
		-0.5, -0.5, -0.5, -1.0, 0.0, 0.0, 0.0, 1.0,
		-0.5, -0.5, -0.5, -1.0, 0.0, 0.0, 0.0, 1.0,
		-0.5, -0.5, 0.5, -1.0, 0.0, 0.0, 0.0, 0.0,
		-0.5, 0.5, 0.5, -1.0, 0.0, 0.0, 1.0, 0.0,

		0.5, 0.5, 0.5, 1.0, 0.0, 0.0, 1.0, 0.0,
		0.5, 0.5, -0.5, 1.0, 0.0, 0.0, 1.0, 1.0,
		0.5, -0.5, -0.5, 1.0, 0.0, 0.0, 0.0, 1.0,
		0.5, -0.5, -0.5, 1.0, 0.0, 0.0, 0.0, 1.0,
		0.5, -0.5, 0.5, 1.0, 0.0, 0.0, 0.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 0.0, 0.0, 1.0, 0.0,

		-0.5, -0.5, -0.5, 0.0, -1.0, 0.0, 0.0, 1.0,
		0.5, -0.5, -0.5, 0.0, -1.0, 0.0, 1.0, 1.0,
		0.5, -0.5, 0.5, 0.0, -1.0, 0.0, 1.0, 0.0,
		0.5, -0.5, 0.5, 0.0, -1.0, 0.0, 1.0, 0.0,
		-0.5, -0.5, 0.5, 0.0, -1.0, 0.0, 0.0, 0.0,
		-0.5, -0.5, -0.5, 0.0, -1.0, 0.0, 0.0, 1.0,

		-0.5, 0.5, -0.5, 0.0, 1.0, 0.0, 0.0, 1.0,
		0.5, 0.5, -0.5, 0.0, 1.0, 0.0, 1.0, 1.0,
		0.5, 0.5, 0.5, 0.0, 1.0, 0.0, 1.0, 0.0,
		0.5, 0.5, 0.5, 0.0, 1.0, 0.0, 1.0, 0.0,
		-0.5, 0.5, 0.5, 0.0, 1.0, 0.0, 0.0, 0.0,
		-0.5, 0.5, -0.5, 0.0, 1.0, 0.0, 0.0, 1.0,
	}
	lightMapVs = `#version 330 core
layout (location = 0) in vec3 position;
layout (location = 1) in vec3 normal;
layout (location = 2) in vec2 texCoord;

out vec3 FragPos;
out vec3 Normal;
out vec2 TexCoords;

uniform mat4 model;
uniform mat4 view;
uniform mat4 projection;

void main()
{
    FragPos = vec3(model * vec4(position, 1.0));
    Normal = mat3(transpose(inverse(model))) * normal;  
    TexCoords = texCoord;
    
    gl_Position = projection * view * vec4(FragPos, 1.0);
}`
	lightMapFs = `#version 330 core
out vec4 FragColor;

struct Material {
    sampler2D diffuse;
    sampler2D specular;    
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
in vec2 TexCoords;
  
uniform vec3 viewPos;
uniform Material material;
uniform Light light;

void main()
{
    // ambient
    vec3 ambient = light.ambient * texture(material.diffuse, TexCoords).rgb;
  	
    // diffuse 
    vec3 norm = normalize(Normal);
    vec3 lightDir = normalize(light.position - FragPos);
    float diff = max(dot(norm, lightDir), 0.0);
    vec3 diffuse = light.diffuse * diff * texture(material.diffuse, TexCoords).rgb;  
    
    // specular
    vec3 viewDir = normalize(viewPos - FragPos);
    vec3 reflectDir = reflect(-lightDir, norm);  
    float spec = pow(max(dot(viewDir, reflectDir), 0.0), material.shininess);
    vec3 specular = light.specular * spec * texture(material.specular, TexCoords).rgb;  
        
    vec3 result = ambient + diffuse + specular;
    FragColor = vec4(result, 1.0);
} `
	diffuseMap  = "assets/textures/container2.png"
	specularMap = "assets/textures/container2_specular.png"
)

type Material1 struct {
	*Lighting
	l   *canvas3d.Light
	tex *canvas3d.Texture

	sliderSpecular *base_ui.Vec3Slider
	sliderDiffuse  *base_ui.Vec3Slider
	sliderAmbient  *base_ui.Vec3Slider
	menu           *fyne.Container
}

func (l *Material1) InitChapterContent(c *base_ui.ChapterContent) {
	l.l.Specular = l.sliderSpecular.GetPos()
	l.l.Diffuse = l.sliderDiffuse.GetPos()
	l.l.Ambient = l.sliderAmbient.GetPos()

	c.Canvas3d().SetShaderConfig(0, lightCube_vs, lightCube_fs)
	c.Canvas3d().AppendObj(0, l.lightCoord)
	c.Canvas3d().AppendObj(0, l.lightVert)

	c.Canvas3d().SetShaderConfig(1, lightMapVs, lightMapFs)
	c.Canvas3d().AppendObj(1, l.cubeVert)
	c.Canvas3d().AppendObj(1, l.cubeCoord)
	c.Canvas3d().AppendRenderFunc(1, func(painter context.Painter) {
		painter.UniformVec3("viewPos", l.cubeCoord.Position)
		painter.Uniform1f("material.shininess", 64.0)
	})
	c.Canvas3d().AppendObj(1, l.l)
	c.Canvas3d().AppendObj(1, l.tex)
}

func (l *Material1) InitParamsContent(c *base_ui.ParamsContent) {
	c.Add(l.menu)
}

func NewMaterial1() base_ui.IChapter {
	l1 := &Lighting{

		lightCoord: canvas3d.NewCoordinate(),
		lightVert:  canvas3d.NewVertexFloat32Array(),
		cubeCoord:  canvas3d.NewCoordinate(),
		cubeVert:   canvas3d.NewVertexFloat32Array()}
	l1.cubeVert.Arr = vert2
	l1.cubeVert.PositionSize = []int{3, 0}
	l1.cubeVert.NormalSize = []int{3, 3}
	l1.cubeVert.TexCoordSize = []int{2, 6}

	l1.lightVert.Arr = vert
	l1.lightVert.PositionSize = []int{3, 0}

	l := &Material1{
		Lighting:       l1,
		l:              canvas3d.NewLight(),
		tex:            canvas3d.NewTexture(),
		sliderSpecular: base_ui.NewVec3Slider("specular", mgl32.Vec3{1, 1, 1}),
		sliderDiffuse:  base_ui.NewVec3Slider("diffuse", mgl32.Vec3{0.5, 0.5, 0.5}),
		sliderAmbient:  base_ui.NewVec3Slider("ambient", mgl32.Vec3{0.2, 0.2, 0.2}),
	}
	l.menu = container.NewVBox(
		l.sliderAmbient.GetRenderObj(),
		l.sliderDiffuse.GetRenderObj(),
		l.sliderSpecular.GetRenderObj(),
	)
	l.cubeCoord.Rotate(45, mgl32.Vec3{1.0, 0.3, 0.5}.Normalize())
	l.cubeCoord.Scale(0.5, 0.5, 0.5)

	l.lightCoord.TranslateVec3(lightPos)
	l.lightCoord.Scale(0.1, 0.1, 0.1)

	l.l.Position = mgl32.Vec3{0.6, 0.5, 1}
	l.tex.AppendPathWithCustomAttr("material.diffuse", diffuseMap)
	l.tex.AppendPathWithCustomAttr("material.specular", specularMap)
	return l
}
