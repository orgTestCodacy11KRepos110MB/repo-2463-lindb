FROM ubuntu

RUN mkdir -p /bin/lindb
WORKDIR /bin/lindb

COPY ./bin .

RUN ./lind broker init-config \
    && ./lind storage init-config \
    && sed -i 's:localhost:etcd:g' broker.toml \
    && sed -i 's:localhost:etcd:g' storage.toml

EXPOSE 9000
EXPOSE 2891