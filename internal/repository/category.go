package repository

import (
	"errors"
	"log"

	"github.com/Masterminds/squirrel"
	"gitlab.com/p9359/backend-prob/febry-go/internal/dto"
	"gitlab.com/p9359/backend-prob/febry-go/internal/model"
)

func (br *bookRepository) CategoryDetail(getCategoryID dto.GetCategoryID) (model.Category, error) {
	category := model.Category{
		Category: &model.Category{
			Category: &model.Category{},
		},
	}

	err := mysqlQB().
		Select("c3.id", "c3.category_id", "c3.name", "c2.id", "c2.category_id", "c2.name", "c1.id", "c1.category_id", "c1.name").
		From("categories c3").
		Where(squirrel.Eq{"c3.id": getCategoryID.CategoryID}).
		LeftJoin("categories as c2 ON c2.id=c3.category_id").
		LeftJoin("categories as c1 ON c1.id=c2.category_id").
		Limit(1).
		Scan(&category.ID, &category.CategoryID, &category.Name, &category.Category.ID, &category.Category.CategoryID, &category.Category.Name, &category.Category.Category.ID, &category.Category.Category.CategoryID, &category.Category.Category.Name)

	if err != nil {
		log.Printf("cannot Get Category -> Error: %v", err)
		return model.Category{}, errors.New("something wrong happened")
	} else {
		log.Printf("Success Get Category")
	}

	return category, nil
}
