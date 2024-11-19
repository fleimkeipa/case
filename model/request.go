package model

type InternalRequest struct {
	Pagination  PaginationOpts
	Method      string
	Paths       []string
	Headers     map[string]string
	Body        interface{}
	QueryParams map[string]string
}

type PaginationOpts struct {
	Page int
	Size int
}

type Filter struct {
	Value    string
	IsSended bool
}
