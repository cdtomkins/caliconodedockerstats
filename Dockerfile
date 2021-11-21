# syntax=docker/dockerfile:1
FROM golang:1.17.3 AS caliconodedockerstatsbuilder
WORKDIR /caliconodedockerstats-builddir
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN CGO_ENABLED=0 go build -o ./caliconodedockerstats

FROM alpine:latest
COPY --from=caliconodedockerstatsbuilder /caliconodedockerstats-builddir/caliconodedockerstats ./
EXPOSE 9088
ENTRYPOINT [ "./caliconodedockerstats" ]
