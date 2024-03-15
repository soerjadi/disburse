# SIMPLE GO MOCKING REST API
A RESTful API example that hit mocking endpoint using **gorilla/mux**

## RUN
before run this project please make sure copy this config file.
```bash
$ cp files/config-example.ini files/config.ini
```
change the database host, username, password with your local configuration. After that please execute this to apply the database schema
```bash
$ psql {database schema} < script/schema.sql
```
to run this project please execute this command
```bash
$ cd cmd/rest
$ go run app.go
```

## Structure
```
├── cmd
│   ├── rest
│   │   └── app.go          // main app
├── files
│   └── config-example.go   // configuration example
├── internal
│   ├── config              // Configuration handler
│   ├── delivery            // Our API core handlers
│   │   └── transaction     // transaction api handler. domain based
│   ├── mocks               // Mock 
│   ├── model               // Model for this application
│   ├── pkg                 // Internal package helper
│   ├── repository          // Repository layer
│   │   └── transaction     // transaction repository. domain based
│   └── usecase             // Usecase layer
│   │   └── transaction     // transaction usecase. domain based
└── script                  // other script like sql, etc
```

## API
#### /transaction/account/check
* `POST` : account validation 

request body
```json
{
    "account_number": "989898",
    "bank_name": "BCA"
}
```

#### /transaction/disburse
* `POST` : transfer/disburse to specific account

request body
```json
{
    "unique_number": 123,
    "account": {
        "id": "7cef7d6d-dae1-42db-a6ae-d3005c9c5636",
        "name": "Kristin Bednar",
        "account_number": "10384770",
        "origin": "Mandiri"
    },
    "amount": 5000
}
```

#### /transaction/callback
* `POST` : handle callback after transfer done

request body
```json
{
    "transaction_id": "4753ff4329b37f3ffacf0700b341567d18d3e767adf4a9910ff4422eb63db540", 
    "status": "success"
}
```

## MOCKING ENDPOINT

#### https://1zeq3.wiremockapi.cloud/v8/vbank/transfers
* `POST`   

request body
```json
{
    "account": {
        "id": "7cef7d6d-dae1-42db-a6ae-d3005c9c5636",
        "account_number": "10384770",
        "name": "Kristin Bednar",
        "origin": "Mandiri"
    },
    "amount": 5000
}
```

#### https://1zeq3.wiremockapi.cloud/v8/vbank/search/account
* `POST`

request body
```json
{    
    "account_number": "10384770"
}
```