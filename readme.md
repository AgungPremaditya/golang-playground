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
