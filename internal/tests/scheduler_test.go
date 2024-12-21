package tests

import (
	"context"
	"project/internal/database"
	"project/internal/scheduler"
	"testing"
	"time"
)

func TestSendMessages(t *testing.T) {
	// Insert mock data into the database
	db := database.DB
	_, err := db.Exec(context.Background(), "INSERT INTO messages (id, content, recipient, status) VALUES (1, 'Test Message', '+1234567890', 'unsent')")
	if err != nil {
		t.Fatalf("Failed to insert mock data: %v", err)
	}

	scheduler.SendMessages(context.Background())

	// Verify the status update
	var status string
	err = db.QueryRow(context.Background(), "SELECT status FROM messages WHERE id = 1").Scan(&status)
	if err != nil {
		t.Fatalf("Failed to fetch message status: %v", err)
	}

	if status != "sent" {
		t.Errorf("Expected status 'sent', got '%s'", status)
	}
}
func TestStartScheduler(t *testing.T) {
	// Insert mock data into the database
	db := database.DB
	_, err := db.Exec(context.Background(), "INSERT INTO messages (id, content, recipient, status) VALUES (1, 'Test Message', '+1234567890', 'unsent')")
	if err != nil {
		t.Fatalf("Failed to insert mock data: %v", err)
	}

	// Start the scheduler
	scheduler.StartScheduler()

	// Wait for a short duration to allow the scheduler to run
	time.Sleep(3 * time.Second)

	// Verify the status update
	var status string
	err = db.QueryRow(context.Background(), "SELECT status FROM messages WHERE id = 1").Scan(&status)
	if err != nil {
		t.Fatalf("Failed to fetch message status: %v", err)
	}

	if status != "sent" {
		t.Errorf("Expected status 'sent', got '%s'", status)
	}

	// Stop the scheduler
	scheduler.StopScheduler()
}
