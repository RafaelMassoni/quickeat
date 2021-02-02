package services

import "github.com/jmoiron/sqlx"

type All struct {
	User UserService
}

func NewServices(db *sqlx.DB) All {
	return All{
		User: NewUserService(db),
	}
}
