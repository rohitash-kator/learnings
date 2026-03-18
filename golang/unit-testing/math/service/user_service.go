package service

import "example.com/testing-add/dao"

type UserService struct {
	userDao dao.UserDaoInterface
}

func (u UserService) GetUser(ID int) (dao.User, error) {
	return u.userDao.GetUser(ID)
}
