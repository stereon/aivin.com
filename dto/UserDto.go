package dto

import (
	"github.com/stereon/aivin.com/model"
)

type UserDto struct {
	Username  string
	Telephone string
}

func UserTodo(user model.User) UserDto {
	return UserDto{
		Username: user.Fusername,
		Telephone: user.Ftelphone,
	}
}