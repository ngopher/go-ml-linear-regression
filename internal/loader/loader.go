package loader

import (
	"encoding/csv"
	"os"
	"strconv"
)

// Housing struct to hold data for each house
type Housing struct {
	Size  float64
	Rooms float64
	Price float64
}

// LoadData
// loads data from CSV file
func LoadData(filePath string) ([]Housing, error) {

	// Open CSV file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}

	// Create new CSV reader
	reader := csv.NewReader(file)

	// Read all rows into structs
	var houses []Housing
	for {
		row, err := reader.Read()
		if err != nil {
			break
		}

		size, _ := strconv.ParseFloat(row[0], 64)
		rooms, _ := strconv.ParseFloat(row[1], 64)
		price, _ := strconv.ParseFloat(row[2], 64)

		house := Housing{Size: size, Rooms: rooms, Price: price}
		houses = append(houses, house)
	}

	return houses, nil
}

// PredictPrice
// predicts the price based on linear regression
func PredictPrice(size, rooms float64) float64 {

	// Hard coded model coefficients
	// Can train model to learn these
	var coeffSize = 0.2
	var coeffRooms = 0.15
	var intercept float64 = 100000

	// Linear regression formula
	price := (coeffSize * size) + (coeffRooms * rooms) + intercept

	return price
}
