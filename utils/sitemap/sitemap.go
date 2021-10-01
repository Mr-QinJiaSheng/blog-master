package sitemap

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

var url = make(map[string]bool)

func Sitemap(dir, baseUrl string) {
	//var dir string
	//var baseUrl string

	//flag.StringVar(&dir, "dir", "", "is ok")
	//flag.StringVar(&baseUrl, "url", "", "is ok")
	//flag.Parse()

	if dir == "" || baseUrl == "" {
		fmt.Printf("%s", "dir or url no allowed null")
		return
	}

	c := colly.NewCollector()

	c.OnHTML("a", func(e *colly.HTMLElement) {
		href := e.Attr("href")
		if (url[href] == false && strings.Contains(href, baseUrl)) || strings.HasPrefix(href, "/") {

			if strings.HasPrefix(href, "/") {
				href = baseUrl + href
			}
			url[href] = true
			e.Request.Visit(href)
		}
	})

	c.Visit(baseUrl)

	write(url, dir)

}

func write(url map[string]bool, path string) {

	uniqurl := make(map[string]bool, len(url))

	for k := range url {
		index := strings.Index(k, "?")

		if index != -1 {
			k = k[:index]
		}

		if _, ok := uniqurl[k]; !ok {
			uniqurl[k] = true
		}

	}

	//fmt.Println(uniqurl)

	path += "/sitemap.xml"

	var str = `<?xml version="1.0" encoding="UTF-8"?>
<urlset xmlns="http://www.sitemaps.org/schemas/sitemap/0.9"
       xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
       xsi:schemaLocation="http://www.sitemaps.org/schemas/sitemap/0.9
                           http://www.sitemaps.org/schemas/sitemap/0.9/sitemap.xsd">
`

	for k := range uniqurl {
		str += `
		<url>
			<loc>` + k + `</loc>
			<priority>0.9</priority>
			<lastmod>` + time.Now().Format("2006-01-02 15:04:05") + `</lastmod>
			<changefreq>weekly</changefreq>
		</url>
`
	}

	str += `
</urlset>
`
	outputFile, outputError := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0666)
	if outputError != nil {
		fmt.Printf("An error occurred with file opening or creation，%s\n", outputError.Error())
		return
	}
	defer outputFile.Close()

	outputWriter := bufio.NewWriter(outputFile)
	outputWriter.WriteString(str)
	outputWriter.Flush()
}
