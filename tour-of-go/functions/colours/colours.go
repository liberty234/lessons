package colours

import "fmt"

func GetColors(color1, color2, color3, color4, color5 string) (string, string, string, string, string) {
	blueColor := fmt.Sprintf("Color1: %s", color1)
	greenColor := fmt.Sprintf("Color1: %s", color2)
	whiteColor := fmt.Sprintf("Color1: %s", color3)
	redColor := fmt.Sprintf("Color1: %s", color4)
	yellowColor := fmt.Sprintf("Color1: %s", color5)
	return blueColor, greenColor, whiteColor, redColor, yellowColor

}
