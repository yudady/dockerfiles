FROM sameersbn/bind:9.16.1-20200524
USER root

RUN apt update
RUN apt list --upgradable
RUN apt upgrade -y
RUN apt install iputils-ping -y
RUN apt install bind9-utils -y
