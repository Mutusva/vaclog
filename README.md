# Vaccination Log (Vaclog)

This application is used to track vaccination records for farm animal

## Running the application

### Requirements
- `go 1.19 or latest` or `docker running on your machine`

- Export Environment variable
```
Environment variable:
 APPLICATION_HOST=0.0.0.0
 APPLICATION_PORT=8080
 IMMUDB_API_KEY=<your IMMUDB API Key>
 # this depends on your collection
 IMMUDB_BASE_URL=https://vault.immudb.io/ics/api/v1/ledger/default/collection

```

Run the following command:
```command
  go run ./cmd/vaclog/main.go
```

- you can now send request to the endpoints define in the swagger documentaion which can be viewed 
  by copying the contents of the swagger.json/ swagger.yaml and paste on [swagger editor](https://editor.swagger.io/)
- The following endponts are defined
```endpoints
	GetAllRecords => /api/v1/records [GET]
	CreateRecord  => /api/v1/records [PUT]
	SearchRecord  => /api/v1/records/:documentID [GET]
```

### Using Docker
A docker file is include in this repo:
Run the following commands:
```docker
  docker build -t <your tag> .
  docker run -p 8080:8080 <your tag>
```

### Makefile
- Some useful commands are defined in the Makefile

Endpoint documentation found here: [swag docs](docs%2Fswag)