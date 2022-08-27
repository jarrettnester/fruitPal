package main

import (
	"encoding/json"
	"fmt"
	"github.com/jarrettnester/fruitPal/apiServiceLayer"
	"net/http"
	"strconv"
)

func GetLatestCommoditiesHandler(w http.ResponseWriter, r *http.Request) {

	availableCommodities, err := apiServiceLayer.RetrieveCommoditiesOnLatestSnapshot()

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	availableCommodityBytes, err := json.Marshal(availableCommodities)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(availableCommodityBytes)
}

func GetTotalCost(w http.ResponseWriter, r *http.Request) {

	var price float64
	var commodityId int
	var tonnage int
	var err error

	price, err = strconv.ParseFloat(r.URL.Query().Get("price"), 64)
	if err != nil {
		return
	}

	commodityId, err = strconv.Atoi(r.URL.Query().Get("commodity_id"))
	if err != nil {
		return
	}

	tonnage, err = strconv.Atoi(r.URL.Query().Get("tonnage"))
	if err != nil {
		return
	}

	totalCostData, err := apiServiceLayer.RetrieveTotalCost(float32(price), commodityId, tonnage)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	totalCostDataBytes, err := json.Marshal(totalCostData)

	if err != nil {
		fmt.Println(fmt.Errorf("Error: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(totalCostDataBytes)
}
