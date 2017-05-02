# apstlib

app store data capture tool

## Required

- go 1.8
- GCP Account and Project
- Cloud SDK

# How to Build

## Install

```console
$ go get github.com/itsubaki/apstlib
```

## Make GAE Application

```go
// main.go
package example

import (
    "github.com/itsubaki/apstlib"
)

func init() {
    apstlib.AppEngine()
}
```

```yaml
# app.yaml
runtime: go
api_version: go1

handlers:
- url: /.*
  script: _go_app
```

```yaml
# cron.yaml
- description: "Capture Review"
  url: /review/capture?id=${IOS_APP_ID}
  schedule: every 1 hours
- description: "Capture Ranking"
  url: /ranking/capture
  schedule: every 24 hours
```

## Deploy

```console
$ ls
app.yaml cron.yaml main.go
$ gcloud app deploy app.yaml --project ${PROJECT_ID}
```

## API Example

### Ranking

```console
$ curl "https://${PROJECT_ID}.appspot.com/ranking"
$ curl "https://${PROJECT_ID}.appspot.com/ranking/search?id=${IOS_APP_ID}"
$ curl "https://${PROJECT_ID}.appspot.com/ranking/search?query=${IOS_APP_ARTIST}"
```

### Review

```console
$ curl "https://${PROJECT_ID}.appspot.com/review"
$ curl "https://${PROJECT_ID}.appspot.com/review/search?id=${IOS_APP_ID}&query=Rating:5"
```
