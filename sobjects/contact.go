package sobjects

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"

	"github.com/grokify/simplego/os/osutil"
)

type ContactSet struct {
	IdSet      IdSet              `xml:"-"`
	Records    []Contact          `json:"records,omitempty" xml:"records"`
	RecordsMap map[string]Contact `xml:"-"`
}

func NewContactSet() ContactSet {
	set := ContactSet{
		IdSet:      NewIdSet(),
		Records:    []Contact{},
		RecordsMap: map[string]Contact{}}
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

func NewContactSetFromJSONResponse(resp *http.Response) (ContactSet, error) {
	set := NewContactSet()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return set, err
	}
	err = json.Unmarshal(bytes, &set)
	return set, err
}

func (set *ContactSet) ReadJsonFilesFromDir(dir string) error {
	entries, err := osutil.ReadDirMore(dir, regexp.MustCompile(`(?i)\.json$`), false, true, true)
	if err != nil {
		return err
	}
	filepaths := osutil.DirEntrySlice(entries).Names(dir, true)
	for _, filepath := range filepaths {
		contact, err := NewContactFromJsonFile(filepath)
		if err == nil && len(contact.Id) > 0 {
			set.Records = append(set.Records, contact)
		}
	}
	return nil
}

func (set *ContactSet) Inflate() {
	for _, record := range set.Records {
		if len(record.Id) > 0 {
			set.IdSet.AddId(record.Id)
			set.RecordsMap[record.Id] = record
		}
		if len(record.AccountId) > 0 {
			set.IdSet.AddId(record.AccountId)
		}
	}
}

func (set *ContactSet) GetContactByName(name string) (Contact, error) {
	for _, contact := range set.Records {
		if contact.Name == name {
			return contact, nil
		}
	}
	return Contact{}, errors.New(fmt.Sprintf("Could not found Contact by name [%v]", name))
}

func (set *ContactSet) GetContactById(id string) (Contact, error) {
	for _, contact := range set.Records {
		if contact.Id == id {
			return contact, nil
		}
	}
	return Contact{}, errors.New(fmt.Sprintf("Could not found Contact by id [%v]", id))
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

func ContactEmailOrId(contact Contact) string {
	emailOrId := ""
	if len(strings.TrimSpace(contact.Email)) > 0 {
		emailOrId = contact.Email
	} else {
		emailOrId = contact.Id
	}
	return strings.TrimSpace(emailOrId)
}

func ContactsEmailOrId(contacts []Contact) []string {
	emailOrIds := []string{}
	for _, contact := range contacts {
		emailOrId := ContactEmailOrId(contact)
		if len(emailOrId) > 0 {
			emailOrIds = append(emailOrIds, emailOrId)
		}
	}
	return emailOrIds
}

func ContactsEmailOrIdString(contacts []Contact, sep string) string {
	return strings.Join(ContactsEmailOrId(contacts), sep)
}

func ContactIdOrEmail(contact Contact) string {
	idOrEmail := ""
	if len(strings.TrimSpace(contact.Id)) > 0 {
		idOrEmail = contact.Id
	} else {
		idOrEmail = contact.Email
	}
	return strings.TrimSpace(idOrEmail)
}

func ContactsIdOrEmail(contacts []Contact) []string {
	idOrEmails := []string{}
	for _, contact := range contacts {
		idOrEmail := ContactIdOrEmail(contact)
		if len(idOrEmail) > 0 {
			idOrEmails = append(idOrEmails, idOrEmail)
		}
	}
	return idOrEmails
}

func ContactsIdOrEmailString(contacts []Contact, sep string) string {
	return strings.Join(ContactsIdOrEmail(contacts), sep)
}
