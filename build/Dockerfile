FROM golang:1.24-alpine AS builder

WORKDIR /workdir

COPY go.mod go.sum ./

RUN go mod download && go mod verify

COPY . .

RUN go clean -cache -modcache

RUN CGO_ENABLED=0 go build -o ./app ./main.go

FROM alpine AS runner

COPY --from=builder /workdir/app /binary

COPY build/entrypoint.sh entrypoint.sh

RUN  chmod +x /binary && \
     chmod +x ./entrypoint.sh

CMD ["/bin/sh", "-c", "./entrypoint.sh"]
