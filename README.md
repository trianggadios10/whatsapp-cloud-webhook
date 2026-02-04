# WhatsApp Cloud API Webhook (Go)

A WhatsApp Cloud API webhook server written in Go.

## Features

- Webhook verification endpoint
- Webhook event handler
- Signature validation
- Support for all message types (text, image, audio, video, document, location, interactive)
- Message status tracking
- Auto-reply functionality
- Template message support

## Quick Start

### 1. Clone and configure

```bash
cp .env.example .env
```

Edit `.env` with your credentials:

```env
PORT=8080
FACEBOOK_APP_SECRET=your_app_secret_here
WEBHOOK_VERIFY_TOKEN=your_verify_token_here
WHATSAPP_ACCESS_TOKEN=your_access_token_here
WHATSAPP_PHONE_NUMBER_ID=your_phone_number_id_here
```

### 2. Run the server

```bash
go run main.go
```

### 3. Expose with ngrok (for development)

```bash
ngrok http 8080
```

### 4. Configure webhook in Meta Developer Portal

See [setup-developer-facebook-com/README.md](setup-developer-facebook-com/README.md) for detailed Facebook setup instructions.

## Environment Variables

| Variable | Required | Description |
|----------|----------|-------------|
| `WEBHOOK_VERIFY_TOKEN` | Yes | Token for webhook verification |
| `FACEBOOK_APP_SECRET` | No | For signature validation (recommended) |
| `WHATSAPP_ACCESS_TOKEN` | No | For sending messages |
| `WHATSAPP_PHONE_NUMBER_ID` | No | For sending messages |
| `PORT` | No | Server port (default: 8080) |

## Where to Find Your Credentials

| Credential | Where to Find |
|------------|---------------|
| App Secret | App Settings > Basic > App Secret |
| Access Token | WhatsApp > API Setup > Generate access token |
| Phone Number ID | WhatsApp > API Setup > Phone number ID |
| Verify Token | You create this yourself |

## API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/webhook` | Webhook verification |
| POST | `/webhook` | Receive webhook events |
| GET | `/health` | Health check |
| GET | `/` | Server info page |

## Project Structure

```
whatsapp-cloud-webhook/
├── main.go
├── go.mod
├── .env.example
├── setup-developer-facebook-com/
│   └── README.md          # Facebook Developer Portal setup guide
└── README.md
```

## Running the Server

### Using PowerShell (Windows)

```powershell
$env:WEBHOOK_VERIFY_TOKEN="your_verify_token"
$env:FACEBOOK_APP_SECRET="your_app_secret"
$env:WHATSAPP_ACCESS_TOKEN="your_access_token"
$env:WHATSAPP_PHONE_NUMBER_ID="your_phone_number_id"
go run main.go
```

### Using Bash (Linux/macOS)

```bash
export WEBHOOK_VERIFY_TOKEN="your_verify_token"
export FACEBOOK_APP_SECRET="your_app_secret"
export WHATSAPP_ACCESS_TOKEN="your_access_token"
export WHATSAPP_PHONE_NUMBER_ID="your_phone_number_id"
go run main.go
```

## Exposing Locally with ngrok

For local development, use ngrok to make your server publicly accessible:

```bash
ngrok http 8080
```

ngrok will output a public URL like:

```
Forwarding    https://abc123.ngrok-free.app -> http://localhost:8080
```

Use `https://abc123.ngrok-free.app/webhook` as your Callback URL in the Meta Configuration page.

## Testing

### Test verification endpoint

```bash
curl "http://localhost:8080/webhook?hub.mode=subscribe&hub.verify_token=your_token&hub.challenge=test123"
```

Should return: `test123`

### Test health endpoint

```bash
curl http://localhost:8080/health
```

Should return: `OK`

## Example Output

```
PS C:\Users\user\whatsapp-cloud-webhook> go run .\main.go
2026/02/04 16:40:48 Server starting on http://localhost:8080
2026/02/04 16:40:48 Webhook endpoint: http://localhost:8080/webhook
2026/02/04 16:40:48 Verify token: your-verify-token
2026/02/04 16:41:56 Verification request - mode: subscribe, token: your-verify-token
2026/02/04 16:41:56 Webhook verified successfully
2026/02/04 16:43:46 Received webhook: object=whatsapp_business_account
2026/02/04 16:43:46 Field: messages
2026/02/04 16:43:46 Message from YourContactName (628xxxxxxxxxx)
2026/02/04 16:43:46 ID: wamid.XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX
2026/02/04 16:43:46 Type: text
2026/02/04 16:43:46 Timestamp: 1770198225
2026/02/04 16:43:46 Text: testing
```

## Auto-Reply Rules

The server includes auto-reply functionality. Edit `handleAutoReply` in `main.go` to customize:

| Trigger | Response |
|---------|----------|
| "hello", "hi", "halo" | Sends hello_world template |
| "help", "bantuan" | Sends help menu |
| "info" | Sends info message |

## Sending Messages

### Send Text Message

```go
sendTextMessage("628xxxxxxxxxx", "Hello from Go!")
```

### Send Template Message

```go
sendTemplateMessage("628xxxxxxxxxx", "hello_world", "en_US", nil)
```

### Send Template with Parameters

```go
sendTemplateWithBodyParams("628xxxxxxxxxx", "order_update", "en_US", []string{"#12345"})
```

## Troubleshooting

**Webhook verification fails:**
- Make sure your server is running and publicly accessible
- Verify token must match exactly (case-sensitive)
- Server must respond with hub.challenge value

**Not receiving messages:**
- Subscribe to the `messages` webhook field
- Check that you're using the correct phone number

**Invalid signature errors:**
- Verify FACEBOOK_APP_SECRET matches your App Secret
- Make sure you're reading the raw request body

**Access token expired:**
- Temporary tokens expire after ~24 hours
- Generate a new one from the API Setup page

## License

MIT
