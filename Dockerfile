FROM golang:1.9

ENV SOURCES /go/src/customer-consumer

ENV GOOGLE_CLOUD_PROJECT chefd-dev-190417
ENV NETSUITE_MESSAGE_BETA netsuite-message-beta
ENV NETSUITE_PRODUCT_BETA netsuite-product-beta
ENV PULL_FOR_NETSUITE_PRODUCT_BETA netsuite-product-beta-subscriber

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go install

ENV PORT 8082
EXPOSE 8082

ENTRYPOINT customer-consumer

