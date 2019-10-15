#!/bin/bash
curl -s https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json \
	| jq -r --arg HERO_ID "$HERO_ID" '.[] | select(.id == 1) | .connections.relatives'