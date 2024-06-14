#!/bin/bash

docker-compose -f service_test.yml build
docker-compose -f service_test.yml up -d

sleep 40

output1=$(grpcurl -d '{"username": "bob", "text":"abcd"}' -plaintext 0.0.0.0:11777 posts.PostsService/CreatePost)
if [ $? -ne 0 ]; then
  docker-compose -f service_test.yml down > /dev/null 2>&1
  exit 1
fi

POST=$(echo $output1 | jq -r '.postId')

output2=$(grpcurl -d "{\"username\": \"bob\", \"post_id\":\"$POST\"}" -plaintext 0.0.0.0:11777 posts.PostsService/GetPost)
if [ $? -ne 0 ]; then
  docker-compose -f service_test.yml down > /dev/null 2>&1
  exit 1
fi

text=$(echo $output2 | jq -r '.text')

if [ "$text" == "abcd" ]; then
  echo "Test passed"
else
  echo "Test failed"
  docker-compose -f service_test.yml down > /dev/null 2>&1
  exit 1
fi

docker-compose -f service_test.yml down > /dev/null 2>&1
