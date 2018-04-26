# RESTful API with golang and mysql
This is simple example restful api server using [gin][gn1] framework.

## Install and Run

Install go and mysql for ubuntu.
```sh
$ sudo add-apt-repository ppa:gophers/go
$ sudo apt-get update
$ sudo apt-get install golang-stable
$ sudo apt-get install mysql-server
```

Install go and mysql for macOS.
```sh
$ brew install go
$ brew install mysql
```

Set your paths.
```sh
$ export GOROOT=/usr/local/opt/go/libexec
$ export GOPATH=$HOME/.go
$ export PATH=$PATH:$GOROOT/bin:$GOPATH/bin
```

Start your mysql and create `dummy_db` database.
```sh
$ create database dummy_db;
```

Create `user_review` table and create `ratingTrigger`. Scripts are available at `table_script.sql` and `rating_trigger.sql` in this repository.

Build and run.
```sh
$ go get "github.com/go-sql-driver/mysql"
$ go get "github.com/gin-gonic/gin"

$ go build base.go
$ ./base
```

## API Endpoints
- http://localhost:3000/user-review/{id}
	- `GET`: get user_review record by id
- http://localhost:3000/user-reviews
	- `GET`: get all user_review records
- http://localhost:3000/user-review
	- `POST`: create new user_review record
- http://localhost:3000/user-review?id={id}
	- `PUT`: update a user_review record by id
	- `DELETE`: delete user_review record by id

[gn1]: <http://gin-gonic.github.io/gin/>
