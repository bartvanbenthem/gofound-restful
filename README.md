# gofound-blogger
RESTful Blog webservice


### example requests
```bash
# test GET
curl -X GET http://localhost:4000/v1/healthcheck
curl -X GET http://localhost:4000/v1/posts/123

# test POST
BODY='{"title":"testing","content":"test content to display","author":"bartb","img_urls":["https://img.nl/01", "https://img.nl/02"]}'
curl -i -d "$BODY" localhost:4000/v1/posts

# test POST empty body error
curl -X POST localhost:4000/v1/posts

# test PUT
BODY='{"title":"updated test","content":"updated test content","author":"bartb","img_urls":["https://img.nl/98", "https://img.nl/99"]}'
curl -X PUT -d "$BODY" localhost:4000/v1/posts/1

# test delete
curl -X DELETE localhost:4000/v1/posts/1

```
