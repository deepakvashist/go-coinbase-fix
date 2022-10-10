# go-fix

Golang sample project for testing [FIX protocol](http://www.fixprotocol.org/) APIs.

## Usage

1) Read Coinbases [FIX API](https://docs.cloud.coinbase.com/exchange/docs/connectivity) documentation

2) Read [Go quickfix](https://github.com/quickfixgo/quickfix) documentation

3) Install [Stunnel](https://www.stunnel.org/) to creating a SSL tunnel (Local connection vs Coiunbase FIX API)

```shell
brew install stunnel
```

4) Download Coinbase FIX API TLS certificate:

```shell
openssl s_client -showcerts -connect fix-public.sandbox.exchange.coinbase.com:4198 < /dev/null | openssl x509 -outform PEM > coinbase.pem
```

5) Export the required env vars:

```shell
export COINBASE_API_KEY=""
export COINBASE_API_KEY_PASSPHRASE=""
export COINBASE_API_KEY_SECRET=""
```

6) Create FIX configuration file and update fiel values based on Coninbase API key:

```shell
cp config/client.cfg.template config/client.cfg
```

7) Start Stunnel:

```shell
stunnel config/stunnel.conf
```

8) Install application dependencies:

```shell
go mod install
```

9) Run application:

```shell
go run ./cmd/app.go
```
