# TTC Project

## Project structure

This repo is a monolith structure with loose coupling between each of the 3 react
sites (thus they all have their own package.json etc). This is to make it easy to
distibute the sites to different servers/EB environments if scaling requires it.

The monolith structure just saves us having to just back and forth between repos
when working on the project

## Developing

If you just want to have everything setup and running for you, you can use
`docker-compose up --build` from the top level directory to have all 3 react sites
run, along with the golang backend. Depending on your computer, this take a while to
run the first time, be patient.

After you have run `docker-compose up --build` the first time, you can run it without
the build flag `docker-compose up` and it should speed up significantly.

By default, after successful start of the docker containers the services are hosted at:

port : service

- 3000 : Delegate frontend (TODO)
- 3001 : Delegate storybook (TODO)

- 4000 : Manager frontend
- 4001 : Manager storybook

- 5000 : Admin Frontend

- 8080 : Golang API

### Running only certain services

If you only wish to run the services separately you can run `docker-compose start api_db` to
just run the API database (which is available at port 5430, so as not to interfear with others).
Other services can be run separetely with:

- `docker-compose start api` (for the API only - no DB)
- `docker-compose start api_db` (just for the database at port 5430)
- `docker-compose start admin` (for the admin react)
- `docker-compose start manager` (for the manager react)

## Issues with docker + postgres

Database migrations mixed in with chaning branches can cause your database to be in a weird state
to sort this out just run `docker-compose up --force-recreate --renew-anon-volumes` to put your
database (and everything else) back to its initial state.

## Production
When building into production the docker containers are all split into separate
EB applications so that they can be load balanced and monitored separately.
