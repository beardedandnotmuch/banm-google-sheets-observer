#!/bin/bash

set -euo pipefail

ssh root@sheets.beardedandnotmuch.com \
    'cd banm-google-sheets-observer && docker-compose pull && docker-compose up -d'
