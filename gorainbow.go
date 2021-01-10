package gorainbow

import (
	"fmt"
	"math"
	"strings"
)

func rgb(i int) (int, int, int) {
	var f = 0.1

	return int(math.Sin(f*float64(i)+0)*127 + 128),
		int(math.Sin(f*float64(i)+2*math.Pi/3)*127 + 128),
		int(math.Sin(f*float64(i)+4*math.Pi/3)*127 + 128)
}

// Rainbow function returns a formated string ready to print it to the shell/terminal
func Rainbow(text string) string {
	var rainbowStr []string
	for index, value := range text {
		r, g, b := rgb(index)
		str := fmt.Sprintf("\033[38;2;%d;%d;%dm%c\033[0m", r, g, b, value)
		rainbowStr = append(rainbowStr, str)
	}

	return strings.Join(rainbowStr, "")
}
