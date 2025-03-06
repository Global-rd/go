package parsing

import "encoding/json"

type JsonData struct {
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	UserName     string `json:"userName"`
	PhoneNumber  string `json:"phoneNumber"`
	EmailAddress string `json:"emailAddress"`
}

func parseJsonData(data string) ([]JsonData, error) {
	var jsonData []JsonData
	err := json.Unmarshal([]byte(data), &jsonData)
	if err != nil {
		return nil, err
	}
	return jsonData, nil
}
