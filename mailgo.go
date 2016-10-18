/* Package mailgo is a collection or routines to use to mail message to BlackListMe users. */
// currently using SendGrid's Go Library, but could use another vendor
// https://github.com/sendgrid/sendgrid-go

package mailgo

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"os"
)

//type mailer interface {
//	send() int
//}

var blmname = "BlackListMeAdmin"
var blmemail = "gladwig@gladworx.com"

type Session struct  {
	Fname string
	Lname string
	Email string
	URL string
	subject string
	message string
}

func NewSession(email, firstname, lastname, URL string) *Session  {
    sess := Session{
	    Fname:firstname,
	    Lname:lastname,
	    URL:URL,
	    Email:email}
    return &sess
}

func CompleteBlacklist(ses *Session)(error) {
	ses.subject = "BlackListMe Blacklist Confirmation"
	ses.message = fmt.Sprintf(`Welcome to BlackListMe 
	        This email is to confirm the request to add the email address %s to the BlackListMe blacklist.
		Click this URL to confirm blacklisting this email address %s 
		---------------------------------------------------------------------------------- 
		You are receiving this email because you recently requested an email address be blacklisted 
		at BlackListMe.net. If this wasn't you, please ignore this email` , ses.Email, ses.URL)
	return(sendreq(ses))
}

func ConfirmRegistration(ses *Session)(error) {
	ses.subject = "BlackListMe Registration Confirmation"
	ses.message = fmt.Sprintf(`Welcome to BlackListMe 
                Click this URL to confirm registration. %s 
		----------------------------------------------------------------------------------
		You are receiving this email because you recently created a new BlackListMe account 
		or added a new email address. If this wasn't you, please ignore this email. `, ses.URL)
	return(sendreq(ses))
}

func ConfirmEmailChange( ses *Session, oldaddress string) (error){
	ses.subject = "BlackListMe Email address change Confirmation"
	ses.message = fmt.Sprintf(`Greetings from BlackListMe 
        Click this URL to confirm an email address change from %s to %s 
        %s  
	----------------------------------------------------------------------------------
	You are receiving this email because you requested an email address change on your BlackListMe.net account.
	If this wasn't you, please ignore this email and the action will not be completed. `, 
		oldaddress, ses.Email, ses.URL)
	return(sendreq(ses))
}

func ConfirmPasswordChange( ses *Session) (error){
	ses.subject = "BlackListMe Password Change Confirmation"
	ses.message = fmt.Sprintf(`Greetings from BlackListMe 
        Click this URL to confirm a password change %s
        ----------------------------------------------------------------------------------
	You are receiving this email because you requested a password change on your BlackListMe.net account.
	If this wasn't you, please ignore this email and the action will not be completed.`, ses.URL)

	return(sendreq(ses))
}

func ConfirmEmailAddressRegistration( ses *Session) (error){
	ses.subject = "BlackListMe Email Registration Confirmation"
	ses.message = fmt.Sprintf(`Welcome to BlackListMe
        Click this URL to confirm registration. %s 
	---------------------------------------------------------------------------------- 
	You are receiving this email because you registered with BlackListMe.org 
	If this wasn't you, please ignore this email and the action will not be completed.`, ses.URL)
	return(sendreq(ses))
}

func ConfirmEmailAddressDeRegistration( ses *Session) (error){
	ses.subject = "BlackListMe Email Deregistration Confirmation"
	ses.message = fmt.Sprintf(`Greetings from  BlackListMe 
        Click this URL to confirm an email deregistration. %s
	----------------------------------------------------------------------------------
	You are receiving this email because you requested an email be deregistered at BlackListMe.net
	If this wasn't you, please ignore this email, and the action will not be completed `, ses.URL)
	return(sendreq(ses))
}

func NoticePasswordHasChanged( ses *Session) (error){
	ses.subject = "BlackListMe Password Change Confirmation"
	ses.message = fmt.Sprintf(`Greetings from  BlackListMe 
        Click this URL to confirm a password change %s
	----------------------------------------------------------------------------------
	You are receiving this email because you requested a password change at BlackListMe.net
	If this wasn't you, please ignore this email, and the action will not be completed `, ses.URL)
	return(sendreq(ses))
}

func sendreq(ses *Session)(error) {
	from := mail.NewEmail(blmname, blmemail)
	subject := ses.subject
	to := mail.NewEmail(ses.Fname+ses.Lname, ses.Email)
	content := mail.NewContent("text/plain", ses.message)
	m := mail.NewV3MailInit(from, subject, to, content)

	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", 
		"https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	response, err := sendgrid.API(request)
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
	} else {
		fmt.Println("success")
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
	error := err
	return(error)
}

func init() {}
