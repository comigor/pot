#!/bin/bash

curl --location --request PUT 'http://localhost:8080/v1/blog/some-id-of-the-blog-post' \
  --header 'Content-Type: application/json' \
  --data-raw '{
    "authorId": "afikrim",
    "title": "My Updated blog",
    "content": "Hello World"
  }'
