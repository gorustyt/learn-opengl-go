package canvas3d

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gorustyt/fyne/v2"
	"github.com/gorustyt/fyne/v2/internal/painter/gl"
)

const (
	LightConstant  = "light.constant"
	LightLinear    = "light.linear"
	LightQuadratic = "light.quadratic"
	LightCutOff    = "light.cutOff"
	LightDirection = "light.direction"
	LightPosition  = "light.position"
)

type Light struct {
	Position  mgl32.Vec3
	Direction mgl32.Vec3
	*Material
}

func (m *Light) Init(p *gl.Painter3D) {

}

func (m *Light) After(p *gl.Painter3D) {

}

func (m *Light) Draw(ctx *gl.Painter3D, pos fyne.Position, frame fyne.Size) {
	m.Material.Draw(ctx, pos, frame)
	ctx.Uniform3f(ctx.GetUniformLocation(ctx.Program(), LightDirection), m.Position)
}

func NewLight() *Light {
	return &Light{
		Material: NewMaterial(),
	}
}

type PointLight struct {
	*Light
	Constant  float32
	Linear    float32
	Quadratic float32
}

func NewPointLight() *PointLight {
	return &PointLight{
		Light: NewLight(),
	}
}
func (m *PointLight) Init(p *gl.Painter3D) {

}

func (m *PointLight) After(p *gl.Painter3D) {

}
func (m *PointLight) Draw(ctx *gl.Painter3D, pos fyne.Position, frame fyne.Size) {
	m.Light.Draw(ctx, pos, frame)
	ctx.Uniform1f(ctx.GetUniformLocation(ctx.Program(), LightConstant), m.Constant)
	ctx.Uniform1f(ctx.GetUniformLocation(ctx.Program(), LightLinear), m.Linear)
	ctx.Uniform1f(ctx.GetUniformLocation(ctx.Program(), LightQuadratic), m.Quadratic)
	ctx.Uniform3f(ctx.GetUniformLocation(ctx.Program(), LightPosition), m.Position)
}

type SpotLight struct {
	*Light
	CutOff float32
}

func NewSpotLight() *SpotLight {
	return &SpotLight{
		Light: NewLight(),
	}
}
func (m *SpotLight) Init(p *gl.Painter3D) {

}

func (m *SpotLight) After(p *gl.Painter3D) {

}
func (m *SpotLight) Draw(ctx *gl.Painter3D, pos fyne.Position, frame fyne.Size) {
	m.Light.Draw(ctx, pos, frame)
	ctx.Uniform1f(ctx.GetUniformLocation(ctx.Program(), LightCutOff), m.CutOff)
	ctx.Uniform3f(ctx.GetUniformLocation(ctx.Program(), LightPosition), m.Position)
}

type DirectionLight struct {
	*Light
}

func (m *DirectionLight) Init(p *gl.Painter3D) {

}
func (m *DirectionLight) After(p *gl.Painter3D) {

}
func NewDirectionLight() *DirectionLight {
	return &DirectionLight{
		Light: NewLight(),
	}
}
func (m *DirectionLight) Draw(ctx *gl.Painter3D, pos fyne.Position, frame fyne.Size) {
	m.Light.Draw(ctx, pos, frame)
}
