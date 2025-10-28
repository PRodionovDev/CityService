FROM golang:1.25.2
WORKDIR /app
COPY ../go.mod ./
RUN go mod download
RUN go get -u github.com/gin-gonic/gin
RUN go get -u gorm.io/gorm
RUN go get -u gorm.io/driver/sqlite
RUN go get github.com/joho/godotenv
RUN go get github.com/gosimple/slug
COPY .. .
RUN go build -o cmd/main ./cmd/
EXPOSE 8080
CMD ["./cmd/main"]