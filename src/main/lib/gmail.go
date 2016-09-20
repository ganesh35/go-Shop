package lib
import (
    "fmt"
    "net/smtp"
    "net/mail"
)
func SendEmail(
    smtpSettings TypeSmtpSettings, 
    mailSettings TypeMailSettings, 
    toEmail string, 
    subject string, 
    body string,
) (err error) {
	defer CatchPanic(&err, "sendEmail")
	// authentication configuration
	
	emailauth := smtp.PlainAuth("", smtpSettings.Username, smtpSettings.Password, smtpSettings.Host)
	
    sender :=  mailSettings.Sender_email// 
 	receivers := []string{
 		toEmail,
 	}

	from := mail.Address{mailSettings.Sender_name, mailSettings.Sender_email}
    to   := mail.Address{"", toEmail}
    
    // Setup headers
    headers := make(map[string]string)
    headers["From"] = from.String()
    headers["To"] = to.String()
    headers["Subject"] = subject
    headers["MIME-version"] = "1.0"
    headers["Content-Type"] = `text/html; charset="UTF-8"`

    // Setup message
    message := ""
    for k,v := range headers {
        message += fmt.Sprintf("%s: %s\r\n", k, v)
    }
    message += "\r\n" + body



    // send out the email
    err = smtp.SendMail( fmt.Sprintf("%s:%d", smtpSettings.Host, smtpSettings.Port), //convert port number from int to string
    	emailauth,
        sender,
        receivers,
        []byte(message),
    )
    return err 
}
