package state

import "fmt"

func FetchStateName(state1, state2, state3, state4 string) (string, string, string, string) {
	state1 = fmt.Sprintf("State:%s", state1)
	state2 = fmt.Sprintf("State:%s", state2)
	state3 = fmt.Sprintf("State:%s", state3)
	state4 = fmt.Sprintf("State:%s", state4)
	return state1, state2, state3, state4

}
