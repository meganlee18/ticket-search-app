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

func DisplayTickets() []Ticket {
	var obj []Ticket

	data := ReadFile("./tickets/tickets.json")
	err := json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
	}

	return obj
}

func DisplayUsers() []User {
	var obj []User

	data := ReadFile("./tickets/users.json")

	err := json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(obj)
	return obj
}

func DisplayOrganizations() []Organizations {
	var obj []Organizations

	data := ReadFile("./tickets/organizations.json")
	err := json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
	}

	return obj
}

func ReadFile(path string) []byte {
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


