# GIN JWT SAMPLE

Sample jwt implementation in golang.

## Deps

- `gin` : web framework
- `golang-jwt` : jwt lib for golang
- `gorm` : database orm
- `logrus` : logging formatter

## Setup

- `make environment` : setup environment
- `make run` : running application

## Endpoints

1. Signup

```bash
curl --location 'http://host:port/signup' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "example@youremail.com",
    "password": "123456"
}'
```

2. Signin

```bash
curl --location 'http://host:port/signin' \
--header 'Content-Type: application/json' \
--data-raw '{
    "email": "example@youremail.com",
    "password": "123456"
}'
```

3. Profile

```bash
curl --location 'http://host:port/me' \
--header 'Authorization: eyJhbGciOiJIUzI.eyJleHAiOjE2.pE858hCGksgx'
```
