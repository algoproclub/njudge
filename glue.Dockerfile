ARG PROJECT_NAME

FROM ${PROJECT_NAME}-base AS njudge-bin

FROM debian:trixie-slim
ENV TZ=Europe/Budapest
RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates tzdata pandoc \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=njudge-bin /app/njudge ./njudge
COPY configs/docker/glue.yaml ./glue.yaml

CMD ["./njudge", "glue"]
