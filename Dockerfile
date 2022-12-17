# Build the project
FROM golang:1.19 as builder

LABEL org.opencontainers.image.source="https://github.com/KevinMidboe/planetposen-images"

WORKDIR /go/src/github.com/kevinmidboe/planetposen-images
ADD . .

RUN make build
# RUN make test

# Create production image for application with needed files
FROM iron/go

EXPOSE 8000

RUN apk add --no-cache ca-certificates

RUN mkdir -p /etc/secrets
VOLUME /etc/secrets

RUN mkdir -p /var/log/planetposen-logs
VOLUME /var/log/planetposen-logs

COPY --from=builder /go/src/github.com/kevinmidboe/planetposen-images .

ENV GOOGLE_APPLICATION_CREDENTIALS=/etc/secrets/sa-key.json

CMD ["./main"]