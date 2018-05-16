package main

import (
	"fmt"
	"io/ioutil"

	log "github.com/sirupsen/logrus"

	"github.com/grokify/go-salesforce/apex"
	"github.com/grokify/go-salesforce/sobjects"
)

func main() {
	inputBodyFile := "input.md"
	outputApexFile := "output.apex"

	bodyBytesMd, err := ioutil.ReadFile(inputBodyFile)
	if err != nil {
		log.Fatal(err)
	}

	bodyTmpl := apex.MarkdownToApexEmailHtml(bodyBytesMd)

	to := []sobjects.Contact{{Email: "alice@example.com"}, {Email: "bob@example.com"}}
	cc := []sobjects.Contact{{Email: "carol@example.com"}, {Email: "dan@example.com"}}
	bcc := []sobjects.Contact{{Email: "erin@example.com"}, {Email: "frank@example.com"}}
	sep := ";"

	emailsData := []map[string]string{{
		"to_":       sobjects.ContactsIdOrEmailString(to, sep),
		"cc_":       sobjects.ContactsIdOrEmailString(cc, sep),
		"bcc_":      sobjects.ContactsIdOrEmailString(bcc, sep),
		"CODE_URL":  "https://github.com/grokify/go-salesforce/apex",
		"FROM_NAME": "grokify"}}

	subjectTmpl := "My Demo Subject"

	apexCode := apex.ApexEmailsSliceTemplate(
		emailsData, subjectTmpl, bodyTmpl,
		"sender@example.com", "Example Sender User")

	err = ioutil.WriteFile(outputApexFile, []byte(apexCode), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DONE")
}
