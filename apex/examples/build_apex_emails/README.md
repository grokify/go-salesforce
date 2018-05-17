# Apex Codegen for Sending Email

Sending email using Salesforce can be useful in the following scenarios:

* You have access to Salesforce
* You want to send from your email address in Salesforce
* You have access to execute Apex code, e.g. via [Developer Workbench](https://workbench.developerforce.com/)

This program allows you to auto-generate Apex code using Go. This project builds Apex code that can send HTML email via Salesforce `Messaging.SingleEmailMessage` and `Messaging.sendEmail(emails)`. It has the following features:

* Will accept simple subject and body templates
* Can convert Markdown body to Salesforce HTML body
* Can prioritize Contact.Id over Contact.Email for sending so messages can be associated with contact object
* Will automatically use `setTargetObjectId` vs. `setToAddresses`