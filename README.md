## Go Tic Tac Toe

### Set Up Your Go Environment
* (skip this if you have a working Go installation)
1. Install [Go](https://golang.org/doc/install)
2. All Go code resides in a single directory called a [workspace](https://golang.org/doc/code.html#Workspaces). Go expects your workspace to be `$HOME/go`. Create this directory. If you want your go workspace to be located in a different directory, you'll have to [set a `GOPATH` environment variable](https://github.com/golang/go/wiki/SettingGOPATH).

### Download the Package
Since the repository is private, you'll need to change your `.gitconfig` to force it to use SSH:

`git config --global url."git@github.com:".insteadOf "https://github.com/"`

Then
`$ go get github.com/aaizuss/go-ttt`

It downloads files from this repo to your local directory, `$GOPATH/src/github.com/aaizuss/go-ttt`

### Play the Game
```
$ cd $GOPATH/src/github.com/aaizuss/go-ttt
$ go run play.go
```

### Run the tests
```
$ cd $GOPATH/src/github.com/aaizuss/go-ttt
$ go test ./...
```
will run all the unit tests

Alternatively, you can navigate to each subdirectory and run `$ go test`

There is a computer vs computer test that takes a little longer to run, so it's in its own file (`/game/ai_test.go`). To run that test, add `-tags=aitest` to the test command.

### Idiomatic Go
I referred to the following resources to write Go like Gophers do:
* [Effective Go](https://golang.org/doc/effective_go.html)
* [Go Style Guide](https://github.com/golang/go/wiki/CodeReviewComments)
* Naming
  * [Names](https://talks.golang.org/2014/names.slide)
  * [Name Style](https://talks.golang.org/2014/organizeio.slide#21)
* Testing
  * [Table Driven Tests](https://dave.cheney.net/2013/06/09/writing-table-driven-tests-in-go)
