# Go gin Study

## [환경]
```
Docker Container
go version go1.16.3

ENV     GOROOT /goroot
ENV     GOPATH /gopath
ENV     PATH $GOROOT/bin:$GOPATH/bin:$PATH
```

## 설치 Package
```
# gin Web Framework
go get -u github.com/gin-gonic/gin
https://github.com/gin-gonic/gin


# Log
go get -u github.com/sirupsen/logrus
https://github.com/sirupsen/logrus


# ORM 
go get -u gorm.io/gorm
https://gorm.io/

go get -u gorm.io/driver/sqlite  (Dev)


# Excel 
https://github.com/qax-os/excelize


# Read ini
go get -u github.com/go-ini/ini
https://github.com/go-ini/ini
```

TODO
``` 
# 프레임워크 미들웨어 개발해보자
gin framework middleware

# ORM을 연동하자
gorms
```