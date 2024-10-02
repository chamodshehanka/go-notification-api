package notification

import "fmt"

type EmailNotifier struct {
}

func (e *EmailNotifier) SendNotification(message string) error {
	// Implement email notification logic here
	fmt.Println("Sending email notification:", message)
	return nil
}
