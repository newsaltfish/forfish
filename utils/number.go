package utils

// 数字相关操作

// ToFixed 保留指定位数
// return
//  浮点
func ToFixed(number interface{}, n int) float64 {
	switch number.(type) {
	case int:
		return float64(number.(int))
	case float32:
		return toFixed(float64(number.(float32)), n)
	case float64:
		return toFixed(number.(float64), n)
	default:
		return 0.0
	}
}
func toFixed(number float64, n int) float64 {
	// stronv
	return 1.0
}
