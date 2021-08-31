### To-do API 
To-do api specification

<br>

| Task | URL | Method | Response code | Response |
|:----:|:---:|:------:|:-------------:|:--------:|
| Create an entry | / | POST | 201 | Created object |
| Check API health | /health| GET | 200 | API is up |
| Read all entries | / | GET | 200 | All entries |
| Read entries by id | /:id=value | GET | 200 | To-do belonging to that id |
| Read entries by id | /:username=value | GET | 200 | All entries from that username |
| Update to-do | /edit/:id | UPDATE | 200 | status | 
| Delete entries | /:id | DELETE | 200 | status |
