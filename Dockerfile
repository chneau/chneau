# build stage
FROM golang:alpine AS build-env
ENV BASE /go/src/github.com/chneau/chneau
ADD . $BASE
RUN cd $BASE && CGO_ENABLED=0 go build -o /chneau -ldflags '-s -w -extldflags "-static"'

FROM alpine AS prod-ready
COPY --from=build-env /chneau /chneau
RUN mkdir /data
ENTRYPOINT [ "/chneau" ]
