package database

import (
	"github.com/Aoi-hosizora/RBAC-learn/src/model/po"
	"github.com/Aoi-hosizora/RBAC-learn/src/util"
)

type UserRepository struct {
	_list []*po.User
}

func NewUserRepository() *UserRepository {
	pass, _ := util.AuthUtil.EncryptPassword("123456")
	return &UserRepository{_list: []*po.User{
		{ID: 1, Name: "User1", Password: pass, Role: "admin"},
		{ID: 2, Name: "User2", Password: pass, Role: "normal"},
		{ID: 3, Name: "User3", Password: pass, Role: "normal"},
		{ID: 4, Name: "User4", Password: pass, Role: "normal"},
	}}
}

func (u *UserRepository) AddUser(user *po.User) {
	u._list = append(u._list, user)
}

func (u *UserRepository) QueryById(id uint32) *po.User {
	for _, user := range u._list {
		if user.ID == id {
			return user
		}
	}
	return nil
}

func (u *UserRepository) QueryAll() []*po.User {
	return u._list
}
