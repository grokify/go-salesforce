package sobjects

import (
	"encoding/json"
	"io/ioutil"
)

type Contact struct {
	Id         string
	AccountId  string
	Department string
	Email      string
	Fax        string
	FirstName  string
	LastName   string
	Name       string
}

func NewContactFromJson(bytes []byte) (Contact, error) {
	obj := Contact{}
	err := json.Unmarshal(bytes, &obj)
	return obj, err
}

func NewContactFromJsonFile(filepath string) (Contact, error) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return Contact{}, err
	}
	return NewContactFromJson(bytes)
}
