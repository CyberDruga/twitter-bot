#!/usr/bin/env nu

#Helper script to manage rules 
def main [] {}

# Gets a list of rules 
def "main get_rules" [] {
	http get https://api.twitterapi.io/oapi/tweet_filter/get_rules --headers { X-API-Key: $"(cat ./.api-key-twitchapi)" }
}

# Deletes a rule 
def "main delete_rule" [
	rule: string # Rule to be deleted
] {
	{ "rule_id": $"($rule)" } 
		| to json
		| http delete https://api.twitterapi.io/oapi/tweet_filter/delete_rule --headers { X-API-Key: $"(cat ./.api-key-twitchapi)" }
}

def "main info" [] {
	http get "https://api.twitterapi.io/oapi/my/info" --headers { X-API-Key: $"(cat ./.api-key-twitchapi)"}
}

