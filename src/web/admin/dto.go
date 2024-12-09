package admin

import "otumian-empire/go-ecom/src/model"

type LoginDto struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type CreatedAdminDto struct {
	FullName string `json:"fullName" binding:"required,min=7,max=55"`
	Email    string `json:"email" binding:"required,email"`
	Role     string `json:"role" binding:"required"`
}

type UpdateProfileDto struct {
	FullName string `json:"fullName" binding:"required,min=7,max=55"`
}

type UpdateAdminRoleDto struct {
	Role string `json:"role" binding:"required"`
}

func (dto UpdateAdminRoleDto) IsValid() bool {
	return dto.Role != SUPER_ADMIN || dto.Role != MODERATOR
}

type IdDto struct {
	Id int `uri:"id" binding:"required"`
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
