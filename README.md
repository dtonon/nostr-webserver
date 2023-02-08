# Nostr Web Server + Client

This is a proof of concept to show how to build a self-contained web server running a Nostr web client. The script is multi platform (Linux, macOS and Windows). The repo contains a [Snort.social](https://github.com/v0l/snort) snapshot (v. 2023-02-08) as an example but the exact procedure can be applied to every other web clients: Iris, Hamstr, etc.

When the compiled app is executed a web server starts listening on port 8080 and the default web browser is opened to localhost:8080.

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

## Releases

Under /release you can find the Linux, macOS and Windows executables for Snort (2023-02-08).