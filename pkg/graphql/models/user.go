package models

import "quickeat/pkg/entity"

func NewUser(u entity.User) *User {
	result := &User{
		ID:          u.Id,
		FirstName:   u.FirstName,
		LastName:    u.LastName,
		PhoneNumber: u.PhoneNumber,
		Email:       u.Email,
	}

	return result
}
