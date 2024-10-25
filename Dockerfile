FROM alpine:latest AS builder
RUN apk add --update nodejs npm
COPY --from=golang:1.22.4-alpine /usr/local/go/ /usr/local/go/
ENV PATH="/usr/local/go/bin:${PATH}"
WORKDIR /workspace
COPY . .
RUN npm ci
RUN ./scripts/build

FROM alpine:latest AS runner
COPY --from=builder /workspace/build/minimalist-go-react-app /
EXPOSE 8080
CMD ["/minimalist-go-react-app"]
