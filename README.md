# To-do API 
A to-do API for all your reminders 

# Stack of the project:
- Golang with Gin Gonic framework
- SQLite as the DB
- Gorm as the ORM
- MVC pattern
<br>

# Installation
To use the API on windows, you can simply run the go-to-do.exe file

Or you can just simply run the following command inside the project's folder:

```sh
go mod tidy
```
Then start the project
```sh
go run main.go
```

## To-do API specification

| Task | URL | Method | Response code | Response |
|:----:|:---:|:------:|:-------------:|:--------:|
| Create an entry | / | POST | 201 | Created object |
| Check API health | /health| GET | 200 | API is up |
| Read all entries | / | GET | 200 | All entries |
| Read entries by id | /:id=value | GET | 200 | To-do belonging to that id |
| Read entries by id | /:username=value | GET | 200 | All entries from that username |
| Update to-do | /edit/:id | UPDATE | 200 | status | 
| Delete entries | /:id | DELETE | 200 | status |
