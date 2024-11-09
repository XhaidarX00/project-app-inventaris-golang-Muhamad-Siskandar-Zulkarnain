package manage

// Struct untuk Data Item
type Item struct {
	ID             int    `json:"id"`
	Name           string `json:"name"`
	Category       string `json:"category"`
	PhotoURL       string `json:"photo_url"`
	Price          int    `json:"price"`
	PurchaseDate   string `json:"purchase_date"`
	TotalUsageDays int    `json:"total_usage_days"`
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

// Response untuk POST /api/items
type PostItemResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Data    Item   `json:"data"`
}

// Response untuk GET /api/items/{id}
type GetItemByIDResponse struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"status_code"`
	Data       Item   `json:"data"`
	Message    string `json:"message,omitempty"`
}

// Response untuk PUT /api/items/{id}
type PutItemResponse struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Data       Item   `json:"data"`
}

// Response untuk DELETE /api/items/{id}
type DeleteItemResponse struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
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
	Message string            `json:"message,omitempty"`
}
