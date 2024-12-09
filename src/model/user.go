package model

type User struct {
	Id         int    `db:"id" json:"id"`
	FullName   string `db:"full_name" json:"fullName"`
	Email      string `db:"email" json:"email"`
	Password   string `db:"password" json:"password"`
	IsBlocked  bool   `db:"is_blocked" json:"isBlocked"`
	IsVerified bool   `db:"is_verified" json:"isVerified"`
	CreatedAt  string `db:"created_at" json:"createdAt"`
	UpdatedAt  string `db:"updated_at" json:"updatedAt"`
}
