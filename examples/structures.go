package main

import (
	"fmt"
	"time"
)

func sample() int {
	return 1
}

type Notification struct {
	messageID int
	message   string
	createdAt time.Time
	// sample    sample (this won't work because sample is not a type)
}

func (notification Notification) sample() int {
	return 1
}

func addNotification(notification *Notification) {
	notification.messageID = 10
	notification.message = "Record Saved Successfully"
	notification.createdAt = time.Now()
}

func main() {
	var notification Notification
	var notificationPtr *Notification

	notificationPtr = &notification
	addNotification(&notification)
	fmt.Printf("message id - %d", (*notificationPtr).messageID)
	fmt.Printf("message id - %s", (*notificationPtr).message)
	fmt.Printf("message id - %s", (*notificationPtr).createdAt)
	fmt.Printf("sample - %d", (*notificationPtr).sample())
}
