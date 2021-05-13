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

func findUsers(path string) []User {
	var obj []User

	data := readFile(path)
	err := json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
	}

	return obj
}

func DisplayUsersBasedOnSearchOptions(path string, searchValue int) map[string]interface{} {
	tickets := findUsers(path)
	var results []User
	var outcome map[string]interface{}

	for _, ticket := range tickets {
		if ticket.ID == searchValue {
			results = append(results, ticket)
		}
	}

	for _, result := range results {
		outcome = map[string]interface{}{
			"_id":             result.ID,
			"active":          result.Active,
			"alias":           result.Alias,
			"created_at":      result.CreatedAt,
			"email":           result.Email,
			"external_id":     result.ExternalId,
			"last_login_at":   result.LastLoginAt,
			"locale":          result.Locale,
			"name":            result.Name,
			"organization_id": result.OrganizationID,
			"phone":           result.Phone,
			"role":            result.Role,
			"shared":          result.Shared,
			"signature":       result.Signature,
			"suspended":       result.Suspended,
			"tags":            result.Tags,
			"timezone":        result.Timezone,
			"url":             result.Url,
			"verified":        result.Verified,
		}
	}

 	return sortMapByKey(outcome)
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
