FROM golang:1.10
WORKDIR /go/src/github.com/nimeshjm/fund-data
RUN go get -d -v golang.org/x/net/html
COPY main.go etfsnapshot.go ./
RUN go get github.com/patrickmn/go-cache
RUN go get github.com/gorilla/mux
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /go/src/github.com/nimeshjm/fund-data/main .
CMD ["./main"]
