package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strconv"
)

// Nested slice
type Company struct {
	Id       int      `json:"id,omitempty"`
	Name     string   `json:"name"`
	Founders []string `json:"founders"`
	Year     int      `json:"year"`
}

// With Go, we can not define slice/array/map as const
// meaning it will be errored with "const companies []Company = []Company{...}"
// https://stackoverflow.com/a/41289218
func GetCompanyList() []Company {
	c := []Company{
		{Id: 1, Name: "Google", Founders: []string{"Larry Page", "Sergey Brin"}, Year: 1998},
		{Id: 2, Name: "Amazon", Founders: []string{"Jeff Bezos"}, Year: 1995},
		{Id: 3, Name: "Facebook", Founders: []string{"Mark Zuckerberg"}, Year: 2004},
		{Id: 4, Name: "Apple", Founders: []string{"Steve Jobs", "Steve Wozniak", "Ronald Wayne"}, Year: 2004},
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
		resp, isCreated := AddCompany(r)
		if isCreated {
			json.NewEncoder(w).Encode(resp)
		} else {
			fmt.Println("Failed to create company into database records.")
			json.NewEncoder(w).Encode(nil)
		}

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
		fmt.Println("Got POST request to /api/v1/companies/{id}, Method Not Allowed")
		json.NewEncoder(w).Encode(nil)

	case http.MethodPatch:
		fmt.Println("PATCH company to /api/v1/companies/{id}")
		resp, isUpdated := UpdateCompany(r)
		if isUpdated {
			w.WriteHeader(http.StatusAccepted)
			json.NewEncoder(w).Encode(resp)
		} else {
			fmt.Println("Failed to update company from database records.")
			json.NewEncoder(w).Encode(nil)
		}

	case http.MethodDelete:
		fmt.Println("DELETE company to /api/v1/companies/{id}")
		t := GetCompany(r)
		isDeleted := DeleteCompany(t)
		if isDeleted {
			w.WriteHeader(http.StatusNoContent)
			json.NewEncoder(w).Encode(nil)
		} else {
			fmt.Println("Failed to delete company from database records.")
			json.NewEncoder(w).Encode(t)
		}

	default:
		fmt.Printf("Got [ %s] to /api/v1/companies, Method Not Allowed\n", r.Method)
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(nil)
	}

}

// each CRUDs methods
func GetAllCompanies() []Company {
	fmt.Println(">>> Querying all companies from database ...")
	db, err := GetDatabaseInstance()
	if err != nil {
		fmt.Printf("Failed to connect %s database\n", DB_TYPE)
		fmt.Println(err.Error())
	}

	// Run queries
	q := "SELECT * FROM companies;"
	rows, e := db.Query(q)
	if e != nil {
		fmt.Printf("Failed to execute query: [ %s ]\n", q)
		fmt.Println(e)
	}

	// iterate and build response body
	var companies []Company
	for rows.Next() {
		// each field should be defined as variables to target table
		var id int
		var name string
		var founder string
		var year int
		err := rows.Scan(&id, &name, &founder, &year)
		if err != nil {
			fmt.Printf("Failed to get record\n")
		}
		c := Company{
			Id:       id,
			Name:     name,
			Founders: []string{founder},
			Year:     year,
		}
		companies = append(companies, c)
	}

	// GetEstablishedHYear()
	return companies
}

func GetCompany(req *http.Request) Company {
	// Fetch company_id from URL path params with string, and cast it to int
	fmt.Printf("The id of company: [ %s ]\n", req.PathValue("id"))
	company_id, _ := strconv.Atoi(req.PathValue("id"))

	if company_id == 0 || company_id < 0 {
		fmt.Printf("company_id [ %v ] invalid, should provide non-negative values\n", company_id)
		return Company{
			Id:       0,
			Name:     "",
			Founders: []string{},
			Year:     0,
		}
	}

	db, err := GetDatabaseInstance()
	if err != nil {
		fmt.Printf("Failed to connect %s database\n", DB_TYPE)
		fmt.Println(err.Error())
	}

	company := db.QueryRow("SELECT * FROM `companies` WHERE id = ?;", company_id)
	var id int
	var name string
	var founder string
	var year int
	e := company.Scan(&id, &name, &founder, &year)
	if e != nil {
		if e == sql.ErrNoRows {
			fmt.Printf("Records for company_id=%v not found\n", company_id)
		}
		fmt.Printf("Failed to get record\n")
	}

	return Company{
		Id:       id,
		Name:     name,
		Founders: []string{founder},
		Year:     year,
	}

}

func AddCompany(req *http.Request) (Company, bool) {
	var rbody Company
	if err := json.NewDecoder(req.Body).Decode(&rbody); err != nil {
		fmt.Println("Failed to parse POST request")
		fmt.Println(err)
		return rbody, false
	}
	fmt.Println("POST request, got: ")
	fmt.Println(rbody)

	db, e := GetDatabaseInstance()
	if e != nil {
		fmt.Printf("Failed to connect %s database\n", DB_TYPE)
		fmt.Println(e.Error())
		return rbody, false
	}

	if rbody.Id == 0 {
		companies := GetAllCompanies()
		// Sort slices for calculate newest Company.Id
		sort.SliceStable(companies, func(i int, j int) bool { return companies[i].Id < companies[j].Id })
		rbody.Id = companies[len(companies)-1].Id + 1
	} else {
		q := "SELECT * FROM companies WHERE id = ?;"
		if _, e := db.Query(q, rbody.Id); e != nil {
			fmt.Printf("Company object whose id is [ %v ] has been already existed in the database\n", rbody.Id)
			fmt.Println(e)
		}
		return rbody, false
	}

	_, err := db.Exec(
		"INSERT INTO companies (id, name, founder, year) VALUES (?, ?, ?, ?)",
		rbody.Id,
		rbody.Name,
		rbody.Founders[0],
		rbody.Year,
	)
	if err != nil {
		fmt.Println("Failed to INSERT data into database")
		fmt.Println(err)
		return rbody, false
	}

	return rbody, true
}

func UpdateCompany(req *http.Request) (Company, bool) {
	var rbody Company
	var company_id int
	if err := json.NewDecoder(req.Body).Decode(&rbody); err != nil {
		fmt.Println("Failed to parse PATCH request")
		fmt.Println(err)
		return rbody, false
	}
	_, csterr := fmt.Sscanf(req.PathValue("id"), "%d", &company_id)
	if csterr != nil {
		fmt.Println("Failed to convert company_id as int.")
		fmt.Println(csterr)
		return rbody, false
	}

	fmt.Printf("PATCH request to id = %d, got: ", company_id)
	fmt.Println(rbody)

	db, err := GetDatabaseInstance()
	if err != nil {
		fmt.Printf("Failed to connect %s database\n", DB_TYPE)
		fmt.Println(err.Error())
		return rbody, false
	}

	q := "UPDATE companies SET name = ?, founder = ?, year = ? WHERE id = ?;"
	_, e := db.Exec(
		q,
		rbody.Name,
		rbody.Founders[0],
		rbody.Year,
		company_id,
	)
	if e != nil {
		fmt.Printf("Failed to update records\n")
		fmt.Println(e)
		return rbody, false
	}

	return rbody, true
}

func DeleteCompany(c Company) bool {
	fmt.Println("DELETE request!")

	db, err := GetDatabaseInstance()
	if err != nil {
		fmt.Printf("Failed to connect %s database\n", DB_TYPE)
		fmt.Println(err.Error())
	}

	q := "DELETE FROM companies WHERE id = ?;"
	_, e := db.Exec(q, c.Id)
	if e != nil {
		fmt.Printf("Failed to delete records\n")
		fmt.Println(e)
		return false
	}

	fmt.Printf("Deletion completed: [ %s ]\n", c.Name)
	return true
}

// // implementations of enum
// type CompanyName string
// type GAFA struct {
// 	Name CompanyName
// }

// const (
// 	GOOGLE   = CompanyName("Google")
// 	APPLE    = CompanyName("Apple")
// 	FACEBOOK = CompanyName("Facebook")
// 	AMAZON   = CompanyName("Amazon")
// )

// func GetEstablishedHYear() {
// 	gafa := GAFA{Name: GOOGLE}
// 	fmt.Printf("%s is one of the companies in GAFA \n", gafa.Name)
// }
