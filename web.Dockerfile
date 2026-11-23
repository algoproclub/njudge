ARG PROJECT_NAME

FROM node:26-trixie AS frontend-builder
WORKDIR /build
COPY package.json ./
RUN npm install
COPY src/ ./src
COPY gulpfile.js ./
COPY internal/web/templates/ ./internal/web/templates
RUN npx gulp

FROM ${PROJECT_NAME}-base AS njudge-bin

FROM debian:trixie-slim
ENV TZ=Europe/Budapest
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates tzdata pandoc \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=njudge-bin /app/njudge ./njudge
COPY configs/docker/web.yaml ./web.yaml
COPY --from=frontend-builder /build/static ./static

CMD ["./njudge", "web"]
