package model

type InternalRequest struct {
	Page    int
	Size    int
	Method  string
	Paths   []string
	Headers map[string]string
	Body    interface{}
}

type PaginationOpts struct {
	Page int `json:"page"`
	Size int `json:"size"`
}

type Filter struct {
	Value    string
	IsSended bool
}
