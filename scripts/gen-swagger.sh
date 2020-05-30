#!/bin/bash

swagger validate --stop-on-error ../api/builder-api-swagger.yaml
retVal=$?
if [ $retVal -ne 0 ]; then
  echo "Builder API Swagger YAML is not valid"
  exit $retVal
else
#  swagger generate client -A builderapi -f ../api/builder-api-swagger.yaml -t ../pkg/ \
#    --client-package=builderapiclient
  swagger generate server -f ../api/builder-api-swagger.yaml -t ../internal/api/builder/ \
    --exclude-main # -A builderapi
fi
