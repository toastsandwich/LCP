#!/bin/bash

set -e

# Color definitions
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
CYAN='\033[0;36m'
NC='\033[0m' # No Color

echo -e "${YELLOW}Building....${NC}"

mkdir -p bin

if templ generate; then
    echo -e "${GREEN}Templ generatation complete.${NC}"
else
    echo -e "${RUN}Templ generatation failed.${NC}"
    exit 1
fi

if go build -o bin/server main.go; then
    echo -e "${GREEN}Build complete!${NC}"
    echo -e "${CYAN}Check the bin directory.${NC}"
else
    echo -e "${RED}Build failed!${NC}"
    exit 1
fi
