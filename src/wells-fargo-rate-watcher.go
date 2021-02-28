package main

import (
	"strconv"
	"log"
	"os"
	"github.com/gocolly/colly"
	"github.com/joho/godotenv"
)




func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env, %v", err)
	}

	// Instantiate default collector
	c := colly.NewCollector(
		colly.AllowedDomains("www.wellsfargorelo.com"),

	)



	var rates []string

	// On every a HTML element which has name attribute call callback
	c.OnHTML(`table[class="trTable trassumptiontable"] tbody tr`, func(e *colly.HTMLElement) {
		// Grab all the data from these cells. The first one will be the interest rate.
		rates = append(rates, e.ChildText("td"))
	})



	// Page for 30 year conforming loans
	myURL := "https://www.wellsfargorelo.com/relo/todaysRates_assumptions.page?loanPurpose=Refinance&loanType=1&productName=30-Year+Fixed+Rate&suffix=yourcompany1096"
	c.Visit(myURL)

	// Only do anything if colly successfully got some results back
	if len(rates) > 0 {
		// Trim the percent symbol off and convert string to a float
		stringRate := rates[0][:len(rates[0])-1]
		currentRate, _ := strconv.ParseFloat(stringRate, 32)

		if currentRate < THRESHOLD_RATE {
			emailRate(stringRate)
		} else {
			return
		}
	}
}

func emailRate(rate string) {
	host := os.Getenv("EMAIL_HOST")
	from := os.Getenv("EMAIL_FROM")
	password := os.Getenv("EMAIL_PASSWORD")
	port := os.Getenv("EMAIL_PORT")
	to := os.Getenv("EMAIL_TO")
	THRESHOLD_RATE, _ := strconv.ParseFloat(os.Getenv("THRESHOLD_RATE"))

	msg := "From: " + from + "\n" +
	"To: " + to + "\n" +
	"Subject: Attractive Wells Fargo interest rate - " + rate + "\n\n" +
	"Think about refinancing, please."

	addr := fmt.Sprintf("%s:%s", host, port)

	err := smtp.SendMail(addr,
		smtp.PlainAuth("", from, password, host),
		from, []string{to}, []byte(msg))

	if err != nil {
		log.Printf("SMTP error: %s", err)
		return
	}

}
