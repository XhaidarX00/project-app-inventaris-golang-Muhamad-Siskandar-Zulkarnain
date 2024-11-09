package investment

// Struct untuk Data Total Investasi
type Investment struct {
	TotalInvestment  int `json:"total_investment"`
	DepreciatedValue int `json:"depreciated_value"`
}

// Struct untuk Data Barang Investasi
type ItemInvestment struct {
	ItemID           int    `json:"item_id"`
	Name             string `json:"name"`
	InitialPrice     int    `json:"initial_price"`
	DepreciatedValue int    `json:"depreciated_value"`
	DepreciationRate int    `json:"depreciation_rate"`
}

// Response untuk GET /api/items/investment
type GetInvestmentResponse struct {
	Success    bool       `json:"success"`
	StatusCode int        `json:"status_code,omitempty"`
	Message    string     `json:"message,omitempty"`
	Data       Investment `json:"data,omitempty"` // Data berisi informasi total investasi
}

// Response untuk GET /api/items/investment/{id}
type GetItemInvestmentResponse struct {
	Success    bool           `json:"success"`
	StatusCode int            `json:"status_code,omitempty"`
	Message    string         `json:"message,omitempty"`
	Data       ItemInvestment `json:"data,omitempty"` // Data berisi informasi barang investasi
}
