package dao

import (
	"github.com/Aoi-hosizora/RBAC-learn/src/database"
	"github.com/Aoi-hosizora/RBAC-learn/src/model/po"
	"github.com/Aoi-hosizora/ahlib/xdi"
	"github.com/jinzhu/gorm"
)

type UserRepository struct {
	Db *gorm.DB `di:"~"`
}

func NewUserRepository(dic *xdi.DiContainer) *UserRepository {
	repo := &UserRepository{}
	dic.InjectForce(repo)
	return repo
}

func (u *UserRepository) QueryAll(page int32, limit int32) (int32, []*po.User) {
	total := 0
	users := make([]*po.User, 0)
	u.Db.Model(&po.User{}).Count(&total)
	u.Db.Model(&po.User{}).Limit(limit).Offset((page - 1) * limit).Find(&users)
	return int32(total), users
}

func (u *UserRepository) QueryById(id uint32) *po.User {
	user := &po.User{ID: id}
	rdb := u.Db.Model(user).Where(user).First(user)
	if rdb.RowsAffected == 0 {
		return nil
	}
	return user
}

func (u *UserRepository) Insert(user *po.User) database.DbStatus {
	rdb := u.Db.Model(user).Create(user)
	if rdb.Error != nil || rdb.RowsAffected == 0 {
		return database.DbFailed
	}
	return database.DbSuccess
}

func (u *UserRepository) Update(user *po.User) database.DbStatus {
	rdb := u.Db.Model(user).Update(user)
	if rdb.Error != nil {
		return database.DbFailed
	} else if rdb.RowsAffected == 0 {
		return database.DbNotFound
	}
	return database.DbSuccess
}

func (u *UserRepository) Delete(id uint32) database.DbStatus {
	rdb := u.Db.Model(&po.User{}).Delete(&po.User{ID: id})
	if rdb.Error != nil {
		return database.DbFailed
	} else if rdb.RowsAffected == 0 {
		return database.DbNotFound
	}
	return database.DbSuccess
}
