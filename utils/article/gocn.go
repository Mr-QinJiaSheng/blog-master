package article

import (
	"strings"

	"github.com/gocolly/colly"
)

type Gocn struct {
}

func (gocn Gocn) Get(url string) map[string]interface{} {

	ret := make(map[string]interface{})
	c := colly.NewCollector()
	c.OnHTML("#main .title", func(e *colly.HTMLElement) {
		ret["title"] = strings.TrimSpace(e.Text)
	})
	// On every a element which has href attribute call callback
	c.OnHTML(".col-md-9 .topic-detail .card-body", func(e *colly.HTMLElement) {
		//fmt.Println(e.DOM.Children().Filter(".card").Remove())
		e.DOM.Children().Filter(".toc-container").Remove()
		//d.ChildrenFiltered()
		e.DOM.Children().Filter(".card").Remove()
		ret["content"], _ = e.DOM.Html()
	})
	// Start scraping on https://hackerspaces.org
	c.Visit(url)
	return ret
}
