### To-do API 
To-do api specification

<br>

| Task | URL | Method | Response code | Response |
|:----:|:---:|:------:|:-------------:|:--------:|
| Check API health | /health| GET | 200 | API is up |
| Create an entry | / | POST | 201 | Created object | 
| Read all entries | / | GET | 200 | All entries |
| Read entries by id | /?id=value | GET | 200 | To-do belonging to that id |
| Read entries by id | /?username=value | GET | 200 | All entries from that username | 
| Delete entries | /:id | DELETE | 200 | status |
