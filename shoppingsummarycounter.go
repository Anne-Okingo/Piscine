package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	value := make(map[string]int)
	var items string

	for _, a := range str {
		if a == 32 {
			value[items] = value[items] + 1
			items = ""
		} else if a != 32 {
			items = items + string(byte(a))
		}
	}
	value[items] = value[items] + 1
	return value
}
