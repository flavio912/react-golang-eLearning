FROM node:10

ENV HOME /app

COPY . /react/admin
WORKDIR /app

RUN ["yarn"]
ENTRYPOINT ["yarn", "run", "startNewProd"]

EXPOSE 80
