# VocaGame E-Wallet System

A secure, multi-currency digital wallet system built with Go and PostgreSQL, featuring real-time currency conversion, JWT authentication, and comprehensive transaction management.

## 🚀 Features

- **Multi-Currency Support**: USD, EUR, JPY, IDR with real-time conversion
- **Secure Authentication**: JWT tokens with RSA256 signing
- **Transaction Safety**: Database transactions with pessimistic locking
- **Clean Architecture**: Dependency injection and separation of concerns
- **RESTful API**: Comprehensive endpoints for wallet management
- **Real-time Exchange**: Currency conversion with configurable rates

## 📋 Table of Contents

- [Architecture](#architecture)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Configuration](#configuration)
- [Running the Application](#running-the-application)
- [Testing](#testing)
- [API Documentation](#api-documentation)
- [Database Schema](#database-schema)
- [Design Decisions](#design-decisions)
- [Security Features](#security-features)
- [Deployment](#deployment)

## 🏗️ Architecture

The application follows Clean Architecture principles with the following layers:

```
cmd/http/               # Application entry point
├── main.go
internal/
├── inbound/            # HTTP handlers and controllers
│   ├── http/
│   └── model/
├── outbound/           # Data persistence layer
│   └── repository/
└── usecase/            # Business logic layer
    ├── model/
    └── shared/
pkg/
├── config/             # Configuration management
├── resource/           # External dependencies
└── utils/              # Utility functions
```

### Dependency Flow
```
HTTP → Controller → Use Case → Repository → Database
```

## 📦 Prerequisites

- **Go**: 1.21 or higher
- **PostgreSQL**: 12 or higher
- **Docker & Docker Compose**: For containerized setup
- **Make**: For build automation

## 🛠️ Installation

### Option 1: Local Development

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd vocagame-be-interview
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up PostgreSQL database**
   ```bash
   # Create database
   createdb vocagame_wallet

   # Or using psql
   psql -U postgres -c "CREATE DATABASE vocagame_wallet;"
   ```

4. **Run database migrations**
   ```bash
   make migrate-up
   ```

### Option 2: Docker Setup

1. **Clone and build**
   ```bash
   git clone <repository-url>
   cd vocagame-be-interview
   ```

2. **Start with Docker Compose**
   ```bash
   docker-compose up --build
   ```

## ⚙️ Configuration

### Environment Variables

Create a `.env` file in the root directory:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=vocagame_wallet
DB_SSL_MODE=disable

# Server Configuration
SERVER_HOST=localhost
SERVER_PORT=3030

# JWT Token Configuration
ACCESS_TOKEN_PRIVATE_KEY=LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0t...
ACCESS_TOKEN_PUBLIC_KEY=LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0K...
ACCESS_TOKEN_EXPIRED_IN=15m
ACCESS_TOKEN_MAXAGE=900

REFRESH_TOKEN_PRIVATE_KEY=LS0tLS1CRUdJTiBQUklWQVRFIEtFWS0tLS0t...
REFRESH_TOKEN_PUBLIC_KEY=LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0K...
REFRESH_TOKEN_EXPIRED_IN=168h
REFRESH_TOKEN_MAXAGE=604800

TOKEN_ISSUER=vocagame-wallet
TOKEN_AUDIENCE=vocagame-users
```

### Generate RSA Keys

```bash
# Generate private key
openssl genrsa -out access_token_private.pem 2048

# Generate public key
openssl rsa -in access_token_private.pem -pubout -out access_token_public.pem

# Convert to base64 for .env
base64 -i access_token_private.pem | tr -d '\n'
base64 -i access_token_public.pem | tr -d '\n'
```

## 🚀 Running the Application

### Development Mode

```bash
# Start the application
make run

# Or with hot reload (if you have air installed)
air
```

### Production Mode

```bash
# Build the application
make build

# Run the binary
./bin/app
```

### Available Make Commands

```bash
make build          # Build the application
make run             # Run in development mode
make test            # Run all tests
make test-coverage   # Run tests with coverage report
make migrate-up      # Run database migrations
make migrate-down    # Rollback migrations
make docker-build    # Build Docker image
make docker-run      # Run in Docker container
```

## 🧪 Testing

### Unit Tests

```bash
# Run all tests
make test

# Run specific package tests
go test ./internal/usecase/user/...

# Run with coverage
make test-coverage
```

### Integration Tests

```bash
# Run integration tests
go test ./tests/integration/...
```

### Test Structure

- **Unit Tests**: Business logic validation in `*_test.go` files
- **Integration Tests**: End-to-end API testing in `tests/integration/`
- **Mocks**: Generated mocks for repositories and external dependencies

## 📚 API Documentation

Comprehensive API documentation is available at [`docs/API.md`](docs/API.md).

### Quick Start

1. **Register a user**
   ```bash
   curl -X POST http://localhost:3030/voca-wallets/v1/user/register \
     -H "Content-Type: application/json" \
     -d '{"username":"john","email":"john@test.com","password":"password123"}'
   ```

2. **Login and get token**
   ```bash
   curl -X POST http://localhost:3030/voca-wallets/v1/user/login \
     -H "Content-Type: application/json" \
     -d '{"username":"john","password":"password123"}'
   ```

3. **Create wallet and deposit funds**
   ```bash
   # Create wallet
   curl -X POST http://localhost:3030/voca-wallets/v1/wallet/ \
     -H "Authorization: Bearer <token>" \
     -H "Content-Type: application/json" \
     -d '{"name":"My Wallet"}'

   # Deposit funds
   curl -X POST http://localhost:3030/voca-wallets/v1/wallet/deposit \
     -H "Authorization: Bearer <token>" \
     -H "Content-Type: application/json" \
     -d '{"wallet_id":1,"currency_code":"USD","amount":100}'
   ```

## 🗄️ Database Schema

### Core Entities

```sql
-- Users
CREATE TABLE users (
    id BIGSERIAL PRIMARY KEY,
    username VARCHAR(255) UNIQUE NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash VARCHAR(255) NOT NULL
);

-- Wallets
CREATE TABLE wallets (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT REFERENCES users(id),
    name VARCHAR(255) NOT NULL
);

-- Currencies
CREATE TABLE currencies (
    id BIGSERIAL PRIMARY KEY,
    currency_code VARCHAR(3) UNIQUE NOT NULL,
    currency_name VARCHAR(100) NOT NULL
);

-- Wallet Balances
CREATE TABLE wallet_balances (
    id BIGSERIAL PRIMARY KEY,
    wallet_id BIGINT REFERENCES wallets(id),
    currency_id BIGINT REFERENCES currencies(id),
    balance DECIMAL(20,8) DEFAULT 0
);

-- Transactions
CREATE TABLE transactions (
    id BIGSERIAL PRIMARY KEY,
    transaction_id VARCHAR(255) UNIQUE NOT NULL,
    wallet_id BIGINT REFERENCES wallets(id),
    currency_id BIGINT REFERENCES currencies(id),
    amount DECIMAL(20,8) NOT NULL,
    type VARCHAR(50) NOT NULL,
    description TEXT
);

-- Exchange Rates
CREATE TABLE exchange_rates (
    id BIGSERIAL PRIMARY KEY,
    from_currency_id BIGINT REFERENCES currencies(id),
    to_currency_id BIGINT REFERENCES currencies(id),
    rate DECIMAL(20,8) NOT NULL
);
```

## 🎯 Design Decisions

### 1. **Clean Architecture**
- **Benefit**: Maintainable, testable, and scalable codebase
- **Implementation**: Separated concerns with clear boundaries between layers
- **Trade-off**: Slightly more complex initial setup for better long-term maintainability

### 2. **Dependency Injection with Dig**
- **Benefit**: Loose coupling and easier testing
- **Implementation**: Container-based dependency management
- **Trade-off**: Learning curve for DI patterns

### 3. **Database Transaction Management**
- **Benefit**: ACID compliance and data consistency
- **Implementation**: Pessimistic locking for concurrent operations
- **Trade-off**: Potential performance impact under high concurrency

### 4. **JWT Authentication with RSA256**
- **Benefit**: Stateless authentication and enhanced security
- **Implementation**: Asymmetric key signing for token verification
- **Trade-off**: Key management complexity

### 5. **Multi-Currency Design**
- **Benefit**: Support for international transactions
- **Implementation**: Separate currency and exchange rate tables
- **Trade-off**: Additional complexity in balance calculations

### 6. **Decimal Precision for Financial Data**
- **Benefit**: Accurate financial calculations
- **Implementation**: DECIMAL(20,8) for amounts and rates
- **Trade-off**: Larger storage requirements

## 🔒 Security Features

### Authentication & Authorization
- JWT tokens with RS256 signing
- Access and refresh token pattern
- Token expiration and rotation
- Password hashing with bcrypt

### Input Validation
- Request payload validation
- SQL injection prevention with GORM
- Amount validation for financial operations

### Transaction Security
- Database transactions for consistency
- Pessimistic locking for critical operations
- Balance verification before transactions

## 📞 Support

For support and questions:

- **Issues**: GitHub Issues
- **Documentation**: [`docs/API.md`](docs/API.md)
- **Contact**: [asnurramdhani12@gmail.com]


**Built with ❤️ for VocaGame Technical Interview**
