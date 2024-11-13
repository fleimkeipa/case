package controller

import (
	"fmt"
	"net/http"

	"github.com/fleimkeipa/case/model"
	"github.com/fleimkeipa/case/uc"

	"github.com/labstack/echo/v4"
)

type ProductController struct {
	ProductUC *uc.ProductAPIUC
}

func NewProductController(productUC *uc.ProductAPIUC) *ProductController {
	return &ProductController{ProductUC: productUC}
}

func (rc *ProductController) FindAll(c echo.Context) error {
	opts := rc.getProductsFindOpts(c)

	list, err := rc.ProductUC.FindAll(c.Request().Context(), opts)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, FailureResponse{
			Error:   fmt.Sprintf("Failed to list products: %v", err),
			Message: "There was an issue retrieving products. Please try again.",
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Data:    list,
		Message: "Products retrieved successfully.",
	})
}

func (rc *ProductController) FindOne(c echo.Context) error {
	id := c.Param("id")

	product, err := rc.ProductUC.FindOne(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, FailureResponse{
			Error:   fmt.Sprintf("Failed to retrieve product: %v", err),
			Message: "There was an issue retrieving product. Please try again.",
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Data:    product,
		Message: "Product retrieved successfully.",
	})
}

func (rc *ProductController) getProductsFindOpts(c echo.Context) model.ProductListOpts {
	return model.ProductListOpts{
		PaginationOpts: getPagination(c),
		SuplierID:      getFilter(c, "suplier_id"),
	}
}
