FROM golang:1.11.5 as builder

WORKDIR /go/src/github.com/Habu-Kagumba/meli/consignment-cli

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep

RUN dep init && dep ensure

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo .

RUN rm -rf .netrc

FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN mkdir /app
WORKDIR /app

COPY --from=builder /go/src/github.com/Habu-Kagumba/meli/consignment-cli/consignment-cli .

ADD consignment.json /app/consignment.json
ADD consignment-cli /app/consignment-cli

CMD ["./consignment-cli"]
