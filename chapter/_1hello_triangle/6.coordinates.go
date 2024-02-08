package _1hello_triangle

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gorustyt/fyne/v2"
	"github.com/gorustyt/fyne/v2/canvas3d"
	"github.com/gorustyt/fyne/v2/container"
	"github.com/gorustyt/fyne/v2/data/binding"
	"github.com/gorustyt/fyne/v2/widget"
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
	pathContainer2   = "assets/textures/container.jpg"
	pathVertexShader = "assets/getting_started/6.coordinates/coordinate.frag"
	pathFragShader   = "assets/getting_started/6.coordinates/coordinate.vs"
)

type HelloCoordinates struct {
	vert         *canvas3d.VertexFloat32Array
	tex          *canvas3d.Texture
	coordinate   *canvas3d.Coordinate
	vertexShader string
	fragShader   string
	menu         fyne.CanvasObject
	lastY        float32
}

func NewHelloCoordinates() base_ui.IChapter {
	r := &HelloCoordinates{
		tex:        canvas3d.NewTexture(),
		coordinate: canvas3d.NewCoordinate(),
		vert:       canvas3d.NewVertexFloat32Array(),
	}
	r.vert.Arr = vertices6
	r.vert.PositionSize = []int{3, 0}
	r.vert.TexCoordSize = []int{2, 3}
	r.initMenu()
	r.initShader()
	r.coordinate.TranslateVec3(cubePositions[0])
	r.coordinate.Rotate(20*2, mgl32.Vec3{1.0, 0.3, 0.5})
	r.tex.AppendPath(pathContainer2)
	r.tex.AppendPath(pathAwesomeface)
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
	//gl_Position =vec4(position, 1.0f);
    TexCoord = vec2(texCoord.x, 1.0 - texCoord.y);
}`
	fragShader3 = `#version 330 core
in vec2 TexCoord;

out vec4 color;
uniform float mixParams;
uniform sampler2D texture1;
uniform sampler2D texture2;

void main()
{
    color = mix(texture(texture1, TexCoord), texture(texture2, TexCoord), mixParams);
}`
)

func (t *HelloCoordinates) InitChapterContent(c *base_ui.ChapterContent) {
	c.Painter(0).SetShaderConfig(vertexShader3, fragShader3)
	t.coordinate.UpdateFrameSize(c.WinSize)
	c.Painter(0).AppendObj(t.vert)
	c.Painter(0).AppendObj(t.coordinate)
	c.Painter(0).AppendObj(t.tex)
}

func (t *HelloCoordinates) initMenu() {
	f := 2.
	data := binding.BindFloat(&f)
	s := widget.NewSliderWithData(1, 10, data)
	s.Step = 1
	s.OnChanged = func(f float64) {
		//offsetY := float32(f) - t.lastY
		//t.coordinate.TranslateXYZ(0, offsetY, 0)
		//t.lastY = float32(f)
	}
	s1 := widget.NewSlider(0, 1)
	s1.Step = 0.1
	s1.OnChanged = func(f float64) {
		t.tex.MixParams = float32(f)
	}
	t.menu = container.NewVBox(
		widget.NewLabel("height"),
		s,
		widget.NewLabel("mix"),
		s1,
	)
}

func (t *HelloCoordinates) InitParamsContent(c *base_ui.ParamsContent) {
	c.Add(t.menu)
}
