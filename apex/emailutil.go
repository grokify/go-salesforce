package apex

import (
	"regexp"
	"strings"

	"github.com/grokify/go-salesforce/sobjects"
	"gopkg.in/russross/blackfriday.v2"
)

type EmailPriorityType int

const (
	ContactIdPriority EmailPriorityType = iota
	ContactEmailPriority
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

func (email *ApexEmailInfo) ToMap(emailPriorityType EmailPriorityType) map[string]string {
	data := email.Data
	sep := ";"
	if len(email.To) > 0 {
		if emailPriorityType == ContactIdPriority {
			data["to_"] = sobjects.ContactsIdOrEmailString(email.To, sep)
		} else {
			data["to_"] = sobjects.ContactsEmailOrIdString(email.To, sep)
		}
	}
	if len(email.Cc) > 0 {
		if emailPriorityType == ContactIdPriority {
			data["cc_"] = sobjects.ContactsIdOrEmailString(email.Cc, sep)
		} else {
			data["cc_"] = sobjects.ContactsEmailOrIdString(email.Cc, sep)
		}
	}
	if len(email.Bcc) > 0 {
		if emailPriorityType == ContactIdPriority {
			data["bcc_"] = sobjects.ContactsIdOrEmailString(email.Bcc, sep)
		} else {
			data["bcc_"] = sobjects.ContactsEmailOrIdString(email.Bcc, sep)
		}
	}
	return data
}

type BuildApexEmailRequest struct {
	EmailInfos            []ApexEmailInfo
	SubjectTemplate       string
	BodyTemplate          string
	ReplyToEmail          string
	ReplyToName           string
	RecipientPriorityType EmailPriorityType
}

func BuildApexEmail(req BuildApexEmailRequest) string {
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
