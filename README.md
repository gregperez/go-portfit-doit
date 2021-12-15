<div align="center">

# Golang Portfit Do It

A Golang project for calculate profit between given dates.

[![Build Status](https://github.com/Justintime50/golang-template/workflows/build/badge.svg)](https://github.com/Justintime50/golang-template/actions)
[![Coverage Status](https://s3.amazonaws.com/assets.coveralls.io/badges/coveralls_97.svg)](https://coveralls.io/github/Justintime50/golang-template?branch=main)


</div>

## Run

The dates in format YYYY-M-DD, E.g: from: 2021-8-01 to: 2021-8-31

```
go run . 2021-12-01 2021-12-15
```

## Test and coverage

```
go test -cover -coverprofile=coverage.out 
go tool cover -html=coverage.out -o coverage.html
```