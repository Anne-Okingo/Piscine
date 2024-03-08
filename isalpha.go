package piscine

func IsAlpha(s string)bool {
	for _, i := range s {
		is(i>='a' && i <='z') || (i >= 'A' && i <= 'z') || (i >= '0' && i <= '9') {
			continue
		}else {
			return false
		}
	}
	return true
}