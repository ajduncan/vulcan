FROM golang:alpine

LABEL maintainer="ajduncan@gmail.com"

# RUN apk add --no-cache python3-dev openssl-dev libffi-dev gcc libc-dev

RUN mkdir /app
ADD ../../bin /app/
WORKDIR /app
# RUN make .
RUN adduser -S -D -H -h /app vulcan
USER appuser
CMD ["./vulcan-whatever"]
