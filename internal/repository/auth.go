package repository

import (
	"errors"
	"log"

	"github.com/Masterminds/squirrel"
	"gitlab.com/p9359/backend-prob/febry-go/internal/model"
)

func (br *bookRepository) GetUser(id int) (model.User, error) {
	user := model.User{}

	err := mysqlQB().
		Select("id", "uuid", "name", "pin", "email", "phone", "date_of_birth").
		From("users").
		Where(squirrel.Eq{"id": id}).
		Limit(1).
		Scan(&user.ID, &user.UUID, &user.Name, &user.Pin, &user.Email, &user.Phone, &user.DateOfBirth)

	if err != nil {
		log.Printf("cannot Get stock -> Error: %v", err)
		return model.User{}, errors.New("something wrong happened")
	} else {
		log.Printf("Success Get Stock")
	}

	return user, nil
}
