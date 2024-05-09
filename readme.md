# Email Verifier

## Overview

This project is a simple email verification service implemented in Go. It provides an HTTP endpoint to verify the validity and existence of email addresses. The service checks whether the provided email address matches the basic email pattern, has valid MX (Mail Exchange) records, and valid NS (Name Server) records.

## Installation

To run the Email Verifier service locally, you need to have Go installed on your machine. Clone this repository to your local environment:

```
git clone https://github.com/aykanatcanberk/mail_verifier.git
```

## Usage

1. Start the Email Verifier service by running the following command:

```
go run main.go
```

2. Once the service is running, you can send POST requests to the `/verify-email` endpoint with a JSON payload containing the email address to be verified. Here's an example request:

```json
{
  "email": "example@example.com"
}
```

3. The service will respond with a JSON object indicating whether the email address is valid and existent. Here's an example response:

```json
{
  "email": "example@example.com",
  "valid": true,
  "message": ""
}
```

If the email address is invalid or cannot be verified, the `valid` field will be `false`, and the `message` field will contain an error message.

## Postman Usage

Alternatively, you can use Postman to interact with the Email Verifier service:

1. Import the provided Postman collection `email-verifier.postman_collection.json` into your Postman application.
2. Send a POST request to the `/verify-email` endpoint with the email address you want to verify in the request body.
3. Check the response to see if the email address is valid and existent.

## Dependencies

This project relies on the following external packages:

- [net](https://golang.org/pkg/net/)
- [regexp](https://golang.org/pkg/regexp/)
- [strings](https://golang.org/pkg/strings/)
- [encoding/json](https://golang.org/pkg/encoding/json/)
- [log](https://golang.org/pkg/log/)
- [net/http](https://golang.org/pkg/net/http/)
