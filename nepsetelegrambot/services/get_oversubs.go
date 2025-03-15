package services

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"golang.org/x/net/html"
)

type CompanyIssue struct {
	CompanyName     string
	IssueManager    string
	IssuedUnit      string
	NumApplications string
	AppliedUnit     string
	Amount          string
	OpenDate        string
	CloseDate       string
}

func fetchDataFromAPI() (string, error) {
	// Replace with your actual API endpoint
	apiURL := os.Getenv("API_URL")

	resp, err := http.Get(apiURL)
	if err != nil {
		return "", fmt.Errorf("error making API request: %v", err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error reading response body: %v", err)
	}

	return string(body), nil
}

func extractTableData(n *html.Node) []CompanyIssue {
	var companies []CompanyIssue
	var currentCompany CompanyIssue
	var inTbody bool
	var cellCount int

	var traverse func(*html.Node)
	traverse = func(n *html.Node) {
		if n.Type == html.ElementNode {
			if n.Data == "tbody" {
				inTbody = true
			}
		}

		if inTbody && n.Type == html.ElementNode && n.Data == "td" {
			var text string
			for c := n.FirstChild; c != nil; c = c.NextSibling {
				if c.Type == html.TextNode {
					text = strings.TrimSpace(c.Data)
					break
				}
			}

			switch cellCount {
			case 1:
				currentCompany.CompanyName = text
			case 2:
				currentCompany.IssueManager = text
			case 3:
				currentCompany.IssuedUnit = text
			case 4:
				currentCompany.NumApplications = text
			case 5:
				currentCompany.AppliedUnit = text
			case 6:
				currentCompany.Amount = text
			case 7:
				currentCompany.OpenDate = text
			case 8:
				currentCompany.CloseDate = text
			}

			cellCount++
			if cellCount == 10 {
				companies = append(companies, currentCompany)
				currentCompany = CompanyIssue{}
				cellCount = 0
			}
		}

		for c := n.FirstChild; c != nil; c = c.NextSibling {
			traverse(c)
		}

		if n.Type == html.ElementNode && n.Data == "tbody" {
			inTbody = false
		}
	}

	traverse(n)
	return companies
}

func GetIPOOverscribeData(symbol string) string {
	// Fetch data from API
	htmlContent, err := fetchDataFromAPI()
	if err != nil {
		log.Fatalf("Failed to fetch data: %v", err)
	}

	// Parse HTML content
	doc, err := html.Parse(strings.NewReader(htmlContent))
	if err != nil {
		log.Fatalf("Error parsing HTML: %v", err)
	}

	// Extract company information
	companies := extractTableData(doc)

	// Print the extracted information
	for _, company := range companies {
		if strings.Contains(company.CompanyName, symbol) {
			issuedUnit, err := strconv.ParseFloat(strings.ReplaceAll(company.IssuedUnit, ",", ""), 64)
			if err != nil {
				log.Fatalf("Error converting IssuedUnit to float: %v", err)
			}
			appliedUnit, err := strconv.ParseFloat(strings.ReplaceAll(company.AppliedUnit, ",", ""), 64)
			if err != nil {
				log.Fatalf("Error converting AppliedUnit to float: %v", err)
			}
			return fmt.Sprintf("%.2f", appliedUnit/issuedUnit)
		}
	}
	return ""
}
