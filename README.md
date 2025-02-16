# easywebstats
![Static Badge](https://img.shields.io/badge/License-MIT-blue)

Lightweight concurrent stat-collecting gin gonic web service with a simple API.

## Registering statistics

Registering statistics is automatically done when settings values. No initialization is needed, just set the values you wish to track.

Data values are grouped by name and label.

## Reading statistics

When stats are registered using a name and label, you can either read all stats for a given name, or a specific value for a given label and name.

## Examples
### Registering statistics
#### Set numeric value
```shell
# Curl example
curl -X 'POST' \
  'http://localhost:8080/api/v0/register/num/set' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "number_of_visits",
  "label": "about_page",
  "value": 10
}'
# Outputs 10 (Returns the number of visits)
```

#### Increase numeric value
Increase the number of visits on the about page by 1
```shell
# Curl example
curl -X 'POST' \
  'http://localhost:8080/api/v0/register/num/increase' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "number_of_visits",
  "label": "about_page",
  "value": 1
}'
# Outputs 11 (Returns the updated number of visits)
```

### Reading statistics
#### Read numeric values by label
```shell
# Curl example
curl -X 'GET' \
  'http://localhost:8080/api/v0/read/num/number_of_visits/about_page' \
  -H 'accept: application/json'

# Outputs 10
```

#### Read numeric values
```shell
# Curl example
curl -X 'GET' \
  'http://localhost:8080/api/v0/read/num/number_of_visits' \
  -H 'accept: application/json'

# Outputs
# {
#   "about_page": 10,
#   "home_page": 321
# }
```

## Building from source

### Prerequisits
- [Go](https://go.dev/doc/install) v1.24
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