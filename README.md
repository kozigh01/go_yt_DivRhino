# Sample Applications from Youtube "Div Rhino" channel
Youtube: 
    * [How to build a CLI tool with Go and Cobra](https://www.youtube.com/channel/UCpc_4OVIYn-04JjSwUCIa-g)
    * [How to add flags to a CLI tool built with Go and Cobra](https://www.youtube.com/watch?v=kT7Z02bR1IY)

## Resources:
* Cobra: [github](https://github.com/spf13/cobra) | [cobra generator](https://github.com/spf13/cobra/blob/master/cobra/README.md)
* Go project layout: [Mark Wolfe](https://www.wolfe.id.au/2020/03/10/how-do-i-structure-my-go-project/)
* Dad jokes: [icanhazdadjoke api](https://icanhazdadjoke.com/api)
* Go packages:
    * [http](https://pkg.go.dev/net/http)
    * Examples:
        ```bash
        curl -H "Accept: application/json" https://icanhazdadjoke.com/

        curl -H "Accept: text/plain" https://icanhazdadjoke.com/ 

        curl -H "Accept: application/json" "https://icanhazdadjoke.com/search?term=hipster"

        curl -X POST -d '{"query": "query { joke {id joke permalink } }"}' -H "Content-Type: application/json" https://icanhazdadjoke.com/graphql
        ```


## Dadjoke App: Usefule commands:
```bash
# initialize go mod
$ go mod init github.com/kozigh01/go_yt_DivRhino
# after installing go packages with "go get", update the go.mod file and download vendor packages with:
$ go mod tidy
$ go mod vendor

# install cobra
$ go get -u github.com/spf13/cobra

# install cobra generator
$ go get github.com/spf13/cobra/cobra

# initialize cobra app (from project root directory)
$ cobra init --pkg-name github.com/kozigh01/go_yt_DivRhino/cmd/dadjoke cmd/dadjoke
$ go mod vendor
$ go run cmd/dadjoke/main.go

# add a new command to the app
$ cd cmd/dadjoke
$ cobra add random

# install and run the app locally
$ go install # run from the cmd/dadjoke directory
$ $GOPATH/bin/dadjoke random --term cats
```

## Study Buddy App: useful commands:
```bash

```
