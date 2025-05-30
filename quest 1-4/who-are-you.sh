#!/bin/ bash
data=$(curl -s https://01.tomorrow-school.ai/assets/superhero/all.json)
name=$(echo "$data"| jq -r '.[] | select(.id == 70) | .name')
echo "\"$name\""
