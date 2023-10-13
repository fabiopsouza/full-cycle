FROM golang:1.16

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN apt-get update && \
    apt-get install build-essential librdkafka-dev -y

RUN mkdir test
#COPY /Users/fabipereira/go/pkg/mod/github.com/confluentinc/confluent-kafka-go@v1.9.2-go@v1.9.2/ /go/test/
COPY confluent-kafka-go@v1.9.2/ test/

CMD ["tail", "-f", "/dev/null"]