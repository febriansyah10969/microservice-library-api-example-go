package service

import (
	"gitlab.com/p9359/backend-prob/febry-go/internal/model"
)

func (bs *bookService) GetUser(id int) (model.User, error) {
	dao := bs.dao.NewGeneralRepository()
	data, err := dao.GetUser(id)
	if err != nil {
		return model.User{}, err
	}

	return data, nil
}
