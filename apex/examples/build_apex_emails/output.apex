
// BEGIN auto-generated Apex code (https://github.com/grokify/go-salesforce/tree/master/apex)
String emailSubjectTmpl = 'My Demo Subject';
String emailBodyTmpl = '<p>This is a test email template showing basic variable substitution.</p>\n\n<ul>\n<li>Code: {{CODE_URL}}</li>\n</ul>\n\n<p>Thanks,</p>\n\n<p>{{FROM_NAME}}</p>\n';
String replyToEmail = 'sender@example.com';
String senderDisplayName = 'Example Sender User';

List<Map<String,String>> emailsData = new List<Map<String,String>>{new Map<String, String> {'bcc_' => 'erin@example.com;frank@example.com','CODE_URL' => 'https://github.com/grokify/go-salesforce/apex','FROM_NAME' => 'grokify','to_' => 'alice@example.com;bob@example.com','cc_' => 'carol@example.com;dan@example.com'}};
List<Messaging.SingleEmailMessage> emails = new List<Messaging.SingleEmailMessage>();

for (Map<String,String> emailData : emailsData) {
  Messaging.SingleEmailMessage mail = new Messaging.SingleEmailMessage();

  String emailSubject = emailSubjectTmpl;
  String emailBody = emailBodyTmpl;

  for (String emailTmplKey : emailData.keySet()) {
    Integer lastChar = emailTmplKey.charAt(emailTmplKey.length()-1);

    if (lastChar != 95) {
      String emailTmplVal = emailData.get(emailTmplKey);
      emailSubject = emailSubject.replace('{{'+emailTmplKey+'}}', emailTmplVal);
      emailBody = emailBody.replace('{{'+emailTmplKey+'}}', emailTmplVal);
    }
  }

  Messaging.SingleEmailMessage email = new Messaging.SingleEmailMessage();

  String sendTo = emailData.get('to_');
  if (string.isNotBlank(sendTo)) {
    email.setToAddresses(sendTo.split(';'));
    String sendCc = emailData.get('cc_');
    if (string.isNotBlank(sendCc)) {
      email.setCcAddresses(sendCc.split(';'));
    }
    String sendBcc = emailData.get('bcc_');
    if (string.isNotBlank(sendBcc)) {
      email.setBccAddresses(sendBcc.split(';'));
    }
    if (string.isNotBlank(replyToEmail)) {
      email.setReplyTo(replyToEmail);
    }
    if (string.isNotBlank(senderDisplayName)) {
      email.setSenderDisplayName(senderDisplayName);    
    }
    email.setSubject(emailSubject);
    email.setHtmlBody(emailBody);
    emails.add(email);
  }
}

if (emails.size()>0) {
  Messaging.sendEmail(emails);
}
// END auto-generated Apex code (https://github.com/grokify/go-salesforce/tree/master/apex)
