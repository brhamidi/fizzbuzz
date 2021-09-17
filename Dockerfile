FROM golang:1.16 AS deps

WORKDIR /go/src

COPY go.mod ./
COPY go.sum ./

RUN go mod download

FROM deps as build

COPY cmd cmd
COPY pkg pkg
COPY docs docs

RUN go build -o /go/bin/app cmd/fizzbuzz/fizzbuzz.go

FROM gcr.io/distroless/base-debian10

COPY --from=build /go/bin/app /

ENTRYPOINT ["/app"]

# Make Dockerfile et Dockerfile.dev perform make swag to allow user to not have go to building the project