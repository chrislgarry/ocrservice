FROM debian:bullseye-slim
LABEL maintainer="Chris Garry <chrislgarry@gmail.com>"

RUN apt-get update \
    && apt-get install -y \
      ca-certificates \
      golang=2:1.15~1 \
      libtesseract-dev=4.1.1-2.1 \
      tesseract-ocr=4.1.1-2.1

ENV GO111MODULE=on
ENV GOPATH=${HOME}/go
ENV PATH=${PATH}:${GOPATH}/bin

ADD . $GOPATH/src/ocrservice
WORKDIR $GOPATH/src/ocrservice
RUN go get -v ./... && go install .

CMD ["mercariocrservice"]
