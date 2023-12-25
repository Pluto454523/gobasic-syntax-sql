package basicSyntax

import "fmt"

func MapTutorial() {

	countries := map[string]string{}
	countries["th"] = "Thailand"
	countries["en"] = "United State"

	fmt.Println(countries["en"])

	country, msgCountry := countries["asd"]
	if msgCountry {
		fmt.Println(country)
	} else {
		fmt.Println("no vaule")
	}

}
