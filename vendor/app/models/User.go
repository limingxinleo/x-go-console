package models

import (
	"app/container"
	"app/providers"
)

type User struct {
	Id        int64
	Username  string
	Password  string
	Name      string
	Email     string
	RoleId    uint
	CreatedAt string
	UpdatedAt string
}

func (this *User)Find(id int64) *User {
	db := container.DI.Get("db").(*providers.DB)
	row := db.Client.QueryRow("SELECT * FROM user WHERE id = ?", id)
	row.Scan(
		&this.Id,
		&this.Username,
		&this.Password,
		&this.Name,
		&this.Email,
		&this.RoleId,
		&this.CreatedAt,
		&this.UpdatedAt,
	)
	return this
}
