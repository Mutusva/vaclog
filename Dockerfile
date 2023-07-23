FROM golang:1.19-alpine

ENV APPLICATION_HOST=0.0.0.0
ENV APPLICATION_PORT=8080
# Override this, can be store in a secret store or vault
ENV IMMUDB_API_KEY=
# Replace with Url to point to your collection
ENV IMMUDB_BASE_URL=https://vault.immudb.io/ics/api/v1/ledger/default/collection

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN  CGO_ENABLED=0 GOOS=linux go build  -o ./vaclog ./cmd/vaclog/main.go

EXPOSE 8080

# Run
CMD ["/app/vaclog"]