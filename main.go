package main

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

var (
	VerifyToken   = os.Getenv("WEBHOOK_VERIFY_TOKEN")
	AppSecret     = os.Getenv("FACEBOOK_APP_SECRET")
	AccessToken   = os.Getenv("WHATSAPP_ACCESS_TOKEN")
	PhoneNumberID = os.Getenv("WHATSAPP_PHONE_NUMBER_ID")
	Port          = getEnvOrDefault("PORT", "8080")
)

func getEnvOrDefault(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

type WebhookPayload struct {
	Object string  `json:"object"`
	Entry  []Entry `json:"entry"`
}

type Entry struct {
	ID      string   `json:"id"`
	Changes []Change `json:"changes"`
}

type Change struct {
	Value Value  `json:"value"`
	Field string `json:"field"`
}

type Value struct {
	MessagingProduct string    `json:"messaging_product"`
	Metadata         Metadata  `json:"metadata"`
	Contacts         []Contact `json:"contacts,omitempty"`
	Messages         []Message `json:"messages,omitempty"`
	Statuses         []Status  `json:"statuses,omitempty"`
	Errors           []Error   `json:"errors,omitempty"`
}

type Metadata struct {
	DisplayPhoneNumber string `json:"display_phone_number"`
	PhoneNumberID      string `json:"phone_number_id"`
}

type Contact struct {
	Profile Profile `json:"profile"`
	WaID    string  `json:"wa_id"`
}

type Profile struct {
	Name string `json:"name"`
}

type Message struct {
	From        string              `json:"from"`
	ID          string              `json:"id"`
	Timestamp   string              `json:"timestamp"`
	Type        string              `json:"type"`
	Text        *TextMessage        `json:"text,omitempty"`
	Image       *MediaMessage       `json:"image,omitempty"`
	Audio       *MediaMessage       `json:"audio,omitempty"`
	Video       *MediaMessage       `json:"video,omitempty"`
	Document    *MediaMessage       `json:"document,omitempty"`
	Location    *LocationMessage    `json:"location,omitempty"`
	Button      *ButtonMessage      `json:"button,omitempty"`
	Interactive *InteractiveMessage `json:"interactive,omitempty"`
}

type TextMessage struct {
	Body string `json:"body"`
}

type MediaMessage struct {
	ID       string `json:"id"`
	MimeType string `json:"mime_type,omitempty"`
	SHA256   string `json:"sha256,omitempty"`
	Caption  string `json:"caption,omitempty"`
	Filename string `json:"filename,omitempty"`
}

type LocationMessage struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Name      string  `json:"name,omitempty"`
	Address   string  `json:"address,omitempty"`
}

type ButtonMessage struct {
	Text    string `json:"text"`
	Payload string `json:"payload"`
}

type InteractiveMessage struct {
	Type        string       `json:"type"`
	ButtonReply *ButtonReply `json:"button_reply,omitempty"`
	ListReply   *ListReply   `json:"list_reply,omitempty"`
}

type ButtonReply struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type ListReply struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description,omitempty"`
}

type Status struct {
	ID          string  `json:"id"`
	Status      string  `json:"status"`
	Timestamp   string  `json:"timestamp"`
	RecipientID string  `json:"recipient_id"`
	Errors      []Error `json:"errors,omitempty"`
}

type Error struct {
	Code    int             `json:"code"`
	Title   string          `json:"title"`
	Message string          `json:"message,omitempty"`
	Details json.RawMessage `json:"error_data,omitempty"`
}

type SendMessageRequest struct {
	MessagingProduct string           `json:"messaging_product"`
	RecipientType    string           `json:"recipient_type,omitempty"`
	To               string           `json:"to"`
	Type             string           `json:"type"`
	Text             *TextMessage     `json:"text,omitempty"`
	Template         *TemplateMessage `json:"template,omitempty"`
}

type TemplateMessage struct {
	Name       string              `json:"name"`
	Language   *TemplateLanguage   `json:"language"`
	Components []TemplateComponent `json:"components,omitempty"`
}

type TemplateLanguage struct {
	Code string `json:"code"`
}

type TemplateComponent struct {
	Type       string              `json:"type"`
	SubType    string              `json:"sub_type,omitempty"`
	Index      *int                `json:"index,omitempty"`
	Parameters []TemplateParameter `json:"parameters,omitempty"`
}

type TemplateParameter struct {
	Type     string         `json:"type"`
	Text     string         `json:"text,omitempty"`
	Currency *CurrencyValue `json:"currency,omitempty"`
	DateTime *DateTimeValue `json:"date_time,omitempty"`
	Image    *MediaValue    `json:"image,omitempty"`
	Document *MediaValue    `json:"document,omitempty"`
	Video    *MediaValue    `json:"video,omitempty"`
}

type CurrencyValue struct {
	FallbackValue string `json:"fallback_value"`
	Code          string `json:"code"`
	Amount1000    int64  `json:"amount_1000"`
}

type DateTimeValue struct {
	FallbackValue string `json:"fallback_value"`
}

type MediaValue struct {
	ID   string `json:"id,omitempty"`
	Link string `json:"link,omitempty"`
}

type SendMessageResponse struct {
	MessagingProduct string `json:"messaging_product"`
	Contacts         []struct {
		Input string `json:"input"`
		WaID  string `json:"wa_id"`
	} `json:"contacts"`
	Messages []struct {
		ID string `json:"id"`
	} `json:"messages"`
}

func main() {
	if VerifyToken == "" {
		log.Println("WARNING: WEBHOOK_VERIFY_TOKEN not set, using default")
		VerifyToken = "my_verify_token"
	}

	http.HandleFunc("/webhook", webhookHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/", homeHandler)

	addr := ":" + Port
	log.Printf("Server starting on http://localhost%s", addr)
	log.Printf("Webhook endpoint: http://localhost%s/webhook", addr)
	log.Printf("Verify token: %s", VerifyToken)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}

func webhookHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		handleVerification(w, r)
	case http.MethodPost:
		handleWebhookEvent(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleVerification(w http.ResponseWriter, r *http.Request) {
	mode := r.URL.Query().Get("hub.mode")
	token := r.URL.Query().Get("hub.verify_token")
	challenge := r.URL.Query().Get("hub.challenge")

	log.Printf("Verification request - mode: %s, token: %s", mode, token)

	if mode == "subscribe" && token == VerifyToken {
		log.Println("Webhook verified successfully")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(challenge))
		return
	}

	log.Println("Webhook verification failed")
	http.Error(w, "Forbidden", http.StatusForbidden)
}

func handleWebhookEvent(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error reading body: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	if AppSecret != "" {
		signature := r.Header.Get("X-Hub-Signature-256")
		if !validateSignature(body, signature) {
			log.Println("Invalid signature")
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}
	}

	var payload WebhookPayload
	if err := json.Unmarshal(body, &payload); err != nil {
		log.Printf("Error parsing payload: %v", err)
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	processWebhookPayload(payload)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("EVENT_RECEIVED"))
}

func validateSignature(body []byte, signature string) bool {
	if signature == "" {
		return false
	}

	signature = strings.TrimPrefix(signature, "sha256=")
	mac := hmac.New(sha256.New, []byte(AppSecret))
	mac.Write(body)
	expectedSignature := hex.EncodeToString(mac.Sum(nil))

	return hmac.Equal([]byte(signature), []byte(expectedSignature))
}

func processWebhookPayload(payload WebhookPayload) {
	log.Printf("Received webhook: object=%s", payload.Object)

	for _, entry := range payload.Entry {
		for _, change := range entry.Changes {
			log.Printf("Field: %s", change.Field)

			for _, msg := range change.Value.Messages {
				processMessage(msg, change.Value.Contacts)
			}

			for _, status := range change.Value.Statuses {
				processStatus(status)
			}

			for _, err := range change.Value.Errors {
				processError(err)
			}
		}
	}
}

func handleAutoReply(to, senderName, messageText string) {
	lowerText := strings.ToLower(messageText)

	switch {
	case strings.Contains(lowerText, "hello") || strings.Contains(lowerText, "hi") || strings.Contains(lowerText, "halo"):
		if err := sendTextMessage(to, fmt.Sprintf("Hello %s! Welcome to our service. How can I help you today?", senderName)); err != nil {
			log.Printf("Failed to send hello message: %v", err)
		}

	case strings.Contains(lowerText, "help") || strings.Contains(lowerText, "bantuan"):
		msg := fmt.Sprintf("Hi %s! How can I help you today?\n\n1. Type 'info' for information\n2. Type 'contact' for contact details\n3. Type 'hello' for a greeting", senderName)
		if err := sendTextMessage(to, msg); err != nil {
			log.Printf("Failed to send help message: %v", err)
		}

	case strings.Contains(lowerText, "info"):
		if err := sendTextMessage(to, "This is a WhatsApp Cloud API webhook demo.\n\nBuilt with Go."); err != nil {
			log.Printf("Failed to send info message: %v", err)
		}

	default:
		log.Printf("No auto-reply rule matched for: %s", messageText)
	}
}

func processMessage(msg Message, contacts []Contact) {
	senderName := "Unknown"
	for _, contact := range contacts {
		if contact.WaID == msg.From {
			senderName = contact.Profile.Name
			break
		}
	}

	log.Printf("Message from %s (%s)", senderName, msg.From)
	log.Printf("ID: %s", msg.ID)
	log.Printf("Type: %s", msg.Type)
	log.Printf("Timestamp: %s", msg.Timestamp)

	switch msg.Type {
	case "text":
		if msg.Text != nil {
			log.Printf("Text: %s", msg.Text.Body)
			go handleAutoReply(msg.From, senderName, msg.Text.Body)
		}
	case "image":
		if msg.Image != nil {
			log.Printf("Image ID: %s", msg.Image.ID)
			log.Printf("Caption: %s", msg.Image.Caption)
		}
	case "audio":
		if msg.Audio != nil {
			log.Printf("Audio ID: %s", msg.Audio.ID)
		}
	case "video":
		if msg.Video != nil {
			log.Printf("Video ID: %s", msg.Video.ID)
		}
	case "document":
		if msg.Document != nil {
			log.Printf("Document ID: %s", msg.Document.ID)
			log.Printf("Filename: %s", msg.Document.Filename)
		}
	case "location":
		if msg.Location != nil {
			log.Printf("Location: %f, %f", msg.Location.Latitude, msg.Location.Longitude)
			log.Printf("Name: %s", msg.Location.Name)
		}
	case "button":
		if msg.Button != nil {
			log.Printf("Button: %s (payload: %s)", msg.Button.Text, msg.Button.Payload)
		}
	case "interactive":
		if msg.Interactive != nil {
			if msg.Interactive.ButtonReply != nil {
				log.Printf("Button Reply: %s (id: %s)", msg.Interactive.ButtonReply.Title, msg.Interactive.ButtonReply.ID)
			}
			if msg.Interactive.ListReply != nil {
				log.Printf("List Reply: %s (id: %s)", msg.Interactive.ListReply.Title, msg.Interactive.ListReply.ID)
			}
		}
	default:
		log.Printf("Unhandled message type: %s", msg.Type)
	}
}

func processStatus(status Status) {
	log.Printf("Status update for message %s", status.ID)
	log.Printf("Status: %s", status.Status)
	log.Printf("Recipient: %s", status.RecipientID)
	log.Printf("Timestamp: %s", status.Timestamp)

	if len(status.Errors) > 0 {
		for _, err := range status.Errors {
			log.Printf("Error: [%d] %s - %s", err.Code, err.Title, err.Message)
		}
	}
}

func processError(err Error) {
	log.Printf("Webhook Error: [%d] %s", err.Code, err.Title)
	if err.Message != "" {
		log.Printf("Message: %s", err.Message)
	}
	if len(err.Details) > 0 {
		log.Printf("Details: %s", string(err.Details))
	}
}

func sendTextMessage(to, message string) error {
	if AccessToken == "" || PhoneNumberID == "" {
		return fmt.Errorf("WHATSAPP_ACCESS_TOKEN and WHATSAPP_PHONE_NUMBER_ID must be set")
	}

	url := fmt.Sprintf("https://graph.facebook.com/v18.0/%s/messages", PhoneNumberID)

	request := SendMessageRequest{
		MessagingProduct: "whatsapp",
		RecipientType:    "individual",
		To:               to,
		Type:             "text",
		Text:             &TextMessage{Body: message},
	}

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	req, err := http.NewRequest("POST", url, strings.NewReader(string(jsonBody)))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API error: %s - %s", resp.Status, string(body))
	}

	var response SendMessageResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}

	if len(response.Messages) > 0 {
		log.Printf("Message sent successfully, ID: %s", response.Messages[0].ID)
	}

	return nil
}

func sendTemplateMessage(to, templateName, languageCode string, components []TemplateComponent) error {
	if AccessToken == "" || PhoneNumberID == "" {
		return fmt.Errorf("WHATSAPP_ACCESS_TOKEN and WHATSAPP_PHONE_NUMBER_ID must be set")
	}

	url := fmt.Sprintf("https://graph.facebook.com/v18.0/%s/messages", PhoneNumberID)

	request := SendMessageRequest{
		MessagingProduct: "whatsapp",
		To:               to,
		Type:             "template",
		Template: &TemplateMessage{
			Name:       templateName,
			Language:   &TemplateLanguage{Code: languageCode},
			Components: components,
		},
	}

	jsonBody, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("failed to marshal request: %w", err)
	}

	log.Printf("Sending template message to %s: %s", to, templateName)

	req, err := http.NewRequest("POST", url, strings.NewReader(string(jsonBody)))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+AccessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("API error: %s - %s", resp.Status, string(body))
	}

	var response SendMessageResponse
	if err := json.Unmarshal(body, &response); err != nil {
		return fmt.Errorf("failed to parse response: %w", err)
	}

	if len(response.Messages) > 0 {
		log.Printf("Template message sent successfully, ID: %s", response.Messages[0].ID)
	}

	return nil
}

func sendHelloWorldTemplate(to string) error {
	return sendTemplateMessage(to, "hello_world", "en_US", nil)
}

func sendTemplateWithBodyParams(to, templateName, languageCode string, params []string) error {
	var parameters []TemplateParameter
	for _, p := range params {
		parameters = append(parameters, TemplateParameter{Type: "text", Text: p})
	}

	components := []TemplateComponent{
		{Type: "body", Parameters: parameters},
	}

	return sendTemplateMessage(to, templateName, languageCode, components)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	html := `<!DOCTYPE html>
<html>
<head>
    <title>WhatsApp Webhook Server</title>
    <style>
        body { font-family: Arial, sans-serif; max-width: 800px; margin: 50px auto; padding: 20px; }
        h1 { color: #25D366; }
        code { background: #f4f4f4; padding: 2px 6px; border-radius: 4px; }
        pre { background: #f4f4f4; padding: 15px; border-radius: 8px; overflow-x: auto; }
        .status { padding: 10px; border-radius: 8px; margin: 20px 0; }
        .status.ok { background: #d4edda; color: #155724; }
    </style>
</head>
<body>
    <h1>WhatsApp Webhook Server</h1>
    <div class="status ok">Server is running!</div>

    <h2>Endpoints</h2>
    <ul>
        <li><code>GET /webhook</code> - Webhook verification</li>
        <li><code>POST /webhook</code> - Receive webhook events</li>
        <li><code>GET /health</code> - Health check</li>
    </ul>

    <h2>Configuration</h2>
    <p>Set these environment variables:</p>
    <pre>
WEBHOOK_VERIFY_TOKEN=your_verify_token
FACEBOOK_APP_SECRET=your_app_secret
WHATSAPP_ACCESS_TOKEN=your_access_token
WHATSAPP_PHONE_NUMBER_ID=your_phone_number_id
PORT=8080 (optional)
    </pre>
</body>
</html>`
	w.Write([]byte(html))
}
