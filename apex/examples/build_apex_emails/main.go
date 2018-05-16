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

	emailsData := []apex.ApexEmailInfo{{
		To:  []sobjects.Contact{{Email: "alice@example.com"}, {Email: "bob@example.com"}},
		Cc:  []sobjects.Contact{{Email: "carol@example.com"}, {Email: "dan@example.com"}},
		Bcc: []sobjects.Contact{{Email: "erin@example.com"}, {Email: "frank@example.com"}},
		Data: map[string]string{
			"CODE_URL":  "https://github.com/grokify/go-salesforce/tree/master/apex",
			"FROM_NAME": "grokify"}}}

	bodyBytesMd, err := ioutil.ReadFile(inputBodyFile)
	if err != nil {
		log.Fatal(err)
	}

	req := apex.BuildApexEmailRequest{
		EmailInfos:            emailsData,
		SubjectTemplate:       "My Demo Subject",
		BodyTemplate:          apex.MarkdownToApexEmailHtml(bodyBytesMd),
		ReplyToEmail:          "sender@example.com",
		ReplyToName:           "Example Sender User",
		RecipientPriorityType: apex.ContactEmailPriority}

	apexCode := apex.BuildApexEmail(req)

	err = ioutil.WriteFile(outputApexFile, []byte(apexCode), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("DONE")
}
