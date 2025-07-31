# VocaGame E-Wallet API Documentation

## Overview
VocaGame E-Wallet is a multi-currency digital wallet system that allows users to manage funds across different currencies (USD, EUR, JPY, IDR) with real-time currency conversion.

## Base URL
```
http://localhost:3030/voca-wallets/v1
```

## Authentication
Most endpoints require JWT authentication. Include the token in the Authorization header:
```
Authorization: Bearer <your_jwt_token>
```

## 📦 Postman Collection
For easier API testing, you can import our pre-configured Postman collection:

**Collection File:** `postman/VocaGame-E-Wallet.postman_collection.json`

### How to Import:
1. Open Postman
2. Click "Import" button
3. Select the collection file from the `postman/` directory
4. The collection includes all endpoints with sample requests
5. Set up environment variables for `base_url` and `auth_token`

### Environment Variables:
- `base_url`: `http://localhost:3030/voca-wallets/v1`
- `auth_token`: Your JWT token (obtained from login endpoint)

This collection includes pre-configured requests for all API endpoints with proper headers and sample payloads.

## API Endpoints

### 1. User Management

#### Register User
```http
POST /user/register
```

**Request Body:**
```json
{
  "username": "johndoe",
  "email": "john@example.com",
  "password": "securePassword123"
}
```

**Response:**
```json
{
  "status": 201,
  "message": "User registered successfully",
  "data": {
    "username": "johndoe",
    "email": "john@example.com"
  },
  "error": null
}
```

#### Login User
```http
POST /user/login
```

**Request Body:**
```json
{
  "username": "johndoe",
  "password": "securePassword123"
}
```

**Response:**
```json
{
  "status": 200,
  "message": "Login successful",
  "data": {
    "access_token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9...",
    "token_type": "Bearer",
    "expires_in": 900
  },
  "error": null
}
```

#### Get User Profile
```http
GET /user/profile
```
**Headers:** `Authorization: Bearer <token>`

**Response:**
```json
{
  "status": 200,
  "message": "Profile retrieved successfully",
  "data": {
    "id": 1,
    "username": "johndoe",
    "email": "john@example.com",
    "wallets": [
      {
        "id": 1,
        "name": "Primary Wallet"
      }
    ]
  },
  "error": null
}
```

### 2. Wallet Management

#### Create Wallet
```http
POST /wallet/
```
**Headers:** `Authorization: Bearer <token>`

**Request Body:**
```json
{
  "name": "Primary Wallet"
}
```

**Response:**
```json
{
  "status": 201,
  "message": "Wallet created successfully",
  "data": {
    "id": 1,
    "name": "Primary Wallet"
  },
  "error": null
}
```

#### Get Wallet Balance
```http
GET /wallet/balance/:wallet_id
```
**Headers:** `Authorization: Bearer <token>`

**Response:**
```json
{
  "status": 200,
  "message": "Wallet balance retrieved successfully",
  "data": {
    "name": "Primary Wallet",
    "balance": [
      {
        "currency_code": "USD",
        "balance": 100.50
      },
      {
        "currency_code": "EUR",
        "balance": 85.75
      }
    ]
  },
  "error": null
}
```

#### Get Total Balance (Converted)
```http
GET /wallet/balance/:wallet_id/total?currency_code=USD
```
**Headers:** `Authorization: Bearer <token>`

**Query Parameters:**
- `currency_code` (optional): Target currency for conversion. Default: USD

**Response:**
```json
{
  "status": 200,
  "message": "Total wallet balance retrieved successfully",
  "data": {
    "name": "Primary Wallet",
    "currency_code": "USD",
    "total": 215.25
  },
  "error": null
}
```

### 3. Transaction Operations

#### Deposit Funds
```http
POST /wallet/deposit
```
**Headers:** `Authorization: Bearer <token>`

**Request Body:**
```json
{
  "wallet_id": 1,
  "currency_code": "USD",
  "amount": 100.00
}
```

**Response:**
```json
{
  "status": 201,
  "message": "Wallet deposit successfully",
  "data": {
    "trx_id": "DEPOSIT-1-1627849200000",
    "balance": 200.50
  },
  "error": null
}
```

#### Withdraw Funds
```http
POST /wallet/withdraw
```
**Headers:** `Authorization: Bearer <token>`

**Request Body:**
```json
{
  "wallet_id": 1,
  "currency_code": "USD",
  "amount": 50.00
}
```

**Response:**
```json
{
  "status": 201,
  "message": "Wallet withdraw successfully",
  "data": {
    "trx_id": "WITHDRAWAL-1-1627849200000",
    "balance": 150.50
  },
  "error": null
}
```

#### Transfer Between Wallets
```http
POST /wallet/transfer
```
**Headers:** `Authorization: Bearer <token>`

**Request Body:**
```json
{
  "from_wallet_id": 1,
  "to_wallet_id": 2,
  "from_currency_code": "USD",
  "to_currency_code": "EUR",
  "amount": 100.00
}
```

**Response:**
```json
{
  "status": 201,
  "message": "Wallet transfer successfully",
  "data": {
    "transaction_id": "TRANSFER-1-1627849200000",
    "status": "success"
  },
  "error": null
}
```

#### Make Payment
```http
POST /wallet/payment
```
**Headers:** `Authorization: Bearer <token>`

**Request Body:**
```json
{
  "wallet_id": 1,
  "currency_code": "USD",
  "amount": 25.99,
  "description": "Coffee purchase"
}
```

**Response:**
```json
{
  "status": 201,
  "message": "Payment processed successfully",
  "data": {
    "transaction_id": "PAYMENT-1-1627849200000",
    "status": "success"
  },
  "error": null
}
```

## Error Responses

### Standard Error Format
```json
{
  "status": 400,
  "message": "Validation failed",
  "data": null,
  "error": {
    "message": "Username is required"
  }
}
```

### Common HTTP Status Codes

| Status Code | Description |
|-------------|-------------|
| 200 | Success |
| 201 | Created |
| 400 | Bad Request (validation errors) |
| 401 | Unauthorized (invalid/missing token) |
| 404 | Not Found |
| 422 | Unprocessable Entity (business logic errors) |
| 500 | Internal Server Error |

### Common Error Messages

| Error | Description |
|-------|-------------|
| `ErrInvalidAmount` | Amount must be greater than 0 |
| `ErrInsufficientBalance` | Insufficient balance for this operation |
| `ErrCurrencyNotFound` | Currency code not supported |
| `ErrWalletNotFound` | Wallet not found or access denied |
| `ErrUserNotFound` | User not found |
| `ErrIdentityAlreadyExists` | Username or email already exists |

## Supported Currencies

| Code | Name |
|------|------|
| USD | United States Dollar |
| EUR | Euro |
| JPY | Japanese Yen |
| IDR | Indonesian Rupiah |

## Exchange Rates
The system uses predefined exchange rates for currency conversion:

- USD → EUR: 1.1
- USD → JPY: 110.0
- USD → IDR: 15,300.0
- (Plus reverse rates and cross-currency rates)

## Rate Limiting
- 100 requests per minute per user
- 1000 requests per minute for unauthenticated endpoints

## Example Usage Flow

> 💡 **Quick Start Tip**: Import the Postman collection (`postman/VocaGame-E-Wallet.postman_collection.json`) for ready-to-use API requests with proper authentication and sample data.

### Manual cURL Examples:

1. **Register a new user**
   ```bash
   curl -X POST http://localhost:3030/voca-wallets/v1/user/register \
     -H "Content-Type: application/json" \
     -d '{"username":"john","email":"john@test.com","password":"password123"}'
   ```

2. **Login to get token**
   ```bash
   curl -X POST http://localhost:3030/voca-wallets/v1/user/login \
     -H "Content-Type: application/json" \
     -d '{"username":"john","password":"password123"}'
   ```

3. **Create a wallet**
   ```bash
   curl -X POST http://localhost:3030/voca-wallets/v1/wallet/ \
     -H "Authorization: Bearer <token>" \
     -H "Content-Type: application/json" \
     -d '{"name":"My Wallet"}'
   ```

4. **Deposit funds**
   ```bash
   curl -X POST http://localhost:3030/voca-wallets/v1/wallet/deposit \
     -H "Authorization: Bearer <token>" \
     -H "Content-Type: application/json" \
     -d '{"wallet_id":1,"currency_code":"USD","amount":100}'
   ```

5. **Check balance**
   ```bash
   curl -X GET http://localhost:3030/voca-wallets/v1/wallet/balance/1 \
     -H "Authorization: Bearer <token>"
   ```
