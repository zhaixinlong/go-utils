package util

import (
	"strconv"

	"github.com/shopspring/decimal"
)

// 向上取整
func DecimalRoundPoint(decimalValue float64, point int32) float64 {
	value, _ := decimal.NewFromFloat(decimalValue).RoundUp(point).Float64()
	return value
}

// 向下取整
func decimalRoundFloorPoint(decimalValue float64, point int32) float64 {
	value, _ := decimal.NewFromFloat(decimalValue).RoundFloor(point).Float64()
	return value
}

func DecimalRoundFloorTwoPoint(decimalValue float64) float64 {
	return decimalRoundFloorPoint(decimalValue, 2)
}

func DecimalRoundFloorThreePoint(decimalValue float64) float64 {
	return decimalRoundFloorPoint(decimalValue, 3)
}

func AddFloatForString(num1, num2 string) float64 {
	d1, _ := strconv.ParseFloat(num1, 64)
	d2, _ := strconv.ParseFloat(num2, 64)
	decimalD1 := decimal.NewFromFloat(d1)
	decimalD2 := decimal.NewFromFloat(d2)
	decimalResult := decimalD1.Add(decimalD2)
	float64Result, _ := decimalResult.Float64()
	return float64Result
}

func BcDivFloat(d1, d2 float64) float64 {
	decimalD1 := decimal.NewFromFloat(d1)
	decimalD2 := decimal.NewFromFloat(d2)
	decimalResult := decimalD1.Div(decimalD2)
	float64Result, _ := decimalResult.Float64()
	return float64Result
}

func AddFloat(d1, d2 float64) float64 {
	decimalD1 := decimal.NewFromFloat(d1)
	decimalD2 := decimal.NewFromFloat(d2)
	decimalResult := decimalD1.Add(decimalD2)
	float64Result, _ := decimalResult.Float64()
	return float64Result
}

// 乘法
func MulFloat(d1, d2 float64) float64 {
	decimalD1 := decimal.NewFromFloat(d1)
	decimalD2 := decimal.NewFromFloat(d2)
	decimalResult := decimalD1.Mul(decimalD2)
	float64Result, _ := decimalResult.Float64()
	return float64Result
}
