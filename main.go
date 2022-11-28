package main

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"html/template"
	"os"
	"strings"
)

func stripContent(content string) string {
	content = strings.Replace(content, " ", "-", -1)
	content = strings.Replace(content, "?", "", -1)
	content = strings.Replace(content, "!", "", -1)
	content = strings.Replace(content, ".", "", -1)
	content = strings.Replace(content, ",", "", -1)
	content = strings.Replace(content, ":", "", -1)
	content = strings.Replace(content, ";", "", -1)
	content = strings.Replace(content, "(", "", -1)
	content = strings.Replace(content, ")", "", -1)
	content = strings.Replace(content, "'", "", -1)
	content = strings.Replace(content, "\"", "", -1)
	content = strings.Replace(content, "’", "", -1)
	content = strings.Replace(content, "‘", "", -1)
	content = strings.Replace(content, "“", "", -1)
	content = strings.Replace(content, "”", "", -1)
	content = strings.Replace(content, "¡", "", -1)
	content = strings.Replace(content, "¿", "", -1)
	content = strings.ToLower(content)
	return content
}

func main() {

	results, err := os.ReadDir("/Users/john/lingospring-tatoeba-html/data")
	if err != nil {
		fmt.Println(err)
	}

	for _, result := range results {
		tgtLangName := strings.Split(result.Name(), "-")[1]
		tgtLangName = strings.Split(tgtLangName, " -")[0]
		tgtLangName = strings.TrimSuffix(tgtLangName, " ")
		tgtLangName = strings.Replace(tgtLangName, " ", "-", -1)

		srcLangName := strings.Split(result.Name(), "-")[1]
		srcLangName = strings.Split(srcLangName, " -")[0]
		srcLangName = strings.TrimSuffix(srcLangName, " ")
		srcLangName = strings.Replace(srcLangName, " ", "-", -1)

		completeFilePath := "/Users/john/lingospring-tatoeba-html/data/" + result.Name()
		file, err := os.Open(completeFilePath)
		if err != nil {
			panic(err)
		}

		reader := csv.NewReader(file)
		reader.FieldsPerRecord = -1

		counter := 0
		for {
			record, err := reader.Read()
			if err != nil {
				break
			}
			lineParts := strings.Split(record[0], "\t")
			if len(lineParts) < 4 {
				fmt.Errorf("%v skipping as less than 4", completeFilePath)
				continue
			}
			if counter > 10000 {
				break
			}
			firstLangContent := lineParts[1]
			firstLangContentUrl := stripContent(firstLangContent)
			secondLangContent := lineParts[3]
			//secondLangContentUrl := stripContent(secondLangContent)
			srcLine := "How to say \"" + firstLangContent + "\" in " + srcLangName
			tgtContent := secondLangContent

			data := map[string]interface{}{
				"srcLine":    srcLine,
				"tgtContent": tgtContent,
			}
			var contentPageString = `<!doctype html><html lang='en'><head><meta charset='utwf-8'>  <title>LingoSpring</title></head><body><header
>  <h1><a href='/'>LingoSpring</a></h1></header><main>  <h2>{{.srcLine}}</h2><p>{{.tgtContent}}</p></main><footer></footer></body></html>`
			firstString := getHtmlString(contentPageString, data)
			filePath := "/Users/john/lingospring-tatoeba-html/html/how-to-say-" + firstLangContentUrl + "-in-" + tgtLangName + ".html"
			writeHtmlStringToFile(filePath, firstString)
			counter++
		}

	}

}

func writeHtmlStringToFile(filePath string, html string) bool {
	err := os.WriteFile(filePath, []byte(html), 0644)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func createSrcTgtFile(record []string, baseFilePath string, srcLangName string, inString string) bool {
	lineParts := strings.Split(record[0], "\t")
	if len(lineParts) < 3 {
		fmt.Errorf("%v lineParts is less than 3", baseFilePath)
		return false
	}
	secondLangContent := stripContent(lineParts[3])
	htmlFileName := baseFilePath + secondLangContent + inString + srcLangName
	fmt.Println(htmlFileName)
	return true
}

func getHtmlString(a string, b map[string]interface{}) string {

	tmpl := template.Must(template.New("content-page.html").Parse(a))
	buf := &bytes.Buffer{}
	err := tmpl.Execute(buf, b)
	if err != nil {
		panic(err)
	}
	s := buf.String()
	return s
}
