FROM golang:1.24 as builder

WORKDIR /opt/app

COPY . .

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o flowbot -ldflags "-X github.com/sysatom/framework/version.Buildstamp=`date -u '+%Y-%m-%dT%H:%M:%SZ'` -X github.com/sysatom/framework/version.Buildtags=`git describe --tags`" github.com/sysatom/framework/cmd


FROM alpine:latest

RUN apk add --no-cache bash curl tzdata

WORKDIR /opt/app

COPY --from=builder /opt/app/flowbot .
COPY deployments/entrypoint.sh .

RUN chmod +x flowbot && chmod +x entrypoint.sh

ENTRYPOINT ./entrypoint.sh
EXPOSE 6060
