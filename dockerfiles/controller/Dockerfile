FROM alpine:latest

RUN apk --no-cache add ca-certificates
COPY aam-controller /usr/local/bin/

CMD ["aam-controller"]
