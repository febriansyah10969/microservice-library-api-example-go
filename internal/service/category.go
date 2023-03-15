package service

import (
	"fmt"

	"gitlab.com/p9359/backend-prob/febry-go/internal/dto"
)

func (bs *bookService) CategoryDetail(categoryID dto.GetCategoryID) ([]dto.CategoryDetailResponse, error) {
	categoryDetail := []dto.CategoryDetailResponse{}

	dao := bs.dao.NewGeneralRepository()
	result, err := dao.CategoryDetail(categoryID)
	if err != nil {
		return categoryDetail, err
	}

	if *result.ID > 0 {
		p1 := dto.CategoryDetailResponse{
			ID:   *result.ID,
			Name: *result.Name,
		}

		fmt.Println(result)

		if result.Category.ID != nil {
			p2 := dto.CategoryDetailResponse{
				ID:       *result.Category.ID,
				Name:     *result.Category.Name,
				ParentID: result.Category.CategoryID,
			}

			if result.Category.Category.ID != nil {
				p3 := dto.CategoryDetailResponse{
					ID:       *result.Category.Category.ID,
					Name:     *result.Category.Category.Name,
					ParentID: result.Category.Category.CategoryID,
				}

				categoryDetail = append(categoryDetail, p3)
			}

			categoryDetail = append(categoryDetail, p2)
		}

		categoryDetail = append(categoryDetail, p1)
	}

	return categoryDetail, nil
}
