package apex

import (
	"regexp"
	"strings"

	"github.com/grokify/go-salesforce/sobjects"
	"gopkg.in/russross/blackfriday.v2"
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

func ContactIdOrEmail(contact sobjects.Contact) string {
	if len(strings.TrimSpace(contact.Id)) > 0 {
		return strings.TrimSpace(contact.Id)
	}
	return strings.TrimSpace(contact.Email)
}

func ContactsIdOrEmail(contacts []sobjects.Contact) []string {
	idOrEmails := []string{}
	for _, contact := range contacts {
		idOrEmail := ContactIdOrEmail(contact)
		if len(idOrEmail) > 0 {
			idOrEmails = append(idOrEmails, idOrEmail)
		}
	}
	return idOrEmails
}

func ContactsIdOrEmailString(contacts []sobjects.Contact) string {
	return strings.Join(ContactsIdOrEmail(contacts), ";")
}
