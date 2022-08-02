#!/bin/bash

echo '****** Updating Validators List ******'
echo ''

VALIDATORS_QUERY=$(nomic validators | tee . validators_list_new.yaml )

echo $VALIDATOR_QUERY
echo ''
echo '****** Saved updated list of validators! ******'
