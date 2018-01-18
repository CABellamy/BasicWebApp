#!/usr/bin/env bash
docker kill basicwebapp && docker rm basicwebapp
docker build -t basicwebapp .
docker run -it -d --rm --name basicwebapp -p 8000:8000 basicwebapp:latest
java -Dwebdriver.chrome.driver=chromedriver -jar selenium-server.jar
