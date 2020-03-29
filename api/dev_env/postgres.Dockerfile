# DockerFile for creating PostgreSQL database
FROM postgres 
ENV POSTGRES_PASSWORD thd7sds3928dxxkjaKi8
ENV POSTGRES_DB postgresdb 
ENV POSTGRES_USER ttc-api
COPY ./dev_env/init.sql /docker-entrypoint-initdb.d/