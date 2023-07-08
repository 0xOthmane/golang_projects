

```console
go mod init github.com/0xOthmane/golang_projects/rss_aggregator
go get github.com/joho/godotenv
go mod tidy
go mod vendor
go build && ./rss_aggregator
```

Use of **chi** lightweight, idiomatic and composable router for building Go HTTP services 

```console
go get github.com/go-chi/chi
go get github.com/go-chi/cors
go mod vendor
go mod tidy
go mod vendor
```


Create a Router and a Server (port:8080 defined in .env file).

Define helper functions that write an HTTP response with: status code, JSON body, content type: `application/json`.

Test the helper function with sending a request (GET/POST/..) to http://localhost:8080/v1/healthz. Or Specify what type of request you allow by changing `v1Router.HandleFunc()` to `v1.Get()`