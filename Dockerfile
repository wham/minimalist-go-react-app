# syntax=docker/dockerfile:1
FROM alpine:latest as builder
RUN apk add --update nodejs npm
COPY --from=golang:1.22.4-alpine /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"

COPY . /workspace
WORKDIR /workspace
RUN npm ci --omit=dev
RUN go run ./cmd/build/main.go
RUN ls -la ./build
RUN go build -o ./build/minimalist-go-react-app ./cmd/server/main.go

FROM alpine:latest as runner
COPY --from=builder /workspace/build/minimalist-go-react-app /app/
WORKDIR /app
EXPOSE 8080
#CMD ["sh", "-c", "sleep 10000000 && ./kaja-twirp"]
CMD ["./minimalist-go-react-app"]
