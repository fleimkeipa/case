package model

type (
	Product struct {
		Barcode           string         `json:"barcode"`
		Title             string         `json:"title"`
		ProductMainID     string         `json:"productMainId"`
		BrandID           int            `json:"brandId"`
		CategoryID        int            `json:"categoryId"`
		Quantity          int            `json:"quantity"`
		StockCode         string         `json:"stockCode"`
		DimensionalWeight int            `json:"dimensionalWeight"`
		Description       string         `json:"description"`
		CurrencyType      string         `json:"currencyType"`
		ListPrice         float64        `json:"listPrice"`
		SalePrice         float64        `json:"salePrice"`
		VatRate           int            `json:"vatRate"`
		CargoCompanyID    int            `json:"cargoCompanyId"`
		DeliveryOption    DeliveryOption `json:"deliveryOption"`
		Images            []Image        `json:"images"`
		Attributes        []Attribute    `json:"attributes"`
	}
	DeliveryOption struct {
		DeliveryDuration int    `json:"deliveryDuration"`
		FastDeliveryType string `json:"fastDeliveryType"`
	}
	Image struct {
		URL string `json:"url"`
	}
	Attribute struct {
		AttributeID          int    `json:"attributeId"`
		AttributeValueID     int    `json:"attributeValueId,omitempty"`
		CustomAttributeValue string `json:"customAttributeValue,omitempty"`
	}
)

type ProductsResponse struct {
	Page          int `json:"page"`
	Size          int `json:"size"`
	TotalElements int `json:"totalElements"`
	TotalPages    int `json:"totalPages"`
	Content       []struct {
		Approved   bool `json:"approved"`
		Archived   bool `json:"archived"`
		Attributes []struct {
			AttributeID      int    `json:"attributeId"`
			AttributeName    string `json:"attributeName"`
			AttributeValue   string `json:"attributeValue"`
			AttributeValueID int    `json:"attributeValueId"`
		} `json:"attributes"`
		Barcode           string `json:"barcode"`
		Brand             string `json:"brand"`
		BrandID           int    `json:"brandId"`
		CategoryName      string `json:"categoryName"`
		CreateDateTime    int64  `json:"createDateTime"`
		Description       string `json:"description"`
		DimensionalWeight int    `json:"dimensionalWeight"`
		HasActiveCampaign bool   `json:"hasActiveCampaign"`
		ID                string `json:"id"`
		Images            []struct {
			URL string `json:"url"`
		} `json:"images"`
		LastUpdateDate      int64         `json:"lastUpdateDate"`
		ListPrice           float64       `json:"listPrice"`
		Locked              bool          `json:"locked"`
		OnSale              bool          `json:"onSale"`
		PimCategoryID       int           `json:"pimCategoryId"`
		PlatformListingID   string        `json:"platformListingId"`
		ProductCode         int           `json:"productCode"`
		ProductContentID    int           `json:"productContentId"`
		ProductMainID       string        `json:"productMainId"`
		Quantity            int           `json:"quantity"`
		SalePrice           float64       `json:"salePrice"`
		StockCode           string        `json:"stockCode"`
		StockUnitType       string        `json:"stockUnitType"`
		SupplierID          int           `json:"supplierId"`
		Title               string        `json:"title"`
		VatRate             int           `json:"vatRate"`
		Rejected            bool          `json:"rejected"`
		RejectReasonDetails []interface{} `json:"rejectReasonDetails"`
		Blacklisted         bool          `json:"blacklisted"`
		HasHTMLContent      bool          `json:"hasHtmlContent"`
		ProductURL          string        `json:"productUrl"`
		LockReason          string        `json:"lockReason,omitempty"`
		BlacklistReason     string        `json:"blacklistReason,omitempty"`
		DeliveryDuration    int           `json:"deliveryDuration,omitempty"`
	} `json:"content"`
}
