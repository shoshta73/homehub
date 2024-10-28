#!/bin/bash

set -e
set -x

rm -rf docs/README.md
cp README.md docs/README.md

rm -rf docs/CONTRIBUTING.md
cp CONTRIBUTING.md docs/CONTRIBUTING.md

mdbook serve .
