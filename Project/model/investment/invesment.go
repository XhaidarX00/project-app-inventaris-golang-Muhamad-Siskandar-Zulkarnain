package investment

// Struct untuk Data Total Investasi
type Investment struct {
	TotalInvestment  float64 `json:"total_investment"`
	DepreciatedValue float64 `json:"depreciated_value"`
}

// Struct untuk Data Barang Investasi
type ItemInvestment struct {
	ItemID           int     `json:"item_id"`
	Name             string  `json:"name"`
	InitialPrice     float64 `json:"initial_price"`
	DepreciatedValue float64 `json:"depreciated_value"`
	DepreciationRate float64 `json:"depreciation_rate"`
}

// Response untuk GET /api/items/investment
type GetInvestmentResponse struct {
	Success    bool        `json:"success"`
	StatusCode int         `json:"status_code,omitempty"`
	Data       interface{} `json:"data,omitempty"`
}

// Response untuk GET /api/items/investment/{id}
type GetItemInvestmentResponse struct {
	Success    bool           `json:"success"`
	StatusCode int            `json:"status_code,omitempty"`
	Message    string         `json:"message,omitempty"`
	Data       ItemInvestment `json:"data,omitempty"` // Data berisi informasi barang investasi
}
