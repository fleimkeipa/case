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

// FindAll godoc
//
//	@Summary		Retrieve a list of products from the supplier API
//	@Description	Retrieves a list of products from the supplier API and returns it
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			supplier_id	path		string			true	"SupplierID"
//	@Param			page		query		int				false	"Page number"
//	@Param			size		query		int				false	"Page size"
//	@Success		200			{object}	SuccessResponse	"List of products"
//	@Failure		500			{object}	FailureResponse	"Bad request or error message"
//	@Router			/products [get]
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

// FindOne godoc
//
//	@Summary		Retrieve a product by Product Main ID
//	@Description	Retrieves a product from the supplier API and returns it
//	@Tags			products
//	@Accept			json
//	@Produce		json
//	@Param			id			path		string	true	"ProductMainID"
//	@Param			supplier_id	path		string	true	"SupplierID"
//	@Success		200			{object}	SuccessResponse{data=model.Product}
//	@Failure		500			{object}	FailureResponse
//	@Router			/products/{id} [get]
func (rc *ProductController) FindOne(c echo.Context) error {
	productMainID := c.Param("id")
	suplierID := c.QueryParam("supplier_id")

	product, err := rc.ProductUC.FindOne(c.Request().Context(), suplierID, productMainID)
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
		SuplierID:      getFilter(c, "supplier_id"),
	}
}
