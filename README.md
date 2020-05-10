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
New User ID is: c886118b-cbfc-48b1-a69b-2966ee52d949
New User ID is: 396b4912-b964-4aa1-bef4-604e746f2608
New User ID is: 7cb2a117-e470-4a3d-a132-4cc894cc23f8
New User ID is: 4f8e9082-e616-48ad-9baf-1931c32d8e25
New User ID is: 26766cb0-ec6d-4b28-8066-7dbbb278e012
Closing DB!
```

## TODO

- [x] Setup env configs and build DB connection
- [x] Create table (User) with columns (name, bio, age, location, id as UUID)
- [x] Write a function that creates dummy users on demand
- [ ] Main file should run as CMD which shows the total count of user available in the DB and further takes free text as input and finds users from the User table based on the matches.
