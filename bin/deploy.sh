#!/bin/bash

set -e

TAG=$CI_PIPELINE_ID
PREFIX="$EB_APPLICATION"

zip app.zip ./* -r

aws s3 cp app.zip s3://$EB_BUCKET/$PREFIX/app-$CI_PIPELINE_ID.zip
rm app.zip

echo "Creating new EB version"
echo $CI_PIPELINE_ID
aws elasticbeanstalk create-application-version \
    --application-name $EB_APPLICATION \
      --version-label $CI_PIPELINE_ID \
        --source-bundle S3Bucket=$EB_BUCKET,S3Key=$PREFIX/app-$CI_PIPELINE_ID.zip

echo "Updating EB environment"
aws elasticbeanstalk update-environment \
    --environment-name $EB_ENV \
      --version-label $CI_PIPELINE_ID