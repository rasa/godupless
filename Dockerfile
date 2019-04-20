FROM golang:alpine as builder
LABEL maintainer="Ross Smith II <ross@smithii.com>"

ENV PATH /go/bin:/usr/local/go/bin:$PATH
ENV GOPATH /go

#RUN	apk add --no-cache \
#	ca-certificates

COPY . /go/src/github.com/rasa/godupless

RUN set -x \
	&& apk add --no-cache --virtual .build-deps \
		git \
		gcc \
		libc-dev \
		libgcc \
		make \
	&& cd /go/src/github.com/rasa/godupless \
	&& make vendor static \
	&& mv godupless /usr/bin/godupless \
	&& apk del .build-deps \
	&& rm -rf /go \
	&& echo "Build complete."

FROM alpine:latest

COPY --from=builder /usr/bin/godupless /usr/bin/godupless
# COPY --from=builder /etc/ssl/certs/ /etc/ssl/certs

ENTRYPOINT [ "godupless" ]
CMD [ "--help" ]
