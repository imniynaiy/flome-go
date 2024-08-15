FROM alpine
COPY _output/main /opt/flome/main
COPY configs /opt/flome/configs
WORKDIR /opt/flome
EXPOSE 8081
ENTRYPOINT ["/opt/flome/main"]