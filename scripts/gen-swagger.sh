#!/bin/bash

swagger validate --stop-on-error ../api/builder-api-swagger.yaml
retVal=$?
if [ $retVal -ne 0 ]; then
  echo "Builder API Swagger YAML is not valid"
  exit $retVal
else
  swagger generate server -A builderapi -f ../api/builder-api-swagger.yaml -t ../internal/ --exclude-main
  swagger generate client -A builderapi -f ../api/builder-api-swagger.yaml -t ../pkg/ \
    --skip-models --existing-models=../internal/models \
    --client-package=builderapiclient
fi
