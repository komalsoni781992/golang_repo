package main

import "fmt"

type Notifier interface {
	SendEmail(m string) string
	SendSMS(m string) string
}

type NotificationService struct {
	service string
}

func (s NotificationService) SendEmail(m string) string {
	fmt.Println("Service name: ", s.service, " sending email :", m)
	return "mail sent"
}

func (s NotificationService) SendSMS(m string) string {
	fmt.Println("Service name: ", s.service, " sending sms :", m)
	return "sms sent"
}

func main() {
	var s Notifier = NotificationService{service: "HelloService"}
	fmt.Println(s.SendEmail("How are you?"))
	fmt.Println(s.SendSMS("Trying to reach out to you!"))
}
