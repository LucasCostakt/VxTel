FROM golang

COPY . /go/src/app

WORKDIR /go/src/app

COPY go.mod .

RUN go mod download

COPY . .

WORKDIR /go/src/app/cmd

RUN go build -o api

CMD ./api

EXPOSE 5000