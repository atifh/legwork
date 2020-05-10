legwork
===

A simple command line tool written in Go to perform PosgreSQL Full Text search.


### Setup database

Install Postgres:

```
$ brew install postgresql
$ brew services start postgresql
```

Create PG extension

```
psql -h -d DB_NAME -U DB_USER -c 'CREATE EXTENSION if not exists "uuid-ossp"'
```

## Setting up


``` sh
$ git clone git@github.com:atifh/legwork.git
$ cd legwork
$ cp application.yml.sample application.yml
```

Note: Please update DB credentials in the application.yml file.

``` sh
$ go install
$ go build main.go
$ ./main
Applying database migrations!
Ran all migrations
Closing DB!
```

## TODO

- [x] Setup env configs and build DB connection
- [x] Create table (User) with columns (name, bio, age, location, id as UUID)
- [ ] Insert 100 dummy data into User table
- [ ] Main file should run as CMD which shows the total count of user available in the DB and further takes free text as input and finds users from the User table based on the matches.
