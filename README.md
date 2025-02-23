# easywebstats
![Static Badge](https://img.shields.io/badge/License-MIT-blue) ![TeamCity Full Build Status](https://img.shields.io/teamcity/build/e/Easywebstats_BuildAndDeploy?server=https%3A%2F%2Fcicd.r59q.com) ![Docker Image Size](https://img.shields.io/docker/image-size/r59q/easywebstats)

Lightweight concurrent stat-collecting gin gonic web service with a simple API. Stats are queried using REST and are exported as prometheus metrics.

## Registering statistics

Registering statistics is automatically done when settings values. No initialization is needed, just set the values you wish to track.

Data values are grouped by name and label.

Note: **Stats are stored in-memory, which means for now, all data is lost when application is rebooted**

## Reading statistics

When stats are registered using a name and label, you can either read all stats for a given name, or a specific value for a given label and name.

# Documentation
## Examples
### Registering statistics
#### Set numeric value
```shell
# Curl example
curl -X 'POST' \
  'http://localhost:8080/api/v1/register/num/set' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "number_of_visits",
  "label": "about_page",
  "value": 10
}'
# Outputs (Returns the updated number of visits)
# {
#   "value": 10
# }
```

#### Increase numeric value
Increase the number of visits on the about page by 1
```shell
# Curl example
curl -X 'POST' \
  'http://localhost:8080/api/v1/register/num/increase' \
  -H 'accept: application/json' \
  -H 'Content-Type: application/json' \
  -d '{
  "name": "number_of_visits",
  "label": "about_page",
  "value": 1
}'
# Outputs (Returns the updated number of visits)
# {
#   "value": 11
# }
```

### Reading statistics
#### Read numeric values by label
```shell
# Curl example
curl -X 'GET' \
  'http://localhost:8080/api/v1/read/num/number_of_visits/about_page' \
  -H 'accept: application/json'

# Outputs
# {
#   "value": 10
# }

```

#### Read numeric values
```shell
# Curl example
curl -X 'GET' \
  'http://localhost:8080/api/v1/read/num/number_of_visits' \
  -H 'accept: application/json'

# Outputs
# {
#   "about_page": 10,
#   "home_page": 321
# }
```

### Demo
A demo is available at [ews-demo.r59q.com](https://ews-demo.r59q.com/swagger/index.html). Check it out and try the api!

## Deploying
### Docker
Standalone container
```shell
docker run --name easywebstats -p 8080:8080 -d r59q/easywebstats
```

Part of a docker compose
```yaml
# docker-compose.yml
services:
  webstats:
    image: r59q/easywebstats
    ports:
      - "8080:8080"
    restart: always
```
```shell
docker compose up -d
```

Swagger docs available at http://localhost:8080/swagger/index.html

### Using binary
#### Standalone

You can run the binary directly. [Build it directly from source](https://github.com/r59q/easywebstats?tab=readme-ov-file#building-from-source) or get it from the [releases](https://github.com/r59q/easywebstats/releases) page

There's no configuration needed, just run the binary

```shell
./easywebstats
```

#### As a service
Install the application as a systemd service

Download the binary and create a systemd service

```shell
sudo vi /etc/systemd/system/easywebstats.service
```

Create a user and fill out the service file

```properties
# easywebstats.service
[Unit]
Description=EasyWebStats
After=network.target

[Service]
Type=simple
# Create a user
User=r59q 
Group=r59q
WorkingDirectory=/usr/local/bin
ExecStart=/usr/local/bin/easywebstats
Restart=always
RestartSec=5
Environment="GIN_MODE=release"
Environment="EWS_PORT=8080"

[Install]
WantedBy=multi-user.target
```

Start the service
```shell
sudo systemctl daemon-reload
sudo systemctl start easywebstats.service
sudo systemctl enable easywebstats.service
```
Confirm it's running
```shell
sudo systemctl status easywebstats.service
```

The port can be changed by changing the EWS_PORT environment variable.

## Security

Currently there's no built-in security, it's expected to be done on a network layer. This may change in the future.

## Building
### Building using dockerfile
```shell
docker build -t easywebstats . -f build/Dockerfile
docker run --name easywebstats -p 8080:8080 -d easywebstats
```
Swagger docs available at http://localhost:8080/swagger/index.html

### Building from source

#### Prerequisits
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

## Prometheus Export

All numeric stats are exported for prometheus to scrape. Visit the `/metrics` url