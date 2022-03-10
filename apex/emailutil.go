package apex

import (
	"regexp"
	"strings"

	"github.com/grokify/go-salesforce/sobjects"
	mu "github.com/grokify/mogo/type/maputil"
	"github.com/russross/blackfriday/v2"
)

type EmailPriorityType int

const (
	// Setting ContactPriorityID will enable
	// Salesforce to associate the email with
	// the contact object and is generally
	// preferred/
	ContactPriorityID EmailPriorityType = iota
	ContactPriorityEmail
)

type ApexEmailRequestOpts struct {
	EmailInfos            []ApexEmailInfo
	SubjectTemplate       string
	BodyTemplate          string
	ReplyToEmail          string
	ReplyToName           string
	RecipientPriorityType EmailPriorityType
}

func BuildApexEmail(req ApexEmailRequestOpts) string {
	data := []map[string]string{}
	for _, info := range req.EmailInfos {
		data = append(data, info.ToMap(req.RecipientPriorityType))
	}
	return ApexEmailsSliceTemplate(data,
		req.SubjectTemplate,
		req.BodyTemplate,
		req.ReplyToEmail,
		req.ReplyToName)
}

const (
	To_             = "to_"
	Cc_             = "cc_"
	Bcc_            = "bcc_"
	TargetObjectId_ = "targetobjectid_"
	Sep             = ";"
)

var rxEscapeSingleQuote = regexp.MustCompile(`(^|[^\\])'`)

func EscapeSingleQuote(s string) string {
	return rxEscapeSingleQuote.ReplaceAllString(s, "${1}\\'")
}

func MarkdownToApexEmailHtml(bytes []byte) string {
	return StringToApexStringSimple(string(blackfriday.Run(bytes)))
}

func StringToApexStringSimple(s string) string {
	return strings.Replace(s, "\n", `\n`, -1)
}

type ApexEmailInfo struct {
	To   []sobjects.Contact
	Cc   []sobjects.Contact
	Bcc  []sobjects.Contact
	Data map[string]string
}

func NewApexEmailInfo() ApexEmailInfo {
	return ApexEmailInfo{
		To:   []sobjects.Contact{},
		Cc:   []sobjects.Contact{},
		Bcc:  []sobjects.Contact{},
		Data: map[string]string{}}
}

func mergeContacts(raw string, contacts []sobjects.Contact, emailPriorityType EmailPriorityType, sep string) string {
	emailAddrs := []string{}

	raw = strings.TrimSpace(raw)
	if len(raw) > 0 {
		emailAddrs = append(emailAddrs, strings.Split(raw, sep)...)
	}

	if emailPriorityType == ContactPriorityID {
		emailAddrs = append(emailAddrs, sobjects.ContactsIDOrEmailString(contacts, sep))
	} else {
		emailAddrs = append(emailAddrs, sobjects.ContactsEmailOrIDString(contacts, sep))
	}

	emailAddrsCanonical := []string{}
	emailAddrsSeen := map[string]int{}

	for _, emailAddr := range emailAddrs {
		emailAddr = strings.TrimSpace(emailAddr)
		if len(emailAddrs) == 0 {
			continue
		}
		if _, ok := emailAddrsSeen[emailAddr]; !ok {
			emailAddrsCanonical = append(emailAddrsCanonical, emailAddr)
			emailAddrsSeen[emailAddr] = 1
		}
	}
	return strings.TrimSpace(strings.Join(emailAddrsCanonical, sep))
}

func (email *ApexEmailInfo) ToMap(emailPriorityType EmailPriorityType) map[string]string {
	data := email.Data
	data[To_] = mergeContacts(mu.MapSSValOrEmpty(data, To_), email.To, emailPriorityType, Sep)
	data[Cc_] = mergeContacts(mu.MapSSValOrEmpty(data, Cc_), email.Cc, emailPriorityType, Sep)
	data[Bcc_] = mergeContacts(mu.MapSSValOrEmpty(data, Bcc_), email.Bcc, emailPriorityType, Sep)
	if len(data[To_]) > 0 && !strings.Contains(data[To_], "@") && !strings.Contains(data[To_], Sep) {
		data[TargetObjectId_] = data[To_]
		data[To_] = ""
	}
	return data
}
