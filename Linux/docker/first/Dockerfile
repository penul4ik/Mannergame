FROM ubuntu:22.04
LABEL author=Penul4ik
RUN apt-get update
RUN apt-get install nginx -y
RUN apt-get clean && rm -rf /var/lib/apt/lists/*
WORKDIR /var/www/html
COPY index.html .
COPY script.sh /opt/script.sh
RUN chmod +x /opt/script.sh
ENV OWNER=Penul4ik
ENV TYPE=demo
EXPOSE 80
CMD ["/opt/script.sh"]
