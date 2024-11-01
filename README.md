# microblog


docker build -t   microblog-command-service  -f cmd/command_service/Dockerfile .

docker tag  microblog-command-service  sebarray/microblog-command-service
