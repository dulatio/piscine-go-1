#!/bin/bash
curl -s https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json \
	| jq --arg HERO_ID "$HERO_ID" -r '.[] | select(.id == $HERO_ID) | .connections.relatives'