FROM golang:1.15.7-buster as build

ENV CGO_ENABLED="0" GOOS="linux" GOARCH="amd64"

WORKDIR /build

ADD . /build

RUN go build -o ports cmd/main.go

FROM scratch
USER 10001:10001

COPY --from=build /build/ports /ports

ENTRYPOINT ["/ports"]
