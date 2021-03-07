FROM golang:1.14.3 as BUILDER

WORKDIR /build

COPY ./go.mod ./
COPY ./go.sum ./

RUN go mod download

COPY . ./

ENV CGO_ENABLED=0
RUN go build -ldflags="-s -w -X github.com/drakejin/fiber-aws-serverless/config.Release=$(git log --format='%H' -n 1)"  -o ./myapp ./main.go

FROM alpine:3.12

WORKDIR /

COPY --from=BUILDER /build/myapp /bin/myapp

ENTRYPOINT ["myapp", "-h"]