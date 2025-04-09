#!/bin/bash

set -a
source .env
set +a

# Add terraform commands
terraform apply
