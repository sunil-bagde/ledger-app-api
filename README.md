# Ledger app

**Ledger simplifies financial management by offering features for budgeting, expense tracking, bill management, and reporting.**

### Setup a DB via docker db

```sh
docker run -itd -e POSTGRES_USER=ledgerapp -e POSTGRES_PASSWORD=ledgerapp -p 5432:5432  --name postgresql postgres
```

### Migrations

```sh
make migrate
```

### Seed

```sh
make seed
```

### `go run main.go` or `go run .`

Runs the app in development mode.<br>
Open [http://localhost:3000](http://localhost:3000) to view it in the browser.

Design a system called a ledger system

- [x] Deposit Money
- [x] Withdraw Money
- [x] Transfer money from one account to another account
- [x] Transfer money from one account of one user to another user
- [] Check Balance
- [] Get Balance
- [] Transaction history
- [] Multi currency

### leader should have

### App strcuture

```
.
├── core
│   ├── account
│   ├── transaction
│   └── user
├── database
├── handler
│   ├── account
│   ├── transaction
│   └── user
├── migrations
├── scripts
│   ├── migrate
│   └── seed
├── store
│   ├── account
│   ├── transaction
│   └── user
└── tmp

```
