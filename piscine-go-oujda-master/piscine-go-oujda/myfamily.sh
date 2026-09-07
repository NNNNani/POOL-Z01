#! /bin/bash
curl https://learn.zone01oujda.ma/assets/superhero/all.json | jq -r --arg hid $HERO_ID ' .[] | select( .id==($hid|tonumber)) | (.connections.relatives |gsub("\n"; "\\n"))' | tr '"' ' '