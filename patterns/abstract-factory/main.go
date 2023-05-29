package main

import (
	"errors"
	"fmt"
	"golang-concepts-apis/patterns/abstract-factory/types"
)

type INotificationFactory interface {
	SendNotification()
	GetSender() types.ISender
}

const SMS_NOTIFICATION_TYPE = "SMS"
const EMAIL_NOTIFICATION_TYPE = "Email"

func getNotificationFactory(notificationType string) (INotificationFactory, error) {
	if notificationType == SMS_NOTIFICATION_TYPE {
		return &types.SMSNotification{}, nil
	}
	if notificationType == EMAIL_NOTIFICATION_TYPE {
		return &types.EmailNotification{}, nil
	}

	return nil, errors.New("unsupported notification type")
}

func sendNotification(f INotificationFactory) {
	f.SendNotification()
}

func getMethod(f INotificationFactory) {
	fmt.Println(f.GetSender().GetSenderMethod())
}

func main() {
	smsFactory, _ := getNotificationFactory("SMS")
	emailFactory, _ := getNotificationFactory("aaa")

	sendNotification(smsFactory)
	sendNotification(emailFactory)

	getMethod(smsFactory)
	getMethod(emailFactory)
}
