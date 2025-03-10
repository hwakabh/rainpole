package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Nested slice
type Company struct {
	Name     string   `json:"name"`
	Founders []string `json:"founders"`
	Year     int      `json:"year"`
}

func GetAllCompanies(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">>> Getting all Companies...")

	companies := []Company{
		{Name: "Google", Founders: []string{"Larry Page", "Sergey Brin"}, Year: 1998},
		{Name: "Amazon", Founders: []string{"Jeff Bezos"}, Year: 1995},
		{Name: "Facebook", Founders: []string{"Mark Zuckerberg"}, Year: 2004},
		{Name: "Apple", Founders: []string{"Steve Jobs", "Steve Wozniak", "Ronald Wayne"}, Year: 2004},
	}

	for i, v := range companies {
		fmt.Println(i, v.Name)
		fmt.Printf("%s is founded in %v by %s\n", companies[i].Name, companies[i].Year, companies[i].Founders)
	}

	fmt.Println("Getting establish year of Google ...")
	GetEstablishedHYear()

	json.NewEncoder(w).Encode(companies)
}

// implementations of enum
type CompanyName string
type GAFA struct {
	Name CompanyName
}

const (
	GOOGLE   = CompanyName("Google")
	APPLE    = CompanyName("Apple")
	FACEBOOK = CompanyName("Facebook")
	AMAZON   = CompanyName("Amazon")
)

func GetEstablishedHYear() {
	gafa := GAFA{Name: GOOGLE}
	fmt.Printf("%s is one of the companies in GAFA \n", gafa.Name)
}
