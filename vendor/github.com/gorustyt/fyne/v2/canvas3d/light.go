package canvas3d

import (
	"github.com/go-gl/mathgl/mgl32"
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
	m.Material.Init(p)
	p.Uniform3f(p.GetUniformLocation(p.Program(), LightDirection), m.Position)
}

func (m *Light) After(p *gl.Painter3D) {

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
	m.Light.Init(p)
	p.Uniform1f(p.GetUniformLocation(p.Program(), LightConstant), m.Constant)
	p.Uniform1f(p.GetUniformLocation(p.Program(), LightLinear), m.Linear)
	p.Uniform1f(p.GetUniformLocation(p.Program(), LightQuadratic), m.Quadratic)
	p.Uniform3f(p.GetUniformLocation(p.Program(), LightPosition), m.Position)
}

func (m *PointLight) After(p *gl.Painter3D) {

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
	m.Light.Init(p)
	p.Uniform1f(p.GetUniformLocation(p.Program(), LightCutOff), m.CutOff)
	p.Uniform3f(p.GetUniformLocation(p.Program(), LightPosition), m.Position)
}

func (m *SpotLight) After(p *gl.Painter3D) {

}

type DirectionLight struct {
	*Light
}

func (m *DirectionLight) Init(p *gl.Painter3D) {
	m.Light.Init(p)
}
func (m *DirectionLight) After(p *gl.Painter3D) {

}
func NewDirectionLight() *DirectionLight {
	return &DirectionLight{
		Light: NewLight(),
	}
}
