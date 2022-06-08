FROM postgres:14.3-bullseye
COPY ./dbinit/* ./docker-entrypoint-initdb.d/
COPY ./postgres.sh .
