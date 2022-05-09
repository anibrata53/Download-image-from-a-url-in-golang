//this function fetches all image urls from the given url and stores them in result
package controllers

import (
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"golang.org/x/net/html"
)

func FindImages1(url string) string { //[]string
	result := make([]string, 0)
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()
	// reads html as a slice of bytes
	html1, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	new_html := string(html1) // convert slice of bytes to string

	//fmt.Printf("%s\n", new_html)

	//parsing
	doc, err := html.Parse(strings.NewReader(new_html))
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" {
			for _, img := range n.Attr {
				if img.Key == "src" {
					result = append(result, img.Val)

				}
			}

		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	//fmt.Println(result)
	images := result
	//we pass the image urls to DownloadImages1 which is used to
	//concurrently download all images with valid urls
	//and returns a completion message after the images
	//have been successfully downloaded
	message := DownloadImages1(images)
	return message

}
