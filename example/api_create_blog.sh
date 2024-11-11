#!/bin/bash

curl --location --request POST 'http://localhost:8080/v1/blog' \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "authorId": "afikrim",
    "title": "My first blog",
    "content": "Hello World"
  }'
