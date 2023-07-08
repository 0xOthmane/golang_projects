go mod init github.com/0xOthmane/golang_projects/rss_project
go get github.com/joho/godotenv
go mod tidy
go mod vendor
go build && ./rss_aggregator