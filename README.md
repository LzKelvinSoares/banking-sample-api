# Banking Sample API

A Go HTTP API for customer lookup, account creation, and account transactions. The service uses SQL Server by default and delegates token authorization to a separate authentication API.

## Requirements

- Go 1.27.1 or later
- SQL Server
- An authentication API compatible with the `/auth/verify` endpoint

## Configuration

The application loads a `.env` file from the working directory. Create one with the following variables and replace the example values with your local configuration:

```dotenv
SERVER_ADDRESS=0.0.0.0
SERVER_PORT=8080
DB_USER=your_database_user
DB_PASSWORD=your_database_password
DB_ADDRESS=localhost
DB_PORT=1433
DB_NAME=banking
AUTH_API=localhost:8081
```

`AUTH_API` is the host and optional port of the authentication service. This API calls `http://<AUTH_API>/auth/verify` for each request. Requests to the banking API must include an `Authorization` header containing a token, for example `Authorization: Bearer YOUR_TOKEN`.

## Database setup

Create the `banking` database in SQL Server, then run [`resources/sqlserver/01-banking.sql`](resources/sqlserver/01-banking.sql) against it to create and seed the tables. **The script drops existing banking tables before recreating them**, so do not run it against data you need to keep.

The SQL Server schema seeds sample customers and accounts. The SQL scripts under `resources/mysql` are also provided, but the application currently selects SQL Server when it starts.

## Run locally

From the repository root:

```powershell
go mod download
go run .
```

The server listens at the address composed from `SERVER_ADDRESS` and `SERVER_PORT` (for example, `http://localhost:8080`).

Run the test suite with:

```powershell
go test ./...
```

## Run with Docker

Build the image:

```powershell
docker build -t banking-sample-api .
```

Run it with the environment file and a reachable SQL Server/auth service:

```powershell
docker run --rm --env-file .env -p 8080:8080 banking-sample-api
```

The container exposes port `8080`; set `SERVER_PORT=8080` when using this command.

## API

All routes require a valid authorization token. JSON responses use `Content-Type: application/json`.

| Method | Path | Description |
| --- | --- | --- |
| `GET` | `/customers` | List customers |
| `GET` | `/customers?status=active` | List customers filtered by `active` or `inactive` status |
| `GET` | `/customers/{id}` | Get one customer |
| `POST` | `/customers/{id}/account` | Create an account for a customer |
| `POST` | `/customers/{id}/account/{account_id}/{transaction_type}` | Deposit to or withdraw from an account |

### List customers

```bash
curl -H "Authorization: Bearer YOUR_TOKEN" \
  "http://localhost:8080/customers?status=active"
```

Use `status=active` or `status=inactive`; omit the parameter to list all customers.

### Get a customer

```bash
curl -H "Authorization: Bearer YOUR_TOKEN" \
  http://localhost:8080/customers/2000
```

Customer responses contain `customer_id`, `full_name`, `city`, `zipcode`, `date_of_birth`, and `status`.

### Create an account

The account type must be `checking` or `saving`, and the initial amount must be at least `5000`.

```bash
curl -X POST \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"account_type":"saving","amount":5000}' \
  http://localhost:8080/customers/2000/account
```

### Make a transaction

The transaction type in the URL must be `deposit` or `withdrawal`. The amount must not be negative.

```bash
curl -X POST \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{"amount":250}' \
  http://localhost:8080/customers/2000/account/95470/deposit
```

Successful account creation returns an `account_id`. A successful transaction returns a `transaction_id` and the resulting `balance`.
