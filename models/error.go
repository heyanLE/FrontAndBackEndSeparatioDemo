package models

import "errors"

var(
	UsernameNotFind = errors.New("Models Error : Username Not Find ")
	PasswordError	= errors.New("Models Error : Password Error ")
	UserExist 		= errors.New("Models Error : User Exist ")
)
