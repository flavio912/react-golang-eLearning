FROM node:10-alpine

ENV HOME /app

WORKDIR /app

ENTRYPOINT ["sh", "./bin/run_dev.sh" ]
EXPOSE 4000