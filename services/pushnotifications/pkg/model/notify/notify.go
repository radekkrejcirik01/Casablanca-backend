package notify

import (
	"log"

	"github.com/appleboy/go-fcm"
)

const client = "AAAAafwM0nY:APA91bHLAtUNExAqr_Z-xlwFaKE_9AzOuKCWYJ7FWwVadQFABmOhyNNlayojY3fAgQ9RcHmAyFYxqOyfFE3YWybY2-8BzgaEQBUXT7MI7NUGcYFdjsb4Bo7FQkoFkoakufS8oLGjOV3W"

type NotifyDevice struct {
	Token string
}

func Notify(t *NotifyDevice) error {
	// Create the message to be sent.
	msg := &fcm.Message{
		To: t.Token,
		Data: map[string]interface{}{
			"foo": "bar",
		},
		Notification: &fcm.Notification{
			Title: "title",
			Body:  "body",
		},
	}

	// Create a FCM client to send the message.
	client, err := fcm.NewClient(client)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	// Send the message and receive the response without retries.
	response, err := client.Send(msg)
	if err != nil {
		log.Fatalln(err)
		return err
	}

	log.Printf("%#v\n", response)

	return nil
}
