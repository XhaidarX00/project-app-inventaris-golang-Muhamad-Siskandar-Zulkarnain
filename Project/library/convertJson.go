package library

import (
	"encoding/json"
	"net/http"
)

func ConvertJson(w http.ResponseWriter, r *http.Request, data interface{}) {
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		StrucToJson(w, BadRequest)
		return
	}
}
