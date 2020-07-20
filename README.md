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
  * Terminal command: go get github.com/gin-gonic/gin
After the installation, your directory should contain two files namely go.mod (or mod.mod) and go.sum.

## Setting Up the Server
* Go through the main.go file. In that file, the gin server is initialized.
* To test it, run the command *go run main.go* and then open up your browser and visit *127.0.0.1:8080* to look at the data.

## Setting Up the Database
The database setup is done entirely in the **models** directory.
For this particular project, I am using PostgreSQL. You can follow the instructions to install it [here](https://www.digitalocean.com/community/tutorials/how-to-install-and-use-postgresql-on-ubuntu-18-04).
* Create a database: You can find out how to do it in the link presented above.
* Command: *go get -u github.com/lib/pq* (This is done to install the necessary drivers for connecting the server to the database).

## RESTful Routes
For this purposes, we will be doing all the coding in the **controllers** directory. The *userAuth.go* file contains the code for both user registration and login.
Finally, we are done. Now let's test the API.
Fire up your POSTMAN and send a POST request to *127.0.0.1:8080/register* and send the Username and Password in JSON format in the **Body->Raw** section and see if it sends you the necessary data. Do the same for login after this.
That is all for this project. Hope this helps!
