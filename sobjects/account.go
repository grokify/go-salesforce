package sobjects

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Account struct {
	Id   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type AccountSet struct {
	IdSet   IdSet     `xml:"-"`
	Records []Account `json:"records,omitempty" xml:"records"`
}

func NewAccountSetFromJSONResponse(resp *http.Response) (AccountSet, error) {
	set := AccountSet{Records: []Account{}}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return set, err
	}
	err = json.Unmarshal(bytes, &set)
	return set, err
}

func (set *AccountSet) GetAccountByName(name string) (Account, error) {
	for _, act := range set.Records {
		if act.Name == name {
			return act, nil
		}
	}
	return Account{}, errors.New(fmt.Sprintf("Could not found Account by name [%v]", name))
}

func (set *AccountSet) GetAccountById(id string) (Account, error) {
	for _, act := range set.Records {
		if act.Id == id {
			return act, nil
		}
	}
	return Account{}, errors.New(fmt.Sprintf("Could not found Account by id [%v]", id))
}
