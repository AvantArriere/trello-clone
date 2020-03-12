FROM golang:1.12.1-stretch as packages
LABEL maintainer="leeyjeen@naver.com"

# set go module mode on without setting GOPATH
ENV GO111MODULE=on

WORKDIR /usr/src/app
COPY . .
RUN go build -o trello_api

FROM golang:1.12.1-stretch as stage

# install netcat & jq
# remove apt install lists
RUN apt-get update && \
    apt-get -y install netcat && \
    apt-get -y install jq && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /usr/src/app

COPY --from=packages /usr/src/app/scripts/entrypoint.sh /usr/src/app/scripts/entrypoint.sh
COPY --from=packages /usr/src/app/trello_api /usr/src/app/trello_api

RUN chmod +x ./scripts/entrypoint.sh

CMD ["sh", "./scripts/entrypoint.sh"]