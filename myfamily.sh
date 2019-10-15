#!/bin/bash
curl -s https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json \
	| jq -r --arg MANID "$HERO_ID" '.[] | select(.id == $MANID) | .connections.relatives'