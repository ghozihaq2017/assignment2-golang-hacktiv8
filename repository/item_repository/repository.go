package item_repository

import (
	"assignment2-golang-hacktiv8/entity"
	"assignment2-golang-hacktiv8/pkg/errs"
)

type Repository interface {
	GetItemsByCodes(itemCodes []string) ([]entity.Item, errs.Error)
}
