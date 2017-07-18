# goherokuname

[![Build Status](https://travis-ci.org/ucirello/goherokuname.svg?branch=master)](https://travis-ci.org/ucirello/goherokuname)

Heroku-like Random Names in Go

To install use:

`go get -u cirello.io/goherokuname/...`

It contais two dictionaries, one small and one big. If you want to compile with the big one, you are going to need to use build tags:

`go build -tags herokuComplete`

or

`go get -u -tags 'herokuComplete' cirello.io/goherokuname/...`

The big dictionary is actually from Wordnet of Princeton. Please, be sure to agree with their license before using it.

See documentation at [http://godoc.org/cirello.io/goherokuname](http://godoc.org/cirello.io/goherokuname).
