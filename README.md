# goserver

`goserver: A light waight Web Server`

This repository contains a collection of Go programs and libraries that demonstrate how we can use go lang as webserver or APIserver

> We have include login and signup functionality within this project 

## Prerequisites

- Install Golang : Please follow [Setup.md](./Steup.md) to install golang and standard setup while working with go.
- Mysql : Please find SQL scema file in repository 

### Packages
We have used below packages with-in project

[Go-MySQL-Driver](https://github.com/go-sql-driver/mysql)
Go MySQL Driver is an implementation of Go's database/sql/driver interface. You only need to import the driver and can use the full database/sql API then.

[gorilla/mux](https://github.com/gorilla/mux)
Package gorilla/mux implements a request router and dispatcher for matching incoming requests to their respective handler.

[securecookie](https://github.com/gorilla/securecookie)
securecookie encodes and decodes authenticated and optionally encrypted cookie values.
Secure cookies can't be forged, because their values are validated using HMAC. When encrypted, the content is also inaccessible to malicious eyes.
-   [Go-MySQL-Driver](https://github.com/go-sql-driver/mysql)
-	[gorilla/mux](https://github.com/gorilla/mux)
-	[securecookie](https://github.com/gorilla/securecookie)
-	[package bcrypt](https://godoc.org/golang.org/x/crypto/bcrypt)

## start the project 

Please follow [Setup.md](./Steup.md) to install golang and standard setup while working with go.
```sh
cd src
git clone https://github.com/AbhiAdr/goserver.git
cd goserver
go get -u github.com/go-sql-driver/mysql github.com/gorilla/mux github.com/gorilla/securecookie golang.org/x/crypto/bcrypt
```
next step will be import sql file , which will create `project` database and `users` table in `mysql`.

```sh
go build 
or
go install
```
will start the goserver open http://localhost:8086/ to see the webapp.

### main files

- mail.go : Init Db And Secure Cookies , Set Static Files Path and start the webserver.
- routes.go : to create a new route or API.
- views/src : this `folder` containes the view part in react.
- handlers : routes will redirect to this `folder`.

since we are using react on front end side please install `node` and `webpack` if you are changing any thing in views/src files.
