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

	notificationID := "new_follow"

	err := notificationapi.Init(clientID, clientSecret)
	if err != nil {
		return err
	}

	//mergeTags is to pass dynamic values into the notification design.
	mergeTags := make(map[string]interface{}) // Change to map[string]interface{}
	mergeTags["comment"] = "Build something great :)"
	mergeTags["commentId"] = "commentId-1234-abcd-wxyz"

	err = notificationapi.Send(
		notificationapi.SendRequest{
			//The ID of the notification you wish to send. You can find this
			//value from the dashboard.
			NotificationId: notificationID,
			//The user to send the notification to.
			User: notificationapi.User{
				Id:    "chamodshehanka",
				Email: "hcsperera@gmail.com",
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
