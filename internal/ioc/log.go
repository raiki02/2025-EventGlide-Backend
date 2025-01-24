package ioc

type LogHdl interface {
}

type Log struct {
}

func NewLog() *Log {
	return &Log{}
}
