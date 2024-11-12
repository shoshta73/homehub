set -x
set -e

docker build -f ./docker/frontend.Dockerfile .
docker build -f ./docker/backend.Dockerfile .
