FROM golang:1.19-alpine as builder

RUN mkdir /build
ADD . /build/
WORKDIR /build/

RUN go build -o gowebapp

FROM alpine:3.17

ARG UID=1000
ARG GID=1000

RUN addgroup -g 1000 "${GID}" && adduser -u "${UID}" -G "${GID}" -h /home/1000 -D "${UID}" -s /sbin/nologin

# Run container as non root user
USER 1000

COPY --from=builder /build/gowebapp /app/

WORKDIR /app
CMD ["./gowebapp"]