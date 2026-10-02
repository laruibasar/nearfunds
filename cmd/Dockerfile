# For improve security, we could use base images we control.
FROM golang:1.26-alpine AS build

WORKDIR /src

COPY go.mod go.sum ./
# If we have some online repo where there is the packages we consider safe to use,
# we can set go env GOPROXY to limit.
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /fund-service ./cmd 

FROM alpine:3.22

RUN apk add --no-cache ca-certificates tzdata \
  && adduser -D -u 10001 appuser

COPY --from=build /fund-service /usr/local/bin/fund-service

USER appuser 
EXPOSE 8080

ENTRYPOINT ["/usr/local/bin/fund-service"]
