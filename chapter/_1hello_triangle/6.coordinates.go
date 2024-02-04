package _1hello_triangle

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gorustyt/fyne/v2/canvas3d"
	"github.com/gorustyt/learn-opengl-go/ui/base_ui"
)

var (
	vertices6 = []float32{
		-0.5, -0.5, -0.5, 0.0, 0.0,
		0.5, -0.5, -0.5, 1.0, 0.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		-0.5, 0.5, -0.5, 0.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 0.0,

		-0.5, -0.5, 0.5, 0.0, 0.0,
		0.5, -0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 1.0,
		0.5, 0.5, 0.5, 1.0, 1.0,
		-0.5, 0.5, 0.5, 0.0, 1.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,

		-0.5, 0.5, 0.5, 1.0, 0.0,
		-0.5, 0.5, -0.5, 1.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 1.0,
		-0.5, -0.5, -0.5, 0.0, 1.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,
		-0.5, 0.5, 0.5, 1.0, 0.0,

		0.5, 0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, -0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, 0.5, 0.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 0.0,

		-0.5, -0.5, -0.5, 0.0, 1.0,
		0.5, -0.5, -0.5, 1.0, 1.0,
		0.5, -0.5, 0.5, 1.0, 0.0,
		0.5, -0.5, 0.5, 1.0, 0.0,
		-0.5, -0.5, 0.5, 0.0, 0.0,
		-0.5, -0.5, -0.5, 0.0, 1.0,

		-0.5, 0.5, -0.5, 0.0, 1.0,
		0.5, 0.5, -0.5, 1.0, 1.0,
		0.5, 0.5, 0.5, 1.0, 0.0,
		0.5, 0.5, 0.5, 1.0, 0.0,
		-0.5, 0.5, 0.5, 0.0, 0.0,
		-0.5, 0.5, -0.5, 0.0, 1.0,
	}
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
)

const (
	pathAwesomeface  = "assets/textures/awesomeface.png"
	pathContainer2   = "assets/textures/container2.png"
	pathVertexShader = "assets/getting_started/6.coordinates/coordinate.frag"
	pathFragShader   = "assets/getting_started/6.coordinates/coordinate.vs"
)

type HelloCoordinates struct {
	vert         *canvas3d.VertexFloat32Array
	tex          *canvas3d.Texture
	coordinate   *canvas3d.Coordinate
	vertexShader string
	fragShader   string
}

func NewHelloCoordinates() base_ui.IChapter {
	r := &HelloCoordinates{
		tex:        canvas3d.NewTexture(),
		coordinate: canvas3d.NewCoordinate(),
		vert:       canvas3d.NewVertexFloat32Array(),
	}
	r.vert.Arr = vertices6
	r.vert.PositionSize = []int{3, 2}
	r.vert.TexCoordSize = []int{2, 3}
	return r
}

func (t *HelloCoordinates) initShader() {
	//if t.vertexShader != "" && t.fragShader != "" {
	//	return
	//}
	//data, err := os.ReadFile(pathVertexShader)
	//if err != nil {
	//	panic(err)
	//}
	//t.vertexShader = string(data)
	//
	//data, err = os.ReadFile(pathFragShader)
	//if err != nil {
	//	panic(err)
	//}
	//t.fragShader = string(data)
}

var (
	vertexShader3 = `
#version 330 core
layout (location = 0) in vec3 position;
layout (location = 1) in vec2 texCoord;

out vec2 TexCoord;

uniform mat4 model;
uniform mat4 view;
uniform mat4 projection;

void main()
{
    gl_Position = projection * view * model * vec4(position, 1.0f);
    TexCoord = vec2(texCoord.x, 1.0 - texCoord.y);
}`
	fragShader3 = `#version 330 core
in vec2 TexCoord;

out vec4 color;

uniform sampler2D texture1;
uniform sampler2D texture2;

void main()
{
    color = mix(texture(texture1, TexCoord), texture(texture2, TexCoord), 0.2);
}`
)

func (t *HelloCoordinates) InitChapterContent(c *base_ui.ChapterContent) {
	t.initShader()
	c.SetShaderConfig(vertexShader3, fragShader3)
	t.coordinate.UpdateFrameSize(c.WinSize)
	t.coordinate.TranslateVec3(cubePositions[0])
	t.tex.AppendPath(pathAwesomeface)
	t.tex.AppendPath(pathContainer2)
	c.AppendObj(t.coordinate)
	c.AppendObj(t.tex)
	c.AppendObj(t.vert)
}
func (t *HelloCoordinates) InitParamsContent(c *base_ui.ParamsContent) {

}
