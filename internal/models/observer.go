package models

type Observer interface {
	Notify(data interface{})
}
