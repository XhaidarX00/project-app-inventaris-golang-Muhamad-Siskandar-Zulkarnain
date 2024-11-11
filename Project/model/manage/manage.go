package manage

import "database/sql"

// Struct untuk Data Item
type Item struct {
	ID               int            `json:"id"`
	Name             string         `json:"name"`
	Category         int            `json:"category_id"`
	CategoryName     string         `json:"category_name"`
	PhotoURL         sql.NullString `json:"photo_url"`
	Price            float64        `json:"price"`
	PurchaseDate     string         `json:"purchase_date"`
	Depreciated_rate float64        `json:"depreciated_rate"`
	TotalUsageDays   int            `json:"total_usage_days"`
}

type ResponseItem struct {
	ID             int     `json:"id"`
	Name           string  `json:"name"`
	Category       string  `json:"category"`
	PhotoURL       string  `json:"photo_url"`
	Price          float64 `json:"price"`
	PurchaseDate   string  `json:"purchase_date"`
	TotalUsageDays int     `json:"total_usage_days"`
}

// Response untuk GET /api/items
type GetItemsResponse struct {
	Success    bool   `json:"success"`
	Page       int    `json:"page"`
	Limit      int    `json:"limit"`
	TotalItems int    `json:"total_items"`
	TotalPages int    `json:"total_pages"`
	Data       []Item `json:"data"`
	Message    string `json:"message,omitempty"`
}

// Struct untuk Data Item yang Perlu Diganti
type ReplacementItem struct {
	ID                  int    `json:"id"`
	Name                string `json:"name"`
	Category            string `json:"category"`
	PurchaseDate        string `json:"purchase_date"`
	TotalUsageDays      int    `json:"total_usage_days"`
	ReplacementRequired bool   `json:"replacement_required"`
}

// Response untuk GET /api/items/replacement-needed
type GetReplacementNeededResponse struct {
	Success bool              `json:"success"`
	Data    []ReplacementItem `json:"data"`
}
