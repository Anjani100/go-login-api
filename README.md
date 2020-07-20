# go-login-api
A Rest API for login and registration using Golang and Gin framework.

## Prerequisites
* Have *Go* installed in your machine.
* Have a general understanding of the Go language and RESTful API.

## Installation
* Initializing a Go module:
  * Command: go mod init
  * If the above command throws an error: use the following command instead: go mod init github.com/<Your_Github_UserName>/<Your_Repository_Name>
* Installing some dependencies: We are going to use Gin framework for building this API
  * Terminal command: go get github.com/gin-gonic/gin github.com/jinzhu/gorm
After the installation, your directory should contain two files namely go.mod (or mod.mod) and go.sum.

## Setting Up the Server
* Go through the main.go file. In that file, the gin server is initialized.
* To test it, run the command *go run main.go* and then open up your browser and visit *127.0.0.1:8080* to look at the data.
