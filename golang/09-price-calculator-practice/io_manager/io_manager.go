package io_manager

type IOManager interface {
	Read() ([]string, error)
	Write(data interface{}) error
}
