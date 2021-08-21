package render_types


type Window interface {
	IsOpen() bool
	Close()
	OnUpdate()

	IsVSync() bool
	SetVSync(bool)

	GetBaseWindow() interface{}
}