# goherokuname
Heroku-like Random Names in Go

To install use:

`go get -u github.com/ccirello/goherokuname/...`

It contais two dictionaries, one small and one big. If you want to compile with the big one, you are going to need to use build tags:

`go build -tags herokuComplete`

The big dictionary is actually from Wordnet of Princeton. Please, be sure to agree with their license before using it.

See documentation at [http://godoc.org/github.com/ccirello/goherokuname](http://godoc.org/github.com/ccirello/goherokuname).