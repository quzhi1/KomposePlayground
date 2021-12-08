# Copied from https://github.com/andrew-jung/gcp-pubsub-emulator/blob/main/Dockerfile
FROM google/cloud-sdk:alpine AS builder
LABEL stage=builder
RUN gcloud components install pubsub-emulator

FROM openjdk:jre-alpine

COPY --from=builder /google-cloud-sdk/platform/pubsub-emulator /pubsub-emulator

RUN apk --update --no-cache add tini bash
ENTRYPOINT ["/sbin/tini", "--"]
CMD /pubsub-emulator/bin/cloud-pubsub-emulator --host=0.0.0.0 --port=8085
EXPOSE 8085
