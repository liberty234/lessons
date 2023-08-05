package names

import "fmt"

func getNameListed() {
	NameListed := []string{"Iredia Ebikade", "Peace Ebikade", "Faith Ebikade", "Liberty Ebikade"}

	NameListed[2] = "Unity Ebikade"

	NameListed = append(NameListed, "Princess Ebikade", "Best Ebikade")

	for _, v := range NameListed {
		fmt.Println(v)

	}

}
