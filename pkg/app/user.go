package app

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID             int      `json:"_id"`
	Url            string   `json:"url"`
	ExternalId     string   `json:"external_id"`
	Name           string   `json:"name"`
	Alias          string   `json:"alias"`
	CreatedAt      string   `json:"created_at"`
	Active         bool     `json:"active"`
	Verified       bool     `json:"verified"`
	Shared         bool     `json:"shared"`
	Locale         string   `json:"locale"`
	Timezone       string   `json:"timezone"`
	LastLoginAt    string   `json:"last_login_at"`
	Email          string   `json:"email"`
	Phone          string   `json:"phone"`
	Signature      string   `json:"signature"`
	OrganizationID int      `json:"organization_id"`
	Tags           []string `json:"tags"`
	Suspended      bool     `json:"suspended"`
	Role           string   `json:"role"`
}


func findUsers(path string) []User {
	var obj []User

	data := readFile(path)
	err := json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
	}

	return obj
}

func DisplayUsersBasedOnSearchOptions(path string, searchValue int) string {
	tickets := findUsers(path)
	var results []User
	var users string

	for _, ticket := range tickets {
		if ticket.ID == searchValue {
			results = append(results, ticket)
		}
	}

	//TODO: Change return type to map and include dd ticket fields as well
	for _, result := range results {
		users = fmt.Sprintf("%v\n %s\n %s\n %s\n %s\n %s\n %t\n %t\n %t\n %s\n %s\n %s\n %s\n %s\n %s\n %v\n %v\n %t\n %s\n", result.ID, result.Url, result.ExternalId, result.Name, result.Alias, result.CreatedAt, result.Active, result.Verified, result.Shared, result.Locale, result.Timezone, result.LastLoginAt, result.Email, result.Phone, result.Signature, result.OrganizationID, result.Tags, result.Suspended, result.Role)
	}

	fmt.Println(users)
	return users
}
