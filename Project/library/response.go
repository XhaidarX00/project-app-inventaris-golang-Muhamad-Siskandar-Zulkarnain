package library

import (
	"encoding/json"
	"main/model/manage"
	"main/model/response"
	"net/http"
)

// OKRequest - Status 200 OK
func OKRequest(msg string, data interface{}) response.Response {
	return response.Response{
		Success:    true,
		StatusCode: http.StatusOK,
		Message:    msg,
		Data:       data,
	}
}

// CreatedRequest - Status 201 Created
// Digunakan ketika resource berhasil dibuat.
func CreatedRequest(data interface{}) response.Response {
	return response.Response{
		Success:    true,
		StatusCode: http.StatusCreated,
		Message:    "Resource created successfully",
		Data:       data,
	}
}

// BadRequest - Status 400 Bad Request
// Digunakan ketika permintaan tidak valid atau memiliki parameter yang salah.
func BadRequest() response.Response {
	return response.Response{
		Success:    false,
		StatusCode: http.StatusBadRequest,
		Message:    "Invalid input parameters",
		Data:       nil,
	}
}

func BadRequest2(msg string) response.Response {
	return response.Response{
		Success:    false,
		StatusCode: http.StatusBadRequest,
		Message:    msg,
		Data:       nil,
	}
}

// UnauthorizedRequest - Status 401 Unauthorized
// Digunakan ketika autentikasi pengguna diperlukan tetapi gagal.
func UnauthorizedRequest() response.Response {
	return response.Response{
		Success:    false,
		StatusCode: http.StatusUnauthorized,
		Message:    "Authentication failed",
		Data:       nil,
	}
}

// ForbiddenRequest - Status 403 Forbidden
// Digunakan ketika pengguna tidak memiliki izin untuk mengakses sumber daya.
func ForbiddenRequest() response.Response {
	return response.Response{
		Success:    false,
		StatusCode: http.StatusForbidden,
		Message:    "You do not have permission to access this resource",
		Data:       nil,
	}
}

// NotFoundRequest - Status 404 Not Found
// Digunakan ketika resource yang diminta tidak ditemukan.
func NotFoundRequest(msg string) response.Response {
	return response.Response{
		Success:    false,
		StatusCode: http.StatusNotFound,
		Message:    msg,
		Data:       nil,
	}
}

// ConflictRequest - Status 409 Conflict
// Digunakan ketika permintaan menyebabkan konflik dengan status saat ini.
func ConflictRequest() response.Response {
	return response.Response{
		Success:    false,
		StatusCode: http.StatusConflict,
		Message:    "Resource conflict occurred",
		Data:       nil,
	}
}

// InternalServerError - Status 500 Internal Server Error
// Digunakan ketika terjadi kesalahan di server yang tidak diketahui.
func InternalServerError(msg string) response.Response {
	return response.Response{
		Success:    false,
		StatusCode: http.StatusInternalServerError,
		Message:    msg,
		Data:       nil,
	}
}

// ServiceUnavailableRequest - Status 503 Service Unavailable
// Digunakan ketika server tidak dapat memproses permintaan karena beban atau pemeliharaan.
func ServiceUnavailableRequest() response.Response {
	return response.Response{
		Success:    false,
		StatusCode: http.StatusServiceUnavailable,
		Message:    "Service is temporarily unavailable",
		Data:       nil,
	}
}

// MethodNotAllowed - Status 405 Method Not Allowed
// Digunakan ketika metode HTTP yang digunakan tidak diizinkan untuk resource tertentu.
func MethodNotAllowed() response.Response {
	return response.Response{
		Success:    false,
		StatusCode: http.StatusMethodNotAllowed,
		Message:    "Method not allowed for this resource",
		Data:       nil,
	}
}

func PageResponse(text string, limit, page, totalItems, totalPages int, data interface{}) response.PaginationResponse {
	return response.PaginationResponse{
		StatusCode: http.StatusOK,
		Message:    text,
		Page:       page,
		Limit:      limit,
		TotalItems: totalItems,
		TotalPages: totalPages,
		Data:       data,
	}
}

func JsonResponse(w http.ResponseWriter, response interface{}) {
	json.NewEncoder(w).Encode(response)
}

func ResponseToJson(w http.ResponseWriter, msg string, statusCode int, data interface{}) {
	if statusCode != 0 {
		response := OKRequest(msg, data)
		json.NewEncoder(w).Encode(response)

	} else {
		response := response.Response{
			Success:    false,
			StatusCode: statusCode,
			Message:    msg,
			Data:       nil,
		}
		json.NewEncoder(w).Encode(response)
	}
}

func StrucToJson(w http.ResponseWriter, function func() response.Response) {
	response := function()
	json.NewEncoder(w).Encode(response)
}

func StrucToJson2(w http.ResponseWriter, msg string, function func(msg string) response.Response) {
	response := function(msg)
	json.NewEncoder(w).Encode(response)
}

func ManageItemsResponse(item *manage.Item) manage.ResponseItem {
	return manage.ResponseItem{
		ID:             item.ID,
		Name:           item.Name,
		Category:       item.CategoryName,
		PhotoURL:       item.PhotoURL.String,
		Price:          item.Price,
		PurchaseDate:   item.PurchaseDate,
		TotalUsageDays: item.TotalUsageDays,
	}
}
