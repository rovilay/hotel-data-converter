FROM golang:latest

RUN mkdir /app

ENV SORT=

ADD . /app

WORKDIR /app

RUN go mod download

RUN go build -o bin/main cmd/app/main.go cmd/app/app.go

CMD ["sh", "-c", "./bin/main -file=${FILE} -sort=${SORT}", "/bin/bash"]
