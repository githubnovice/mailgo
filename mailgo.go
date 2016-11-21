/* Package mailgo is a collection or routines to use to mail message to BlackListMe users. */
// currently using SendGrid's Go Library, but could use another vendor
// https://github.com/sendgrid/sendgrid-go

package mailgo

import (
	"fmt"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	//"github.com/goinggo/tracelog"
	"github.com/golang/glog"
//	"github.com/mailgo/nongit.go"
	"os"
	"flag"
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

// Non-account based actions

//ConfirmEmailAddressBlacklist sends an email to confirm a non-account based user blacklist request
func ConfirmEmailAddressBlacklist(ses *Session)(error) {
	ses.subject = "BlackListMe blacklist an email address confirmation"
	ses.message = fmt.Sprintf(`Welcome to BlackListMe 
        This email is to confirm the request to include %s in the BlackListMe blacklist.
        Click this URL to confirm %s 
 
        Regards,
 
          The BlackListMe Team
        ---------------------------------------------------------------------------------- 
        You are receiving this email because you recently requested an email address be blacklisted 
        at BlackListMe.net. If this wasn't you, please ignore this email` , ses.Email, ses.URL)
	return(sendreq(ses))
}

//ConfirmEmailAddressUnBlacklist sends an email to confirm a non-account based user blacklist request
func ConfirmEmailAddressUnBlacklist(ses *Session)(error) {
	ses.subject = "Remove an email address from BlackListMe blacklist confirmation"
	ses.message = fmt.Sprintf(`Greetings from BlackListMe 
        This email is to confirm the request to remove %s from the BlackListMe blacklist.
        Click this URL to confirm  %s 
 
        Regards,
 
           The BlackListMe Team
        --------------------------------------------------------------------------------
        You are receiving this email because you recently requested an email address removal from the
        BlackListMe.net blacklist. If this wasn't you, please ignore this email. No action will be taken.`, 
		ses.Email, ses.URL)
	return(sendreq(ses))
}

// Account based actions

// ConfirmRegistration sends an email to an account based user when they sign up
func ConfirmRegistration(ses *Session)(error) {
	ses.subject = "BlackListMe registration confirmation"
	ses.message = fmt.Sprintf(`Welcome to BlackListMe 
        Click this URL to confirm registration. %s 
 
        Regards,
 
           The BlackListMe Team

        --------------------------------------------------------------------------------
        You are receiving this email because you recently created a new BlackListMe account 
        or added a new email address. If this wasn't you, please ignore this email. 
        The account registration will not complete.`, 
		ses.URL)
	return(sendreq(ses))
}

// ConfirmEmailChangeAddress sends an email to an account based user when they change an account email address
// this will check to make sure they have the address
func ConfirmEmailChangeAddress(ses *Session) (error){
	ses.subject = "BlackListMe account email address verification"
	ses.message = fmt.Sprintf(`Greetings from BlackListMe 
        We need to verify your mail to finish updating your email.
        Click this URL to confirm an email address %s  
 
        Regards,
 
          The BlackListMe Team
        --------------------------------------------------------------------------------
        You are receiving this email because you requested an email address change on your BlackListMe.net account.
        If this wasn't you, please ignore this email and the action will not be completed. `, 
		ses.Email, ses.URL)
	return(sendreq(ses))
}

// NotifyEmailChange sends an email to an account based user when they have changed an email address
// this is sent once the address has been verified with ConfirmEmailChangeVerify Address
func NotifyEmailAddressChange(ses *Session, oldaddress string) (error){
	ses.subject = "BlackListMe account mail address change notification"
	ses.message = fmt.Sprintf(`Greetings from BlackListMe 
        Your email address has been changed 
        from %s 
        to %s 
 
        Regards,
 
            The BlackListMe Team`,
		oldaddress, ses.Email)
	sendreq(ses)
	sesc := ses;
	sesc.Email = oldaddress
	return(sendreq(sesc))
}

// NotifyPasswordChange sends an email to an account based user when they have changed a password
func NotifyPasswordChange( ses *Session) (error){
	ses.subject = "BlackListMe account password change notification"
	ses.message = fmt.Sprintf(`Greetings from BlackListMe 
        Your BlackListMe account password has been recently changed.
 
        Regards,
 
           The BlackListMe Team`)
	return(sendreq(ses))
}

func sendreq(ses *Session)(error) {
	from := mail.NewEmail(blmname, blmemail)
	subject := ses.subject
	to := mail.NewEmail(ses.Fname+ses.Lname, ses.Email)
	content := mail.NewContent("text/plain", ses.message)
	m := mail.NewV3MailInit(from, subject, to, content)

//	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", 
//		"https://api.sendgrid.com")
	request := sendgrid.GetRequest(sendgrid_api_key, "/v3/mail/send", 
		"https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)

	messageo := fmt.Sprintf("Method = %s \n", request.Method)
	messageo += fmt.Sprintf("BaseURL = %s \n", request.BaseURL)
	messageo += fmt.Sprintf("Headers = %s \n", request.Headers)
	messageo += fmt.Sprintf("QueryParams = %s \n", request.QueryParams)
	messageo += fmt.Sprintf("Body = %s \n", request.Body)
	glog.Infof("preparing to send the following message to sendgrid %s \n", messageo)
	glog.Flush()

	response, err := sendgrid.API(request)
	if err != nil {
		message := fmt.Sprintf("Error in sending mail to sendgrid, err = %d", err)
		glog.Error(message)
		glog.Flush()
	} else {
		message := fmt.Sprintf("err = %d \n", err)
		message += fmt.Sprintf("StatusCode = %d \n", response.StatusCode)
		message += fmt.Sprintf("Body = %s \n", response.Body)
		message += fmt.Sprintf("Headers = %s \n", response.Headers)
		glog.Info("Success in sending mail to sendgrid %s \n", message)
		glog.Flush()
	}
	error := err
	return(error)
}

func init() {
	flag.Usage = usage
	flag.Parse()  // must be called for command line flags to work
}

func usage() {
	fmt.Fprintf(os.Stderr, "usage: example -stderrthreshold=[INFO|WARN|FATAL] -log_dir=[string]\n", )
	flag.PrintDefaults()
	os.Exit(2)
}
