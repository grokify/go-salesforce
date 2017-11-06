package workbench

import (
	"encoding/xml"
	"io/ioutil"
)

type Account struct {
	Id   string `xml:"Id"`
	Name string `xml:"Name"`
	Type string `xml:"type"`
}

type AccountsSet struct {
	Accounts []Account `xml:"records"`
}

func ReadAccountsXML(filepath string) (AccountsSet, error) {
	qres := AccountsSet{}

	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return AccountsSet{}, err
	}

	err = xml.Unmarshal(bytes, &qres)
	return qres, err
}
