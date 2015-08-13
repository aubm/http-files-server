FROM golang:1.4.2-onbuild

RUN mkdir /data

VOLUME ["/data"]

CMD ["app", "/data", "azerty1234"]

EXPOSE 8888