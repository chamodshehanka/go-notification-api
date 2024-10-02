package notification

import (
	"fmt"
	"github.com/chamodshehanka/go-notification-api/internal/configs"
	notificationapi "github.com/notificationapi-com/notificationapi-go-server-sdk"
	"github.com/sirupsen/logrus"
)

type InAppNotifier struct{}

func (i *InAppNotifier) SendNotification(message string) error {
	clientID := configs.GetConfig().NotificationAPIConfig.ClientID
	clientSecret := configs.GetConfig().NotificationAPIConfig.ClientSecret

	err := notificationapi.Init(clientID, clientSecret)
	if err != nil {
		return err
	}

	//mergeTags is to pass dynamic values into the notification design.
	mergeTags := make(map[string]interface{}) // Change to map[string]interface{}
	mergeTags["item"] = "Krabby Patty Burger"
	mergeTags["address"] = "124 Conch Street"
	mergeTags["orderId"] = "1234567890"

	err = notificationapi.Send(
		notificationapi.SendRequest{
			//The ID of the notification you wish to send. You can find this
			//value from the dashboard.
			NotificationId: "order_tracking",
			//The user to send the notification to.
			User: notificationapi.User{
				Id:    "spongebob.squarepants",
				Email: "spongebob@squarepants.com",
			},
			MergeTags: mergeTags,
		},
	)
	if err != nil {
		logrus.Errorf("Error occurred while sending notification: %v", err)
		return err
	}

	// Implement in-app notification logic here
	fmt.Println("Sending in-app notification:", message)
	return nil
}
