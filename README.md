# Mitre Att&ck Framework Scraper

<div align="center">
  <img name="mit-license" src="https://img.shields.io/badge/License-MIT-yellow.svg">
  <img name="version" src="https://img.shields.io/badge/version-1.0.0-blue">
  <img name="issues" src="https://img.shields.io/github/issues/ChistaDev/Chista">
  <img name="maintained-yes" src="https://img.shields.io/badge/Maintained%3F-yes-green.svg">
  <img name="made-with-go" src="https://img.shields.io/badge/Made%20with-Go-1f425f.svg">
  <img name="open-source" src="https://badges.frapsoft.com/os/v1/open-source.svg?v=103">
  <img name="goreport" src="https://goreportcard.com/badge/gojp/goreportcard">
</div>

<p align="right"><i>v1.0</i></p>

## Table of Contents  

- [Basic Usage](#basic-usage)
- [Sample Outputs](#sample-outputs)

<h3>Prerequisite</h3>
- go1.21

- `7777` and `7778` ports should be available

<h2>Basic Usage</h2>
You can use pre-built binaries or you can build the project and use. It's up to your choice!

### API Usage
You just send http GET request for `http://0.0.0.0:7777/api/v1/apt_feed` endpoint. 
```
http://0.0.0.0:7777/api/v1/apt_feed?aptFeed=endpoints
http://0.0.0.0:7777/api/v1/apt_feed?aptFeed=tactics
http://0.0.0.0:7777/api/v1/apt_feed?aptFeed=technics
http://0.0.0.0:7777/api/v1/apt_feed?aptFeed=mitigations
http://0.0.0.0:7777/api/v1/apt_feed?aptFeed=relationships
```

### Using API Withot Building
If you already installed go you can use this command to run the API.

```
go run main.go
```

<h3>Building & Running from source</h3>

**1. Clone the repository**
```sh
git clone https://github.com/grealyve/mitre-cti-scraper.git
```
**2. Build & Run the API application**

First, open a Command Prompt/Terminal. Then execute the following commands.
- For Windows:
```sh
go build -o mitre-cti-scraper.exe
./mitre-cti-scraper.exe
```
- For Linux: 
```sh
go build -o mitre-cti-scraper
./mitre-cti-scraper
```
NOTE: If you cannot execute the command in Linux, you should give execute permission yourself on the file. You can use `chmod +x mitre-cti-scraper`.
After running the API server, you'll see the following output.
```
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /api/v1/apt_feed          --> github.com/grealyve/mitre-cti-scraper/controller.GetAptFeed (5 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on localhost:7777
```

### Sample Outputs:
![Screenshot_1](https://github.com/grealyve/mitre-cti-scraper/assets/41903311/da032a73-a068-4b49-addd-1be2ddb80513)
![Screenshot_2](https://github.com/grealyve/mitre-cti-scraper/assets/41903311/c8738643-46ba-44f7-86d5-a16cba7ca927)

