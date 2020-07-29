FROM golang@sha256:1e9c36b3fd7d7f9ab95835fb1ed898293ec0917e44c7e7d2766b4a2d9aa43da6 as builder

WORKDIR $GOPATH/src/mypackage/myapp/

COPY go.mod .

ENV GO111MODULE=on
RUN go mod download
RUN go mod verify

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build \
      -ldflags='-w -s -extldflags "-static"' -a \
      -o /go/bin/seqsy .

FROM gcr.io/distroless/static@sha256:c6d5981545ce1406d33e61434c61e9452dad93ecd8397c41e89036ef977a88f4

COPY --from=builder /go/bin/seqsy /go/bin/seqsy

EXPOSE 8080

ENTRYPOINT ["/go/bin/seqsy"]