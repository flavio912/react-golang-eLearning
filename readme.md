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
run, along with the golang backend.

By default, after successful start of the docker containers the services are hosted at:

port : service

- 3000 : Delegate frontend (TODO)
- 3001 : Delegate storybook (TODO)

- 4000 : Manager frontend
- 4001 : Manager storybook

- 5000 : Admin Frontend

- 8080 : Golang API

## Devserver

When building into devserver the docker environments are run on a multicontainer
docker EB application to be cheap.

## Production
When building into production the docker containers are all split into separate
EB applications so that they can be load balanced and monitored separately.
