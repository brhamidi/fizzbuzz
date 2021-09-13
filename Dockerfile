FROM golang:1.16 AS deps

WORKDIR /go/src

COPY Makefile ./
COPY go.mod ./
# COPY go.sum ./

RUN go mod download

FROM deps as build

COPY *.go ./
RUN go build -o /go/bin/app

FROM gcr.io/distroless/base-debian10

COPY --from=build /go/bin/app /

ENTRYPOINT ["/app"]
