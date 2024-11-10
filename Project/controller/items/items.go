package items

import (
	"encoding/json"
	"fmt"
	"main/library"
	"main/model/manage"
	"main/model/response"
	"main/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func GetItemsPaginated(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		page := chi.URLParam(r, "page")
		limit := chi.URLParam(r, "limit")
		if page == "" || limit == "" {
			library.StrucToJson2(w, "page atau limit tidak boleh kosong", library.BadRequest2)
			return
		}

		pageInt, err := strconv.Atoi(page)
		if err != nil {
			library.StrucToJson2(w, "Page tidak valid", library.BadRequest2)
			return
		}

		limitInt, err := strconv.Atoi(limit)
		if err != nil {
			library.StrucToJson2(w, "Limit tidak valid", library.BadRequest2)
			return
		}

		var data response.PaginationResponse
		data.Page = pageInt
		data.Limit = limitInt

		err = service.ServiceF.GetItemsPaginatedService(&data)
		if err != nil {
			library.StrucToJson2(w, "Gagal mengambil data barang inventaris", library.InternalServerError)
			return
		}

		library.JsonResponse(w, data)
		return
	}

	library.StrucToJson(w, library.MethodNotAllowed)
}

func AddInventoryItemHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		categoryIDStr := r.FormValue("category_id")
		photoURL := r.FormValue("photo_url")
		priceStr := r.FormValue("price")
		purchaseDateStr := r.FormValue("purchase_date")

		// Validasi dan konversi data
		categoryID, err := strconv.Atoi(categoryIDStr)
		if err != nil {
			library.ResponseToJson(w, "Invalid category_id", http.StatusBadRequest, nil)
			return
		}

		price, err := strconv.ParseFloat(priceStr, 64)
		if err != nil {
			library.ResponseToJson(w, "Invalid price", http.StatusBadRequest, nil)
			return
		}

		// purchaseDate, err := time.Parse("2006-01-02", purchaseDateStr)
		if err != nil {
			library.ResponseToJson(w, "Invalid purchase_date format. Use YYYY-MM-DD", http.StatusBadRequest, nil)
			return
		}

		// Membuat instance InventoryItem
		item := manage.Item{
			Name:         name,
			Category:     categoryID,
			PhotoURL:     photoURL,
			Price:        price,
			PurchaseDate: purchaseDateStr,
		}

		err = service.ServiceF.AddInventoryItemService(&item)
		if err != nil {
			library.StrucToJson2(w, "Barang tidak ditemukan", library.NotFoundRequest)
			return
		}

		result := library.ManageItemsResponse(&item)
		library.ResponseToJson(w, "Barang berhasil ditambahkan", http.StatusOK, result)
	} else {
		library.StrucToJson(w, library.MethodNotAllowed)
	}

}

func GetInventoryItemByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := chi.URLParam(r, "id")
		if id == "" {
			library.StrucToJson2(w, "ID item tidak ditemukan", library.BadRequest2)
			return
		}

		// Konversi id menjadi integer
		idInt, err := strconv.Atoi(id)
		if err != nil {
			library.StrucToJson2(w, "ID item tidak valid", library.BadRequest2)
			return
		}

		var item manage.Item
		item.ID = idInt

		err = service.ServiceF.GetInventoryItemByIdService(&item)
		if err != nil {
			library.StrucToJson2(w, "Barang tidak ditemukan", library.NotFoundRequest)
			return
		}
		result := library.ManageItemsResponse(&item)
		library.ResponseToJson(w, "Berhasil Mengambil Barang", http.StatusOK, result)
	} else {
		library.StrucToJson(w, library.MethodNotAllowed)
	}
}

func UpdateInventoryItemByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		id := chi.URLParam(r, "id")
		if id == "" {
			library.StrucToJson2(w, "ID item tidak ditemukan", library.BadRequest2)
			return
		}

		// Konversi id menjadi integer
		idInt, err := strconv.Atoi(id)
		if err != nil {
			library.StrucToJson2(w, "ID item tidak valid", library.BadRequest2)
			return
		}

		var item manage.Item
		err = json.NewDecoder(r.Body).Decode(&item)
		if err != nil {
			library.StrucToJson(w, library.BadRequest)
			return
		}

		item.ID = idInt

		err = service.ServiceF.UpdateInventoryItemByIdService(&item)
		if err != nil {
			fmt.Println(err.Error())
			library.StrucToJson2(w, "Barang tidak ditemukan", library.NotFoundRequest)
			return
		}
		result := library.ManageItemsResponse(&item)
		library.ResponseToJson(w, "Barang berhasil diperbarui", http.StatusOK, result)
	} else {
		library.StrucToJson(w, library.MethodNotAllowed)
	}
}

func DeleteInventoryItemByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "DELETE" {
		id := chi.URLParam(r, "id")
		if id == "" {
			library.StrucToJson2(w, "ID item tidak ditemukan", library.BadRequest2)
			return
		}

		// Konversi id menjadi integer
		idInt, err := strconv.Atoi(id)
		if err != nil {
			library.StrucToJson2(w, "ID item tidak valid", library.BadRequest2)
			return
		}

		var item manage.Item
		item.ID = idInt

		err = service.ServiceF.DeleteInventoryItemByIdService(idInt)
		if err != nil {
			fmt.Println(err.Error())
			library.StrucToJson2(w, "Barang tidak ditemukan", library.NotFoundRequest)
			return
		}
		library.ResponseToJson(w, "Barang berhasil dihapus", http.StatusOK, nil)
	} else {
		library.StrucToJson(w, library.MethodNotAllowed)
	}
}
