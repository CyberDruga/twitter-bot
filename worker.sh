#!/usr/bin/env nix-shell
#! nix-shell -p websocat -i bash

cleanup () {
	echo
	echo cleaning up
	pkill -P "$$"
	exit
}

trap cleanup EXIT SIGINT

send_webhook() {

	url="${1/x\.com/fixupx.com}"

	message="<ping> $url" # <- message goes here

	curl -X POST "$(cat ./.webhook-url)" --data-raw "{\"content\": \"$message\"}" -H 'content-type: application/json' 
}

args=(
	-B 2000000
	"wss://ws.twitterapi.io/twitter/tweet/websocket"
	-H "x-api-key:$(cat ./.api-key-twitchapi)"
)

websocat "${args[@]}"  | while read -r message ; do 
	event_type=$(jq ' .event_type ' -r <<< "$message")
	case "$event_type" in
		tweet)

			[[ "$(jq ' .rule_id ' -r <<< "$message")" != "$(cat .rule_id)" ]] && continue

			echo "sending messages"
			jq ' .tweets[] | .url ' -r <<< "$message" \
				| tac \
				| while read -r line ; do send_webhook "$line" ; done
			;;
		*)
			echo "$message"
			;;
	esac
done 

