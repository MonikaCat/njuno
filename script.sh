#!/bin/bash

echo '****** Updating Validators List ******'
echo ''

VALIDATORS_QUERY=$(nomic validators | sed 's/\- /- validator: \n  validator_address: /g' | sed 's/\	VOTING POWER/  voting_power/' | sed 's/\	MONIKER/  moniker/' | sed 's/\	COMMISSION/  commission/' |  sed 's/\	IDENTITY/  identity/' |  sed 's/\	DETAILS:/  details: /' | tee . validators_list.yaml )

echo $VALIDATOR_QUERY
echo ''
echo '****** Saved updated list of validators! ******'
