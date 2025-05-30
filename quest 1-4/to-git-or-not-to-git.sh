#!/bin/bash
data=$(curl -s https://01.tomorrow-school.ai/assets/superhero/all.json)
name=$(jq -r '.[] | select(.id == 170) | .name' <<< "$data")
durability=$(jq -r '.[] | select(.id == 170) | .powerstats.power' <<< "$data")
gender=$(jq -r '.[] | select(.id == 170) | .appearance.gender' <<< "$data")
printf "$name\n$durability\n$gender\n"