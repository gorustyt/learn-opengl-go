package canvas3d

import (
	"github.com/gorustyt/fyne/v2"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/gorustyt/fyne/v2/internal/painter/gl"
)

const (
	ProjectName = "project"
	ViewName    = "view"
	ModelName   = "model"
)

var _ gl.Canvas3D = (*Coordinate)(nil)

type Coordinate struct {
	Project mgl32.Mat4
	View    mgl32.Mat4
	Model   mgl32.Mat4
	*ViewConfig
	*ProjectConfig
	*ModelConfig
}

func (c Coordinate) Init(p *gl.Painter3D) {

}

func (c Coordinate) After(p *gl.Painter3D) {

}

func (c Coordinate) Draw(ctx *gl.Painter3D, pos fyne.Position, frame fyne.Size) {
	ctx.UniformMatrix4fv(ctx.Program(), ProjectName, c.Project)
	ctx.UniformMatrix4fv(ctx.Program(), ViewName, c.View)
	ctx.UniformMatrix4fv(ctx.Program(), ModelName, c.Model)
}

func NewCoordinate() *Coordinate {
	return &Coordinate{
		Project:       mgl32.Ident4(),
		View:          mgl32.Ident4(),
		Model:         mgl32.Ident4(),
		ModelConfig:   NewModelConfig(),
		ProjectConfig: &ProjectConfig{},
		ViewConfig:    &ViewConfig{},
	}
}

type ViewConfig struct {
	Eye    mgl32.Vec3
	Center mgl32.Vec3
	Up     mgl32.Vec3
}

type ProjectConfig struct {
	Near      float32
	Far       float32
	FrameSize fyne.Size
	Angle     float32
}

type ModelConfig struct {
	mat mgl32.Mat4
}

func NewModelConfig() *ModelConfig {
	return &ModelConfig{
		mat: mgl32.Ident4(),
	}
}
func (m *ModelConfig) TranslateXYZ(x, y, z float32) {
	m.mat = m.mat.Mul4(mgl32.Translate3D(x, y, z))
}

func (m *ModelConfig) TranslateVec3(vec mgl32.Vec3) {
	m.mat = m.mat.Mul4(mgl32.Translate3D(vec.X(), vec.Y(), vec.Z()))
}

func (m *ModelConfig) Rotate(angle float32, axis mgl32.Vec3) {
	m.mat = m.mat.Mul4(mgl32.HomogRotate3D(mgl32.DegToRad(angle), axis))
}

func (m *ModelConfig) Scale(x, y, z float32) {
	m.mat = m.mat.Mul4(mgl32.Scale3D(x, y, z))
}
