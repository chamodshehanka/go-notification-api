package notification

import "fmt"

type SMSNotifier struct{}

func (s *SMSNotifier) SendNotification(message string) error {
	// Implement SMS notification logic here
	fmt.Println("Sending SMS notification:", message)
	return nil
}
