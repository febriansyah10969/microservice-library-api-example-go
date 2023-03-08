package repository

import (
	"encoding/base64"
	"encoding/json"
	"log"

	"github.com/Masterminds/squirrel"
	"gitlab.com/p9359/backend-prob/febry-go/internal/helper"
)

type curs struct {
	Id               int  `json:"id"`
	PointsToNextItem bool `json:"_pointsToNextItems"`
}

func Paginate(sc, sb squirrel.SelectBuilder, p helper.InPage) (squirrel.SelectBuilder, *helper.Pagination) {
	pag := helper.Pagination{}
	if len(p.Perpage) == 0 || p.Perpage == "0" {
		p.Perpage = "10"
	}

	if p.Cursor != nil {
		cursor := decodeCursor(*p.Cursor)
		if len(*p.Cursor) == 0 {
			sb = sb.Suffix("ORDER BY pr.id DESC LIMIT ?", p.Perpage)
		} else {
			if cursor.PointsToNextItem {
				sb = sb.Where(squirrel.Lt{"pr.id": cursor.Id}).Suffix("ORDER BY pr.id DESC LIMIT ?", p.Perpage)
			} else {
				sb = sb.Where(squirrel.Gt{"pr.id": cursor.Id}).Suffix("ORDER BY pr.id ASC LIMIT ?", p.Perpage)
			}
		}
	} else {
		sb = sb.Suffix("ORDER BY pr.id DESC LIMIT ?", p.Perpage)
	}

	return sb, &pag
}

func decodeCursor(cursor string) curs {
	tmp := curs{}
	decodedByte, _ := base64.StdEncoding.DecodeString(cursor)
	json.Unmarshal(decodedByte, &tmp)

	return tmp
}

func encodeCursor(c curs) string {
	marshal, _ := json.Marshal(c)
	return base64.StdEncoding.EncodeToString(marshal)
}

func hasMorePages(sb squirrel.SelectBuilder) bool {
	var count int
	var result bool
	if err := sb.Suffix("LIMIT 1").Scan(&count); err != nil {
		log.Println("has more paginate error: ", err)
		return false
	}

	if count == 0 {
		result = false
	} else {
		result = true
	}

	return result
}

func isFirstPage(sb squirrel.SelectBuilder) bool {
	var count int
	var result bool
	if err := sb.Suffix("LIMIT 1").Scan(&count); err != nil {
		log.Println(err)
		return false
	}

	if count == 0 {
		result = true
	} else {
		result = false
	}

	return result
}
