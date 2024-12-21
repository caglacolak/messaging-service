package messaging

import (
	"context"
	"errors"
	"net/http"
	"encoding/json"

	"project/internal/database"
	"project/internal/models"
)

func GetSentMessages(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	rows, err := database.DB.Query(ctx, "SELECT id, content, recipient, sent_at FROM messages WHERE status = 'sent'")
	if err != nil {
		http.Error(w, "Failed to fetch messages", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var messages []map[string]interface{}
	for rows.Next() {
		var id int
		var content, recipient, sentAt string
		if err := rows.Scan(&id, &content, &recipient, &sentAt); err != nil {
			continue
		}
		messages = append(messages, map[string]interface{}{
			"id":        id,
			"content":   content,
			"recipient": recipient,
			"sent_at":   sentAt,
		})
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(messages)
}

// GetUnsentMessages retrieves unsent messages from the database.
func GetUnsentMessages(ctx context.Context) ([]models.Message, error) {
	rows, err := database.DB.Query(ctx, "SELECT id, content, recipient, status, sent_at FROM messages WHERE status != 'sent'")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []models.Message
	for rows.Next() {
		var message models.Message
		err := rows.Scan(&message.ID, &message.Content, &message.Recipient, &message.Status, &message.SentAt)
		if err != nil {
			return nil, err
		}
		messages = append(messages, message)
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	return messages, nil
}

// MarkMessageAsSent updates the status of a message to 'sent' in the database.
func MarkMessageAsSent(ctx context.Context, messageID int) error {
	result, err := database.DB.Exec(ctx, "UPDATE messages SET status = 'sent', sent_at = NOW() WHERE id = $1", messageID)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return errors.New("no rows updated")
	}

	return nil
}
