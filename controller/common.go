package controller

import (
	"strconv"

	"github.com/fleimkeipa/case/model"

	"github.com/labstack/echo/v4"
)

func getPagination(c echo.Context) model.PaginationOpts {
	pageQ := c.QueryParam("page")
	sizeQ := c.QueryParam("size")

	page, _ := strconv.Atoi(pageQ)

	if page == 0 {
		page = 30
	}

	size, _ := strconv.Atoi(sizeQ)

	return model.PaginationOpts{
		Page: page,
		Size: size,
	}
}

func getFilter(c echo.Context, query string) model.Filter {
	param := c.QueryParam(query)
	if param == "" {
		return model.Filter{}
	}

	return model.Filter{
		IsSended: true,
		Value:    param,
	}
}
