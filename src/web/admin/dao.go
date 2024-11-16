package admin

import (
	"database/sql"
	"fmt"
	"otumian-empire/go-ecom/src/config"
	"otumian-empire/go-ecom/src/model"
)

type Dao struct {
	*sql.DB
}

func NewDao(db *sql.DB) Dao {
	return Dao{db}
}

func (db Dao) FindOneById(id config.IdType) (model.Admin, error) {
	row := db.QueryRow("SELECT id, email, full_name, is_blocked, role, created_at, updated_at FROM admin WHERE id=$1", id)
	if row.Err() != nil {
		return model.Admin{}, fmt.Errorf("record not found")
	}

	var admin model.Admin

	err := row.Scan(
		&admin.Id,
		&admin.Email,
		&admin.FullName,
		&admin.IsBlocked,
		&admin.Role,
		&admin.CreatedAt,
		&admin.UpdatedAt)

	if err != nil {
		fmt.Println(err)
		return model.Admin{}, fmt.Errorf("record not found, %v", err)
	}

	return admin, nil
}

func (db Dao) FindOneByEmail(email string) (model.Admin, error) {
	row := db.QueryRow("SELECT id, email, password, full_name, is_blocked, role, created_at, updated_at FROM admin WHERE email=$1", email)

	var admin model.Admin

	err := row.Scan(
		&admin.Id,
		&admin.Email,
		&admin.Password,
		&admin.FullName,
		&admin.IsBlocked,
		&admin.Role,
		&admin.CreatedAt,
		&admin.UpdatedAt)
	fmt.Println("row")
	fmt.Println(admin)
	fmt.Println("row")
	if err != nil {
		return model.Admin{}, fmt.Errorf("record not found, %v", err)
	}
	fmt.Println("Admin read")
	fmt.Println(admin)
	return admin, nil
}

func (db Dao) Create(payload CreatedAdminDto) error {
	// TODO: generate a random password
	row, err := db.Exec("INSERT INTO admin (email, full_name, role, password) VALUES($1, $2, $3, $4)", payload.Email, payload.FullName, payload.Role, "")

	if err != nil {
		return fmt.Errorf("an error occurred inserting , %v", err)
	}

	if count, err := row.RowsAffected(); err != nil {
		return fmt.Errorf("an error occurred inserting , %v", err)
	} else if count < 1 {
		return fmt.Errorf("no new record was created")
	}

	return nil
}

func (db Dao) Update(userId config.IdType, field, value string) error {
	// TODO: generate a random password
	sql := fmt.Sprintf("UPDATE admin set %v=$1 WHERE id=$2", field)
	row, err := db.Exec(sql, value, userId)
	if err != nil {
		return fmt.Errorf("an error occurred updating , %v", err)
	}

	if count, err := row.RowsAffected(); err != nil {
		return fmt.Errorf("an error occurred inserting , %v", err)
	} else if count < 1 {
		return fmt.Errorf("record was not updated")
	}

	return nil
}
