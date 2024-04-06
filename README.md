
This document provides an overview of the security considerations and implementations within the project. It covers authentication, authorization, encryption, rate limiting, and audit logging.

## Security Measures Implemented

### 1. Authentication and Authorization

- **Authentication**:
    - Users are authenticated using a username-password combination.
    - User credentials are stored securely on the server.
    - Passwords are hashed for storage and comparison.
    - JSON Web Tokens (JWT) are issued upon successful authentication.

- **Authorization**:
    - Access control lists (ACLs) are used to determine which users have access to specific resources.
    - Each user has a set of allowed paths defined in the ACL.
    - Requests are only processed if the user has access to the requested resource.

### 2. Token-based Authentication

- **Encryption Algorithm**:
    - The project supports two cryptographic algorithms: HMAC and ECDSA.
    - HMAC (Hash-based Message Authentication Code) and ECDSA (Elliptic Curve Digital Signature Algorithm) are used for token generation and validation.
    - A private ECDSA key is generated at initialization for signing JWT tokens.
    - The encryption key for HMAC is kept secret.

### 3. Rate Limiting

- Requests from each remote address are rate-limited to prevent abuse or DoS attacks.
- Exceeding the rate limit results in a 429 Too Many Requests HTTP response.

### 4. Audit Logging

- Each request is logged for auditing purposes.
- Logs include information such as timestamp, request method, requested path, user, and IP address.
- Logs are stored in a separate file (`audit-log.txt`) for traceability.

### 5. OAuth2 Integration
- OAuth2 authentication is supported for third-party authentication.
- OAuth2 endpoints and credentials are configured in the `oauthConfig` variable.
- Endpoints for OAuth2 authorization and callback are defined.

### 6. Security Best Practices

- **Content Security Policy (CSP)**:
    - The `Content-Security-Policy` header is set to define the allowed sources for scripts, stylesheets, images, etc.
    - Content Security Policy headers are set to restrict resources loaded by the browser.
    - Default-src is set to 'self' to only allow resources from the same origin.
    - Additional directives can be added based on specific requirements.


### Deployment with Nginx and Apache Servers

This project can be deployed with NGINX or Apache web servers acting as reverse proxies to handle incoming requests and route them to the Go application container.



#### Disclaimer
This project serves as a learning exercise and may not be suitable for production use without further hardening and security assessment. Always follow security best practices and consult security professionals for guidance.