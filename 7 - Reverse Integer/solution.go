import "math"

func reverse(x int) int {
	var PMX = math.MaxInt32
	var MMX = math.MinInt32

	var re int = 0
	for ; x != 0; x /= 10 {
		var d = x % 10

		if x > 0 {
			if re > (PMX-d)/10 {
				return 0
			}
		} else {
			if re < (MMX-d)/10 {
				return 0
			}
		}

		re = re*10 + d
	}

	return re
}