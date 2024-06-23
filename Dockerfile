FROM golang:1.22.4-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

COPY . ./


RUN go build -o /pricefetcher

EXPOSE 3000

CMD [ "/pricefetcher" ]