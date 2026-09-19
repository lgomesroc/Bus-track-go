FROM golang:1.26-bookworm AS builder

WORKDIR /app

COPY backend/go.mod backend/go.sum ./backend/
RUN cd backend && go mod download

COPY backend ./backend

RUN cd backend && CGO_ENABLED=1 go build -o /app/bustrack .

FROM debian:bookworm-slim

RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates libaio1 unzip wget \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

RUN wget -q https://download.oracle.com/otn_software/linux/instantclient/2326000/instantclient-basic-linux.x64-23.26.0.0.0.zip \
    && unzip instantclient-basic-linux.x64-23.26.0.0.0.zip -d /opt/oracle \
    && rm instantclient-basic-linux.x64-23.26.0.0.0.zip

ENV LD_LIBRARY_PATH=/opt/oracle/instantclient_23_26

COPY --from=builder /app/bustrack /app/bustrack
COPY backend/start.sh /app/start.sh

RUN chmod +x /app/start.sh

EXPOSE 8080

CMD ["/app/start.sh"]
