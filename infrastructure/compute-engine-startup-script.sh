#!/bin/bash
export INSTANCE_EXTERNAL_IP=`curl -H "Metadata-Flavor: Google" http://metadata.google.internal/computeMetadata/v1/instance/network-interfaces/0/access-configs/0/external-ip`

echo "startup script"