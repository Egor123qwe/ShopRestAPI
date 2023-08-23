package users

import model "ShopRestAPI/internal/models/users"

type UserRepository interface {
	Create(u *model.Users) error
	FindByEmail(email string) (*model.Users, error)
}
