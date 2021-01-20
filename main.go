package main

import (
	"fmt"
	"strings"

	"github.com/tebeka/selenium"
)

func main() {
	// Connect to the WebDriver instance running locally.
	port := 4444
	caps := selenium.Capabilities{"browserName": "firefox"}
	wd, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d", port))
	if err != nil {
		panic(err)
	}
	defer wd.Quit()
	// Navigate to the simple playground interface.
	if err := wd.Get("https://joelholmes.dev/"); err != nil {
		panic(err)
	}
	// Wait for the program to finish running and get the output.
	description, err := wd.FindElement(selenium.ByCSSSelector, ".site-description > p")
	if err != nil {
		panic(err)
	}
	descriptionText, err := description.Text()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s", strings.Replace(descriptionText, "\n\n", "\n", -1))
}
