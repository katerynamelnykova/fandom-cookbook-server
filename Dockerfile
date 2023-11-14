FROM mongo:latest

RUN apt-get update
RUN apt-get -y install vim

COPY ./db /dbbackup

EXPOSE 27017
