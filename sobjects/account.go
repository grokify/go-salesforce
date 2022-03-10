package sobjects

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Account struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

type AccountSet struct {
	IDSet   IDSet     `xml:"-"`
	Records []Account `json:"records,omitempty" xml:"records"`
}

func NewAccountSetFromJSONResponse(resp *http.Response) (AccountSet, error) {
	set := AccountSet{Records: []Account{}}

	bytes, err := io.ReadAll(resp.Body)
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
	return Account{}, fmt.Errorf("could not found Account by name [%v]", name)
}

func (set *AccountSet) GetAccountByID(id string) (Account, error) {
	for _, act := range set.Records {
		if act.ID == id {
			return act, nil
		}
	}
	return Account{}, fmt.Errorf("could not found Account by id [%v]", id)
}
