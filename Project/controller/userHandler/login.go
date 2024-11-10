package usershandler

import (
	"encoding/json"
	"main/library"
	users "main/model/user"
	"main/service"
	"main/validation"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var user users.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			library.StrucToJson(w, library.BadRequest)
			return
		}

		columnsToCheck := []string{"Username", "Password"}
		if !validation.Validation(w, &user, columnsToCheck) {
			return
		}

		err = service.ServiceF.LoginService(&user)
		if err != nil {
			library.StrucToJson(w, library.UnauthorizedRequest)
			return
		}

		library.ResponseToJson(w, "Succes Login", http.StatusOK, user)

	} else {
		library.StrucToJson(w, library.MethodNotAllowed)
	}
}
