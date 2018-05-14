package main

import (
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"

	"github.com/grokify/go-salesforce/apex"
	"github.com/grokify/go-salesforce/sobjects"
)

func main() {
	bodyFile := "email.md"

	bodyBytesMd, err := ioutil.ReadFile(bodyFile)
	if err != nil {
		log.Fatal(err)
	}

	bodyTmpl := apex.MarkdownToApexEmailHtml(bodyBytesMd)
	fmt.Println(bodyTmpl)

	to := []sobjects.Contact{
		{Email: "alice@example.com"}, {Email: "bob@example.com"}}
	cc := []sobjects.Contact{
		{Email: "carol@example.com"}, {Email: "dan@example.com"}}
	bcc := []sobjects.Contact{
		{Email: "erin@example.com"}, {Email: "frank@example.com"}}

	email := map[string]string{
		"to_":       apex.ContactsIdOrEmailString(to),
		"cc_":       apex.ContactsIdOrEmailString(cc),
		"bcc_":      apex.ContactsIdOrEmailString(bcc),
		"CODE_URL":  "https://github.com/grokify/go-salesforce/apex",
		"FROM_NAME": "grokify"}

	msmss := map[string]map[string]string{"first": email}

	subjectTmpl := "My Demo Subject"

	apexCode := apex.ApexEmailsTemplate(
		msmss, subjectTmpl, bodyTmpl,
		"sender@example.com", "Example Sender User")

	fmt.Println(apexCode)

	fmt.Println("DONE")
}
