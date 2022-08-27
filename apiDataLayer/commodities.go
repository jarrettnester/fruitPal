package apiDataLayer

import (
	"context"
	"github.com/georgysavva/scany/sqlscan"
	_ "github.com/lib/pq"
)

const (
	COMMODITY_QUERY = `
		SELECT id, name FROM commodities WHERE id IN (
    		SELECT commodity_id FROM data_snapshot_values WHERE data_snapshot_id = (
        		SELECT MAX(id) FROM data_snapshots
			)
		)`
)

type Commodity struct {
	Base
	Name string
}

func RetrieveCommoditiesOnLatestSnapshot() (commodities []Commodity, err error) {

	DB, err := RetrieveDatabase()
	if err != nil {
		return
	}

	sqlscan.Select(
		context.Background(),
		DB,
		&commodities,
		COMMODITY_QUERY)

	return
}

