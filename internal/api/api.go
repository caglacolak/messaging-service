package api

import (
	"net/http"
	"project/internal/messaging"
	"project/internal/scheduler"
)

var isProcessing bool

// StartServer initializes the HTTP server
func StartServer() error {
	http.HandleFunc("/sent-messages", messaging.GetSentMessages)
	http.HandleFunc("/start-messages", StartScheduler)
	http.HandleFunc("/stop-messages", StopScheduler)
	return http.ListenAndServe(":8080", nil)
}

// GetSentMessages retrieves a list of sent messages.
// @Summary Retrieve sent messages
// @Description Returns a list of all messages that have been sent.
// @Tags Messaging
// @Accept json
// @Produce json
// @Success 200  {object} models.Message
// @Failure 500
// @Router /sent-messages [get]
func GetSentMessages(w http.ResponseWriter, r *http.Request) {
	if isProcessing {
		http.Error(w, "Processing already started", http.StatusBadRequest)
		return
	}
	isProcessing = true
	go scheduler.StartScheduler()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Automatic message sending started."))
}

// StartScheduler starts the automatic message sending scheduler.
// @Summary Start the scheduler
// @Description Starts the scheduler that sends unsent messages every 2 minutes.
// @Tags Scheduler
// @Accept json
// @Produce json
// @Success 200 
// @Failure 500 
// @Router /start-messages [get]
// StartSendMessages starts the scheduler
func StartScheduler(w http.ResponseWriter, r *http.Request) {
	if isProcessing {
		http.Error(w, "Processing already started", http.StatusBadRequest)
		return
	}
	isProcessing = true
	go scheduler.StartScheduler()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Automatic message sending started."))
}

// StopScheduler stops the automatic message sending scheduler.
// @Summary Stop the scheduler
// @Description Stops the currently running scheduler for sending messages.
// @Tags Scheduler
// @Accept json
// @Produce json
// @Success 200 
// @Failure 500 
// @Router /stop-messages [get]
// StopScheduler stops the scheduler
func StopScheduler(w http.ResponseWriter, r *http.Request) {
	if !isProcessing {
		http.Error(w, "Processing not running", http.StatusBadRequest)
		return
	}
	isProcessing = false
	go scheduler.StopScheduler()
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Automatic message sending stopped."))
}