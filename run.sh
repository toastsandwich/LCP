#!/bin/bash

# Color definitions
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${YELLOW}Starting server....${NC}"

if ./bin/server; then 
    echo -e "${GREEN}Server started successfully!${NC}"
else
    echo -e "${RED}Error occurred while starting the server!!!${NC}"
fi
