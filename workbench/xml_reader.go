package workbench

import (
	"encoding/xml"
	"os"
)

type Account struct {
	ID   string `xml:"Id"`
	Name string `xml:"Name"`
	Type string `xml:"type"`
}

type AccountsSet struct {
	Accounts []Account `xml:"records"`
}

func ReadAccountsXML(filepath string) (AccountsSet, error) {
	qres := AccountsSet{}

	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return AccountsSet{}, err
	}

	err = xml.Unmarshal(bytes, &qres)
	return qres, err
}
