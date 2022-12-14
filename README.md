# Safulet/bitshares-go
[![Go Report Card](https://goreportcard.com/badge/github.com/Safulet/bitshares-go)](https://goreportcard.com/report/github.com/Safulet/bitshares-go)
[![GoDoc](https://godoc.org/github.com/Safulet/bitshares-go?status.svg)](https://godoc.org/github.com/Safulet/bitshares-go)
[![Build Status](https://travis-ci.org/Safulet/bitshares-go.svg?branch=master)](https://travis-ci.org/Safulet/bitshares-go)


Golang RPC (via websockets) client library for [Bitshares](https://bitshares.org/) and [OpenLedger](https://openledger.io) in particular

## Usage

```go
import "github.com/Safulet/bitshares-go"
```

## Example
```go
client, err := NewClient("wss://bitshares.openledger.info/ws")

// retrieve the current global_property_object
props, err := client.Database.GetDynamicGlobalProperties()

// lookup symbols ids
symbols, err := client.Database.LookupAssetSymbols("OPEN.SCR", "USD")
require.NoError(t, err)

openSCR := symbols[0].ID
USD := symbols[1].ID

// retrieve 5 last filled orders
orders, err := client.History.GetFillOrderHistory(openSCR, USD, 5)

// set a block applied callback
client.Database.SetBlockAppliedCallback(func(blockID string, err error) {
    log.Println(blockID)
})

// cancel all callbacks
client.Database.CancelAllSubscriptions()

```
## Status
The project is in active development but should not be used in production yet.

## Supported operations
 - Transfer
 - LimitOrderCreate
 - LimitOrderCancel

