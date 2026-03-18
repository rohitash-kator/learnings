package service

import (
	"fmt"
	"testing"

	"example.com/testing-add/dao"
	"github.com/stretchr/testify/assert"
)

type UserDaoMockFoundUser struct{}

func (u *UserDaoMockFoundUser) GetUser(ID int) (dao.User, error) {
	return dao.User{
		ID:   ID,
		Name: "Bruce",
	}, nil
}

type UserDaoMockNotFoundUser struct{}

func (u *UserDaoMockNotFoundUser) GetUser(ID int) (dao.User, error) {
	return dao.User{}, fmt.Errorf("User not found")
}

func TestUserService_GetUser(t *testing.T) {
	type fields struct {
		userDao dao.UserDaoInterface
	}

	type args struct {
		ID int
	}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    dao.User
		wantErr error
	}{
		{name: "should return user when found", fields: fields{userDao: &UserDaoMockFoundUser{}}, args: args{10}, want: dao.User{ID: 10, Name: "Bruce"}},
		{name: "should return error when not found", fields: fields{userDao: &UserDaoMockNotFoundUser{}}, args: args{21}, wantErr: fmt.Errorf("User not found")},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			u := UserService{
				userDao: testCase.fields.userDao,
			}

			got, err := u.GetUser(testCase.args.ID)
			assert.Equal(t, testCase.want, got)
			assert.Equal(t, testCase.wantErr, err)
		})
	}
}
