FROM alpine
COPY _output/dal /opt/dal/dal
COPY configs /opt/dal/configs
WORKDIR /opt/dal
EXPOSE 8081
ENTRYPOINT ["/opt/dal/dal"]