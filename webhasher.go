// +build ignore
// wget --no-check-certificate --quiet  --method POST   --timeout=0  --header 'Content-Type: application/json'  --body-data '[ {"phrase": "test"}, {"phrase": "test_2"}]'    'http://localhost:8080/get-phrase-hash'



package main

import (
    "crypto/sha256"
    "encoding/json"
    "encoding/binary"
    "log"
    "net/http"
)

type Hashed struct{
    Phrase string `json:"phrase"`
    Hash int64    `json:"hash"`
}

const SizeOfInt64 = 8

func handler(w http.ResponseWriter, r *http.Request) {
    var hashed_phrases []Hashed
    if err := json.NewDecoder(r.Body).Decode(&hashed_phrases); err != nil {
        log.Println(err)
        return
    }
    
    
    for i := 0; i < len(hashed_phrases); i++{
        h := sha256.Sum256([]byte(hashed_phrases[i].Phrase))
        hashed_phrases[i].Hash, _ = binary.Varint(h[0:SizeOfInt64])
    }
    
    if err := json.NewEncoder(w).Encode(hashed_phrases); err != nil {
        log.Println(err)
        return
    }
    
}

func main() {
    http.HandleFunc("/get-phrase-hash/", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
