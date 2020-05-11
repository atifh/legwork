legwork
===

A simple command line tool written in Go to perform PosgreSQL Full Text search.


### Setup database

Install Postgres:

```
$ brew install postgresql
$ brew services start postgresql
```

Create DB

``` sh
$ createdb DB_NAME
```

Create PG extension

```
$ psql -h -d DB_NAME -U DB_USER -c 'CREATE EXTENSION if not exists "uuid-ossp"'
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
Creating 5 dummy users
New User ID is: 37726cba-a502-4c71-b7f3-b6cbef8c066b
tvs vector updated for User ID: 37726cba-a502-4c71-b7f3-b6cbef8c066b
New User ID is: ec7fc1ee-dfd5-478b-8c0e-b732a8f242ae
tvs vector updated for User ID: ec7fc1ee-dfd5-478b-8c0e-b732a8f242ae
New User ID is: 02512198-7f47-4073-9505-8a45d36f4b9f
tvs vector updated for User ID: 02512198-7f47-4073-9505-8a45d36f4b9f
New User ID is: b512c0d2-dced-4df1-a439-d4ab54803aee
tvs vector updated for User ID: b512c0d2-dced-4df1-a439-d4ab54803aee
New User ID is: 58dbc7a6-1d81-44cb-ab51-74ead0e6ae0b
tvs vector updated for User ID: 58dbc7a6-1d81-44cb-ab51-74ead0e6ae0b
Closing DB!
```

## TODO

- [x] Setup env configs and build DB connection
- [x] Create table (User) with columns (name, bio, age, location, id as UUID)
- [x] Write a function that creates dummy users on demand
- [ ] Main file should run as CMD which shows the total count of user available in the DB and further takes free text as input and finds users from the User table based on the matches.
