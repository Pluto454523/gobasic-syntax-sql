package syntax

import "fmt"

func showAllCountries(x map[string]string) {
	fmt.Printf("\nstatus length: %d\n", len(x))
	// Iterate over the map
	for key, value := range x {
		fmt.Printf("%#v -> %#v\n", key, value)
	}
}

func MapTutorial() {

	// ** EXAMPLE 1# MAKE MAP: make(map[key]value)
	countries := make(map[string]string)
	countries["th"] = "Thailand"
	countries["en"] = "English"
	countries["us"] = "United State"

	// ? Check map nil value
	if countries == nil {
		fmt.Println("countries is empty.")
	} else {
		fmt.Println("countries = ", countries)
	}

	// ? call by key
	key := "en"
	fmt.Printf("countries[%q] = %#v\n", key, countries[key])

	// ? delete key
	delete(countries, key)

	// ? Checking if a key exits
	if country, ok := countries[key]; ok {
		fmt.Printf("countries[%q] = %#v\n", key, country)
	} else {
		fmt.Printf("countries[%q] = no value\n", key)
	}

	// ** EXAMPLE 2# Make map by value: map[key]value{key: value}
	status := map[int]string{
		200: "OK",
		308: "Permanent Redirect",
		400: "Bad Request",
		500: "Internal Server Error",
	}

	fmt.Printf("\nstatus length: %d\n", len(status))

	// ** Iterate over the map
	for key, value := range status {
		fmt.Printf("%#v -> %#v\n", key, value)
	}

}
