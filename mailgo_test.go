package mailgo

import (
	//"github.com/githubnovice/mailgo"
	"fmt"
	"testing"
)

func TestConfirmEmailAddressBlacklist(t *testing.T) {
	s := Session {Email: "gladwig@gladworx.com", Fname: "Frank", Lname:"Ziffle", URL:"http://BlackListme/Confirm"}
	//fmt.Printf("(%v, %T)\n", s, s)
	err := ConfirmEmailAddressBlacklist(&s)
	if err != nil {
		t.Errorf("ConfirmEmailAddressBlacklist Failed\n")
	}
}

func TestConfirmEmailAddressUnBlacklist(t *testing.T) {
	s := Session {Email: "gladwig@gladworx.com", Fname: "Frank", Lname:"Ziffle", URL:"http://BlackListme/Confirm"}
	//fmt.Printf("(%v, %T)\n", s, s)
	err := ConfirmEmailAddressUnBlacklist(&s)
	if err != nil {
		t.Errorf("ConfirmEmailUnAddressBlacklist Failed\n")
	}
}

func TestConfirmRegistration(t *testing.T) {
	s := Session {Email: "gladwig@gladworx.com", Fname: "Frank", Lname:"Ziffle", URL:"http://BlackListme/Confirm"}
	err := ConfirmRegistration(&s)
	if err != nil {
		t.Errorf("Registration Confirmation Failed\n")
	}
}

func TestConfirmEmailChangeAddress(t *testing.T) {
	s := Session {Email: "gladwig@gladworx.com", Fname: "Frank", Lname:"Ziffle", URL:"http://BlackListme/Confirm"}
	err := ConfirmEmailChangeAddress(&s)
	if err != nil {
		t.Errorf("Email Change Address Confirmation Failed\n")
	}
}

func TestNotifyEmailAddressChange(t *testing.T) {
	s := Session {Email: "gladwig@gladworx.com", Fname: "Frank", Lname:"Ziffle", URL:"http://BlackListme/Confirm"}
	err := NotifyEmailAddressChange(&s, "geoff_ladwig@comcast.net")
	if err != nil {
		t.Errorf("Email Change Address Confirmation Failed\n")
	}
}

func TestNotifyPasswordChange(t *testing.T) {
	s := Session {Email: "gladwig@gladworx.com", Fname: "Frank", Lname:"Ziffle", URL:"http://BlackListme/Confirm"}
	err := NotifyPasswordChange(&s)
	if err != nil {
		t.Errorf("Password Change Confirmation Failed\n")
	}
}




