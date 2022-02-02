FROM golang:1.16
WORKDIR go/

RUN apt-get update
RUN apt-get install vim -y

RUN ls
RUN go mod init main
RUN go get -v "github.com/go-sql-driver/mysql"
COPY ./scripts/ /go/go
RUN go build -o /golang_app_main
CMD [ "/golang_app_main" ]