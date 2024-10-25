FROM alpine:latest AS builder
RUN apk add --update nodejs npm
COPY --from=golang:1.22.4-alpine /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"
ARG APP_NAME
COPY . /workspace/$APP_NAME
WORKDIR /workspace/$APP_NAME
RUN npm ci
RUN ./scripts/test
RUN mkdir -p ./build
RUN ./scripts/build

# FROM alpine:latest AS runner
# COPY --from=builder /workspace/build/minimalist-go-react-app /app/
# WORKDIR /app
# EXPOSE 8080
# CMD ["./minimalist-go-react-app"]
