FROM ubuntu:22.04
LABEL author=Penul4ik
RUN apt-get update
RUN apt-get install apache2 -y
EXPOSE 80
CMD ["apache2ctl","-D","FOREGROUND"]
