package main

import (
	"fmt"
	"regexp"
	"strconv"
)

var hasil []string

func KonversiNumberToRomawi(input int) []string {
	var satuan, _ = regexp.Compile(`(\d)`)
	var puluhan, _ = regexp.Compile(`(\d\d)`)
	var ratusan, _ = regexp.Compile(`(\d\d\d)`)
	var ribuan, _ = regexp.Compile(`(\d\d\d\d)`)

	var s string = strconv.Itoa(input)

	for true {
		if ribuan.MatchString(s) {
			var n, _ = strconv.Atoi(s)
			var x = n
			fmt.Println(x)
	
			for x >= 1000 {
				hasil = append(hasil, "M")
				x -= 1000
			}
			s = strconv.Itoa(x)
		}
		if ratusan.MatchString(s) {
			var n, _ = strconv.Atoi(s)
			var x = n
			fmt.Println(x)
	
			for x >= 100 {
				if x >= 900 {
					hasil = append(hasil, "C M")
					x -= 900
				} else if x >= 500 {
					hasil = append(hasil, "D")
					x -= 500
				} else if x >= 400 {
					hasil = append(hasil, "C D")
					x -= 400
				}else {
					hasil = append(hasil, "C")
					x -= 100
				}
			}
			s = strconv.Itoa(x)
		}
		if puluhan.MatchString(s) {
			var n, _ = strconv.Atoi(s)
			var x = n

			fmt.Println(x)
	
			for x >= 10 {
				if x >= 90 {
					hasil = append(hasil, "X C")
					x -= 90
				} else if x >= 50 {
					hasil = append(hasil, "L")
					x -= 50
				} else if x >= 40 {
					hasil = append(hasil, "X L")
					x -= 40
				}else {
					hasil = append(hasil, "X")
					x -= 10
				}
			}
			s = strconv.Itoa(x)
		}
		if satuan.MatchString(s) {
			var n, _ = strconv.Atoi(s)
			var x = n
	
			for x > 0 {
				if x >= 9 {
					hasil = append(hasil, "I X")
					x -= 9
				} else if x >= 5 {
					hasil = append(hasil, "V")
					x -= 5
				} else if x >= 4 {
					hasil = append(hasil, "I V")
					x -= 4
				}else {
					hasil = append(hasil, "I")
					x -= 1
				}
			}
			break
		}
	}
	return hasil
}

func main() {
	fmt.Println(KonversiNumberToRomawi(3999))
}
