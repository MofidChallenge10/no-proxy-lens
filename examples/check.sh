#!/usr/bin/env sh
set -eu
go run ./cmd/no-proxy-lens --env 'localhost,.internal.example,10.0.0.0/8' https://api.internal.example:8443/health
