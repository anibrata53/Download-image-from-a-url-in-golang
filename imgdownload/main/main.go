//given a url,fetch all images and download them concurrently
package main

import (
	"fmt"
	"imgdownload/controllers"
)

func main() {
	url := "https://golang.org"
	fmt.Println("Processing...")
	message := controllers.FindImages(url)
	fmt.Println(message)

}
