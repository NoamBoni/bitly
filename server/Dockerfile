FROM golang:1.19-alpine

WORKDIR /usr/app

COPY . .

RUN go get . 

CMD ["go","run","."]
