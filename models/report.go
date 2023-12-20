package models

type GetListClientReportRequest struct {
	DateFrom string `json:"date_from"`
	DateTo   string `json:"date_to"`
	Query    string `json:"-"`
}

type GetListClientReportResponse struct {
	Count   int       `json:"count"`
	Clients []*Client `json:"clients"`
}

type GetAllSaleReportRequest struct {
	DateFrom string `json:"date_from"`
	DateTo   string `json:"date_to"`
	Query    string `json:"-"`
}

type GetAllSaleReportResponse struct {
	Count       int           `json:"count"`
	SaleReports []*SaleReport `json:"all_sale_report"`
}

type SaleReport struct {
	Name     string  `json:"name"`
	Quantity int64   `json:"quantity"`
	Price    float64 `json:"price"`
}
