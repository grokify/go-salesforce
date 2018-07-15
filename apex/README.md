# Apex Codegen for Sending Email

Salesforce supports sending email from its system. You can do this programmatically or via the UI.

Some benefits of using Apex to send email with Salesforce:

* You don't need to set up a separate MX record as Salesforce can send as your user's email address.
* You can customize email by programmatically creating email
* You can send email via [Developer Workbench](https://workbench.developerforce.com/)
* Emails can be automatically attached to Salesforce Contacts
* Emails do not count towards quota if you use TargetObjectId

This program allows you to auto-generate Apex code using Go. This project builds Apex code that can send HTML email via Salesforce `Messaging.SingleEmailMessage` and `Messaging.sendEmail(emails)`. It has the following features:

* Will accept simple subject and body templates
* Can convert Markdown body to Salesforce HTML body
* Can prioritize Contact.Id over Contact.Email for sending so messages can be associated with contact object
* Automatically use `setTargetObjectId` vs. `setToAddresses`
* Support `Contact.Id` or `Contact.Email` preference for sending to (a) attach to contact object and for (b) testing.

## Example Useful SOQL

SELECT Id, Email, SenderEmail, FirstName, LastName, Name, UserType FROM User WHERE Email = 'myemail'

vyshakh.babji@ringcentral.com