FROM mysql:latest
USER root
RUN apt-get update && apt-get install -y --no-install-recommends \
	g++ \
	gcc \
	libc6-dev \
	vim \
	curl \
	wget
RUN cd / &&  wget --no-check-certificate https://dl.google.com/go/go1.9.4.linux-amd64.tar.gz
RUN cd / && tar -xvf go1.9.4.linux-amd64.tar.gz && mv go /usr/local
RUN cd / && rm -f go1.9.4.linux-amd64.tar.gz
RUN mkdir -p "/go" "/go/bin" && chmod -R 777 "/go"
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
COPY ./bin/ /go/bin/
COPY ./src/goserver/public /go/bin/public 
COPY ./src/goserver/views /go/bin/views
COPY ./src/goserver/docker/docker_mjolnir/init.sh /init.sh 
WORKDIR $GOPATH
CMD ["sh","/init.sh"]