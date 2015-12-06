FROM golang:1.4.2-onbuild

RUN mkdir /data
ADD run.sh /run.sh

ENV TOKEN azerty1234

VOLUME ["/data"]

CMD ["sh", "/run.sh"]

EXPOSE 8888
