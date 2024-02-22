FROM golang:1.21.6-alpine AS builder
RUN apk update && apk add --no-cache git ca-certificates


WORKDIR /go/awsUtils
COPY . .

RUN go install
RUN go build -o /go/awsUtils/bin/awsUtils

FROM scratch

# Pour les certificats
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ 

COPY --from=builder /go/awsUtils/bin/awsUtils .

CMD ["/awsUtils"]