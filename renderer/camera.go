package renderer

import (
	"fmt"
	math32 "github.com/go-gl/mathgl/mgl32"
)

type baseCamera struct {
	viewMatrix, projectionMatrix, viewProjectionMatrix math32.Mat4
	cameraPosition                                     math32.Vec3
	nearPlane, farPlane                                float32
}

type PerspectiveCamera struct {
	baseCamera
	FixedTarget    math32.Vec3
	HasFixedTarget bool

	Yaw, Roll, Pitch                       float32
	FieldOfView, AspectRatio               float32
	CameraUp, CameraDirection, CameraRight math32.Vec3
	WorldUp                                math32.Vec3
}

type OrthographicCamera struct {
	baseCamera
	rotation, zoom float32
}

func CreateOrthographicCamera(left, right, bottom, top float32) OrthographicCamera {
	projMatrix := math32.Ortho(left, right, bottom, top, -1.0, 1.0)
	viewMatrix := math32.Ident4()
	viewProj := projMatrix.Mul4(viewMatrix)
	return OrthographicCamera{
		baseCamera: baseCamera{
			viewMatrix:           viewMatrix,
			projectionMatrix:     projMatrix,
			viewProjectionMatrix: viewProj,
			cameraPosition:       math32.Vec3{0, 0, 0},
			nearPlane:            -1.0,
			farPlane:             1.0,
		},
		rotation: 0,
		zoom:     0,
	}

}

func (oc *OrthographicCamera) SetPosition(position math32.Vec3) {
	oc.cameraPosition = position
	oc.calculateViewMatrix()
}

func (oc OrthographicCamera) GetPosition() math32.Vec3 {
	return oc.cameraPosition
}

func (oc *OrthographicCamera) Move(dx, dy, dz float32) {
	oc.cameraPosition[0] += dx
	oc.cameraPosition[1] += dy
	oc.cameraPosition[2] += dz
	oc.calculateViewMatrix()
}

func (oc *OrthographicCamera) Rotate(dDegrees float32) {
	oc.rotation += dDegrees
	oc.calculateViewMatrix()
}

//Angles
func (oc *OrthographicCamera) SetRotation(a float32) {
	oc.rotation = a
	oc.calculateViewMatrix()
}

func (oc OrthographicCamera) GetRotation() float32 {
	return oc.rotation
}

func (oc *OrthographicCamera) SetZoom(a float32) {
	oc.zoom = a
	oc.calculateViewMatrix()
	panic("not implemented")
}

func (oc *OrthographicCamera) SetFrustum(left, right, bottom, top, far, near float32) {
	oc.projectionMatrix = math32.Ortho(left, right, bottom, top, near, far)
	oc.viewProjectionMatrix = oc.projectionMatrix.Mul4(oc.viewMatrix)
}

func (oc OrthographicCamera) GetZoom() float32 {
	return oc.zoom
}

func (oc OrthographicCamera) GetViewMatrix() math32.Mat4 {
	return oc.viewMatrix
}
func (oc OrthographicCamera) GetProjectionMatrix() math32.Mat4 {
	return oc.projectionMatrix
}
func (oc OrthographicCamera) GetViewProjectionMatrix() *math32.Mat4 {
	return &oc.viewProjectionMatrix
}

func (oc *OrthographicCamera) calculateViewMatrix() {
	IM := math32.Ident4()
	zAxis := math32.Vec3{0, 0, 1}

	rotationMatrix := math32.HomogRotate3D(math32.DegToRad(oc.rotation), zAxis)
	translationMatrix := math32.Translate3D(oc.cameraPosition[0], oc.cameraPosition[1], oc.cameraPosition[2])
	IM = IM.Mul4(rotationMatrix).Mul4(translationMatrix)

	oc.viewMatrix = IM.Inv()
	oc.viewProjectionMatrix = oc.projectionMatrix.Mul4(oc.viewMatrix)
	fmt.Println(oc.viewMatrix)
}
