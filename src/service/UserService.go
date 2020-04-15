package service

import (
	"github.com/Aoi-hosizora/RBAC-learn/src/database"
	"github.com/Aoi-hosizora/RBAC-learn/src/model/po"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/jinzhu/gorm"
)

type UserService struct {
	Db *gorm.DB `di:"~"`
}

func NewUserService(dic *xdi.DiContainer) *UserService {
	srv := &UserService{}
	dic.MustInject(srv)
	return srv
}

func (u *UserService) QueryAll(page int32, limit int32) (int32, []*po.User) {
	total := 0
	users := make([]*po.User, 0)
	u.Db.Model(&po.User{}).Count(&total)
	u.Db.Model(&po.User{}).Limit(limit).Offset((page - 1) * limit).Find(&users)
	return int32(total), users
}

func (u *UserService) QueryById(id uint32) (*po.User, bool) {
	user := &po.User{ID: id}
	rdb := u.Db.Model(user).Where(user).First(user)
	if rdb.RowsAffected == 0 {
		return nil, false
	}
	return user, true
}

func (u *UserService) Insert(user *po.User) database.DbStatus {
	rdb := u.Db.Model(user).Create(user)
	if rdb.Error != nil || rdb.RowsAffected == 0 {
		return database.DbFailed
	}
	return database.DbSuccess
}

func (u *UserService) Update(user *po.User) database.DbStatus {
	rdb := u.Db.Model(user).Update(user)
	if rdb.Error != nil {
		return database.DbFailed
	} else if rdb.RowsAffected == 0 {
		return database.DbNotFound
	}
	return database.DbSuccess
}

func (u *UserService) Delete(id uint32) database.DbStatus {
	rdb := u.Db.Model(&po.User{}).Delete(&po.User{ID: id})
	if rdb.Error != nil {
		return database.DbFailed
	} else if rdb.RowsAffected == 0 {
		return database.DbNotFound
	}
	return database.DbSuccess
}
