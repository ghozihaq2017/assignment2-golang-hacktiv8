package item_pg

import (
	"assignment2-golang-hacktiv8/entity"
	"assignment2-golang-hacktiv8/pkg/errs"
	"assignment2-golang-hacktiv8/repository/item_repository"
	"database/sql"
)

type itemPG struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) item_repository.Repository {
	return &itemPG{
		db: db,
	}
}

func (itemPG *itemPG) GetItemsByCodes(itemCodes []string) ([]entity.Item, errs.Error) {
	query := itemPG.generateGetItemsByCodesQuery(len(itemCodes))

	itemCodeDbArg := []any{}

	for _, itemCode := range itemCodes {
		itemCodeDbArg = append(itemCodeDbArg, itemCode)
	}

	rows, err := itemPG.db.Query(query, itemCodeDbArg...)

	if err != nil {
		return nil, errs.NewInternalServerError("something went wrong")
	}

	items := []entity.Item{}

	for rows.Next() {

		item := entity.Item{}

		err = rows.Scan(&item.ItemId, &item.ItemCode, &item.Quantity, &item.Description, &item.OrderId, &item.CreatedAt)

		if err != nil {
			return nil, errs.NewInternalServerError("something went wrong")
		}

		items = append(items, item)
	}

	return items, nil
}
