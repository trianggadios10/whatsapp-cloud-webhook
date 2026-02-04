# WhatsApp Cloud API — Meta Developer Portal Setup Guide

Step-by-step guide to register a Meta Developer account, create a WhatsApp Business app, generate API credentials, and configure webhooks.

---

## Table of Contents

- [1. Prerequisites](#1-prerequisites)
- [2. Register a Meta Developer Account](#2-register-a-meta-developer-account)
- [3. Create a New App](#3-create-a-new-app)
- [4. Create a Business Portfolio](#4-create-a-business-portfolio)
- [5. Complete App Creation](#5-complete-app-creation)
- [6. Set Up WhatsApp Business Platform](#6-set-up-whatsapp-business-platform)
- [7. Generate an Access Token](#7-generate-an-access-token)
- [8. Configure the Webhook](#8-configure-the-webhook)
- [9. Retrieve the App Secret](#9-retrieve-the-app-secret)
- [10. Credentials Summary](#10-credentials-summary)
- [11. Understanding Message Types and Limitations](#11-understanding-message-types-and-limitations)

---

## 1. Prerequisites

Before you begin, make sure you have:

- A Facebook account (personal account is fine)
- A mobile phone number for SMS verification
- A business email address (Gmail or any email works for development)

---

## 2. Register a Meta Developer Account

If you already have a Meta Developer account, skip to [Step 3](#3-create-a-new-app).

**Step 2.1** — Open your browser and navigate to:

```
https://developers.facebook.com
```

**Step 2.2** — Click the **"Get Started"** button in the top-right corner of the page. If you're not logged into Facebook, you'll be prompted to log in first.

**Step 2.3 — Verify Your Account:**

A dialog titled **"Verify Your Account"** will appear asking for your mobile number.

1. Select your country code from the dropdown (e.g., **Indonesia (+62)**).
2. Enter your mobile phone number in the input field.
3. Click **"Send via SMS"** to receive a verification code.

**Step 2.4 — Enter Verification Code:**

A 6-digit verification code will be sent to your phone via SMS.

1. Enter the 6-digit code in the input fields.
2. Click **"Continue"** (or it may auto-submit once all digits are entered).

**Step 2.5 — Confirm Contact Email:**

A page titled **"Let's confirm your contact info"** will appear showing your email address associated with your Facebook account.

1. Review the email address shown. This will be used for developer communications.
2. If correct, click **"Confirm Email"**.
3. If you want to use a different email, click **"Change email"** and enter a new one.

**Step 2.6 — Select Your Role:**

A page titled **"About you"** will appear asking what best describes you.

1. Select **"Developer"** from the available options (Developer, Marketer, Analyst, etc.).
2. Click **"Complete Registration"**.

Your Meta Developer account is now active.

---

## 3. Create a New App

**Step 3.1** — After registration, you'll land on the **App Dashboard**. It should show **"No apps yet"** if this is your first time.

Click the **"Create App"** button.

**Step 3.2** — A popup may appear describing the new app creation system. Read through it and click **"Continue"** or **"Next"** to proceed.

**Step 3.3 — Enter App Details:**

You'll see an **"App details"** form with the following fields:

1. **App name**: Enter a descriptive name for your app.
   - Example: `my-whatsapp-webhook`
2. **App contact email**: Enter your business email address.
   - Example: `youremail@gmail.com`
3. Click **"Next"** to continue.

**Step 3.4 — Select Use Case:**

A **"Use cases"** page will appear with several options. Look for:

> **Connect with customers through WhatsApp**
> Start a WhatsApp conversation, send notifications, create ads that click-to-WhatsApp and provide support.

1. Select this option by clicking on it.
2. Click **"Next"** to continue.

**Step 3.5 — Connect a Business Portfolio:**

A **"Business"** page will appear asking which business portfolio to connect to your app.

- If you **already have a business portfolio**, select it and click **"Next"**.
- If you see **"No businesses available"**, you need to create one — see [Step 4](#4-create-a-business-portfolio).

---

## 4. Create a Business Portfolio

> **Note:** A Business Portfolio is required for WhatsApp Business API apps. This is where Meta manages your business assets (WhatsApp accounts, Facebook pages, etc.).

**Step 4.1** — On the "Business" page, click **"Create a new one"** (or **"Create a business portfolio"** link).

**Step 4.2** — A popup titled **"Create a business portfolio"** will appear with the following fields:

1. **Business portfolio name**: Enter your company or brand name.
   - Example: `My Company`
   - This name will be publicly visible across Meta, so use your real business name.
   - It cannot contain special characters.
2. **First name**: Enter your first name.
3. **Last name**: Enter your last name.
4. **Business email**: Enter your business contact email.

**Step 4.3** — Click **"Create portfolio"**.

**Step 4.4** — A success message will appear: **"[Your Business Name] was created"**.

It will also prompt you to verify your business. For development/testing purposes:

1. Click **"Verify later"** to skip verification for now.
2. You can always come back to verify when you're ready for production.

> **Important:** Business verification is required later if you want to go live (production mode), submit for App Review, or access data from other businesses.

**Step 4.5** — Back on the **"Business"** page, your newly created portfolio should now appear.

1. Select your business portfolio by clicking the radio button next to it.
2. Click **"Next"** to continue.

---

## 5. Complete App Creation

**Step 5.1 — Publishing Requirements:**

A **"Requirements"** page will appear showing any publishing requirements.

- It will typically say: **"No requirements identified. This may change if you add more to this app."**
- Click **"Next"** to continue.

**Step 5.2 — Overview:**

A summary **"Overview"** page will appear showing all your selections:

- **App Name**: Your app name
- **App Email**: Your contact email
- **Use Case**: Connect with customers through WhatsApp
- **Business**: Your business portfolio name (Unverified business)
- **Requirements**: No requirements for the use cases on this app

Review everything and click **"Go to dashboard"** to create the app.

**Step 5.3** — You'll be redirected to your app's **Dashboard**. The app is now created!

---

## 6. Set Up WhatsApp Business Platform

**Step 6.1** — On the Dashboard, click **"Customize the Connect with customers through WhatsApp use case"** (the first item in the list).

**Step 6.2** — You'll be taken to the **"Customize use case"** page. On the left sidebar, you'll see:

- Permissions and features
- **Quickstart** (selected by default)
- API Setup
- Configuration
- Resources

**Step 6.3 — WhatsApp Business Platform Quickstart:**

A card will appear showing:

- **Select a business portfolio**: Your business name should be pre-selected.
- A note that you'll receive a **WhatsApp test phone number** to send messages to a maximum of **5 phone numbers**.

Click **"Continue"** to proceed.

**Step 6.4** — A welcome page will appear: **"Welcome to the WhatsApp Business Platform"**. This is your main hub for WhatsApp API management.

**Step 6.5** — Click **"API Setup"** in the left sidebar to access your API credentials.

**Step 6.6 — API Setup Page:**

This is a critical page containing your API credentials:

- **Access Token**: An empty field with a **"Generate access token"** button
- **Send and receive messages** section:
  - **Step 1: Select phone numbers**
    - **From**: A dropdown showing your test number (e.g., `Test number: +1 555 123 4567`)
    - **Phone number ID**: A numeric ID (e.g., `123456789012345`) — **copy and save this**
    - **WhatsApp Business Account ID**: A numeric ID (e.g., `987654321098765`) — **copy and save this**

> **Important:** Write down or copy the **Phone number ID** and **WhatsApp Business Account ID**.

---

## 7. Generate an Access Token

**Step 7.1** — On the **API Setup** page, click the **"Generate access token"** button.

**Step 7.2** — A new browser window will open with a **Facebook Login for Business** dialog.

**Page 1 — Continue as [Your Name]:**

- The dialog will show: "Continue as [Your Name]?"
- Click **"Continue as [Your Name]"** to proceed.

**Step 7.3 — Choose WhatsApp Accounts:**

You'll see two options:

1. **Opt in to all current and future WhatsApp accounts** — Recommended for simplicity
2. **Opt in to current WhatsApp accounts only** — Only gives access to accounts you explicitly select

Click **"Continue"**.

**Step 7.4 — Review Access Request:**

This page summarizes the permissions being granted:

- **Manage your WhatsApp accounts**
- **Manage and access conversations in WhatsApp**

Click **"Save"** to confirm.

**Step 7.5** — The popup window will close automatically, and you'll be returned to the **API Setup** page. The **Access Token** field should now be populated with a long token string.

> **Important:** Copy this access token immediately and store it securely. This is a **temporary token** that expires after ~24 hours.

---

## 8. Configure the Webhook

**Step 8.1** — In the left sidebar, click **"Configuration"**.

**Step 8.2** — The **Webhook** section will appear with the following fields:

- **Callback URL**: The public URL where your server will receive webhook events
  - Example: `https://yourdomain.com/webhook`
  - For local development with ngrok: `https://abc123.ngrok-free.app/webhook`
- **Verify token**: A secret string that **you create yourself**

**Step 8.3** — Enter your **Callback URL** and **Verify token**.

> **Important:** Your server must be running and publicly accessible **before** you click "Verify and save".

**Step 8.4** — Click **"Verify and save"**.

Meta will send a verification request to your callback URL:

```
GET https://yourdomain.com/webhook?hub.mode=subscribe&hub.verify_token=YOUR_VERIFY_TOKEN&hub.challenge=CHALLENGE_STRING
```

Your server must:
1. Check that `hub.mode` is `subscribe`
2. Check that `hub.verify_token` matches your verify token
3. Respond with the `hub.challenge` value as plain text with HTTP status `200`

**Step 8.5 — Subscribe to Webhook Fields:**

After saving the webhook, you need to subscribe to webhook fields to receive events.

1. On the **Configuration** page, scroll down to the **Webhook fields** section
2. Click the **"Manage"** button
3. Toggle the **"Subscribe"** switch next to **messages** to enable it
4. Click **"Done"** to save

| Field | Description | Subscribe? |
|-------|-------------|------------|
| `messages` | Incoming messages and message status updates | **Yes (Required)** |
| `message_template_status_update` | Template approval/rejection notifications | Optional |
| `account_alerts` | Account-level alerts and warnings | Optional |

> **Important:** You must subscribe to the `messages` field to receive incoming messages.

---

## 9. Retrieve the App Secret

The **App Secret** is needed to verify the authenticity of webhook payloads.

**Step 9.1** — In the left sidebar, click **"App settings"** (the gear icon near the bottom).

**Step 9.2** — Click **"Basic"** under App settings.

**Step 9.3** — The **Basic Settings** page will display:

- **App ID**: Your app's unique identifier (e.g., `1234567890123456`)
- **App secret**: Hidden behind dots with a **"Show"** button

**Step 9.4** — Click **"Show"** next to the App secret. You may be prompted to enter your Facebook password.

**Step 9.5** — Copy the **App Secret** and store it securely.

> **Security Warning:** Never expose your App Secret in client-side code, public repositories, or logs.

---

## 10. Credentials Summary

After completing all the steps above, you should have the following credentials:

| Credential | Where to Find |
|------------|---------------|
| **App ID** | App Settings > Basic |
| **App Secret** | App Settings > Basic > Click "Show" |
| **Access Token** | WhatsApp > API Setup > Generate access token |
| **Phone Number ID** | WhatsApp > API Setup |
| **WABA ID** | WhatsApp > API Setup |
| **Verify Token** | You created this yourself |

---

## 11. Understanding Message Types and Limitations

### Test Account Limitations

With a test WhatsApp Business Account, you can only send messages to **up to 5 phone numbers** that you have added to your allowed list.

**To add a phone number to the allowed list:**

1. Go to **WhatsApp > API Setup**
2. Find **"Step 2: Send messages with the API"** section
3. Click the **"To"** dropdown
4. Select **"Manage phone number list"**
5. Click **"Add phone number"**
6. Enter the recipient's phone number (with country code, e.g., `628xxxxxxxxxx`)
7. They will receive a verification code via WhatsApp
8. Enter the code to verify

### 24-Hour Messaging Window

WhatsApp has a **24-hour customer service window**:

- When a user sends you a message, a 24-hour window opens
- During this window, you can send **any type of message** (text, image, etc.) to that user
- After 24 hours, you can only send **template messages** to re-initiate the conversation

### Template Messages vs Text Messages

| Message Type | When to Use | Allowed List Required? (Test Account) |
|--------------|-------------|---------------------------------------|
| **Template Message** | Initiating conversation (user hasn't messaged you first) | Yes |
| **Text Message** | Replying within 24-hour window after user messages you | No |

### Common Error: "Recipient phone number not in allowed list"

Error code: `131030`

This error occurs when:
- You try to send a **template message** to a number not in your allowed list
- You try to send a message **outside** the 24-hour window

**Solution:**
- For replies within 24 hours: Use **text messages** instead of templates
- For initiating conversations: Add the number to your allowed list, or go to production

### Common Error: "Business account is restricted from messaging users in this country"

Error code: `130497`

This error occurs when:
- The **test phone number** (US-based +1 555...) cannot send messages to certain countries
- Indonesia and some other countries have this restriction for test accounts

**Solution:**

1. **Add your own phone number** (Recommended)
   - Go to **WhatsApp > API Setup**
   - Click **"Add phone number"**
   - Verify your own business phone number
   - Use that number instead of the test number

2. **Complete business verification**
   - Go to **Meta Business Suite > Settings > Business Info**
   - Complete the verification process
   - This may unlock more countries for messaging

3. **Test with a supported country**
   - Use a phone number from US, UK, or other supported countries for testing

### Going to Production (Unlimited Recipients)

To send messages to any phone number without restrictions:

1. Complete **business verification** in Meta Business Suite
2. Add a **real phone number** (not the test number)
3. Get your app **reviewed and approved** by Meta
4. Create and get **message templates approved**

---

## Quick Reference: Key URLs

| Page | URL |
|------|-----|
| Meta for Developers | `https://developers.facebook.com` |
| Your App Dashboard | `https://developers.facebook.com/apps/{YOUR_APP_ID}/dashboard/` |
| API Setup | `https://developers.facebook.com/apps/{YOUR_APP_ID}/use_cases/customize/wa-dev-console/` |
| Webhook Configuration | WhatsApp > Configuration |
| App Settings (Basic) | `https://developers.facebook.com/apps/{YOUR_APP_ID}/settings/basic/` |
| WhatsApp API Docs | `https://developers.facebook.com/docs/whatsapp/cloud-api` |

---

## Troubleshooting

**Webhook verification fails:**
- Make sure your server is running and publicly accessible before clicking "Verify and save"
- Verify token must match exactly (case-sensitive)
- Server must respond with `200 OK` and the `hub.challenge` value

**Not receiving messages:**
- Ensure you have subscribed to the `messages` webhook field
- Check that your app is using the correct phone number

**Access token expired:**
- Temporary tokens expire after ~24 hours
- Generate a new one from the API Setup page
