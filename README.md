# caliconodedockerstats
A simple Golang tool to pull calico-node's pull_count from Docker Hub and present it via a Prometheus target.
## Documentation:
This documentation is a work-in-progress and should not yet be considered complete. Hopefully it is still of some use in this state.
## Building:
```
git pull && docker build -t caliconodedockerstats-`git log --pretty=format:'%h' -n 1` .
```
## Testing:
Run it headless like this, grabbing the container ID:
```
export CALICONODEDOCKERSTATS_ATTR_NAME="pull_count" && export CALICONODEDOCKERSTATS_TARGET_NAME="https://hub.docker.com/v2/repositories/calico/node/" && export CID=$(docker run --detach --env-file ./env.list caliconodedockerstats-`git log --pretty=format:'%h' -n 1`)
```
Then, just grab your docker container's IP from the container ID and curl port 9088, like this:
```
curl http://`docker inspect -f '{{range.NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $CID`:9088/metrics
```
Finally, you can stop it like this:
```
docker stop $CID
```
