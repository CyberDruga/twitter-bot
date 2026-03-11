#!/usr/bin/env bash

# curl -X GET "https://scrapebadger.com/v1/twitter/users/cyberdruga/latest_tweets" \

args=(
  --get 
  --data-urlencode "query=from:cyberdruga" 
  --data-urlencode "query_type=Latest" 
  -H "x-api-key: $(cat .api-keys-scrapebadgers)" 
  "https://scrapebadger.com/v1/twitter/tweets/advanced_search" 
)

curl "${args[@]}" | tee result.json

  # --data "query_type=Latest" \
