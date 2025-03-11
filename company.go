package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

// Nested slice
type Company struct {
	Name     string   `json:"name"`
	Founders []string `json:"founders"`
	Year     int      `json:"year"`
}

// With Go, we can not define slice/array/map as const
// meaning it will be errored with "const companies []Company = []Company{...}"
func GetCompanyList() []Company {
	c := []Company{
		{Name: "Google", Founders: []string{"Larry Page", "Sergey Brin"}, Year: 1998},
		{Name: "Amazon", Founders: []string{"Jeff Bezos"}, Year: 1995},
		{Name: "Facebook", Founders: []string{"Mark Zuckerberg"}, Year: 2004},
		{Name: "Apple", Founders: []string{"Steve Jobs", "Steve Wozniak", "Ronald Wayne"}, Year: 2004},
	}
	return c
}

func GetAllCompanies(w http.ResponseWriter, r *http.Request) {
	fmt.Println(">>> Getting all Companies...")
	companies := GetCompanyList()

	for i, v := range companies {
		fmt.Println(i, v.Name)
		fmt.Printf("%s is founded in %v by %s\n", companies[i].Name, companies[i].Year, companies[i].Founders)
	}

	fmt.Println("Getting establish year of Google ...")
	GetEstablishedHYear()

	json.NewEncoder(w).Encode(companies)
}

func GetCompany(w http.ResponseWriter, r *http.Request) {
	// fmt.Println(r)
	fmt.Printf("The id of company: [ %s ]\n", r.PathValue("id"))

	// Fetch company_id from URL path params with string, and cast it to int
	i, _ := strconv.Atoi(r.PathValue("id"))

	companies := GetCompanyList()

	if len(companies)-1 < i {
		fmt.Printf("company_id [ %v ] is invalid \n", i)
		// returning "null" in response
		json.NewEncoder(w).Encode(nil)
	} else {
		resp := companies[i]
		json.NewEncoder(w).Encode(resp)
	}

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
