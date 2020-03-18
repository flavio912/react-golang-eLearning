FROM node:10

ENV HOME /app

COPY . /
WORKDIR /app

COPY bin/docker-entrypoint.sh /usr/local/bin
RUN ["chmod", "+x", "/usr/local/bin/docker-entrypoint.sh"]

ENTRYPOINT ["/usr/local/bin/docker-entrypoint.sh"]

EXPOSE 80
