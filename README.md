# go-fix

Golang sample project for testing [FIX protocol](http://www.fixprotocol.org/) APIs.

## Usage

1) Read Coinbases [FIX API](https://docs.cloud.coinbase.com/exchange/docs/connectivity) documentation

2) Read [Go quickfix](https://github.com/quickfixgo/quickfix) documentation

3) Export the required env vars:

```shell
export COINBASE_API_KEY=""
export COINBASE_API_KEY_PASSPHRASE=""
export COINBASE_API_KEY_SECRET=""
```

4) Create FIX configuration file and update fiel values based on Coninbase API key:

```shell
cp config/client.cfg.template config/client.cfg
```

5) Install application dependencies:

```shell
go mod install
```

6) Run application:

```shell
go run ./cmd/app.go
```
