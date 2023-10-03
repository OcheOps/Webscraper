Go Web Scraper
This is a simple Go web scraper application that uses the "colly" library to scrape data from a website. In this example, we scrape quotes from the Quotes to Scrape website.

Prerequisites
Before running the application, make sure you have Go installed on your system. You also need to install the "colly" library, which can be done using the following command:

go get github.com/gocolly/colly/v2

git clone this repo 

cd webscraper 

go run scraper.go


The scraper will visit the specified URL (http://quotes.toscrape.com), extract quotes, authors, and tags, and print them to the console.

Customization
You can customize the scraper to work with different websites or extract different types of data by modifying the code in scraper.go. Be sure to review the colly documentation for more advanced usage and customization options.

Legal Considerations
Please be aware of the legal implications of web scraping. Make sure to review and adhere to the website's terms of service and scraping policies. Respect robots.txt and use scraping responsibly.

License
This project is licensed under the MIT License. See the LICENSE file for details.


currently working on Version 2 but you can replace the url and change the functions you want to scrape