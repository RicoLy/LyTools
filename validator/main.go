package main

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Users struct {
	Phone  string `form:"phone" json:"phone" validate:"required"`
	Passwd string `form:"passwd" json:"passwd" validate:"required,max=20,min=6"`
	Code   string `form:"code" json:"code" validate:"required,len=6"`
}

func main() {

	users := &Users{
		Phone:  "",
		Passwd: "123",
		Code:   "123456",
	}
	validate := validator.New()
	err := validate.Struct(users)
	fmt.Println(err) //Key: 'Users.Passwd' Error:Field validation for 'Passwd' failed on the 'min' tag

	fmt.Printf("abdd:%+v\n", uuid.NewString())

	fmt.Printf("users:%+v", users)
}
