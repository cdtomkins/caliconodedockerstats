# syntax=docker/dockerfile:1
FROM golang:1.16 AS /caliconodedockerstats-builder
WORKDIR /caliconodedockerstats-builddir
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o ./caliconodedockerstats

FROM alpine:latest
WORKDIR /caliconodedockerstats
COPY --from=/caliconodedockerstats-builder //caliconodedockerstats-builddir/caliconodedockerstats ./
EXPOSE 9088
CMD [ "./caliconodedockerstats" ]
