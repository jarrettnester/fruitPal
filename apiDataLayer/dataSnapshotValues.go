package apiDataLayer

import (
	"context"
	"github.com/georgysavva/scany/sqlscan"
	"strconv"
)

const (
	DATA_SNAP_VIEW_ONE = `SELECT
(SELECT country_code FROM countries WHERE id = country_id) AS CountryCode,
(variable_cost +`
	DATA_SNAP_VIEW_TWO = `) * `
	DATA_SNAP_VIEW_THREE = ` AS TotalCost
FROM data_snapshot_values WHERE data_snapshot_id = (SELECT MAX(id) FROM data_snapshots) AND commodity_id = `
	DATA_SNAP_VIEW_FOUR = ` ORDER BY TotalCost DESC`
)

type DataSnapshotTotals struct {
	Countrycode string
	Totalcost float32
	Iscurrenttrade bool
}

func RetrieveLatestDataSnapshotTotalView(commodityId int, price float32, tonnage int) (dataSnapshotTotals []DataSnapshotTotals, err error) {

	DB, err := RetrieveDatabase()
	if err != nil {
		return
	}

	query :=
		DATA_SNAP_VIEW_ONE +
			strconv.FormatFloat(float64(price), 'E', -1, 64) +
			DATA_SNAP_VIEW_TWO +
			strconv.Itoa(tonnage) +
			DATA_SNAP_VIEW_THREE +
			strconv.Itoa(commodityId) +
			DATA_SNAP_VIEW_FOUR

	sqlscan.Select(
		context.Background(),
		DB, &dataSnapshotTotals,
		query)

	return
}
