FROM golang:1.12.9-alpine3.10 as build-env

RUN apk update \
    && apk upgrade \
    && apk add --no-cache ca-certificates openssl \
    && update-ca-certificates 2>/dev/null || true

RUN mkdir /{{<service_name>}}
WORKDIR /{{<service_name>}}
RUN apk add git
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o /go/bin/{{<service_name>}}

FROM scratch
COPY --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build-env /go/bin/{{<service_name>}} /go/bin/{{<service_name>}}
ENTRYPOINT ["/go/bin/{{<service_name>}}"]