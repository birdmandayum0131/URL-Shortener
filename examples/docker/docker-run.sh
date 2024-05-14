docker system prune --all --force
docker compose -f ./examples/docker/docker-compose.yaml up -d --force-recreate #--build