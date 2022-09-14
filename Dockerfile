FROM golang:1.18

WORKDIR /app
COPY go.sum go.mod ./
RUN go mod download
COPY . .
RUN go build -o /bin/app ./app/cmd/
ENTRYPOINT ["/bin/app"]
