package apiServiceLayerTests

import (
	"github.com/jarrettnester/fruitPal/apiServiceLayer"
	"testing"
)

func TestRetrieveTotalCostInvalidCommodityId(t *testing.T) {
	_, err := apiServiceLayer.RetrieveTotalCost(34, 0, 45)
	if err == nil {
		t.Error("Should have been an error because commodity id was less than 1")
	}
}

func TestRetrieveTotalCostInvalidPrice(t *testing.T) {
	_, err := apiServiceLayer.RetrieveTotalCost(0, 1, 45)
	if err == nil {
		t.Error("Should have been an error because price is not greater than 0")
	}
}

func TestRetrieveTotalCostInvalidTonnage(t *testing.T) {
	_, err := apiServiceLayer.RetrieveTotalCost(34, 1, 0)
	if err == nil {
		t.Error("Should have been an error because tonnage is less than 1")
	}
}
