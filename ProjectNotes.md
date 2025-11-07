# Ok so do you want the routes to look like?
so the routes should be of following structure

the frontend will make the urls for the dynamic qr code look like `my-website.com/pizza-hut` or  `my-website.com/@ABC-Marketing/pizza-hut` (this may be better kinda like link tree). On the backend side of things this may either relate to a uuid or slug like `qrcode-backend.com/18196f64-5ddc-4860-a471-4a1a2aaf389b`

- `qrcode\uuid` it will redirect to another websiite a simple *get request*
- `qrcode\create_new` a *POST* request which consists of:
    - the uuid

# how will the data models look like?
### qr code

| Field           | Type          | Description                               |
| --------------- | ------------- | ----------------------------------------- |
| id              | string / UUID | Unique short ID used in the QR URL        |
| name            | string        | Optional name like “Summer Campaign 2025” |
| destination_url | string        | The current redirect target               |
| created_at      | datetime      | When the QR code was created              |
| updated_at      | datetime      | When the destination was last changed     |
| active          | boolean       | Whether the QR is currently active        |
| user_id         | foreign key   | Owner of this QR code                     |
| qr_code_image   | base64        | The image of the qrcode stored in base64  |

--- 
### QRScan /  Analytics
| Field       | Type          | Description                        |
| ----------- | ------------- | ---------------------------------- |
| id          | UUID          | Unique ID for each scan event      |
| qr_id       | foreign key   | Which QRCode was scanned           |
| scanned_at  | datetime      | When the scan occurred             |
| ip_address  | string        | IP (can be anonymized for privacy) |
| user_agent  | string        | Browser or app info                |
| device_type | string        | Mobile, desktop, tablet            |
| location    | string / JSON | Country/city (resolved from IP)    |
| referrer    | string        | Source link, if any                |
---
### Users
| Field      | Type     | Description            |
| ---------- | -------- | ---------------------- |
| id         | UUID     | Unique user identifier |
| email      | string   | User’s email           |
| name       | string   | User’s name            |
| created_at | datetime | Account creation date  |


# Considerations
May wanna use