# syntax=docker/dockerfile:1
FROM alpine:latest AS builder
RUN apk add --update nodejs npm
COPY --from=golang:1.22.4-alpine /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"

COPY . /workspace
WORKDIR /workspace
RUN npm ci
RUN ls -la ./scripts
RUN ./scripts/test
RUN mkdir -p ./build
RUN ./scripts/build

FROM alpine:latest AS runner
COPY --from=builder /workspace/build/minimalist-go-react-app /app/
WORKDIR /app
EXPOSE 8080
CMD ["./minimalist-go-react-app"]
