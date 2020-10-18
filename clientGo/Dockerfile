FROM golang

ENV URL http://172.28.1.3:5000/consult

ENV PORT 5050

COPY . /go/src/app

WORKDIR /go/src/app

COPY go.mod .

RUN go mod download

COPY . .

WORKDIR /go/src/app/cmd

RUN go build -o client

CMD ./client

EXPOSE 5050