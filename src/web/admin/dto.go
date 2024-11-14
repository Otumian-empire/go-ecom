package admin

import "otumian-empire/go-ecom/src/model"

type LoginDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreatedAdminDto struct {
	FullName string `json:"fullName"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type UpdateProfileDto struct {
	FullName string `json:"fullName"`
}

type LoginResponse struct {
	Id        int    `db:"id" json:"id"`
	FullName  string `db:"full_name" json:"fullName"`
	Email     string `db:"email" json:"email"`
	Role      string `db:"role" json:"role"`
	CreatedAt string `db:"created_at" json:"createdAt"`
	UpdatedAt string `db:"updated_at" json:"updatedAt"`
	Token     string `json:"token"`
}

func LoginResponseMapper(record model.Admin, token string) LoginResponse {
	return LoginResponse{
		Id:        record.Id,
		FullName:  record.FullName,
		Email:     record.Email,
		Role:      record.Role,
		CreatedAt: record.CreatedAt,
		UpdatedAt: record.UpdatedAt,
		Token:     token,
	}
}
