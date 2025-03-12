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
// https://stackoverflow.com/a/41289218
func GetCompanyList() []Company {
	c := []Company{
		{Name: "Google", Founders: []string{"Larry Page", "Sergey Brin"}, Year: 1998},
		{Name: "Amazon", Founders: []string{"Jeff Bezos"}, Year: 1995},
		{Name: "Facebook", Founders: []string{"Mark Zuckerberg"}, Year: 2004},
		{Name: "Apple", Founders: []string{"Steve Jobs", "Steve Wozniak", "Ronald Wayne"}, Year: 2004},
	}
	return c
}

// Allowing only GET/POST /api/v1/companies
func MultipleCompanyHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		resp := GetAllCompanies()
		json.NewEncoder(w).Encode(resp)
	case http.MethodPost:
		// TODO: Add companies
		fmt.Println("POST company to /api/v1/companies")
		resp := AddCompany()
		json.NewEncoder(w).Encode(resp)
	default:
		fmt.Printf("Got [ %s] to /api/v1/companies, Method Not Allowed", r.Method)
		// http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(nil)
	}
}

// Allowing CRUDs GET/POST/PATCH/DELETE /api/v1/companies/{id}
func SingleCompanyHandler(w http.ResponseWriter, r *http.Request) {
	if r.PathValue("id") == "" {
		fmt.Printf("company_id [ %s ] invalid\n", r.PathValue("id"))
		json.NewEncoder(w).Encode(nil)
	}

	switch r.Method {
	case http.MethodGet:
		resp := GetCompany(r)
		json.NewEncoder(w).Encode(resp)

	case http.MethodPost:
		// TODO: Add companies
		fmt.Println("POST company to /api/v1/companies/{id}")

	case http.MethodPatch:
		// TODO: Update companies
		fmt.Println("PATCH company to /api/v1/companies/{id}")

	case http.MethodDelete:
		// TODO: Delete companies
		fmt.Println("DELETE company to /api/v1/companies/{id}")
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(nil)

	default:
		fmt.Printf("Got [ %s] to /api/v1/companies, Method Not Allowed\n", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(nil)
	}

}

// each CRUDs methods
func GetAllCompanies() []Company {
	fmt.Println(">>> Getting all Companies...")
	companies := GetCompanyList()

	for i, v := range companies {
		fmt.Println(i, v.Name)
		fmt.Printf("%s is founded in %v by %s\n", companies[i].Name, companies[i].Year, companies[i].Founders)
	}

	fmt.Println("Getting establish year of Google ...")
	GetEstablishedHYear()
	return companies
}

func GetCompany(req *http.Request) Company {
	// Fetch company_id from URL path params with string, and cast it to int
	fmt.Printf("The id of company: [ %s ]\n", req.PathValue("id"))
	company_id, _ := strconv.Atoi(req.PathValue("id"))

	companies := GetCompanyList()

	// TODO: need to validate negative values such as -1
	if len(companies)-1 < company_id {
		fmt.Printf("company_id [ %v ] not found\n", company_id)
		// assign zero-values for each type
		return Company{
			Name:     "",
			Founders: []string{},
			Year:     0,
		}
	} else {
		return companies[company_id-1]
	}
}

func AddCompany() Company {
	fmt.Println("POST request!")
	return Company{
		Name:     "",
		Founders: []string{},
		Year:     0,
	}
}

func UpdateCompany() Company {
	fmt.Println("PATCH request!")
	return Company{
		Name:     "",
		Founders: []string{},
		Year:     0,
	}
}

func DeleteCompany() Company {
	fmt.Println("DELETE request!")
	return Company{
		Name:     "",
		Founders: []string{},
		Year:     0,
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
