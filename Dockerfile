FROM golang:1.4.2-onbuild

RUN mkdir /data

ENV TOKEN azerty1234

VOLUME ["/data"]

CMD ["app", "/data"]

EXPOSE 8888
