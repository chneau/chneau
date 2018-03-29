FROM golang:latest
WORKDIR /go/src/github.com/chneau/chneau
COPY . .
RUN go build -o /app

FROM scratch
COPY --from=0 /app /app
COPY public /public
ENTRYPOINT ["/app"]