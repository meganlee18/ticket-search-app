package app

import (
	"encoding/json"
	"fmt"
	"sort"
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

func readUsersFromFile(path string) []User {
	var obj []User

	data := readFile(path)
	err := json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
	}

	return obj
}

func DisplayUsersBasedOnSearchOptions(path string, searchValue int) map[string]interface{} {
	tickets := readUsersFromFile(path)
	users := returnUsers(tickets, searchValue)

	var outcome map[string]interface{}

	for _, user := range users {
		outcome = map[string]interface{}{
			"_id":             user.ID,
			"active":          user.Active,
			"alias":           user.Alias,
			"created_at":      user.CreatedAt,
			"email":           user.Email,
			"external_id":     user.ExternalId,
			"last_login_at":   user.LastLoginAt,
			"locale":          user.Locale,
			"name":            user.Name,
			"organization_id": user.OrganizationID,
			"phone":           user.Phone,
			"role":            user.Role,
			"shared":          user.Shared,
			"signature":       user.Signature,
			"suspended":       user.Suspended,
			"tags":            user.Tags,
			"timezone":        user.Timezone,
			"url":             user.Url,
			"verified":        user.Verified,
		}
	}

 	return sortMapByKey(outcome)
}

func returnUsers(tickets []User, searchValue int) []User {
	var results []User
	for _, ticket := range tickets {
		if ticket.ID == searchValue {
			results = append(results, ticket)
		}
	}

	return results
}

func sortMapByKey(fields map[string]interface{}) map[string]interface{} {
	sortedMap := make(map[string]interface{})
	keys := make([]string, 0, len(fields))
	for k := range fields {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, key := range keys {
		sortedMap[key] = fields[key]
		fmt.Printf("%s: %v\n", key, fields[key])
	}

	return sortedMap
}
