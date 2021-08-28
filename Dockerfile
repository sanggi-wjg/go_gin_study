FROM    ubuntu:20.04

########################################################################################
# docker build -t for_go .
# docker run -d -p 9000:9000 --privileged=true --name for_go for_go /usr/sbin/init
# docker exec -it for_go bash
########################################################################################
ARG     DEBIAN_FRONTEND=noninteractive
ENV     TZ=Asia/Seoul

########################################################################################
# Set Basic 
########################################################################################
RUN     apt-get update && apt-get -y upgrade
RUN     apt install -y build-essential wget curl git

########################################################################################
# Set Entrypoint 
########################################################################################
COPY        entrypoint.sh /entrypoint.sh
RUN         chmod +x /entrypoint.sh
ENTRYPOINT  ["/entrypoint.sh"]

########################################################################################
# Install Go
########################################################################################
RUN     mkdir -p /goroot 
WORKDIR /goroot
RUN     wget https://golang.org/dl/go1.17.linux-amd64.tar.gz
RUN     tar xvzf  go1.17.linux-amd64.tar.gz -C /goroot --strip-components=1

# Set environment variables.
ENV     GOROOT /goroot
ENV     GOPATH /gopath
ENV     PATH $GOROOT/bin:$GOPATH/bin:$PATH

RUN     go get -u github.com/gin-gonic/gin
RUN     go get -u github.com/sirupsen/logrus
RUN     go get -u github.com/go-ini/ini
RUN     go get -u gorm.io/gorm
# Only Dev
RUN     go get -u gorm.io/driver/sqlite


WORKDIR /gopath
CMD     ["bash"]
EXPOSE  9000 9022