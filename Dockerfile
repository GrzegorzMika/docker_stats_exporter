FROM alpine
LABEL author="Grzegorz Mika" source="https://github.com/GrzegorzMika/docker_stats_exporter"
ENTRYPOINT ["/docker_stats_exporter"]
COPY docker_stats_exporter /