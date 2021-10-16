# go-docker

Creating a Simple Golang App
Let’s create a simple Go app that we’ll containerize. Fire up your terminal and type the following command to create a Go project -

$ mkdir go-docker
We’ll use Go modules for dependency management. Change to the root directory of the project and initialize Go modules like so -

$ cd go-docker
$ go mod init github.com/venupammi26/go-docker

#Docker build image

>docker build -t go-docker .

#BAsh into container

>docker run -it go-docker
