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
    FragColor = vec4(1.0,1.0,1.0,1.0); // set all 4 vector values to 1.0
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
	lightCoord *canvas3d.Coordinate
	lightVert  *canvas3d.VertexFloat32Array

	cubeVert  *canvas3d.VertexFloat32Array
	cubeCoord *canvas3d.Coordinate
}

func (l *Lighting) InitChapterContent(c *base_ui.ChapterContent) {
	c.Canvas3d().SetShaderConfig(0, lightCube_vs, lightCube_fs)
	c.Canvas3d().AppendObj(0, l.lightCoord)
	c.Canvas3d().AppendObj(0, l.lightVert)

	c.Canvas3d().SetShaderConfig(1, color_vs, color_fs)
	c.Canvas3d().AppendObj(1, l.cubeVert)
	c.Canvas3d().AppendObj(1, l.cubeCoord)

	c.Canvas3d().AppendRenderFunc(1, func(ctx context.Painter) {
		ctx.UniformVec3("lightColor", mgl32.Vec3{1, 1, 1})
		ctx.UniformVec3("objectColor", mgl32.Vec3{1.0, 0.5, 0.31})
	})
}

func (l *Lighting) InitParamsContent(c *base_ui.ParamsContent) {

}

func NewLighting() base_ui.IChapter {
	l := &Lighting{
		lightCoord: canvas3d.NewCoordinate(),
		lightVert:  canvas3d.NewVertexFloat32Array(),
		cubeCoord:  canvas3d.NewCoordinate(),
		cubeVert:   canvas3d.NewVertexFloat32Array()}
	l.cubeVert.Arr = vert
	l.cubeVert.PositionSize = []int{3, 0}
	l.cubeCoord.Scale(0.5, 0.5, 0.5)

	l.lightVert.Arr = vert
	l.lightVert.PositionSize = []int{3, 0}
	l.lightCoord.TranslateXYZ(0.5, 0.5, 0.5)
	l.lightCoord.Scale(0.1, 0.1, 0.1)
	return l
}
