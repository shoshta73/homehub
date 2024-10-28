#!/bin/bash

set -e
set -x

rm -rf docs/README.md
cp README.md docs/README.md

rm -rf docs/CONTRIBUTING.md
cp CONTRIBUTING.md docs/CONTRIBUTING.md

rm -rf docs/CODE_OF_CONDUCT.md
cp CODE_OF_CONDUCT.md docs/CODE_OF_CONDUCT.md

rm -rf docs/CHANGELOG.md
cp CHANGELOG.md docs/CHANGELOG.md

mdbook build
