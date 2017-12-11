package entity

/**
	读取消息结构
 */
type PowerInfo struct {
	PowerMeterId string
	Time string
	Zxygdn float64
	Zxygdn1 float64
	Zxygdn2 float64
	Zxygdn3 float64
	Zxygdn4 float64
}

var PowerInfoMap map[string] *PowerInfo
