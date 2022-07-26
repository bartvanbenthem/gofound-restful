#!/bin/bash
curl -X GET http://localhost:4000/v1/healthcheck
curl -X GET http://localhost:4000/v1/posts/123

BODY='{"title":"testing","content":"test content to display","author":"bartb","img_urls":["https://img.nl/01", "https://img.nl/02"]}'
curl -i -d "$BODY" localhost:4000/v1/posts

# test empty body error
curl -X POST localhost:4000/v1/posts