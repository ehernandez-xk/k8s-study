FROM alpine:3.10

EXPOSE 8080

RUN apk --no-cache add curl bind-tools

COPY bin/apphostname /usr/apphostname

CMD /usr/apphostname