package sobjects

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"strings"

	"github.com/grokify/mogo/os/osutil"
)

type ContactSet struct {
	IDSet      IDSet              `xml:"-"`
	Records    []Contact          `json:"records,omitempty" xml:"records"`
	RecordsMap map[string]Contact `xml:"-"`
}

func NewContactSet() ContactSet {
	set := ContactSet{
		IDSet:      NewIDSet(),
		Records:    []Contact{},
		RecordsMap: map[string]Contact{}}
	return set
}

func NewContactSetSetFromXML(bytes []byte) (ContactSet, error) {
	set := ContactSet{IDSet: NewIDSet()}
	err := xml.Unmarshal(bytes, &set)
	set.Inflate()
	return set, err
}

func NewContactSetFromXMLFile(filepath string) (ContactSet, error) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return ContactSet{}, err
	}
	return NewContactSetSetFromXML(bytes)
}

func NewContactSetFromJSONResponse(resp *http.Response) (ContactSet, error) {
	set := NewContactSet()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return set, err
	}
	err = json.Unmarshal(bytes, &set)
	return set, err
}

func (set *ContactSet) ReadJSONFilesFromDir(dir string) error {
	entries, err := osutil.ReadDirMore(dir, regexp.MustCompile(`(?i)\.json$`), false, true, true)
	if err != nil {
		return err
	}
	filepaths := entries.Names(dir)
	for _, filepath := range filepaths {
		contact, err := NewContactFromJSONFile(filepath)
		if err == nil && len(contact.ID) > 0 {
			set.Records = append(set.Records, contact)
		}
	}
	return nil
}

func (set *ContactSet) Inflate() {
	for _, record := range set.Records {
		if len(record.ID) > 0 {
			set.IDSet.AddID(record.ID)
			set.RecordsMap[record.ID] = record
		}
		if len(record.AccountID) > 0 {
			set.IDSet.AddID(record.AccountID)
		}
	}
}

func (set *ContactSet) GetContactByName(name string) (Contact, error) {
	for _, contact := range set.Records {
		if contact.Name == name {
			return contact, nil
		}
	}
	return Contact{}, fmt.Errorf("could not found Contact by name [%v]", name)
}

func (set *ContactSet) GetContactByID(id string) (Contact, error) {
	for _, contact := range set.Records {
		if contact.ID == id {
			return contact, nil
		}
	}
	return Contact{}, fmt.Errorf("could not found Contact by id [%v]", id)
}

type Contact struct {
	ID         string
	AccountID  string
	Department string
	Email      string
	Fax        string
	FirstName  string
	LastName   string
	Name       string
}

func NewContactFromJSON(bytes []byte) (Contact, error) {
	obj := Contact{}
	err := json.Unmarshal(bytes, &obj)
	return obj, err
}

func NewContactFromJSONFile(filepath string) (Contact, error) {
	bytes, err := os.ReadFile(filepath)
	if err != nil {
		return Contact{}, err
	}
	return NewContactFromJSON(bytes)
}

func ContactEmailOrID(contact Contact) string {
	emailOrID := ""
	if len(strings.TrimSpace(contact.Email)) > 0 {
		emailOrID = contact.Email
	} else {
		emailOrID = contact.ID
	}
	return strings.TrimSpace(emailOrID)
}

func ContactsEmailOrID(contacts []Contact) []string {
	emailOrIDs := []string{}
	for _, contact := range contacts {
		emailOrID := ContactEmailOrID(contact)
		if len(emailOrID) > 0 {
			emailOrIDs = append(emailOrIDs, emailOrID)
		}
	}
	return emailOrIDs
}

func ContactsEmailOrIDString(contacts []Contact, sep string) string {
	return strings.Join(ContactsEmailOrID(contacts), sep)
}

func ContactIDOrEmail(contact Contact) string {
	idOrEmail := ""
	if len(strings.TrimSpace(contact.ID)) > 0 {
		idOrEmail = contact.ID
	} else {
		idOrEmail = contact.Email
	}
	return strings.TrimSpace(idOrEmail)
}

func ContactsIDOrEmail(contacts []Contact) []string {
	idOrEmails := []string{}
	for _, contact := range contacts {
		idOrEmail := ContactIDOrEmail(contact)
		if len(idOrEmail) > 0 {
			idOrEmails = append(idOrEmails, idOrEmail)
		}
	}
	return idOrEmails
}

func ContactsIDOrEmailString(contacts []Contact, sep string) string {
	return strings.Join(ContactsIDOrEmail(contacts), sep)
}
