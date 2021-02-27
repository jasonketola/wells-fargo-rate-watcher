package main

import (
	


	"github.com/gocolly/colly"
)




func main() {



	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("www.wellsfargorelo.com"),

	)






	var rates []string

	// On every a HTML element which has name attribute call callback
	c.OnHTML(`table[class="trTable trassumptiontable"] tbody tr`, func(e *colly.HTMLElement) {
		// Activate detailCollector if the class contains "product-link product-thumbnail"
		rates = append(rates, e.ChildText("td"))
	})




	myURL := "https://www.wellsfargorelo.com/relo/todaysRates_assumptions.page?loanPurpose=Refinance&loanType=1&productName=30-Year+Fixed+Rate&suffix=yourcompany1096"
	c.Visit(myURL)

	println(rates[0])
	
}
