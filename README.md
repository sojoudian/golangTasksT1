#CRUD API for User management

Run:
go run main.go

Tests:

GET method:```curl -v localhost:3003/users/ -X GET ```

POST method:``` curl -v localhost:3003/users/ -X POST -d '{"first_name": "Albus Percival Wulfric Brian", "last_name": "Dumbledore", "password": "Sherbet lemon"}' ```

DELETE method ``` curl -v localhost:3003/users/1 -X DELETE ```

Then use another get request to confirm delete process

