package catagories

import (
	"encoding/json"
	"log"
	"main/library"
	"main/model/category"
	"main/service"
	"main/validation"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func CategoryHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var catagories []category.Category
		err := service.ServiceF.GetCategoriesService(&catagories)
		if err != nil {
			if err.Error() == "Data not found" {
				library.StrucToJson2(w, "Data not Found", library.NotFoundRequest)
			} else {
				library.StrucToJson2(w, "Gagal mengambil data kategori", library.InternalServerError)
			}
			return
		}

		library.ResponseToJson(w, "Berhasil Mengambil Kategori", http.StatusOK, catagories)
		return
	}

	if r.Method == "POST" {
		var category category.Category
		err := json.NewDecoder(r.Body).Decode(&category)
		if err != nil {
			library.StrucToJson(w, library.BadRequest)
			return
		}

		columnsToCheck := []string{"Name"}
		if !validation.Validation(w, &category, columnsToCheck) {
			return
		}

		err = service.ServiceF.AddCategoryService(&category)
		if err != nil {
			library.StrucToJson2(w, "Gagal mengambil data kategori", library.InternalServerError)
			return
		}

		library.ResponseToJson(w, "Kategori berhasil ditambahkan", http.StatusCreated, category)
	} else {
		library.StrucToJson(w, library.MethodNotAllowed)
	}
}

func GetCategoryByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		id := chi.URLParam(r, "id")

		log.Println(id)
		if id == "" {
			library.StrucToJson2(w, "ID kategori tidak ditemukan", library.BadRequest2)
			return
		}

		// Konversi id menjadi integer
		idInt, err := strconv.Atoi(id)
		if err != nil {
			library.StrucToJson2(w, "ID kategori tidak valid", library.BadRequest2)
			return
		}

		var category category.Category
		category.ID = idInt

		err = service.ServiceF.GetCategoryByIdService(&category)
		if err != nil {
			if err.Error() == "Data not found" {
				library.StrucToJson2(w, "Kategori tidak ditemukan", library.NotFoundRequest)
			} else {
				library.StrucToJson2(w, "Gagal mengambil data kategori", library.InternalServerError)
			}
			return
		}

		library.ResponseToJson(w, "Kategori berhasil diambil", http.StatusOK, category)
		return
	}

	library.StrucToJson(w, library.MethodNotAllowed)
}

func PutCategoryByIdHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "PUT" {
		id := chi.URLParam(r, "id")

		log.Println(id)
		if id == "" {
			library.StrucToJson2(w, "ID kategori tidak ditemukan", library.BadRequest2)
			return
		}

		// Konversi id menjadi integer
		idInt, err := strconv.Atoi(id)
		if err != nil {
			library.StrucToJson2(w, "ID kategori tidak valid", library.BadRequest2)
			return
		}

		var category category.Category
		err = json.NewDecoder(r.Body).Decode(&category)
		if err != nil {
			library.StrucToJson(w, library.BadRequest)
			return
		}

		columnsToCheck := []string{"Name", "Description"}
		if !validation.Validation(w, &category, columnsToCheck) {
			return
		}

		category.ID = idInt

		err = service.ServiceF.PutCategoryByIdService(&category)
		if err != nil {
			if err.Error() == "Data not found" {
				library.StrucToJson2(w, "Kategori tidak ditemukan", library.NotFoundRequest)
			} else {
				library.StrucToJson2(w, "Gagal mengambil data kategori", library.InternalServerError)
			}
			return
		}

		library.ResponseToJson(w, "Kategori berhasil diperbarui", http.StatusOK, category)
		return
	}

	library.StrucToJson(w, library.MethodNotAllowed)
}
