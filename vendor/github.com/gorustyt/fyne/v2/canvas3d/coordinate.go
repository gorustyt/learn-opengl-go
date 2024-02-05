package canvas3d

import "C"
import (
	"github.com/go-gl/mathgl/mgl64"
	"github.com/gorustyt/fyne/v2"
	"math"
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"github.com/gorustyt/fyne/v2/internal/painter/gl"
)

const (
	ProjectName = "projection"
	ViewName    = "view"
	ModelName   = "model"
)
const (
	Forward = iota
	Backward
	Left
	Right
)
const (
	Yaw       = -90.0
	Pitch     = 0.0
	Speed     = 3.0
	Sensitivy = 0.1
)

const (
	Near = 0.1
	Far  = 100
)

var (
	_ gl.Canvas3D              = (*Coordinate)(nil)
	_ fyne.Scrollable          = (*Coordinate)(nil)
	_ fyne.Draggable           = (*Coordinate)(nil)
	_ gl.Canvas3DBeforePainter = (*Coordinate)(nil)
)

type Coordinate struct {
	*ViewConfig
	*ProjectConfig
	*ModelConfig
	frameSize  fyne.Size
	firstMouse bool
}

func (c *Coordinate) UpdateFrameSize(frameSize fyne.Size) {
	if c.frameSize != frameSize {
		c.frameSize = frameSize
		c.LastX = frameSize.Width / 2
		c.LastY = frameSize.Height / 2
		c.firstMouse = true
	}

}

func (c *Coordinate) BeforeDraw(p *gl.Painter3D, pos fyne.Position, frame fyne.Size) {
	c.UpdateFrameSize(frame)
	project := mgl32.Perspective(mgl32.DegToRad(c.Fov), frame.Width/frame.Height, c.Near, c.Far)
	p.UniformMatrix4fv(p.Program(), ProjectName, project)
	fmt.Printf("p.GetUniformMatrix4fv(p.Program(), ProjectNamee)",p.GetUniformMatrix4fv(p.Program(), ProjectName))
	fmt.Println()
	fmt.Println(project)
}

func (c *Coordinate) Dragged(event *fyne.DragEvent) {
	xPos := event.AbsolutePosition.X
	yPos := event.AbsolutePosition.Y
	if c.firstMouse {
		c.LastX = xPos
		c.LastY = yPos
		c.firstMouse = false
	}
	xOffset := float64(xPos - c.LastX)
	yOffset := float64(c.LastY - yPos)
	c.LastX = xPos
	c.LastY = yPos
	xOffset *= c.MouseSensitivity
	yOffset *= c.MouseSensitivity

	c.Yaw += xOffset
	c.Pitch += yOffset

	// Make sure that when pitch is out of bounds, screen doesn't get flipped
	if c.ConstrainPitch {
		if c.Pitch > 89.0 {
			c.Pitch = 89.0
		}
		if c.Pitch < -89.0 {
			c.Pitch = -89.0
		}
	}

}

func (c *Coordinate) DragEnd() {
	x := float32(math.Cos(mgl64.DegToRad(c.Yaw)) * math.Cos(mgl64.DegToRad(c.Pitch)))
	y := float32(math.Sin(mgl64.DegToRad(c.Pitch)))
	z := float32(math.Sin(mgl64.DegToRad(c.Yaw)) * math.Cos(mgl64.DegToRad(c.Pitch)))
	front := mgl32.Vec3{x, y, z}
	front = front.Normalize()
	// Also re-calculate the Right and Up vector
	// Normalize the vectors, because their length gets closer to 0 the more you look up or down which results in slower movement.
	c.Right = front.Cross(c.WorldUp).Normalize()
	c.Up = c.Right.Cross(c.Front).Normalize()
}

func (c *Coordinate) Scrolled(event *fyne.ScrollEvent) {
	yOffset := event.AbsolutePosition.Y
	if c.Fov >= 1.0 && c.Fov <= 45. {
		c.Fov -= yOffset
		if c.Fov <= 1.0 {
			c.Fov = 1.0
		}
		if c.Fov >= 45.0 {
			c.Fov = 45.0
		}
	}

}

func (c *Coordinate) InitOnce(p *gl.Painter3D) {

}

func (c *Coordinate) Init(p *gl.Painter3D) {
	view:=c.GetView()
	p.UniformMatrix4fv(p.Program(), ViewName, view)
	fmt.Printf("p.GetUniformMatrix4fv(p.Program(), ViewName)",p.GetUniformMatrix4fv(p.Program(), ViewName))
	fmt.Println()
	fmt.Println(view)
	p.UniformMatrix4fv(p.Program(), ModelName, c.mat)
	fmt.Printf("p.GetUniformMatrix4fv(p.Program(), ModelName)",p.GetUniformMatrix4fv(p.Program(), ModelName))
	fmt.Println()
	fmt.Println(c.mat)

}

func (c *Coordinate) After(p *gl.Painter3D) {

}

func NewCoordinate() *Coordinate {
	c:=&Coordinate{
		firstMouse:  true,
		ModelConfig: NewModelConfig(),
		ProjectConfig: &ProjectConfig{
			Near: Near,
			Far:  Far,
			Fov:  45,
		},
		ViewConfig: &ViewConfig{
			Yaw:              Yaw,
			Pitch:            Pitch,
			Front:            mgl32.Vec3{0.0, 0.0, -1.0},
			Position:         mgl32.Vec3{0.0, 0.0, 3},
			WorldUp:          mgl32.Vec3{0, 1, 0},
			MovementSpeed:    Speed,
			MouseSensitivity: Sensitivy,
			ConstrainPitch:   true,
		},
	}
	c.Right = c.Front.Cross(c.WorldUp).Normalize()
	c.Up = c.Right.Cross(c.Front).Normalize()
	return c
}

type ViewConfig struct {
	Position mgl32.Vec3
	Front    mgl32.Vec3
	Up       mgl32.Vec3
	Right    mgl32.Vec3
	WorldUp  mgl32.Vec3

	Yaw   float64
	Pitch float64

	LastX float32
	LastY float32

	MovementSpeed    float64
	MouseSensitivity float64

	ConstrainPitch bool
}

func (c *ViewConfig) ProcessKeyboard(direction int, deltaTime float64) {
	velocity := float32(c.MovementSpeed * deltaTime)
	if direction == Forward {
		c.Position = c.Position.Add(c.Front.Mul(velocity))
	}
	if direction == Backward {
		c.Position = c.Position.Sub(c.Front.Mul(velocity))
	}
	if direction == Left {
		c.Position = c.Position.Sub(c.Right.Mul(velocity))
	}
	if direction == Right {
		c.Position = c.Position.Add(c.Right.Mul(velocity))
	}
}

func (c *ViewConfig) GetView() mgl32.Mat4 {
	eye := c.Position
	return mgl32.LookAtV(eye, eye.Add(c.Front), c.Up)
}

type ProjectConfig struct {
	Near float32
	Far  float32
	Fov  float32
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
