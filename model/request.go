package model

type InternalRequest struct {
	Body        interface{}
	Headers     map[string]string
	QueryParams map[string]string
	Method      string
	Paths       []string
	Pagination  PaginationOpts
}

type PaginationOpts struct {
	Page int
	Size int
}

type Filter struct {
	Value    string
	IsSended bool
}
