package sobjects

import (
	"encoding/json"
	"encoding/xml"
	"io/ioutil"
	"path"
	"regexp"

	"github.com/grokify/gotilla/io/ioutilmore"
)

type ContactSet struct {
	IdSet   IdSet
	Records []Contact `xml:"records"`
}

func NewContactSet() ContactSet {
	set := ContactSet{IdSet: NewIdSet(), Records: []Contact{}}
	return set
}

func NewContactSetSetFromXml(bytes []byte) (ContactSet, error) {
	set := ContactSet{IdSet: NewIdSet()}
	err := xml.Unmarshal(bytes, &set)
	set.Inflate()
	return set, err
}

func NewContactSetFromXmlFile(filepath string) (ContactSet, error) {
	bytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return ContactSet{}, err
	}
	return NewContactSetSetFromXml(bytes)
}

func (set *ContactSet) ReadJsonFilesFromDir(dir string) error {
	files, err := ioutilmore.DirEntriesReSizeGt0(dir, regexp.MustCompile(`(?i)\.json$`))
	if err != nil {
		return err
	}
	for _, fi := range files {
		filepath := path.Join(dir, fi.Name())
		contact, err := NewContactFromJsonFile(filepath)
		if err == nil && len(contact.Id) > 0 {
			set.Records = append(set.Records, contact)
		}
	}
	return nil
}

func (set *ContactSet) Inflate() {
	for _, record := range set.Records {
		set.IdSet.AddId(record.Id)
		set.IdSet.AddId(record.AccountId)
	}
}

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
