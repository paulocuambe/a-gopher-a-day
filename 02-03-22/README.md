The goal was to create a very simple API. No need for any kind of setup, just run `go run .` inside this directory and you are good to go.

If you want to change the port in which the app runs, run `export APP_PORT=<PORT_NUMBER>` with your preferred port number before starting the application. The default port is `8080`.

Available endpoints:
* GET `/api`
* GET `/api/users`
* POST `/api/users` 
    * Body: `{username: "username", name: "name"}`


Responses look like this:  
GET `/api`
```json
{
	"name": "Simple Api",
	"version": "0.0.1",
	"author": "paulocuambe"
}
```

GET `/api/users`
```json
[{
	"id": 1677786244,
	"name": "hello",
	"username": "there",
	"created_at": "2023-03-02T21:44:04.726122856+02:00",
	"updated_at": "2023-03-02T21:44:04.726123004+02:00"
}]
```