package users

type RegistrationResponse struct {
	Success    bool   `json:"succes"`
	StatusCode bool   `json:"status_code"`
	Message    string `json:"message,omitempty"`
}

type LoginResponse struct {
	Success    bool   `json:"succes"`
	StatusCode bool   `json:"status_code"`
	Message    string `json:"message,omitempty"`
	Token      string `json:"token"`
}

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Token    string `json:"token"`
}
