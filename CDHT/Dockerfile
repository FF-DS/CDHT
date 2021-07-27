FROM golang:1.16-alpine

WORKDIR /Docker/

COPY go.mod ./
# COPY go.sum ./
RUN go mod download

COPY *.go ./

RUN go build -o /docker-cdht

EXPOSE 443

CMD [ "/docker-cdht" ]

# 