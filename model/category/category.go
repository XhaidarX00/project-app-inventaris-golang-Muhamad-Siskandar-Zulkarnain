package category

// Struct untuk Data Kategori
type Category struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// Response untuk GET /api/categories
type GetCategoriesResponse struct {
	Success    bool       `json:"success"`
	StatusCode int        `json:"status_code"`
	Data       []Category `json:"data,omitempty"` // Data adalah list kategori
	Message    string     `json:"message,omitempty"`
}

// Response untuk GET /api/categories/{id}
type GetCategoryByIDResponse struct {
	Success    bool     `json:"success"`
	StatusCode int      `json:"status_code"`
	Data       Category `json:"data,omitempty"`
	Message    string   `json:"message,omitempty"`
}

// Response untuk POST /api/categories
type PostCategoryResponse struct {
	Success    bool     `json:"success"`
	StatusCode int      `json:"status_code"`
	Message    string   `json:"message"`
	Data       Category `json:"data"`
}

// Response untuk PUT /api/categories/{id}
type PutCategoryResponse struct {
	Success    bool     `json:"success"`
	StatusCode int      `json:"status_code"`
	Message    string   `json:"message"`
	Data       Category `json:"data"`
}

// Response untuk DELETE /api/categories/{id}
type DeleteCategoryResponse struct {
	Success    bool   `json:"success"`
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}
