package model

type (
	Product struct {
		DeliveryOption    DeliveryOption `json:"deliveryOption"`
		Description       string         `json:"description"`
		Title             string         `json:"title"`
		ProductMainID     string         `json:"productMainId"`
		CurrencyType      string         `json:"currencyType"`
		StockCode         string         `json:"stockCode"`
		Barcode           string         `json:"barcode"`
		Attributes        []Attribute    `json:"attributes"`
		Images            []Image        `json:"images"`
		DimensionalWeight int            `json:"dimensionalWeight"`
		Quantity          int            `json:"quantity"`
		ListPrice         float64        `json:"listPrice"`
		SalePrice         float64        `json:"salePrice"`
		VatRate           int            `json:"vatRate"`
		CargoCompanyID    int            `json:"cargoCompanyId"`
		CategoryID        int            `json:"categoryId"`
		SupplierID        int            `json:"supplierId"`
		BrandID           int            `json:"brandId"`
	}
	DeliveryOption struct {
		FastDeliveryType string `json:"fastDeliveryType"`
		DeliveryDuration int    `json:"deliveryDuration"`
	}
	Image struct {
		URL string `json:"url"`
	}
	Attribute struct {
		CustomAttributeValue string `json:"customAttributeValue,omitempty"`
		AttributeID          int    `json:"attributeId"`
		AttributeValueID     int    `json:"attributeValueId,omitempty"`
	}
)

type ProductsResponse struct {
	Content       []Product `json:"content"`
	Page          int       `json:"page"`
	Size          int       `json:"size"`
	TotalElements int       `json:"totalElements"`
	TotalPages    int       `json:"totalPages"`
}

type ProductListOpts struct {
	SuplierID      Filter
	PaginationOpts PaginationOpts
}
