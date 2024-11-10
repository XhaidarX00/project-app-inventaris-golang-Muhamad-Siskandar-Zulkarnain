package invesment

import (
	"log"
	"main/library"
	"main/model/investment"
	"main/service"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

func GetItemsInvesmentHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var invesment investment.Investment
		err := service.ServiceF.GetTotalInvesmentService(&invesment)
		if err != nil {
			log.Println(err.Error())
			if err.Error() == "Data not found" {
				library.StrucToJson2(w, "Data not Found", library.NotFoundRequest)
			} else {
				library.StrucToJson2(w, "Gagal menghitung total investasi", library.InternalServerError)
			}
			return
		}

		result := investment.GetInvestmentResponse{
			Success:    true,
			StatusCode: http.StatusOK,
			Data:       invesment,
		}
		library.ResponseToJson(w, "Berhasil Menghitung Total Invesment", http.StatusOK, result)
		return
	} else {
		library.StrucToJson(w, library.MethodNotAllowed)
	}
}

func GetItemsInvesmentByIdHandler(w http.ResponseWriter, r *http.Request) {
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
		var invesment investment.ItemInvestment
		invesment.ItemID = idInt
		err = service.ServiceF.GetTotalInvesmentByIdService(&invesment)
		if err != nil {
			log.Println(err.Error())
			if err.Error() == "Data not found" {
				library.StrucToJson2(w, "Data not Found", library.NotFoundRequest)
			} else {
				library.StrucToJson2(w, "Barang tidak ditemukan", library.InternalServerError)
			}
			return
		}

		result := investment.GetInvestmentResponse{
			Success:    true,
			StatusCode: http.StatusOK,
			Data:       invesment,
		}
		library.ResponseToJson(w, "Berhasil Menghitung Total Invesment", http.StatusOK, result)
		return
	} else {
		library.StrucToJson(w, library.MethodNotAllowed)
	}
}
