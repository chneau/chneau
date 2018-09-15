# build stage
ARG BASE=/go/src/github.com/chneau/chneau
FROM golang:alpine AS build-env
ARG BASE
ADD . $BASE
RUN cd $BASE && CGO_ENABLED=0 go build -o /chneau -ldflags '-s -w -extldflags "-static"'

FROM alpine AS prod-ready
ARG BASE
ENV PORT 80
COPY --from=build-env $BASE/public /public
COPY --from=build-env /chneau /chneau
HEALTHCHECK --interval=5s --timeout=3s --start-period=1s --retries=3 CMD wget --quiet --tries=1 --spider localhost:$PORT/health || exit 1
ENTRYPOINT [ "/chneau" ]
