FROM golang:1.14-alpine AS build

WORKDIR /go/src/github.com/HackDalton/let-me-in-2
COPY . .

RUN apk add build-base
RUN go get -d -v ./...
RUN go install -v ./...

FROM alpine

COPY --from=build /go/bin/let-me-in-2 ./
COPY ./public ./public

CMD ["./let-me-in-2"]