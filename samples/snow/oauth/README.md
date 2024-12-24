# OAuth Authentication in Go

This example demonstrates how to authenticate with an OAuth 2.0 provider and create an incident using the ServiceNow API in Go.

## Explanation

### creds Struct

The `creds` struct contains the necessary credentials for authenticating with the OAuth provider. It includes fields for the OAuth URL, incident URL, grant type, client ID, client secret, username, and password.

### token Struct

The `token` struct contains the fields of the authentication token returned by the OAuth provider. It includes fields for the access token, refresh token, scope, token type, and expiration time.

### main Function

The `main` function is the entry point of the program. It initializes a logger, calls the `getTokens` function to authenticate with the OAuth provider, and logs the acquired tokens.

### getTokens Function

The `getTokens` function authenticates with the OAuth provider and returns an authentication token. It constructs a credentials string, creates an HTTP POST request, and sends the request to the OAuth URL. The response body is read and unmarshaled into a `token` struct, which is then returned.

### createIncident Function

The `createIncident` function creates an incident in ServiceNow using the provided authentication token and credentials. It constructs a JSON payload with the incident parameters, creates an HTTP POST request, and sends the request to the incident URL. The response body is read and printed.
