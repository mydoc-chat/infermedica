# Infermedica (unofficial)
Go interface to the infermedica REST API

## Description

This is a Go interface to the Infermedica REST API: https://developer.infermedica.com/docs/api

## Installation

```go get github.com/torniker/infermedica```

## Usage examples

#### Fetching symptoms
```go
app := infermedica.NewApp("appid", "appkey", "model")
symptoms, err := app.Symptoms()
if err != nil {
    log.Errorf("Could not fetch symptoms: %v", err)
}
log.Infof("All Symptoms: %v", symptoms)
```


