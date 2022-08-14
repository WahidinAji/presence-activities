
# TODO
* `Logout service`
* `View activities by date`
* `View presences by date`
* `Complete the request documentation`

## User Services
## Register Service
* url
```http
http://127.0.0.1:3000/api/users/register
```
* method
`POST`

* body
```json
{
    "name" : "Your Name",
    "email" : "your.email@mail.com",
    "password" : "yourpassword",
}
```

### Login service
* url
```http
http://127.0.0.1:3000/api/users/login
```
* method
`POST`

* body
```json
{
    "email" : "your.email@mail.com",
    "password" : "yourpassword",
}
```

### Logout service
* Authentication
`Bearer Token : <token>`
* url
```http
http://127.0.0.1:3000/api/users/login
```
* method
`POST`
* body
```json
{
    "user_id" : "your.email@mail.com",
}
```

## Presence Services
### Check In Service
* Authentication
`Bearer Token : <token>`
* url
```http
http://127.0.0.1:3000/api/presences
```
* method
`POST`

* body
```json
{
    "user_id" : "your.email@mail.com",
    "status" : "check_in | check_out",
}
```

