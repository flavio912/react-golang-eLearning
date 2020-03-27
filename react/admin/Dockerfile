FROM node:10

ENV HOME /app

COPY . /app
WORKDIR /app

RUN ["yarn"]
ENTRYPOINT ["yarn", "run", "startNewProd"]

EXPOSE 80
