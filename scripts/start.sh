#!/bin/bash

# install air a hot reloading tool, so we dont
# have to restart the server manually every time we make
# a code change
go install github.com/cosmtrek/air@latest

# run this command, requires GOPATH or GOBIN to be set in your path
air