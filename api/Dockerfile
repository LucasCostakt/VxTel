FROM golang

# ENV GO111MODULE=on \
#     CGO_ENABLED=0 \
#     GOOS=linux \
#     GOARCH=amd64

COPY . /go/src/app

WORKDIR /go/src/app

# RUN go get ./

COPY go.mod .

# COPY go.sum .

RUN go mod download

COPY . .

WORKDIR /go/src/app/cmd

RUN go build -o api

CMD ./api

EXPOSE 5000