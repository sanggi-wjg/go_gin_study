# Go gin Study

## [환경]
```
go version go1.17

GOROOT=/goroot
GOPATH=/gopath
PATH=$GOROOT/bin:$GOPATH/bin:$PATH
```

## 설치 Package
```
# gin Web Framework  https://github.com/gin-gonic/gin
go get -u github.com/gin-gonic/gin

# Log  https://github.com/sirupsen/logrus
go get -u github.com/sirupsen/logrus

# ORM  https://gorm.io/
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
go get -u gorm.io/driver/sqlite  (Dev)

# Excel  https://github.com/qax-os/excelize

# Read ini  https://github.com/go-ini/ini
go get -u github.com/go-ini/ini
```

## TODO
``` 
# 프레임워크 미들웨어 개발해보자
gin framework middleware
```

## MYSQL
```
docker pull mysql:5.7
docker run -d -p 3306:3306 -e MYSQL_ROOT_PASSWORD=rootroot --name mysql_1 mysql:5.7

sudo docker inspect mysql_1 | grep IPAddress
docker exec -it mysql_1 bash

select user,host,authentication_string from mysql.user order by user;
create user 'root'@'172.17.0.1' identified by 'rootroot';
grant all privileges on *.* to 'root'@'172.17.0.1' with grant option;
flush privileges;
```