# Go CRUD Backend Project

This repository contains a demonstration of Go programs with CRUD features using ElephantSQL Database. The project was created with MVC design pattern.

## Technologies covered:

- Go language for backend server
- Gorm
- Gin
- CompileDaemon
- ElephantSQL
- MVC Design Pattern
- REST API

## Author

Marsetio Noorprajuda,

Github: https://github.com/noorprajuda

## How to run the program

```
$ git clone "this go-crud project URL"
$ cd go-crud
$ go build
```

Then use Postman to exercise the REST API Features :

# Create post

```
POST http://localhost:3000/posts
```

Req body

```
{
    Title: string,
    Body: string,
}
```

Response

```
{
    "post": {
        "ID": 3,
        "CreatedAt": "2024-03-18T12:09:31.367693+01:00",
        "UpdatedAt": "2024-03-18T12:09:31.367693+01:00",
        "DeletedAt": null,
        "Title": "from req body2",
        "Body": "from req body2"
    }
}
```

# Read posts

```
GET http://localhost:3000/posts
```

Response

```
{
    "posts": [
        {
            "ID": 1,
            "CreatedAt": "2024-03-18T09:50:30.150008+01:00",
            "UpdatedAt": "2024-03-18T09:50:30.150008+01:00",
            "DeletedAt": null,
            "Title": "FirstPost",
            "Body": "Post body"
        },
        {
            "ID": 3,
            "CreatedAt": "2024-03-18T12:09:31.367693+01:00",
            "UpdatedAt": "2024-03-18T12:09:31.367693+01:00",
            "DeletedAt": null,
            "Title": "from req body2",
            "Body": "from req body2"
        }
    ]
}
```

# Read post

```
GET http://localhost:3000/posts/:id
```

Response

```
{
    "post": {
        "ID": 3,
        "CreatedAt": "2024-03-18T12:09:31.367693+01:00",
        "UpdatedAt": "2024-03-18T12:09:31.367693+01:00",
        "DeletedAt": null,
        "Title": "from req body2",
        "Body": "from req body2"
    }
}
```

# Update post

```
PUT http://localhost:3000/posts/:id
```

Req body

```
{
    Title: string,
    Body: string,
}
```

Response

```
{
    "post": {
        "ID": 3,
        "CreatedAt": "2024-03-18T12:09:31.367693+01:00",
        "UpdatedAt": "2024-03-18T12:13:20.19769+01:00",
        "DeletedAt": null,
        "Title": "Updated Title",
        "Body": "Updated Body"
    }
}
```

# Delete post

```
DELETE http://localhost:3000/posts/:id
```

Response

```
{
    "message": "Post with id 2 has been deleted"
}
```

created on Mon, 18th March 2024 © Marsetio Noorprajuda. All rights reserved.
