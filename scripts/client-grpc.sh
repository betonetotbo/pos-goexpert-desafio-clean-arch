#!/bin/bash
docker run -it --rm -v "$(pwd)/..:/mount:ro" \
    --network=host \
    ghcr.io/ktr0731/evans:latest \
    --package pb \
    --host "$(hostname -I | cut -d' ' -f1)" \
    -r repl