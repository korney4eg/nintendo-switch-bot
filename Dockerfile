FROM golang:alpine AS builder
# Git is required for fetching the dependencies.
WORKDIR $GOPATH/src/mypackage/myapp/
COPY . .
RUN go build -o /go/bin/nintendo-switch-bot

############################
# STEP 2 build a small image
############################

FROM alpine:latest
# Copy our static executable.
RUN apk update && apk add --no-cache apk-cron
COPY --from=builder /go/bin/nintendo-switch-bot /usr/local/bin/
WORKDIR /app
CMD ["nintendo-switch-bot"]
