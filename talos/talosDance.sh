#!/usr/bin/env bash

echo "-= Pulumi up.."
(cd ../pulumi-pxmx; pulumi up --yes)

echo "-= Exporting stack"
STACK_JSON=stack.json
(cd ../pulumi-pxmx; pulumi stack export) > ${STACK_JSON}

CONTROL_PLANE_MAC=$(cat ${STACK_JSON} | jq -r .deployment.resources[2].outputs.networks[0].macaddr)
WORKER_MAC=$(cat ${STACK_JSON} | jq -r .deployment.resources[3].outputs.networks[0].macaddr)

echo "-=-= CONTROL_PLANE_MAC: ${CONTROL_PLANE_MAC}"
echo "-=-= WORKER_MAC: ${WORKER_MAC}"

echo "-= Resolving IP's"



