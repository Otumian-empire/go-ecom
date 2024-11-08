package model

type Admin struct {
	Id        int    `db:"id" json:"id"`
	FullName  string `db:"full_name" json:"fullName"`
	Email     string `db:"email" json:"email"`
	Password  string `db:"password" json:"password"`
	Role      string `db:"role" json:"role"`
	IsBlocked bool   `db:"is_blocked" json:"isBlocked"`
	CreatedAt string `db:"created_at" json:"createdAt"`
	UpdatedAt string `db:"updated_at" json:"updatedAt"`
}
