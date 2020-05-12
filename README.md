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
./main
Applying database migrations!
Ran all migrations
How many users do you want to auto create? (5, 10 ..) >> 5
Creating 5 dummy users
New User ID is: 76eadef2-6267-4f29-8159-fe1108bf2659
tvs vector updated for User ID: 76eadef2-6267-4f29-8159-fe1108bf2659
New User ID is: bc0a1abb-f5e6-4264-afb8-b700a892e0e5
tvs vector updated for User ID: bc0a1abb-f5e6-4264-afb8-b700a892e0e5
New User ID is: 2ea798b4-ec34-41c8-9dc9-56c000bc11df
tvs vector updated for User ID: 2ea798b4-ec34-41c8-9dc9-56c000bc11df
New User ID is: 48eb77bd-2f02-4efe-8a9f-020dd048b000
tvs vector updated for User ID: 48eb77bd-2f02-4efe-8a9f-020dd048b000
New User ID is: d87e8f5a-264a-428f-a15a-acd2a9d7801e
tvs vector updated for User ID: d87e8f5a-264a-428f-a15a-acd2a9d7801e

5 Users in the DB

[
    {
        "Bio": "The sun set; the dusk fell on the stream, and lights began to appear along the shore. The Chapman light–house, a three–legged thing erect on a mud–flat, shone strongly.",
        "Email": "sophiadavis111@test.com",
        "ID": "76eadef2-6267-4f29-8159-fe1108bf2659",
        "Location": "Plympton",
        "Name": "Collarlinen"
    },
    {
        "Bio": "He completely abandoned the child of his marriage with Adelaida Ivanovna, not from malice, nor because of his matrimoni- al grievances, but simply because he forgot him.",
        "Email": "oliviajohnson736@test.org",
        "ID": "bc0a1abb-f5e6-4264-afb8-b700a892e0e5",
        "Location": "San Martin",
        "Name": "Vipernotch"
    },
    {
        "Bio": "I am all in a sea of wonders. I doubt; I fear; I think strange things, which I dare not confess to my own soul. God keep me, if only for the sake of those dear to me!",
        "Email": "isabellamartin425@test.org",
        "ID": "2ea798b4-ec34-41c8-9dc9-56c000bc11df",
        "Location": "Plympton",
        "Name": "Bowfire"
    },
    {
        "Bio": "Then followed a battle of looks between them, but the captain soon knuckled under, put up his weapon, and resumed his seat, grumbling like a beaten dog.",
        "Email": "miawhite466@test.net",
        "ID": "48eb77bd-2f02-4efe-8a9f-020dd048b000",
        "Location": "Plympton",
        "Name": "Carpetsatin"
    },
    {
        "Bio": "Then followed a battle of looks between them, but the captain soon knuckled under, put up his weapon, and resumed his seat, grumbling like a beaten dog.",
        "Email": "averyrobinson231@test.org",
        "ID": "d87e8f5a-264a-428f-a15a-acd2a9d7801e",
        "Location": "Baldock",
        "Name": "Sargentclover"
    }
]
Enter any text to search users: >> dusk Chapman


Found 1 Search Results for dusk Chapman


[
    {
        "Bio": "The sun set; the dusk fell on the stream, and lights began to appear along the shore. The Chapman light–house, a three–legged thing erect on a mud–flat, shone strongly.",
        "Email": "sophiadavis111@test.com",
        "ID": "76eadef2-6267-4f29-8159-fe1108bf2659",
        "Location": "Plympton",
        "Name": "Collarlinen"
    }
]
Closing DB!
```

## TODO

- [x] Setup env configs and build DB connection
- [x] Create table (User) with columns (name, bio, age, location, id as UUID)
- [x] Write a function that creates dummy users on demand
- [x] Main file should run as CMD which shows the total count of user available in the DB and further takes free text as input and finds users from the User table based on the matches.
