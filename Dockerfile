ARG DIR=/src/app
ARG APP=main

FROM golang:1.17-alpine as builder
ARG DIR
ARG APP
WORKDIR ${DIR}
COPY . .
ENV CGO_ENABLED=0
ENV GOOS=linux
RUN go build -a -o ${APP} .

FROM alpine:3.17 as production
ARG DIR
ARG APP
WORKDIR ${DIR}
COPY --from=builder ${DIR}/${APP} .
CMD [ "./main" ]
