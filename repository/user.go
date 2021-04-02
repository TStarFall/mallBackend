package repository

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"mall/model"
	"mall/query"
	"mall/utils"
)

type UserRepository struct {
	DB *gorm.DB
}

type UserRepoInterface interface {
	List(req *query.ListQuery) (users []*model.User, err error)
	GetTotal(req *query.ListQuery) (total int, err error)
	Get(user model.User) (*model.User,error)
	Exist(user model.User) *model.User
	ExistByMobile(mobile string) *model.User
	ExitByUserID(id string) *model.User
	Add(user model.User) (*model.User, error)
	Edit(user model.User) (bool, error)
	Deleted(user model.User) (bool, error)
}

func (repo *UserRepository) List(req *query.ListQuery) (users []*model.User, err error) {
	db := repo.DB
	limit, offset := utils.Page(req.PageSize,req.Page) //分页限制
	if err := db.Order("id desc").Limit(limit).Offset(offset).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
//获取所有用户
func (repo *UserRepository) GetTotal(req *query.ListQuery) (total int, err error) {
	var users []model.User
	db := repo.DB
	if err := db.Find(&users).Count(&total).Error; err != nil {
		return total, err
	}
	return total, nil
}
//单个查找用户
func (repo *UserRepository) Get(user model.User) (*model.User,error) {
	if err := repo.DB.Where(&user).Find(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repo *UserRepository) Exist(user model.User) *model.User {
	var count int
	repo.DB.Find(&user).Where("nick_name = ?", user.NickName).Count(&count)
	if count > 0 {
		return &user
	}
	return nil
}

func (repo *UserRepository) ExistByMobile(mobile string) *model.User {
	var count int
	var user model.User
	repo.DB.Find(&user).Where("mobile = ?", mobile).Count(&count)
	if count > 0 {
		return &user
	}
	return nil
}

func (repo *UserRepository) ExitByUserID(id string) *model.User {
	var user model.User
	repo.DB.Where("user_id = ?", id).First(&user)
	return &user
}
//添加用户
func (repo *UserRepository) Add(user model.User) (*model.User, error) {
	if exist := repo.Exist(user); exist != nil {
		return nil, fmt.Errorf("用户注册已经存在")
	}
	err := repo.DB.Create(&user).Error
	if  err != nil {
		return nil, fmt.Errorf("用户注册失败")
	}
	return &user, nil
}
//编辑用户
func (repo *UserRepository) Edit(user model.User) (bool, error) {
	err := repo.DB.Model(&user).Where("user_id = ?",user.UserId).Updates(map[string]interface{}{
		"nick_name" : user.NickName,
		"mobile" : user.Mobile,
		"address" : user.Address,
	}).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
//删除用户
func (repo *UserRepository) Deleted(user model.User) (bool, error) {
	err := repo.DB.Model(&user).Where("user_id = ?", user.UserId).Update("is_deleted",user.IsDeleted).Error
	if err != nil {
		return false, err
	}
	return true, nil
}




