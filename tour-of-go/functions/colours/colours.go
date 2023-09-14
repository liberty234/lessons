package colours

import "fmt"

func GetColors(color1, color2, color3, color4, color5 string) (string, string, string, string, string) {
	b := fmt.Sprintf("Color1: %s", color1)
	g := fmt.Sprintf("Color1: %s", color2)
	w := fmt.Sprintf("Color1: %s", color3)
	r := fmt.Sprintf("Color1: %s", color4)
	y := fmt.Sprintf("Color1: %s", color5)
	return b, g, w, r, y

}
