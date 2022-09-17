package symbols

import "fmt"

func ConvertTimeToSymbols(h int, m int, s int) [8]symbol {
	// first get all the digits needed, including the leading zeroes
	digits := [8]int{
		h / 10,
		h % 10,
		0, // colon
		m / 10,
		m % 10,
		0, // colon
		s / 10,
		s % 10,
	}

	time := [8]symbol{
		2: colon,
		5: colon,
	}

	for i := 0; i < 8; i++ {
		switch i {
		case 2, 5:
			continue
		default:
			switch digits[i] {
			case 0:
				time[i] = zero
			case 1:
				time[i] = one
			case 2:
				time[i] = two
			case 3:
				time[i] = three
			case 4:
				time[i] = four
			case 5:
				time[i] = five
			case 6:
				time[i] = six
			case 7:
				time[i] = seven
			case 8:
				time[i] = eight
			case 9:
				time[i] = nine
			}
		}
	}

	return time

}

func PrintClock(time [8]symbol) {
	for i := 0; i < len(zero); i++ {
		for j := 0; j < len(time); j++ {
			fmt.Printf("%s ", time[j][i])
		}
		fmt.Printf("\n")
	}
}
