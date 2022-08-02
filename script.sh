#!/bin/bash

echo '****** Updating Validators List ******'
echo ''

VALIDATORS_QUERY=$(nomic validators | tee . validators_list.yaml yamlfmt validators_list.yaml)

echo $VALIDATOR_QUERY
echo ''
echo '****** Saved updated list of validators! ******'
