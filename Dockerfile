# STEP 1 build executable binary
FROM golang:1.11.2 as builder

ADD . /go/src/gitlab.com/terrabot/service-template

RUN cd /go/src/gitlab.com/terrabot/service-template \
    && apt-get update \
    && apt install xz-utils \
    && apt-get install ca-certificates \
    && update-ca-certificates \
    && wget --quiet https://github.com/upx/upx/releases/download/v3.95/upx-3.95-amd64_linux.tar.xz \
    && tar -xf upx-3.95-amd64_linux.tar.xz \
    && upx-3.95-amd64_linux/upx --version \
    && CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -a -installsuffix cgo -o app.bin . \
    && ./upx-3.95-amd64_linux/upx -k --best --ultra-brute -o app app.bin \
    && ./upx-3.95-amd64_linux/upx -t app

# STEP 2 build a small image
# start from scratch
FROM scratch

# Copy certificates
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

# Copy our static executable
COPY --from=builder /go/src/gitlab.com/terrabot/service-template/app .

EXPOSE 8080

CMD ["./app"]