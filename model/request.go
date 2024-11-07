package model

type Request struct {
	Page    int
	Size    int
	Method  string
	Paths   []string
	Headers map[string]string
	Body    interface{}
}
