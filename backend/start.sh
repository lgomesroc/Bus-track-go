#!/bin/sh
set -e

mkdir -p /tmp/oracle-wallet

if [ -f /etc/secrets/bustrack-wallet.zip ]; then
    unzip -q /etc/secrets/bustrack-wallet.zip -d /tmp/oracle-wallet
    export TNS_ADMIN=/tmp/oracle-wallet/bustrack
fi

exec /app/bustrack
