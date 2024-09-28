FROM alpine
LABEL author="Grzegorz Mika"
LABEL org.opencontainers.image.source="https://github.com/GrzegorzMika/docker_stats_exporter"
LABEL org.opencontainers.image.licenses=MIT
LABEL org.opencontainers.image.description="Docker Stats exporter for Prometheus"
ENTRYPOINT ["/docker_stats_exporter"]
COPY docker_stats_exporter /