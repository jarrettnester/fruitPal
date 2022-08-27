package apiServiceLayer

import (
	"errors"
	"github.com/jarrettnester/fruitPal/apiDataLayer"
)

const (
	COMMODITY_ERROR = "Invalid commodity id"
	TONNAGE_ERROR = "Tonnage must be specified and greater than 0"
	PRICE_ERROR = "Price must be greater than 0"
)

func RetrieveCommoditiesOnLatestSnapshot() (fruits []apiDataLayer.Commodity, err error) {

	return apiDataLayer.RetrieveCommoditiesOnLatestSnapshot()

}

func RetrieveTotalCost(price float32, commodity_id int, tonnage int) (totalCostValues []apiDataLayer.DataSnapshotTotals, err error)  {

	if commodity_id < 1 {
		err = errors.New(COMMODITY_ERROR)
		return
	}

	if tonnage < 1 {
		err = errors.New(TONNAGE_ERROR)
		return
	}

	if price <= 0 {
		err = errors.New(PRICE_ERROR)
		return
	}

	return apiDataLayer.RetrieveLatestDataSnapshotTotalView(commodity_id, price, tonnage)

}
