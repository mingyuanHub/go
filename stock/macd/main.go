package main

import (
	"fmt"
)

// 计算指数移动平均线(EMA)
func EMA(values []float64, period int) []float64 {
	if len(values) < period {
		return nil
	}
	ema := make([]float64, len(values))
	multiplier := 2.0 / float64(period+1)
	ema[period-1] = SimpleMA(values[:period], period)

	for i := period; i < len(values); i++ {
		ema[i] = ((values[i]-ema[i-1])*multiplier) + ema[i-1]

		//等同于如下： 前一日EMA（12）×11/13+今日收盘价×2/13
		//ema[i] = ema[i-1]*(1-multiplier) + values[i]*multiplier

		//例如：(12-5.5)*2/13 + 5.5 = 5.5*11/13 + 12*2/13
	}
	return ema
}

// 计算简单移动平均线(SMA)
func SimpleMA(values []float64, period int) float64 {
	sum := 0.0
	for _, v := range values[:period] {
		sum += v
	}
	return sum / float64(period)
}

// 计算MACD
func MACD(values []float64) (DIF []float64, DEA []float64, MACD []float64) {
	if len(values) < 27 { // 至少需要27个数据点（12日EMA + 26日EMA + 9日EMA）
		return nil, nil, nil
	}

	EMA12 := EMA(values, 12)
	EMA26 := EMA(values, 26)

	DIF = make([]float64, len(EMA12))
	for i := 25; i < len(EMA12); i++ { // 从第26个数据点开始，因为EMA需要之前的12和26个数据点
		DIF[i] = EMA12[i] - EMA26[i]
	}

	DEA = EMA(DIF, 9)

	MACD = make([]float64, len(DEA))
	for i := 0; i < len(DEA); i++ {
		MACD[i] = 2*(DIF[i]-DEA[i]) // MACD柱是DIF与DEA之差的2倍，这是为了放大效果，便于观察
	}

	return DIF[25:len(values)], DEA[25:len(values)], MACD[25:len(values)] // 截去最开始无法计算的部分
}

func main() {
	// 示例价格数据
	prices := []float64{}
	for i := 0.0; i < 30; i ++ {
		prices = append(prices, i)
	}

	DIF, DEA, MACD1 := MACD(prices)

	fmt.Println(len(DIF), "DIF:", DIF)
	fmt.Println(len(DEA), "DEA:", DEA)
	fmt.Println(len(MACD1), "MACD:", MACD1)
}