package application

import (
	"fmt"

	"example2.com/m/Users/domain"
)

type GetAllUsers struct {
	db domain.IUser
}

func NewGetAllUsers(db domain.IUser)*GetAllUsers{
	return&GetAllUsers{db: db}
}

func(uc *GetAllUsers)Execute()([]domain.User, error){
	fmt.Println("Si funca")
	result, err := uc.db.GetAllUserFunction()
	return result, err
}