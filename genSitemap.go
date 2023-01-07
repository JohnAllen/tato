package main

import (
	"fmt"
	"os"
)

func main() {

	f, err := os.Create("sitemap.xml")
	if err != nil {
		fmt.Println(err)
	}
	var bpStringOne = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\"><url>\n"
	_, err = f.WriteString(bpStringOne)
	if err != nil {
		fmt.Println("write string error", err)
	}
	files, err := os.ReadDir("static")
	if err != nil {
		fmt.Println("read dir err", err)
	}
	for _, file := range files {
		fileName := file.Name()
		completeUrl := "https://rarelanguages.net/" + fileName
		//	fmt.Println(completeUrl)
		finalString := "<url><loc>" + completeUrl + "</loc><lastmod>2022-01-07</lastmod></url>\n"
		fmt.Println(finalString)
		f.WriteString(finalString)
	}
	f.WriteString("</urlset>")
}
