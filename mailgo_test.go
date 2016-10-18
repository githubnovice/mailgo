package mailgo

import (
	//"github.com/githubnovice/mailgo"
	"fmt"
	"testing"
)

func TestComplete(t *testing.T) {
	s := Session {Email: "gladwig@gladworx.com", Fname: "Frank", Lname:"Ziffle", URL:"http://BlackListme/Confirm"}
	fmt.Printf("(%v, %T)\n", s, s)
	err := CompleteBlacklist(&s)
	if err != nil {
		t.Errorf("TestComplete Failed\n")
	}
}

 func TestRegister(t *testing.T) {
	s := Session {Email: "gladwig@gladworx.com", Fname: "Frank", Lname:"Ziffle", URL:"http://BlackListme/Confirm"}
	err := ConfirmRegistration(&s)
	if err != nil {
		t.Errorf("Registration Confirmation Failed\n")
	}
}

func TestConfirmEmailChange(t *testing.T) {
	s := Session {Email: "gladwig@gladworx.com", Fname: "Frank", Lname:"Ziffle", URL:"http://BlackListme/Confirm"}
	err := ConfirmEmailChange(&s, "oldaddress@yahoo.com")
	if err != nil {
		t.Errorf("Email Change Confirmation Failed\n")
	}
}

func TestConfirmPasswordChange(t *testing.T) {
	s := Session {Email: "gladwig@gladworx.com", Fname: "Frank", Lname:"Ziffle", URL:"http://BlackListme/Confirm"}
	err := ConfirmPasswordChange(&s)
	if err != nil {
		t.Errorf("Password Change Confirmation Failed\n")
	}
}

func TestConfirmEmailAddressRegistration(t *testing.T) {
	s := Session {Email: "gladwig@gladworx.com", Fname: "Frank", Lname:"Ziffle", URL:"http://BlackListme/Confirm"}
	err := ConfirmEmailAddressRegistration(&s)
	if err != nil {
		t.Errorf("Email address Confirmation Failed\n")
	}
}

func TestConfirmEmailDeRegistration(t *testing.T) {
	s := Session {Email: "gladwig@gladworx.com", Fname: "Frank", Lname:"Ziffle", URL:"http://BlackListme/Confirm"}
	err := ConfirmEmailAddressDeRegistration(&s)
	if err != nil {
		t.Errorf("Email address deregistration Confirmation Failed\n")
	}
}


