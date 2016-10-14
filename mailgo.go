// using SendGrid's Go Library
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

//type session struct  {
//	fname string
//	lname string
//	email string
//}

var blmname = "BlackListMeAdmin"
var blmemail = "gladwig@gladworx.com"


func Confirm( emailaddress, firstname, lastname, url string) {
	from := mail.NewEmail(blmname, blmemail)
	fmt.Printf("(%v, %T)\n", from, from)
	subject := "BlackListMe Confirmation"
	to := mail.NewEmail(firstname + lastname, emailaddress)
	message := fmt.Sprintf("Welcome to BlackListMe \n Click this URL to confirm blacklisting this email address %s",url)
	content := mail.NewContent("text/plain", message)
	m := mail.NewV3MailInit(from, subject, to, content)
	fmt.Printf("(%v, %T)\n", m, m)

	request := sendgrid.GetRequest(os.Getenv("SENDGRID_API_KEY"), "/v3/mail/send", "https://api.sendgrid.com")
	request.Method = "POST"
	request.Body = mail.GetRequestBody(m)
	response, err := sendgrid.API(request)
	err = nil
	if err != nil {
		fmt.Println("error")
		fmt.Println(err)
	} else {
		fmt.Println("success")
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}

//func describe(i I) {
//	fmt.Printf("(%v, %T)\n", i, i)
//}
