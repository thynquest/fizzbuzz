package status

import (
	"net/http"

	"github.com/go-chi/render"
)

type HttpResp struct {
	StatusCode int    `json:"-"`
	StatusText string `json:"message,omitempty"`
	ErrorText  string `json:"error,omitempty"`
}

func (hr *HttpResp) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, hr.StatusCode)
	return nil
}

func BadRequest(err error) render.Renderer {
	return &HttpResp{
		StatusCode: http.StatusBadRequest,
		StatusText: "Bad request.",
		ErrorText:  err.Error(),
	}
}

func ErrRender(msg string, err error) render.Renderer {
	return &HttpResp{
		StatusCode: http.StatusUnprocessableEntity,
		StatusText: msg,
		ErrorText:  err.Error(),
	}
}
