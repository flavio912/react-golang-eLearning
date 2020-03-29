#!/bin/bash

go generate ./schema

# Auto Regenerate graphQL schema stuff when it changes
reflex -g '*.graphql' go generate ./schema &

# Run our golang server
fresh