This is a follow-up practice on [Ewan's blog](https://ewanvalentine.io/microservices-in-golang-part-1/).

# How to start service

The system is organized via docker compose. Run following code to launch:
    
    $ docker-compose -p shippy up -d

This will launch:

- consul (dev only): for service discovery
- consignment service
- vessel service

As ip address of consul is important for other client to interact, we are using the default container network for system (i.e. **shippy_default** in `docker network ls` output). So any client running in host network could directly interactive with this system, while client running in container could interative with this system after being added to the **shippy_default** network.

For example:

- For micro related tools, just specify `--registry=consul` when launching (the `--registry_address` is not needed, since docker host has has access to the container network)

- For consul web, just open **localhost:8500**.

- For consignment client:

    - if running directly from host: just specify `--registry=consul`
    - if running within container, run as: `$ docker run -it --rm -e MICRO_REGISTRY=consul -e MICRO_REGISTRY_ADDRESS=dev-consul --network=shippy_default shippy-cli-consignment`. (`dev-consul` is the host name of consul, defined in docker-comopse.yml)
