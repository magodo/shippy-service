This is a follow-up practice on [Ewan's blog](https://ewanvalentine.io/microservices-in-golang-part-1/).

# How to start service

The system is organized via docker compose. Run following code to launch:
    
    $ # launch database first
    $ docker-compose up -d database
    $ # wait a second, then launch the remaining parts
    $ docker-compose -p shippy up -d

# Send request to micro api

    $ curl -XPOST -H 'Content-Type: application/json' -d '{ "service": "shippy.srv.user", "method": "UserService.Create", "request": { "user": { "email": "foo_email", "password": "foo_password", "name": "foobar", "company": "foo_company" } } }'   http://localhost:8080/rpc
