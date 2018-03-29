FROM golang:latest
WORKDIR /go/src/github.com/chneau/chneau
COPY . .

ENV CGO_ENABLED=0
RUN go get -v
RUN go build -o /app -a -ldflags '-extldflags "-static"'
# RUN go build -o /app -a -ldflags '-s -w -extldflags "-static"'

FROM scratch
COPY --from=0 /app /app
COPY public /public
ENTRYPOINT ["/app"]