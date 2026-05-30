# Projector
Projector is a backend REST API that powers our frontend for managing client orders. Authenticated users can perform CRUD operations for manipulating orders stored within a lightweight, central database.

## Installation
### Prerequisites
- Go
- Docker

## Technical Details
Projector was built using the Gin REST library, and relies on SQLite as the database, coupled with GORM as the ORM for handling model definitions and translations to and from the database.

### Running Standalone
Clone the repository.
```
git clone https://github.com/westonpgreene/projector.git
cd projector
```

Install dependencies.
```
go get .
```

Launch the internal Gin server.
```
go run main.go
```

### Using Docker
Build the Projector container image.
```
docker build -t projector .
```

Run the Project container.
```
docker run -p 5000:5000 -v $(pwd)/data:/app/data projector
```

## Usage
Projector can be used to create, retrieve, update, and delete the orders we have stored in the database, implemented using SQLite. In production, we would likely opt to use a more robust technology, likely PostgreSQL.

> If I were pushing this to AWS, I would leverage DynamoDB. But for the purposes of this example, we used SQLite.

1. Generate an API token that you can use to authenticate yourself.
```
curl -X POST http://localhost:5000/keys \
  -H "Content-Type: application/json" \
  -d '{"label": "MY_EXAMPLE_API_KEY"}'
```

This will output your key and timestamp of creation:
```
{
  "created_at": "2026-05-29T21:36:22.263451206-04:00",
  "key": "projector_bb5513c022785970c11c9aa50afcbc3376c055000c7d3ab483d3565e6a29b74e",
  "label": "test"
}
```

Use this key to authenticate all of your requests to the API.

### Creating Orders
```
curl -X POST http://localhost:5000/orders \
  -H "Content-Type: application/json" \
  -H "X-API-Key: projector_bb5513c022785970c11c9aa50afcbc3376c055000c7d3ab483d3565e6a29b74e" \
  -d '{"client_name":"Acme Corp","project_type":"Web","delivery_date":"2026-12-01T00:00:00Z"}'
```

The output from this example:
```
{
  "id": "eeba69b0-80c7-4af0-8904-b260ed307fcc",
  "client_name": "Acme Corp",
  "project_type": "Web",
  "delivery_date": "2026-12-01T00:00:00Z",
  "status": "Pending",
  "created_at": "2026-05-28T21:38:48.691230315-04:00",
  "updated_at": "2026-05-28T21:38:48.691230315-04:00"
}
```

### Retrieving Orders
```
curl -s http://localhost:5000/orders -H "X-API-Key: projector_bb5513c022785970c11c9aa50afcbc3376c055000c7d3ab483d3565e6a29b74e"
```

The output from this example:
```
[
  {
    "id": "eeba69b0-80c7-4af0-8904-b260ed307fcc",
    "client_name": "Acme Corp",
    "project_type": "Web",
    "delivery_date": "2026-12-01T00:00:00Z",
    "status": "Pending",
    "created_at": "2026-05-29T21:38:48.691230315-04:00",
    "updated_at": "2026-05-29T21:38:48.691230315-04:00"
  }
]
```

### Updating Orders
```
curl -s -X PUT http://localhost:5000/orders/eeba69b0-80c7-4af0-8904-b260ed307fcc \
  -H "Content-Type: application/json" \
  -H "X-API-Key: projector_bb5513c022785970c11c9aa50afcbc3376c055000c7d3ab483d3565e6a29b74e" \
  -d '{"status":"In Progress"}'
```

The output from this example:
```
{
  "id": "eeba69b0-80c7-4af0-8904-b260ed307fcc",
  "client_name": "Acme Corp",
  "project_type": "Web",
  "delivery_date": "2026-12-01T00:00:00Z",
  "status": "In Progress",
  "created_at": "2026-05-29T21:38:48.691230315-04:00",
  "updated_at": "2026-05-29T21:38:48.691230315-04:00"
}
```

### Deleting Orders
```
curl -s -X DELETE http://localhost:5000/orders/eeba69b0-80c7-4af0-8904-b260ed307fcc \
  -H "X-API-Key: projector_bb5513c022785970c11c9aa50afcbc3376c055000c7d3ab483d3565e6a29b74e"
```

The output from this example:
```
{
  "message": "order deleted"
}
```
