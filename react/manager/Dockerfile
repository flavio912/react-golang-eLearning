FROM node:10.16-buster-slim

ENV HOME /app

COPY . /app
WORKDIR /app

COPY ./bin/run_staging.sh /scripts/run_staging.sh
RUN ["chmod", "+x", "/scripts/run_staging.sh"]

ENTRYPOINT ["/scripts/run_staging.sh"]

EXPOSE 4000 4001