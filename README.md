# gocal - Example Golang API w/ Google Calendar integration

## Getting Started

First and foremost you will need to [enable the Google Calendar API](https://developers.google.com/calendar/quickstart/go) and download the generated `credentials.json` file to the parent of this repo.

After doing so, running the app for the first time will provide a URL that you must open in your browser. After accepting what is presented, a code will be generated to paste back into the command line that will generate a `token.json` file for authentication.

### Running locally

```sh
go get
go run gocal.go
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
