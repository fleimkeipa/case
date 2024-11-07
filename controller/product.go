package controller

import (
	"fmt"
	"net/http"

	"github.com/fleimkeipa/case/uc"
	"github.com/labstack/echo/v4"
)

type ProductController struct {
	ProductUC *uc.ProductUC
}

func NewProductController(productUC *uc.ProductUC) *ProductController {
	return &ProductController{ProductUC: productUC}
}

func (rc *ProductController) FindAll(c echo.Context) error {
	suplierID := c.QueryParam("suplier_id")

	list, err := rc.ProductUC.FindAll(suplierID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, FailureResponse{
			Error:   fmt.Sprintf("Failed to list deployments: %v", err),
			Message: "There was an issue retrieving deployments. Please try again.",
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Data:    list,
		Message: "Deployments retrieved successfully.",
	})
}

func (rc *ProductController) FindOne(c echo.Context) error {
	id := c.Param("id")

	product, err := rc.ProductUC.FindOne(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, FailureResponse{
			Error:   fmt.Sprintf("Failed to retrieve deployment: %v", err),
			Message: "There was an issue retrieving deployment. Please try again.",
		})
	}

	return c.JSON(http.StatusOK, SuccessResponse{
		Data:    product,
		Message: "Deployment retrieved successfully.",
	})
}
