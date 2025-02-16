# easywebstats
![Static Badge](https://img.shields.io/badge/License-MIT-blue)

Stat-collecting gin gonic web service with a simple API.



## Building from source

### Prerequisits
- [Go](https://go.dev/doc/install) (v1.24+)
- [Swaggo/swag](https://github.com/swaggo/swag?tab=readme-ov-file#getting-started) (generation of swagger documentation)

Clone the repo
```shell
git clone git@github.com:r59q/easywebstats.git && cd easywebstats
```

Install dependencies
```shell
go install
```
Generate swagger docs, using swag *(Optional, already included after cloning)*
```shell
swag init
```
Either run it directly
```shell
go run .
```
Swagger docs available at http://localhost:8080/swagger/index.html

**or** build the binary
```shell
go build .
```
and run it
```shell
./easywebstats
```