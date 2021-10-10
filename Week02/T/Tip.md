# Tip

```shell
docker inspect [OPTIONS] NAME|ID [NAME|ID...]
```



Return low-level information on Docker objects

Docker inspect provides detailed information on constructs controlled by Docker.

By default, `docker inspect` will render results in a JSON array.



```shell
# Get an instance’s IP address
docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' $INSTANCE_ID
 
# Get an instance’s MAC address
docker inspect --format='{{range .NetworkSettings.Networks}}{{.MacAddress}}{{end}}' $INSTANCE_ID
 
# Get an instance’s log path
docker inspect --format='{{.LogPath}}' $INSTANCE_ID

# Get an instance’s image name
docker inspect --format='{{.Config.Image}}' $INSTANCE_ID

# List all port bindings
# You can loop over arrays and maps in the results to produce simple text output:
docker inspect --format='{{range $p, $conf := .NetworkSettings.Ports}} {{$p}} -> {{(index $conf 0).HostPort}} {{end}}' $INSTANCE_ID
 
 

```

