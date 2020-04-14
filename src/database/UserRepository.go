package database

import (
	"github.com/Aoi-hosizora/RBAC-learn/src/model/po"
)

type UserRepository struct {
	_list []*po.User
}

func NewUserRepository() *UserRepository {
	return &UserRepository{_list: []*po.User{
		{ID: 1, Name: "User1", Password: "123456", Role: "Admin"},
		{ID: 2, Name: "User2", Password: "123456", Role: "Normal"},
		{ID: 3, Name: "User3", Password: "123456", Role: "Normal"},
		{ID: 4, Name: "User4", Password: "123456", Role: "Normal"},
	}}
}

func (u *UserRepository) AddUser(user *po.User) {
	u._list = append(u._list, user)
}

func (u *UserRepository) QueryById(id uint32) *po.User {
	for _, u := range u._list {
		if u.ID == id {
			return u
		}
	}
	return nil
}
