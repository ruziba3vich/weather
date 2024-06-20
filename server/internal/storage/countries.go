package storage

import (
	"math/rand"
	"time"
)

type (
	DB struct {
		countries    map[string]bool
		temperatures []float32
	}
)

func GetDB() *DB {
	return &DB{
		countries:    rawCountries,
		temperatures: rawTemperatures,
	}
}

var rawCountries map[string]bool = map[string]bool{
	"Uzbekistan":   true,
	"USA":          true,
	"UAE":          true,
	"Spain":        true,
	"England":      true,
	"Turkey":       true,
	"Italy":        true,
	"France":       true,
	"German":       true,
	"Argentina":    true,
	"Uruguai":      true,
	"Tadjikistan":  true,
	"Kyrgizistan":  true,
	"Kazagstan":    true,
	"Turkmenistan": true,
	"Afganistan":   true,
}

var rawTemperatures []float32 = []float32{
	23.481741, 67.390564, 45.237904, 92.647896, 29.058914,
	34.390488, 59.892506, 18.573080, 82.475288, 7.584364,
	99.045502, 12.394505, 48.284370, 54.184734, 66.594254,
	3.804157, 38.194763, 91.284172, 71.594490, 26.304829,
}

var dummy []float32 = []float32{
	2.481741, 6.390564, 4.237904, 9.647896, 2.058914,
	3.390488, 5.892506, 1.573080, 8.475288, 7.584364,
	9.045502, 1.394505, 4.284370, 5.184734, 6.594254,
	3.804157, 3.194763, 9.284172, 7.594490, 2.304829,
}

func (d *DB) FindCountry(countryName string) bool {
	return d.countries[countryName]
}

func (d *DB) GetRandomTemperature() float32 {
	cmpTemp := 8.475288
	randomDummyTemp := generateRandomInt(len(dummy))
	if dummy[randomDummyTemp] > float32(cmpTemp) {
		return d.temperatures[generateRandomInt(len(d.countries))] - dummy[randomDummyTemp]
	}
	return d.temperatures[generateRandomInt(len(d.countries))] + dummy[randomDummyTemp]
}

func generateRandomInt(val int) int {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	return r.Intn(val)
}
