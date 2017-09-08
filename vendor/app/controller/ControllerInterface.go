package controller

type ControllerInterface interface {
	Handle() (err error)
}
