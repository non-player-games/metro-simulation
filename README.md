# Metro simulation

[![GoDoc](https://godoc.org/github.com/non-player-games/metro-simulation?status.svg)](https://godoc.org/github.com/non-player-games/metro-simulation)[![Build Status](https://travis-ci.org/non-player-games/metro-simulation.svg?branch=master)](https://travis-ci.org/non-player-games/metro-simulation)

Metro simulation in Go for data analytics problem.

## Get started

### Dependencies

* Install [Go](https://golang.org/)
* Install MySQL

### Start command

Before starting running the application, you will need to run through ddl defined
as `schema.sql`.

Please create a database called `metro` for starting point and run the create
table script like below:

```
mysql -u root metro < schema.sql
```

This repo uses the Makerfile to group commands. To get started, run `make run`
