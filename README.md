# Bambu SSDP Relay
Since Bablulab do not support to connect to an other subnet, this small relay was written to have a work around for this problem.

## Build Docker image
```
docker build -f Containerfile -t bambu-relay .
```

## Run container

The container must be started somewhere in the network where the printer is located.

You need also to setup your firewall properly. Babulab need access to the port 8883/tcp (MQTT) and 300/tcp (video) to run properly.


The environment variable DESTINATION_ADDRESS must be set to run the container.
This is the address of the machine where bambuStudio is running on.

A example start of the container could look like this inf you use the self build container:

```
docker run --rm -it -e DESTINATION_ADDRESS=192.168.0.105:2021 -p 2021:2021/udp bambu-relay
```

If you prefer the prebuild container following command should do the job:

```
docker run --rm -it -e DESTINATION_ADDRESS=192.168.0.105:2021 -p 2021:2021/udp ghcr.io/andreasschwalb/babulabrelay:latest
```


