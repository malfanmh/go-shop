package controller

import (
	"fmt"
	"github.com/kataras/iris/context"
)

type errorData struct {
	Code    string `json:"code"`
	Title   string `json:"title"`
	Detail  string `json:"detail"`
}

type paginationData struct {
	Next     string `json:"next,omitempty"`
	Previous string `json:"previous,omitempty"`
}

type Response struct {
	Pagination *paginationData `json:"pagination,omitempty"`
	Errors     *errorData      `json:"errors,omitempty"`
	Data       interface{}     `json:"data,omitempty"`
}

func NewResponse(data interface{}) *Response {
	return &Response{Data: data}
}

func (r *Response) SetPagination(ctx context.Context, page int, next bool) {
	u := ctx.Request().URL
	u.User = nil

	nextp := ""
	if next {
		q := u.Query()
		q.Set("page", fmt.Sprintf("%d", page+1))
		u.RawQuery = q.Encode()
		nextp = u.String()
	}

	prev := ""
	if page > 1 {
		q := u.Query()
		q.Set("page", fmt.Sprintf("%d", page-1))
		u.RawQuery = q.Encode()
		prev = u.String()
	}

	if nextp != "" || prev != "" {
		r.Pagination = &paginationData{nextp, prev}
	}
}

func (r *Response) AddError(code, title, detail string) {
	r.Errors = &errorData{code, title, detail}
}

func (e *errorData) ToString() string {
	str := fmt.Sprintf("%#v", e)
	return str
}

func RenderJSON(ctx context.Context, i interface{},code int){
	ctx.StatusCode(code)
	ctx.JSON(i)
	return
}