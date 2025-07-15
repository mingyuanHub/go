package main

import (
	"fmt"
	"strconv"
	"time"
)

type Kline struct {
	Day     string  `json:"day"`
	Open    float64 `json:"open"`
	Close   float64 `json:"close"`
	Highest float64 `json:"highest"`
	Lowest  float64 `json:"lowest"`
}

var klineData = []*Kline{
	{"2024-04-01", 83.69, 80.60, 83.88, 80.55},
	{"2024-04-02", 80.21, 80.43, 81.12, 79.96},
	{"2024-04-03", 80.43, 78.60, 80.63, 77.81},
	{"2024-04-04", 78.49, 77.02, 79.28, 76.71},
	{"2024-04-05", 77.01, 75.40, 77.79, 75.12},
	{"2024-04-08", 77.18, 80.22, 82.94, 76.91},
	{"2024-04-09", 79.25, 81.70, 83.50, 78.66},
	{"2024-04-10", 81.27, 82.07, 83.00, 80.59},
	{"2024-04-11", 81.49, 80.81, 82.38, 80.71},
	{"2024-04-12", 80.03, 82.25, 84.90, 80.03},
}

const (
	PriceTimeOpen = iota
	PriceTimeClose
)

func main() {
	for _, wdc := range getWeekdayCombine() {
		for _, ptc := range getPriceTimeCombine() {
			getWeekdayKlineData(klineData, wdc[0], wdc[1], ptc[0], ptc[1])
		}
	}
}

func getWeekdayCombine() [][]time.Weekday {
	var arr = []time.Weekday{time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday}
	var cb = [][]time.Weekday{}
	for _, v1 := range arr {
		for _, v2 := range arr {
			cb = append(cb, []time.Weekday{v1, v2})
		}
	}
	return cb
}

func getPriceTimeCombine() [][]int {
	var arr = []int{PriceTimeOpen, PriceTimeClose}
	var cb = [][]int{}
	for _, v1 := range arr {
		for _, v2 := range arr {
			cb = append(cb, []int{v1, v2})
		}
	}
	return cb
}

func getWeekdayKlineData(klineData []*Kline, buyWeekday, sellWeekday time.Weekday, buyPriceTime, sellPriceTime int) []float64 {
	var (
		data    []float64 = []float64{0}
		start   float64
		isStart bool = true
		now     float64
		buy     float64
		sell    float64
		isNext  bool
	)

	for _, k := range klineData {

		if getWeekday(k.Day) == buyWeekday && !isNext {
			if buyPriceTime == PriceTimeOpen {
				buy = k.Open
			} else if buyPriceTime == PriceTimeClose {
				buy = k.Close
			}

			if isStart {
				start = buy
				isStart = false
			}

			isNext = true
			//now = buy
			continue
		}

		if getWeekday(k.Day) == sellWeekday {
			if sellPriceTime == PriceTimeOpen {
				sell = k.Open
			} else if sellPriceTime == PriceTimeClose {
				sell = k.Close
			}

			if buy > 0 && sell > 0 {
				now = floatP4(now + (sell - buy) * 100 / start )
				data = append(data, now)
				buy = 0
				sell = 0
				isNext = false
			}
		}
	}

	//fmt.Println(buyWeekday, sellWeekday, buyPriceTime, sellPriceTime, data)
	return data
}

func getWeekday(date string) time.Weekday {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		panic(err)
	}
	return t.Weekday()
}

func floatP4(v float64) float64 {
	v, _ = strconv.ParseFloat(fmt.Sprintf("%.4f", v), 64)
	return v
}