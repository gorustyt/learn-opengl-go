package _2lighting

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gorustyt/fyne/v2/canvas3d"
	"github.com/gorustyt/fyne/v2/canvas3d/context"
	"github.com/gorustyt/learn-opengl-go/ui/base_ui"
)

var (
	lightCube_vs = `#version 330 core
layout (location = 0) in vec3 position;

uniform mat4 model;
uniform mat4 view;
uniform mat4 projection;

void main()
{
	gl_Position = projection * view * model * vec4(position, 1.0);

}`
	lightCube_fs = `#version 330 core
out vec4 FragColor;

void main()
{
    FragColor = vec4(0,0.5,0,1.0); // set all 4 vector values to 1.0
}`
	color_vs = `#version 330 core
layout (location = 0) in vec3 position;

uniform mat4 model;
uniform mat4 view;
uniform mat4 projection;

void main()
{
	gl_Position = projection * view * model * vec4(position, 1.0);
}`
	color_fs = `#version 330 core
out vec4 FragColor;
  
uniform vec3 objectColor;
uniform vec3 lightColor;

void main()
{
    FragColor = vec4(lightColor * objectColor, 1.0);
}`
	vert = []float32{
		-0.5, -0.5, -0.5,
		0.5, -0.5, -0.5,
		0.5, 0.5, -0.5,
		0.5, 0.5, -0.5,
		-0.5, 0.5, -0.5,
		-0.5, -0.5, -0.5,

		-0.5, -0.5, 0.5,
		0.5, -0.5, 0.5,
		0.5, 0.5, 0.5,
		0.5, 0.5, 0.5,
		-0.5, 0.5, 0.5,
		-0.5, -0.5, 0.5,

		-0.5, 0.5, 0.5,
		-0.5, 0.5, -0.5,
		-0.5, -0.5, -0.5,
		-0.5, -0.5, -0.5,
		-0.5, -0.5, 0.5,
		-0.5, 0.5, 0.5,

		0.5, 0.5, 0.5,
		0.5, 0.5, -0.5,
		0.5, -0.5, -0.5,
		0.5, -0.5, -0.5,
		0.5, -0.5, 0.5,
		0.5, 0.5, 0.5,

		-0.5, -0.5, -0.5,
		0.5, -0.5, -0.5,
		0.5, -0.5, 0.5,
		0.5, -0.5, 0.5,
		-0.5, -0.5, 0.5,
		-0.5, -0.5, -0.5,

		-0.5, 0.5, -0.5,
		0.5, 0.5, -0.5,
		0.5, 0.5, 0.5,
		0.5, 0.5, 0.5,
		-0.5, 0.5, 0.5,
		-0.5, 0.5, -0.5,
	}
)

type Lighting struct {
	light      *canvas3d.Light
	lightCoord *canvas3d.Coordinate
	lightVert  *canvas3d.VertexFloat32Array

	cubeVert  *canvas3d.VertexFloat32Array
	cubeCoord *canvas3d.Coordinate
}

func (l *Lighting) InitChapterContent(c *base_ui.ChapterContent) {
	c.Painter(0).SetShaderConfig(lightCube_vs, lightCube_fs)
	l.lightCoord.UpdateFrameSize(c.WinSize)
	c.Painter(0).AppendObj(l.lightCoord)
	c.Painter(0).AppendObj(l.lightVert)

	c.Painter(1).SetShaderConfig(color_vs, color_fs)
	c.Painter(1).AppendObj(l.cubeVert)
	c.Painter(1).AppendObj(l.cubeCoord)
	l.cubeCoord.Scale(0.5, 0.5, 0.5)
	c.Painter(1).AppendRenderFunc(func(ctx context.Painter) {
		ctx.UniformVec3("lightColor", mgl32.Vec3{1, 1, 1})
		ctx.UniformVec3("objectColor", mgl32.Vec3{1.0, 0.5, 0.31})
	})
}

func (l *Lighting) InitParamsContent(c *base_ui.ParamsContent) {

}

func NewLighting() base_ui.IChapter {
	l := &Lighting{
		light:      canvas3d.NewLight(),
		lightCoord: canvas3d.NewCoordinate(),
		lightVert:  canvas3d.NewVertexFloat32Array(),
		cubeCoord:  canvas3d.NewCoordinate(),
		cubeVert:   canvas3d.NewVertexFloat32Array()}
	l.light.Position = mgl32.Vec3{1.2, 1, 2.0}
	l.cubeVert.Arr = vert
	l.cubeVert.PositionSize = []int{3, 0}

	l.lightVert.Arr = vert
	l.lightVert.PositionSize = []int{3, 0}
	l.lightCoord.TranslateXYZ(0.5, 0.5, 0.5)
	l.lightCoord.Scale(0.1, 0.1, 0.1)
	return l
}
