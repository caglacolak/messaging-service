package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

const webhookURL = "https://webhook.site/410ba03e-9baf-4633-90bc-2346921e6176"

type PostMessageRequest struct {
	To      string `json:"to"`
	Content string `json:"content"`
}

type PostMessageResponse struct {
	Message   string `json:"message"`
	MessageID string `json:"messageId"`
}

// PostMessage sends a message to the webhook API.
func PostMessage(ctx context.Context, to, content string) (*PostMessageResponse, error) {
	body, err := json.Marshal(PostMessageRequest{To: to, Content: content})
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, webhookURL, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-ins-auth-key", "INS.me1x9uMcyYGlhKKQVPoc.bO3j9aZwRTOcA2Ywo")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, errors.New("failed to send message")
	}

	var response PostMessageResponse
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
