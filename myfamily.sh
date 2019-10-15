#!/bin/bash
curl -s https://raw.githubusercontent.com/kigiri/superhero-api/master/api/all.json | jq -r --arg MAN_ID "$HERO_ID" '.[] | select(.id == ($MAN_ID|tonumber)) | .connections.relatives'