package main

import "fmt"

// SMS Email
/* 2*4 */
type INotifiacionFactory interface {
	SendNotification()
	GetSender() ISender
}

type ISender interface {
	GetSenderMethod() string
	GetSenderChanel() string
}

type SMSNotification struct {
}

func (SMSNotification) SendNotification() {
	fmt.Println("Sending Notification via SMS")
}
func (SMSNotification) GetSender() ISender {
	return SMSNotificationSender{}
}

type SMSNotificationSender struct {
}

func (SMSNotificationSender) GetSenderMethod() string {
	return "SMS"
}

func (SMSNotificationSender) GetSenderChanel() string {
	return "Twilio"
}

type EmailNotification struct {
}

func (EmailNotification) SendNotification() {
	fmt.Println("Sending notification via email")
}

func (EmailNotification) GetSender() ISender {
	return EmailNotificationSender{}
}

type EmailNotificationSender struct {
}

func (EmailNotificationSender) GetSenderMethod() string {
	return "Email"
}

func (EmailNotificationSender) GetSenderChanel() string {
	return "SES"
}

func getNotificationFactory(notificationType string) (INotifiacionFactory, error) {
	if notificationType == "SMS" {
		return &SMSNotification{}, nil
	}
	if notificationType == "Email" {
		return &EmailNotification{}, nil
	}
	return nil, fmt.Errorf("no notification type")

}
func SendNotification(f INotifiacionFactory) {
	f.SendNotification()
}
func getMethod(f INotifiacionFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("Email")
	SendNotification(smsFactory)
	SendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)
}
