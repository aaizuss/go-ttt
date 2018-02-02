## Go Tic Tac Toe

### Set Up Your Go Environment
* (skip this if you have a working Go installation)
1. Install [Go](https://golang.org/doc/install)
2. All Go code resides in a single directory called a [workspace](https://golang.org/doc/code.html#Workspaces). Go expects your workspace to be `$HOME/go`. Create this directory. If you want your go workspace to be located in a different directory, you'll have to [set a `GOPATH` environment variable](https://github.com/golang/go/wiki/SettingGOPATH).

### Download the Package
Get my tictactoe code with this command

`$ go get github.com/aaizuss/tictactoe`

The command downloads files from this repo to your local directory `$GOPATH/src/github.com/aaizuss/tictactoe`

### Play the Game
```
$ cd $GOPATH/src/github.com/aaizuss/tictactoe
$ go run main.go
```

### Run the tests
```
$ cd $GOPATH/src/github.com/aaizuss/tictactoe
$go test ./...
```
will run all the unit tests

Alternatively, you can navigate to each subdirectory and run `$ go test`

There is a computer vs computer test, and it takes 2 seconds to run, so it's in its own file (`/game/ai_test.go`). To run that test, add `-tags=aitest` to the test command.

### Idiomatic Go
I'm just starting too learn go, and the following are resources that I use to write Go like Gophers do
* [Effective Go](https://golang.org/doc/effective_go.html)
* [Go Style Guide](https://github.com/golang/go/wiki/CodeReviewComments)
* [Go Readability](https://talks.golang.org/2014/readability.slide)
* Naming
  * [Names](https://talks.golang.org/2014/names.slide)
  * [Name Style](https://talks.golang.org/2014/organizeio.slide#21)
* Testing
  * [Slide from readablity talk](https://talks.golang.org/2014/readability.slide#39)
  * [Table Driven Tests](https://dave.cheney.net/2013/06/09/writing-table-driven-tests-in-go)
