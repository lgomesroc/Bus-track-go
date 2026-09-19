#!/bin/sh
set -e

rm -rf /tmp/oracle-wallet /tmp/bustrack-wallet.zip
mkdir -p /tmp/oracle-wallet

if [ -f /etc/secrets/bustrack-wallet.zip.b64 ]; then
    base64 -d /etc/secrets/bustrack-wallet.zip.b64 > /tmp/bustrack-wallet.zip
    unzip -q /tmp/bustrack-wallet.zip -d /tmp/oracle-wallet
    export TNS_ADMIN=/tmp/oracle-wallet/bustrack
fi

exec /app/bustrack
