# Go-Captcha API

Server to generate and check captcha

This Captcha Generation is using "Inter Font© SIL OpenFontLicense 1.1"

## Requirements

- Redis Server

## How to Use

As example, you can start the go-captcha API like that:

```bash
export API_PORT=10888
export API_BIND=0.0.0.0
export LOGLEVEL=debug
export REDIS_SERVER=localhost:6379

go run init.go app.go
```

If you want to generate a captcha:

```bash
curl -X GET localhost:10888/api/captcha/v0
```

It will give you as result a PNG File and in the HTTP Response Header a "sessiontoken". This token you have to add to your CaptchaCheck Request:

```bash
curl -H 'sessiontoken: <token>' -X POST localhost:10888/api/captcha/v0/<token>
```

Thats it! Pretty easy.