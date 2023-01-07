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

	results, err := os.ReadDir("/root/tato/data")
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

		completeFilePath := "/root/tato/data/" + result.Name()
		file, err := os.Open(completeFilePath)
		if err != nil {
			panic(err)
		}
		fmt.Println("Opened file: ", file.Name())

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
			if counter > 1_000_000 {
				break
			}
			firstLangContent := lineParts[1]
			firstLangContentUrl := stripContent(firstLangContent)
			secondLangContent := lineParts[3]
			srcLine := "How to say \"" + firstLangContent + "\" in " + srcLangName
			tgtContent := secondLangContent

			data := map[string]interface{}{
				"srcLangName": srcLangName,
				"srcLine":     srcLine,
				"tgtContent":  tgtContent,
			}
			var contentPageString = `<!doctype html><html lang='en'><head><meta charset='utwf-8'>  <title>RareLanguages.net</title></head><body><header
>  <h1><a href='/'>RareLanguages.net</a></h1></header><main>  <h2>{{.srcLine}}</h2><p>{{.tgtContent}}</p></main><footer></footer></body></html>`
			firstString := getHtmlString(contentPageString, data)
			filePath := "/root/tato/static/how-to-say-" + firstLangContentUrl + "-in-" + tgtLangName + ".html"
			writeHtmlStringToFile(filePath, firstString)
			counter++
		}

	}

}

func writeHtmlStringToFile(filePath string, html string) bool {
	fmt.Println("Writing to file: ", filePath)
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

	tmpl := template.Must(template.New("temp").Parse(a))
	buf := &bytes.Buffer{}
	err := tmpl.Execute(buf, b)
	if err != nil {
		panic(err)
	}
	s := buf.String()
	return s
}
