legwork
===

A simple command line tool written in Go to perform PosgreSQL Full Text search.


### Setup database

Install Postgres:

```
$ brew install postgresql
$ brew services start postgresql
```

## Setting up

```sh
$ git clone git@github.com:atifh/legwork.git
$ cd legwork
$ cp application.yml.sample application.yml

Please update DB credentials in the application.yml file.

$ go install
$ go build main.go
$ ./main
legwork_dev
Closing DB!
```

## TODO

- [x] Setup env configs and build DB connection
- [ ] Create table (User) with columns (name, bio, age, location, id as UUID)
- [ ] Insert 100 dummy data into User table
- [ ] Main file should run as CMD which shows the total count of user available in the DB and further takes free text as input and finds users from the User table based on the matches.
