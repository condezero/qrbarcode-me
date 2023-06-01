FROM golang as builder
WORKDIR /go/src/
COPY . .
RUN go get .
RUN CGO_ENABLED=0 GOOS=linux go test 
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o qrbarcode-me .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
LABEL author="Carlos Lisandra"

WORKDIR /root/
COPY --from=builder /go/src/ .


EXPOSE 5000
ENTRYPOINT ["./qrbarcode-me", ".production"]