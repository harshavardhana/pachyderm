FROM postgres:9.6.1-alpine
RUN apk update \
  && apk add ca-certificates wget \
  && update-ca-certificates \
  && wget https://github.com/stedolan/jq/releases/download/jq-1.5/jq-linux64 -O /usr/bin/jq \
  && chmod +x /usr/bin/jq
ADD ./to_sql /usr/bin/
