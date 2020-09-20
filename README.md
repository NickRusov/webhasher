# webhasher
A web server for hashing strings

#### Usage: 
Starting server:

`docker run -p 8080:8080 webhasher:latest`


Checking the server works

`wget --no-check-certificate --quiet  --method POST   --timeout=0  --header 'Content-Type: application/json'  --body-data '[ {"phrase": "test"}, {"phrase": "test_2"}]'    'http://<host>:8080/get-phrase-hash'`
