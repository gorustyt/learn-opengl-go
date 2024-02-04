package canvas3d

import (
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gorustyt/fyne/v2"
	"github.com/gorustyt/fyne/v2/internal/painter/gl"
)

var _ gl.Canvas3D = (*Material)(nil)

type Material struct {
	Ambient   mgl32.Vec3
	Diffuse   mgl32.Vec3
	Specular  mgl32.Vec3
	Shininess float32

	AmbientStrength  float32
	DiffuseStrength  float32
	SpecularStrength float32
}

func (m *Material) Init(p *gl.Painter3D) {

}

func (m *Material) After(p *gl.Painter3D) {

}

func NewMaterial() *Material {
	return &Material{
		Shininess:        1,
		AmbientStrength:  1,
		DiffuseStrength:  1,
		SpecularStrength: 1,
	}
}

func (m *Material) Draw(ctx *gl.Painter3D, pos fyne.Position, frame fyne.Size) {
	ctx.Uniform3f(ctx.GetUniformLocation(ctx.Program(), "material.ambient"), m.Ambient)
	ctx.Uniform3f(ctx.GetUniformLocation(ctx.Program(), "material.diffuse"), m.Diffuse)
	ctx.Uniform3f(ctx.GetUniformLocation(ctx.Program(), "material.specular"), m.Specular)

	ctx.Uniform1f(ctx.GetUniformLocation(ctx.Program(), "material.ambient_strength"), m.AmbientStrength)
	ctx.Uniform1f(ctx.GetUniformLocation(ctx.Program(), "material.diffuse_strength"), m.DiffuseStrength)
	ctx.Uniform1f(ctx.GetUniformLocation(ctx.Program(), "material.specular_strength"), m.SpecularStrength)

}
