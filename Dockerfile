FROM core.harbor.ebupt.com/library/builder-golang:1.12-alpine as builder

ADD . /jarvis

RUN cd /jarvis && make

# Pull first-demo into a second stage deploy alpine container
FROM alpine:latest
RUN apk add --no-cache ca-certificates
COPY --from=builder /jarvis/build/bin/jarvis /jarvis/jarvis
EXPOSE 8080/tcp
CMD cd /jarvis && ./jarvis
