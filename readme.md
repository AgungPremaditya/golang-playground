# Movie & Series API with GOLANG

This might be not the best golang code you've seen, because of this is **my first API that I made with golang**. This API is Build Using

- [GoLang v1.19.1](https://go.dev/)
- [gin-gonic](https://gin-gonic.com/)
- [MySQL v8.0](https://www.mysql.com/)

## Development

First you can clone this repository. Then install all package that needed for this API. For installing all package you can use command below:

```sh
go get
```

After installing all package, make sure you already create databases for this API.
**To connect the database with this app you can change database config on file `./controllers/connection.go [on line 11]`.**
Then you can migrate your databases using command below:

```sh
migrate -database "mysql://user:password@tcp(localhost:3306)/movie_api" -path db/migrations up
```

For Running this applications you can run command below:

```sh
go run main.go
```

or you can use [air](https://github.com/cosmtrek/air) for hot-reload. You can use air with installing it on your $gopath/bin. For installing using go, you can follow command below:

```sh
go install github.com/cosmtrek/air@latest
```

Then for run your aplications in hot-reload mode use this command below in your golang application directory.

```sh
air
```

**Notes : Please make sure your $gopath/bin is registered on your local environments**
