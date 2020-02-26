# gocal - Example Golang API

## Getting Started

### Running locally

```sh
go get
go run gocal.go
```

or via Docker:

```sh
docker-compose up
```

#### Example Request

```sh
curl --location --request POST 'http://localhost:3000/api/classes/' \
--header 'Content-Type: text/plain' \
--data-raw '{
  "type": {
    "id": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
    "type": "Training"
  },
  "location": {},
  "room": {
    "id": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
    "name": "Weight Room",
    "capacity": 10,
    "active": true
  },
  "reservations": false,
  "start_datetime": "2020-02-06T15:00:00+0000",
  "end_date": "2020-02-20",
  "recurrence": "string",
  "instructor": {
    "id": "3fa85f64-5717-4562-b3fc-2c963f66afa6",
    "name": "John Wick"
  },
  "attendees": [
    {
      "id": "04b37adf-8170-46c4-9799-d0fcd63af643",
      "name": "Firstname Lastname",
      "email": "firstname.lastname@example.com",
      "checked_in": true
    }
  ],
  "total_attendance": 5
}'
```

### Testing

```sh
go get -t ./...
go test ./...
```
