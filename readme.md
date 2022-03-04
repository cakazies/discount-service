
# Discount Service

This repo about implementation golang connection with Database Postgres and setting configure with in `.toml`

## ERD (Entity Relationship Diagram)

please check this [link](https://dbdiagram.io/d/6220ef7554f9ad109a550dad) for ERD 

## Run Migration with

For migration please check `./schema/`

## How to run

- you can configuration config (**Database**) in folder  `configs/config.toml` and compare your local configs, database, host and depends

you can install dependencies using `go mod tidy`,

## Run Local

`go run main.go`

## Run Docker

- build docker `docker build -t discount-service .`
- Run `docker run -it --rm --name cont-discount-service discount-service`