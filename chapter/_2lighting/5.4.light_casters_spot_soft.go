package _2lighting

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/go-gl/mathgl/mgl64"
	"github.com/gorustyt/fyne/v2"
	"github.com/gorustyt/fyne/v2/canvas3d"
	"github.com/gorustyt/fyne/v2/canvas3d/context"
	"github.com/gorustyt/fyne/v2/container"
	"github.com/gorustyt/fyne/v2/data/binding"
	"github.com/gorustyt/fyne/v2/widget"
	"github.com/gorustyt/learn-opengl-go/ui/base_ui"
	"log/slog"
	"math"
)

var (
	lightCastVs1 = `#version 330 core
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
	lightCastFs1 = `#version 330 core
out vec4 FragColor;

struct Material {
    sampler2D diffuse;
    sampler2D specular;    
    float shininess;
}; 

struct Light {
    vec3 position;  
    vec3 direction;
    float cutOff;
    float outerCutOff;
  
    vec3 ambient;
    vec3 diffuse;
    vec3 specular;
	
    float constant;
    float linear;
    float quadratic;
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
    
    // spotlight (soft edges)
    float theta = dot(lightDir, normalize(-light.direction)); 
    float epsilon = (light.cutOff - light.outerCutOff);
    float intensity = clamp((theta - light.outerCutOff) / epsilon, 0.0, 1.0);
    diffuse  *= intensity;
    specular *= intensity;
    
    // attenuation
    float distance    = length(light.position - FragPos);
    float attenuation = 1.0 / (light.constant + light.linear * distance + light.quadratic * (distance * distance));    
    ambient  *= attenuation; 
    diffuse   *= attenuation;
    specular *= attenuation;   
        
    vec3 result = ambient + diffuse + specular;
    FragColor = vec4(result, 1.0);
} `
	vert4 = []float32{
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
)

type LightSpotSoft struct {
	*Lighting
	l   *canvas3d.SpotLight
	tex *canvas3d.Texture

	sliderSpecular *base_ui.Vec3Slider
	sliderDiffuse  *base_ui.Vec3Slider
	sliderAmbient  *base_ui.Vec3Slider
	cutOffSlider   *widget.Slider
	menu           *fyne.Container
}

func (l *LightSpotSoft) InitChapterContent(c *base_ui.ChapterContent) {
	l.l.Specular = l.sliderSpecular.GetPos()
	l.l.Diffuse = l.sliderDiffuse.GetPos()
	l.l.Ambient = l.sliderAmbient.GetPos()

	l.l.Position = l.cubeCoord.Position
	l.l.Direction = l.cubeCoord.Front

	c.Canvas3d().SetShaderConfig(0, lightCube_vs, lightCube_fs)
	//c.Canvas3d().AppendObj(0, l.lightCoord)
	//c.Canvas3d().AppendObj(0, l.lightVert)

	c.Canvas3d().SetShaderConfig(1, lightCastVs1, lightCastFs1)
	c.Canvas3d().AppendObj(1, l.cubeVert)
	c.Canvas3d().AppendObj(1, l.cubeCoord)
	c.Canvas3d().AppendRenderFunc(1, func(painter context.Painter) {
		painter.UniformVec3("viewPos", l.cubeCoord.Position)
		painter.Uniform1f("material.shininess", 32)
	})
	c.Canvas3d().AppendObj(1, l.l)
	c.Canvas3d().AppendObj(1, l.tex)
}

func (l *LightSpotSoft) InitParamsContent(c *base_ui.ParamsContent) {
	c.Add(l.menu)
}
func (l *LightSpotSoft) initMenu() {
	l.sliderSpecular = base_ui.NewVec3Slider("specular", mgl32.Vec3{1, 1, 1})
	l.sliderDiffuse = base_ui.NewVec3Slider("diffuse", mgl32.Vec3{0.8, 0.8, 0.8})
	l.sliderAmbient = base_ui.NewVec3Slider("ambient", mgl32.Vec3{0.1, 0.1, 0.1})
	value := 12.5
	l.cutOffSlider = widget.NewSliderWithData(0, 360, binding.BindFloat(&value))
	l.cutOffSlider.Step = 1
	l.cutOffSlider.OnChanged = func(f float64) {
		slog.Info("cutoff change", "value", f)
		l.l.CutOff = float32(math.Cos(mgl64.DegToRad(f)))
		base_ui.ChapterRefresh()
	}

	value1 := 17.5
	outerCutOffSlider := widget.NewSliderWithData(0, 360, binding.BindFloat(&value1))
	outerCutOffSlider.Step = 1
	outerCutOffSlider.OnChanged = func(f float64) {
		slog.Info("cutoff change", "value", f)
		l.l.OuterCutOff = float32(math.Cos(mgl64.DegToRad(f)))
		base_ui.ChapterRefresh()
	}
	l.menu = container.NewVBox(
		l.sliderAmbient.GetRenderObj(),
		l.sliderDiffuse.GetRenderObj(),
		l.sliderSpecular.GetRenderObj(),
		widget.NewLabel("cutOff"),
		l.cutOffSlider,
		widget.NewLabel("outerCutOff"),
		outerCutOffSlider,
	)
}

func NewLightSpotSoft() base_ui.IChapter {
	l1 := &Lighting{

		lightCoord: canvas3d.NewCoordinate(),
		lightVert:  canvas3d.NewVertexFloat32Array(),
		cubeCoord:  canvas3d.NewCoordinate(),
		cubeVert:   canvas3d.NewVertexFloat32Array()}
	l1.cubeVert.Arr = vert4
	l1.cubeVert.PositionSize = []int{3, 0}
	l1.cubeVert.NormalSize = []int{3, 3}
	l1.cubeVert.TexCoordSize = []int{2, 6}

	l1.lightVert.Arr = vert
	l1.lightVert.PositionSize = []int{3, 0}

	l := &LightSpotSoft{
		Lighting: l1,
		l:        canvas3d.NewSpotLight(),
		tex:      canvas3d.NewTexture(),
	}
	l.initMenu()
	l.cubeCoord.Rotate(20*2, mgl32.Vec3{1.0, 0.3, 0.5}.Normalize())
	l.cubeCoord.Scale(0.5, 0.5, 0.5)

	l.lightCoord.TranslateVec3(lightPos)
	l.lightCoord.Scale(0.1, 0.1, 0.1)

	l.l.Constant = 1
	l.l.Linear = 0.09
	l.l.Quadratic = 0.032

	l.tex.AppendPathWithCustomAttr("material.diffuse", diffuseMap)
	l.tex.AppendPathWithCustomAttr("material.specular", specularMap)
	return l
}
