#!/bin/bash

echo '****** Updating validators list ******'
echo ''

VALIDATORS_QUERY=$(nomic validators 2>&1 | sed '1 i\
validators: 
' |   tee ${HOME}/.njuno/validators_list.yaml)

echo $VALIDATORS_QUERY
echo ''
echo '****** Saved updated list of validators! ******'