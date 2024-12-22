package scheduler

import (
	"context"
	"fmt"
	"time"

	"project/internal/httpclient"
	"project/internal/messaging"
	"project/internal/redis"
)

var cancelFunc context.CancelFunc // Global variable to store cancel function

// fetch unsent messages and sends them
func SendMessages(ctx context.Context) {
	messages, err := messaging.GetUnsentMessages(ctx)
	if err != nil {
		fmt.Printf("Error retrieving unsent messages: %v\n", err)
		return
	}

	for _, message := range messages {
		resp, err := httpclient.PostMessage(ctx, message.Recipient, message.Content)
		if err != nil {
			fmt.Printf("Error sending message to %s: %v\n", message.Recipient, err)
			continue
		}

		if resp == nil {
			fmt.Printf("Received nil response for message to %s\n", message.Recipient)
			continue
		}
		messaging.MarkMessageAsSent(ctx, message.ID)

		cacheKey := fmt.Sprintf("message:%d", message.ID)
		value := fmt.Sprintf("message ID:%s , time:%s", resp.MessageID, time.Now().String())
		redis.Redis.Set(ctx, cacheKey, value, 0)
	}
}

// starts scheduler to send messages every 2 minutes
func StartScheduler() {
	ctx, cancel := context.WithCancel(context.Background())
	cancelFunc = cancel // Store the cancel function to stop the scheduler
	go func() {
		for {
			select {
			case <-ctx.Done(): // Check if context is canceled
				fmt.Println("Scheduler stopped.")
				return
			default:
				SendMessages(ctx)
				time.Sleep(2 * time.Minute)
			}
		}
	}()
}

// stops the scheduler
func StopScheduler() {
	if cancelFunc != nil {
		cancelFunc() // Call the stored cancel function
		fmt.Println("Scheduler stopping...")
	}
}
