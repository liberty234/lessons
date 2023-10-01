package state

import "fmt"

func FetchStateName(state1, state2, state3, state4 string) (state12 string, state13 string, state14 string, state15 string) {
	state12 = fmt.Sprintf("State:%s", state1)
	state13 = fmt.Sprintf("State:%s", state2)
	state14 = fmt.Sprintf("State:%s", state3)
	state15 = fmt.Sprintf("State:%s", state4)
	return

}
