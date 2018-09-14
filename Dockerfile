# build stage
ARG BASE=/go/src/github.com/chneau/chneau
FROM golang:alpine AS build-env
ARG BASE
ENV BASE $BASE
ADD . $BASE
RUN cd $BASE && CGO_ENABLED=0 go build -o /chneau -ldflags '-s -w -extldflags "-static"'

FROM alpine AS prod-ready
ARG BASE
COPY --from=build-env $BASE/public /public
COPY --from=build-env /chneau /chneau
RUN mkdir /data
ENTRYPOINT [ "/chneau" ]
