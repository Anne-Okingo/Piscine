package piscine

func IterativePower(nb int, power int) int {
	if nb < 0 {
		return 0 
	}else if nb == 0 {
		return 1
	} 
	var result = 1
	for a := 1; a<= power; a++ {
		result = result * nb 
	}
return result
}
