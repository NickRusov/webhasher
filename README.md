# webhasher
A web server for hashing strings

#### Usage: 
`wget --no-check-certificate --quiet  --method POST   --timeout=0  --header 'Content-Type: application/json'  --body-data '[ {"phrase": "test"}, {"phrase": "test_2"}]'    'http://<host>:8080/get-phrase-hash'`

