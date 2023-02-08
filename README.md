# Nostr Web Server

This is a proof of concept to show how to build a multi platform self contained web server running Snort. The exact procedure can be applied to every other client side web application, ex. Iris, Hamstr, etc.

## 1. Install go
Follow the instructions at https://github.com/golang/go

## 2. Checkout the repo

```
git clone git@github.com:dtonon/nostr-webserver.git
````

## 4. Upgrad/modify the installation

If necessary replace the /public dir with an upgraded/modified version of Snort or another build version of the web app you need.

## 3. Build for your platform

```
GOOS=windows GOARCH=amd64 go build -o nostr-webserver-windows-amd64.exe nostr-webserver.go

GOOS=darwin GOARCH=amd64 go build -o nostr-webserver-mac-amd64 nostr-webserver.go

GOOS=linux GOARCH=amd64 go build -o nostr-webserver-linux-amd64 nostr-webserver.go

```

When the compiled app is executed a web server starts listening on port 8080 and the default web browser is opened to localhost:8080.