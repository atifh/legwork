legwork
===

A simple command line tool written in Go to perform PosgreSQL Full Text search.

*NOTE: Make sure you have [installed Go](https://golang.org/doc/install) locally.*


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

Found 1 Search Results for dog blood

[
    {
        "Bio": "One dog rolled before him, well-nigh slashed in half; but a second had him by the thigh, a third gripped his collar be- hind, and a fourth had the blade of the sword between its teeth, tasting its own blood.",
        "Email": "sofiajackson677@example.org",
        "ID": "2aec998e-2194-4bb9-b429-bc7b8df593ce",
        "Location": "Newstead",
        "Name": "Nosecandle"
    }
]
Closing DB!
```

## TODO

- [x] Setup env configs and build DB connection
- [x] Create table (User) with columns (name, bio, age, location, id as UUID)
- [x] Write a function that creates dummy users on demand
- [ ] Main file should run as CMD which shows the total count of user available in the DB and further takes free text as input and finds users from the User table based on the matches.
