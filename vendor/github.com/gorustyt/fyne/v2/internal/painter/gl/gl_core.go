//go:build (!gles && !arm && !arm64 && !android && !ios && !mobile && !js && !test_web_driver && !wasm) || (darwin && !mobile && !ios && !js && !wasm && !test_web_driver)
// +build !gles,!arm,!arm64,!android,!ios,!mobile,!js,!test_web_driver,!wasm darwin,!mobile,!ios,!js,!wasm,!test_web_driver

package gl

import (
	"github.com/go-gl/mathgl/mgl32"
	"image"
	"image/draw"
	"strings"

	"github.com/go-gl/gl/v4.2-core/gl"

	"github.com/gorustyt/fyne/v2"
)

const (
	arrayBuffer           = gl.ARRAY_BUFFER
	bitColorBuffer        = gl.COLOR_BUFFER_BIT
	bitDepthBuffer        = gl.DEPTH_BUFFER_BIT
	clampToEdge           = gl.CLAMP_TO_EDGE
	colorFormatRGBA       = gl.RGBA
	compileStatus         = gl.COMPILE_STATUS
	constantAlpha         = gl.CONSTANT_ALPHA
	float                 = gl.FLOAT
	fragmentShader        = gl.FRAGMENT_SHADER
	front                 = gl.FRONT
	glFalse               = gl.FALSE
	linkStatus            = gl.LINK_STATUS
	one                   = gl.ONE
	oneMinusConstantAlpha = gl.ONE_MINUS_CONSTANT_ALPHA
	oneMinusSrcAlpha      = gl.ONE_MINUS_SRC_ALPHA
	scissorTest           = gl.SCISSOR_TEST
	srcAlpha              = gl.SRC_ALPHA
	staticDraw            = gl.STATIC_DRAW
	texture0              = gl.TEXTURE0
	texture2D             = gl.TEXTURE_2D
	textureMinFilter      = gl.TEXTURE_MIN_FILTER
	textureMagFilter      = gl.TEXTURE_MAG_FILTER
	textureWrapS          = gl.TEXTURE_WRAP_S
	textureWrapT          = gl.TEXTURE_WRAP_T
	triangles             = gl.TRIANGLES
	triangleStrip         = gl.TRIANGLE_STRIP
	unsignedByte          = gl.UNSIGNED_BYTE
	vertexShader          = gl.VERTEX_SHADER
)

const noBuffer = Buffer(0)
const noShader = Shader(0)

type (
	// Attribute represents a GL attribute
	Attribute int32
	// Buffer represents a GL buffer
	Buffer uint32
	// Program represents a compiled GL program
	Program uint32
	// Shader represents a GL shader
	Shader uint32
	// Uniform represents a GL uniform
	Uniform int32
)

var textureFilterToGL = []int32{gl.LINEAR, gl.NEAREST, gl.LINEAR}

func (p *painter) Init() {
	p.ctx = &coreContext{}
	err := gl.Init()
	if err != nil {
		fyne.LogError("failed to initialise OpenGL", err)
		return
	}

	gl.Disable(gl.DEPTH_TEST)
	gl.Enable(gl.BLEND)
	p.logError()
	p.program = p.createProgram("simple")
	p.lineProgram = p.createProgram("line")
	p.rectangleProgram = p.createProgram("rectangle")
	p.roundRectangleProgram = p.createProgram("round_rectangle")
}

type coreContext struct{}

var _ context = (*coreContext)(nil)

func (c *coreContext) ActiveTexture(textureUnit uint32) {
	gl.ActiveTexture(textureUnit)
}

func (c *coreContext) AttachShader(program Program, shader Shader) {
	gl.AttachShader(uint32(program), uint32(shader))
}

func (c *coreContext) BindBuffer(target uint32, buf Buffer) {
	gl.BindBuffer(target, uint32(buf))
}

func (c *coreContext) BindTexture(target uint32, texture Texture) {
	gl.BindTexture(target, uint32(texture))
}

func (c *coreContext) BlendColor(r, g, b, a float32) {
	gl.BlendColor(r, g, b, a)
}

func (c *coreContext) BlendFunc(srcFactor, destFactor uint32) {
	gl.BlendFunc(srcFactor, destFactor)
}

func (c *coreContext) BufferData(target uint32, points []float32, usage uint32) {
	gl.BufferData(target, 4*len(points), gl.Ptr(points), usage)
}

func (c *coreContext) Clear(mask uint32) {
	gl.Clear(mask)
}

func (c *coreContext) ClearColor(r, g, b, a float32) {
	gl.ClearColor(r, g, b, a)
}

func (c *coreContext) CompileShader(shader Shader) {
	gl.CompileShader(uint32(shader))
}

func (c *coreContext) CreateBuffer() Buffer {
	var vbo uint32
	gl.GenBuffers(1, &vbo)
	return Buffer(vbo)
}

func (c *coreContext) CreateProgram() Program {
	return Program(gl.CreateProgram())
}

func (c *coreContext) CreateShader(typ uint32) Shader {
	return Shader(gl.CreateShader(typ))
}

func (c *coreContext) CreateTexture() (texture Texture) {
	var tex uint32
	gl.GenTextures(1, &tex)
	return Texture(tex)
}

func (c *coreContext) DeleteBuffer(buffer Buffer) {
	gl.DeleteBuffers(1, (*uint32)(&buffer))
}

func (c *coreContext) DeleteTexture(texture Texture) {
	tex := uint32(texture)
	gl.DeleteTextures(1, &tex)
}

func (c *coreContext) Disable(capability uint32) {
	gl.Disable(capability)
}

func (c *coreContext) DrawElementsArrays(mode uint32, index []uint32) {
	gl.DrawElements(mode, int32(len(index)), gl.UNSIGNED_INT, gl.Ptr(index))
}

func (c *coreContext) DrawArrays(mode uint32, first, count int) {
	gl.DrawArrays(mode, int32(first), int32(count))
}

func (c *coreContext) Enable(capability uint32) {
	gl.Enable(capability)
}

func (c *coreContext) EnableVertexAttribArray(attribute Attribute) {
	gl.EnableVertexAttribArray(uint32(attribute))
}

func (c *coreContext) GetAttribLocation(program Program, name string) Attribute {
	return Attribute(gl.GetAttribLocation(uint32(program), gl.Str(name+"\x00")))
}

func (c *coreContext) GetError() uint32 {
	return gl.GetError()
}

func (c *coreContext) GetProgrami(program Program, param uint32) int {
	var value int32
	gl.GetProgramiv(uint32(program), param, &value)
	return int(value)
}

func (c *coreContext) GetProgramInfoLog(program Program) string {
	var logLength int32
	gl.GetProgramiv(uint32(program), gl.INFO_LOG_LENGTH, &logLength)
	info := strings.Repeat("\x00", int(logLength+1))
	gl.GetProgramInfoLog(uint32(program), logLength, nil, gl.Str(info))
	return info
}

func (c *coreContext) GetShaderi(shader Shader, param uint32) int {
	var value int32
	gl.GetShaderiv(uint32(shader), param, &value)
	return int(value)
}

func (c *coreContext) GetShaderInfoLog(shader Shader) string {
	var logLength int32
	gl.GetShaderiv(uint32(shader), gl.INFO_LOG_LENGTH, &logLength)
	info := strings.Repeat("\x00", int(logLength+1))
	gl.GetShaderInfoLog(uint32(shader), logLength, nil, gl.Str(info))
	return info
}

func (c *coreContext) GetUniformLocation(program Program, name string) Uniform {
	return Uniform(gl.GetUniformLocation(uint32(program), gl.Str(name+"\x00")))
}

func (c *coreContext) LinkProgram(program Program) {
	gl.LinkProgram(uint32(program))
}

func (c *coreContext) ReadBuffer(src uint32) {
	gl.ReadBuffer(src)
}

func (c *coreContext) ReadPixels(x, y, width, height int, colorFormat, typ uint32, pixels []uint8) {
	gl.ReadPixels(int32(x), int32(y), int32(width), int32(height), colorFormat, typ, gl.Ptr(pixels))
}

func (c *coreContext) Scissor(x, y, w, h int32) {
	gl.Scissor(x, y, w, h)
}

func (c *coreContext) ShaderSource(shader Shader, source string) {
	csources, free := gl.Strs(source + "\x00")
	defer free()
	gl.ShaderSource(uint32(shader), 1, csources, nil)
}

func (c *coreContext) TexImage2D(target uint32, level, width, height int, colorFormat, typ uint32, data []uint8) {
	gl.TexImage2D(
		target,
		int32(level),
		int32(colorFormat),
		int32(width),
		int32(height),
		0,
		colorFormat,
		typ,
		gl.Ptr(data),
	)
}

func (c *coreContext) TexParameteri(target, param uint32, value int32) {
	gl.TexParameteri(target, param, value)
}

func (c *coreContext) Uniform1f(uniform Uniform, v float32) {
	gl.Uniform1f(int32(uniform), v)
}

func (c *coreContext) Uniform2f(uniform Uniform, v0, v1 float32) {
	gl.Uniform2f(int32(uniform), v0, v1)
}
func (c *coreContext) Uniform3f(uniform Uniform, v mgl32.Vec3) {
	gl.Uniform3f(int32(uniform), v[0], v[1], v[2])
}

func (c *coreContext) Uniform4f(uniform Uniform, v0, v1, v2, v3 float32) {
	gl.Uniform4f(int32(uniform), v0, v1, v2, v3)
}

func (c *coreContext) UseProgram(program Program) {
	gl.UseProgram(uint32(program))
}

func (c *coreContext) VertexAttribPointerWithOffset(attribute Attribute, size int, typ uint32, normalized bool, stride, offset int) {
	gl.VertexAttribPointerWithOffset(uint32(attribute), int32(size), typ, normalized, int32(stride), uintptr(offset))
}

func (c *coreContext) Viewport(x, y, width, height int) {
	gl.Viewport(int32(x), int32(y), int32(width), int32(height))
}

func (c *coreContext) UniformMatrix4fv(program Program, name string, mat4 mgl32.Mat4) {
	gl.UniformMatrix4fv(int32(c.GetUniformLocation(program, name)), 1, false, &mat4[0])
}

func (c *coreContext) Uniform1i(program Program, name string, v0 int32) {
	gl.Uniform1i(int32(c.GetUniformLocation(program, name)), v0)
}

func (c *coreContext) EnableDepthTest() {
	gl.Enable(gl.DEPTH_TEST)
	gl.DepthFunc(gl.LESS)
}

func (c *coreContext) DisableDepthTest() {
	gl.Disable(gl.DEPTH_TEST)
}

func (c *coreContext) MakeVaoWithEbo(points []float32, indexs []uint32) Buffer {
	c.CreateBuffer()
	c.BufferData(gl.ARRAY_BUFFER, points, gl.STATIC_DRAW)

	var ebo uint32
	gl.GenBuffers(1, &ebo)
	gl.BindBuffer(gl.ELEMENT_ARRAY_BUFFER, ebo)
	gl.BufferData(gl.ELEMENT_ARRAY_BUFFER, 4*len(indexs), gl.Ptr(indexs), gl.STATIC_DRAW)

	var vao uint32
	gl.GenVertexArrays(1, &vao)
	gl.BindVertexArray(vao)
	return Buffer(ebo)
}

func (c *coreContext) MakeVao(points []float32) Buffer {
	var vbo uint32

	// 在显卡中开辟一块空间，创建顶点缓存对象，个数为1，变量vbo会被赋予一个ID值。
	gl.GenBuffers(1, &vbo)

	// 将 vbo 赋值给 gl.ARRAY_BUFFER，要知道这个对象会被赋予不同的vbo，因此其值是变化的
	// 可选类型：GL_ARRAY_BUFFER, GL_ELEMENT_ARRAY_BUFFER, GL_PIXEL_PACK_BUFFER, GL_PIXEL_UNPACK_BUFFER
	gl.BindBuffer(gl.ARRAY_BUFFER, vbo)
	// 将内存中的数据传递到显卡中的gl.ARRAY_BUFFER对象上，其实是把数据传递到绑定在其上面的vbo对象上。
	// 4*len(points) 代表总的字节数，因为是32位的
	gl.BufferData(gl.ARRAY_BUFFER, 4*len(points), gl.Ptr(points), gl.STATIC_DRAW)

	var vao uint32
	// 创建顶点数组对象，个数为1，变量vao会被赋予一个ID值。
	gl.GenVertexArrays(1, &vao)
	// 后面的两个函数都是要操作具体的vao的，因此需要先将vao绑定到opengl上。
	// 解绑：gl.BindVertexArray(0)，opengl中很多的解绑操作都是传入0
	gl.BindVertexArray(vao)
	return Buffer(vbo)
}

func (c *coreContext) MakeTexture(img image.Image, index uint32) Texture {
	rgba := image.NewRGBA(img.Bounds())
	if rgba.Stride != rgba.Rect.Size().X*4 {
		panic("unsupported stride")
	}

	draw.Draw(rgba, rgba.Bounds(), img, image.Point{0, 0}, draw.Src)
	var te uint32
	gl.GenTextures(1, &te)
	gl.ActiveTexture(index)
	gl.BindTexture(gl.TEXTURE_2D, te)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.NEAREST)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.NEAREST)
	gl.TexImage2D(gl.TEXTURE_2D, 0, gl.RGBA, int32(rgba.Rect.Size().X), int32(rgba.Rect.Size().Y), 0, gl.RGBA, gl.UNSIGNED_BYTE, gl.Ptr(rgba.Pix))
	gl.GenerateMipmap(gl.TEXTURE_2D)
	return Texture(te)
}

func GetTextureByIndex(index int) uint32 {
	switch index {
	case 0:
		return gl.TEXTURE0
	case 1:
		return gl.TEXTURE1
	case 2:
		return gl.TEXTURE2
	case 3:
		return gl.TEXTURE3
	case 4:
		return gl.TEXTURE4
	case 5:
		return gl.TEXTURE5
	case 6:
		return gl.TEXTURE6
	case 7:
		return gl.TEXTURE7
	case 8:
		return gl.TEXTURE8
	case 9:
		return gl.TEXTURE9
	case 10:
		return gl.TEXTURE10
	}
	return 0
}
