#!/bin/sh
set -e

rm -rf /tmp/oracle-wallet /tmp/bustrack-wallet.zip
mkdir -p /tmp/oracle-wallet

if [ -f /etc/secrets/bustrack-wallet.zip.b64 ]; then
    base64 -d /etc/secrets/bustrack-wallet.zip.b64 > /tmp/bustrack-wallet.zip
    unzip -q /tmp/bustrack-wallet.zip -d /tmp/oracle-wallet

    cat > /tmp/oracle-wallet/bustrack/sqlnet.ora <<'EOF'
WALLET_LOCATION = (SOURCE = (METHOD = file) (METHOD_DATA = (DIRECTORY="/tmp/oracle-wallet/bustrack")))
SSL_SERVER_DN_MATCH=yes
EOF

    export TNS_ADMIN=/tmp/oracle-wallet/bustrack
fi

exec /app/bustrack
