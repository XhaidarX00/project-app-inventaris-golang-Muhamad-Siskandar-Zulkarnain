package reminder

import (
	"main/library"
	"main/model/manage"
	"main/service"
	"net/http"
)

func GetItemsReplacementHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var replacement []manage.ReplacementItem
		err := service.ServiceF.GetItemsReplacementService(&replacement)
		if err != nil {
			library.StrucToJson2(w, "Gagal mengambil data barang yang perlu diganti", library.InternalServerError)
			return
		}

		result := manage.GetReplacementNeededResponse{
			Success: true,
			Data:    replacement,
		}
		library.JsonResponse(w, result)
		return
	} else {
		library.StrucToJson(w, library.MethodNotAllowed)
	}
}
