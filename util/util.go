package util

func boolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

func byteAtStr(b byte, s string) int {
	for i, _b := range []byte(s) {
		if _b == b {
			return i
		}
	}
	return -1
}

func inInts(e int, arr []int) bool {
	for _, _e := range arr {
		if e == _e {
			return true
		}
	}
	return false
}

// 258m 258p 258s 12345z 在不考虑国士无双和七对子时为八向听
var chineseShanten = []string{"听牌", "一向听", "两向听", "三向听", "四向听", "五向听", "六向听", "七向听", "八向听"}

func NumberToChineseShanten(num int) string {
	return chineseShanten[num]
}

func CountPairs(tiles34 []int) (pairs int) {
	for _, c := range tiles34 {
		if c >= 2 {
			pairs++
		}
	}
	return
}
