package http

import (
	errorkit "myapp/core/initialize/errors"
	"myapp/core/utils"
	"reflect"
	"strconv"

	"github.com/labstack/echo/v4"
)

type EncodingResponse struct {
	Data     interface{} `json:"data"`
	PageInfo interface{} `json:"pageInfo,omitempty"`
	Error    []error     `json:"errors"`
}

func EncodeJsonResponse(data, pageInfo interface{}, err interface{}) EncodingResponse {
	result := EncodingResponse{
		Data:     data,
		PageInfo: pageInfo,
	}
	if err != nil {
		if reflect.TypeOf(err).Kind() == reflect.Slice {
			result.Error = err.([]error)
		} else {
			result.Error = append(result.Error, err.(error))
		}
	}

	return result
}

func DecodePagingRequest(page, perPage string) (int, int) {
	p := 1
	l := 10
	if pageParse, err := strconv.Atoi(page); err == nil {
		p = pageParse
		if p <= 0 {
			p = 1
		}
	}
	if perPageParse, err := strconv.Atoi(perPage); err == nil {
		l = perPageParse
		if l > 1000 {
			l = 1000
		}
		if l <= 0 {
			l = 10
		}
	}
	return p, l
}

func ErrorInternalServer(c echo.Context, err error) error {
	err = errorkit.InternalServerError(utils.GetError(err).Message)
	return c.JSON(utils.GetError(err).ID, EncodeJsonResponse(nil, nil, err))
}
