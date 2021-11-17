FROM golang:1.16-alpine
WORKDIR /caliconodedockerstats
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY *.go ./
RUN go build -o ./caliconodedockerstats
CMD [ "./caliconodedockerstats" ]