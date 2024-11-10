package usershandler

import (
	"encoding/json"
	"main/library"
	users "main/model/user"
	"main/service"
	"main/validation"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var user users.User

		// Decode JSON dari request body
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			library.StrucToJson(w, library.BadRequest)
			return
		}

		columnsToCheck := []string{"Username", "Password", "Email"}
		if !validation.Validation(w, &user, columnsToCheck) {
			return
		}

		err = service.ServiceF.RegisterService(user)
		if err != nil {
			if err.Error() == "username or email already exists" {
				library.ResponseToJson(w, "username or email already exists", http.StatusConflict, nil)
			} else {
				library.StrucToJson(w, library.UnauthorizedRequest)
			}
			return
		}

		library.ResponseToJson(w, "Success Register", http.StatusCreated, nil)

	} else {
		library.StrucToJson(w, library.MethodNotAllowed)
	}
}
