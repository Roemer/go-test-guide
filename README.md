# go-test-guide
A library / cmd to interact with Test.Guide

## Installation

```
go get github.com/roemer/go-test-guide
```

## Usage

Create a client:
```go
client, err := gotestguide.NewClient("server-url", "token")
if err != nil {
    log.Fatalf("Failed to create client: %v", err)
}
```
