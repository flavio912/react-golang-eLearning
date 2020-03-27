FROM node:10

ENV HOME /app

COPY . /app
WORKDIR /app

COPY ./bin/run_staging.sh /scripts/run_staging.sh
RUN ["chmod", "+x", "/scripts/run_staging.sh"]

RUN ["yarn"]
ENTRYPOINT ["/scripts/run_staging.sh"]

EXPOSE 4000 4001