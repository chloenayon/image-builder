#!/bin/bash

set -euxo pipefail

SONAR_SCANNER_CLI_VERSION=${SONAR_SCANNER_CLI_VERSION:-4.6.2.2472}

export SONAR_SCANNER_OPTS="-Djavax.net.ssl.trustStore=schutzbot/RH-IT-Root-CA.keystore -Djavax.net.ssl.trustStorePassword=$KEYSTORE_PASS"
sudo dnf install -y unzip
curl "https://binaries.sonarsource.com/Distribution/sonar-scanner-cli/sonar-scanner-cli-$SONAR_SCANNER_CLI_VERSION-linux.zip" -o sonar-scanner-cli.zip
unzip -q sonar-scanner-cli.zip

SONAR_SCANNER_CMD="$(pwd)/sonar-scanner-$SONAR_SCANNER_CLI_VERSION-linux/bin/sonar-scanner"

$SONAR_SCANNER_CMD -Dsonar.projectKey=osbuild:image-builder \
                   -Dsonar.sources=. \
                   -Dsonar.host.url=https://sonarqube.corp.redhat.com \
                   -Dsonar.login="$SONAR_SCANNER_TOKEN" \
                   -Dsonar.pullrequest.branch="$CI_COMMIT_BRANCH" \
                   -Dsonar.pullrequest.key="$CI_COMMIT_SHA" \
                   -Dsonar.pullrequest.base="main" \
				   -Dsonar.c.file.suffixes=-

SONARQUBE_URL="https://sonarqube.corp.redhat.com/dashboard?id=osbuild%3Aimage-builder&pullRequest=$CI_COMMIT_SHA"
# Report back to GitHub
curl \
    -u "${SCHUTZBOT_LOGIN}" \
    -X POST \
    -H "Accept: application/vnd.github.v3+json" \
    "https://api.github.com/repos/osbuild/image-builder/statuses/${CI_COMMIT_SHA}" \
    -d '{"state":"success", "description": "SonarQube scan sent for analysis", "context": "SonarQube", "target_url": "'"${SONARQUBE_URL}"'"}'