#!/bin/bash

set -e

TAG=$CI_PIPELINE_ID-api
PREFIX="$EB_APPLICATION"

zip app.zip ./* -r

aws s3 cp app.zip s3://$EB_BUCKET/$PREFIX/app-$TAG.zip
rm app.zip

echo "Creating new EB version"
echo $TAG
aws elasticbeanstalk create-application-version \
    --application-name $EB_APPLICATION \
      --version-label $TAG \
        --source-bundle S3Bucket=$EB_BUCKET,S3Key=$PREFIX/app-$TAG.zip

echo "Updating EB environment"
aws elasticbeanstalk update-environment \
    --environment-name $EB_ENV \
      --version-label $TAG
