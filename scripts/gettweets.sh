#!/usr/bin/env nix-shell
#! nix-shell -p websocat -i bash

args=(
	-B 2000000
	"wss://ws.twitterapi.io/twitter/tweet/websocket"
	-H "x-api-key:$(cat ./.api-key-twitchapi)"
)

websocat "${args[@]}" 

