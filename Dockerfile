FROM golang:1.17-alpine

WORKDIR /src/app

COPY . .

RUN go get -d -v
RUN go build -v

CMD [ "./example" ]
