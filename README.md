# golang interview prep

## Goal of this repo.

This repo contains Golang code that does not follow best practises, contains bugs and security issues. It is intended to
be used as an interview exercise or a practise exercise for jr/mid-level Go engineers.

This repo contains, technically, a functional golang application that receives a request to create a user and stores it
into a postgres Database.

As an exercise, you could try identifying and correcting some of the issues in this repo. This would work particularly
well as a pair programming exercise.

## Getting Started

You can get the database started by running `docker-compose up`

Once running the Go app, you can make a CURL request as follows:

```curl
 curl -X POST -H "Content-Type: application/json" -d '{"username":"john", "password":"secret"}' http://localhost:8080/user
```

## Solutions
you can find the solutions [here](./solutions.md)
