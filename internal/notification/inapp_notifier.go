package notification

import "fmt"

type InAppNotifier struct{}

func (i *InAppNotifier) SendNotification(message string) error {
	// Implement in-app notification logic here
	fmt.Println("Sending in-app notification:", message)
	return nil
}
