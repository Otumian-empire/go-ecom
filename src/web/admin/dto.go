package admin

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreatedAdminDto struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}
