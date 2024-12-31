#!/bin/bash

go build -o crawler

# usage: ./crawler URL maxConcurrency maxPages
./crawler $1 $2 $3
