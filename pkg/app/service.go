package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

type Ticket struct {
	ID             string   `json:"_id"`
	Url            string   `json:"url"`
	ExternalId     string   `json:"external_id"`
	CreatedAt      string   `json:"created_at"`
	Type           string   `json:"type"`
	Subject        string   `json:"subject"`
	Description    string   `json:"description"`
	Priority       string   `json:"priority"`
	Status         string   `json:"status"`
	SubmitterID    int      `json:"submitter_id"`
	AssigneeID     int      `json:"assignee_id"`
	OrganizationId int      `json:"organization_id"`
	Tags           []string `json:"tags"`
	HasIncidents   bool     `json:"has_incidents"`
	DueAt          string   `json:"due_at"`
	Via            string   `json:"via"`
}

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

type Organizations struct {
	ID            int      `json:"_id"`
	Url           string   `json:"url"`
	ExternalId    string   `json:"external_id"`
	Name          string   `json:"name"`
	DomainNames   []string `json:"domain_names"`
	CreatedAt     string   `json:"created_at"`
	Details       string   `json:"details"`
	SharedTickets bool     `json:"shared_tickets"`
	Tags          []string `json:"tags"`
}

func DisplayUsers(path string) []User {
	var obj []User

	data := readFile(path)
	err := json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
	}

	return obj
}

func DisplayUsersBasedOnSearchOptions(path string, searchValue int) string {
	tickets := DisplayUsers(path)
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

func DisplayTicketFields(path string) {
	var result []map[string]interface{}

	data := readFile(path)
	unmarshalledResult, err := unmarshalData(data, result)
	if err != nil {
		fmt.Println("error: ", err)
	}

	sortedFields := displaySortedFields(unmarshalledResult)
	fmt.Println(strings.Join(removeDuplicateValues(sortedFields),  "\n"))
}

func readFile(path string) []byte {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return data
}

func unmarshalData(data []byte, result []map[string]interface{}) ([]map[string]interface{}, error) {
	err := json.Unmarshal(data, &result)
	if err != nil {
		fmt.Println("Error: ", err)
		return nil, err
	}

	return result, err
}

func displaySortedFields(result []map[string]interface{}) []string {
	var fields []string

	for _, v := range result {
		for k := range v {
			fields = append(fields, k)
		}
	}

	sort.Strings(fields)
	return fields
}

func removeDuplicateValues(fields []string) []string {
	keys := make(map[string]bool)
	var list []string

	for _, entry := range fields {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}

	return list
}


