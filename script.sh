#!/bin/bash

echo '****** Updating Validators List ******'
echo ''

QUERY_VALIDATORS_LIST=$(nomic validators | tee . validators_list.yaml)
FORMAT_VALIDATORS_LIST=$(yamlfmt validators_list.yaml)

echo $QUERY_VALIDATORS_LIST
echo $FORMAT_VALIDATORS_LIST
echo ''
echo '****** Saved updated list of validators! ******'
