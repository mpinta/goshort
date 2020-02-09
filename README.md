# goshort
Simple URL shortener written in [React](https://reactjs.org/) and [Golang](https://golang.org/), using [SQLite](https://www.sqlite.org/index.html).

## What is goshort doing?
__Goshort__ shortens given URL into a random character string that is much prettier and easier to remember. When we enter a short URL into the search bar, the goshort redirects us to its source page. Shortener can provide a short URL without a time limit or a short URL, which is valid between 1 and 60 minutes.

## Usage
You can set up goshort either using Docker, or manually building it from source. To install it, firstly `cd` into `$GOPATH/src/github.com/mpinta` and clone the repository:
```
$ git clone https://github.com/mpinta/goshort
```

### Docker
Requirements in order to install project with Docker are:
* [Docker](https://docs.docker.com/install/)
* [Docker Compose](https://docs.docker.com/compose/install/)

The easiest way to install goshort is by running the `docker-compose.yml` file. The file will automatically create `backend` and `web` containers with all the dependencies that the URL shortener needs:
```
$ cd goshort
$ sudo docker-compose up
```

### Building from source
Requirements in order to build project from source are:
* [Golang 1.13.4+](https://golang.org/doc/install)
* [Node.js 12.15+](https://nodejs.org/en/download/) with [npm](https://www.npmjs.com/)

#### Golang packages:
* `github.com/gin-gonic/gin`
* `github.com/gin-contrib/cors`
* `github.com/jinzhu/gorm`
* `github.com/mattn/go-sqlite3`

You can install them with `go get <package>` command, build `backend` and start it:
```
$ cd goshort/backend
$ go get <package>
$ go build -o <name> .
$ ./<name>
```
    
#### Main Node.js dependencies:
* `react`
* `react-bootstrap`
* `react-dom`
* `react-router-dom`
* `react-scripts`

You can install them with `npm install` command and then start the `web` interface:
```
$ cd goshort/web
$ npm install
$ npm start
```

## Other important things
Make sure to create `.env` file inside `web` directory and add new React environment variable by name `REACT_APP_API_URL`, which is indicating `backend` API endpoint address. The default address is following:
```
REACT_APP_API_URL=http://localhost:8080/api
```

SQLite database creates its `.db` file only the first time the `backend` is started. To enable database recreation on every new start, change this part of code in `main.go` file from `data.Create()` to:
```go
err = data.Recreate()
if err != nil {
    exception.FatalInternal(err)
}
```
