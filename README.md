## Envspitter

Envspitter is a simple http server spitting env variables in plain text.
Very basic script to use with docker for example.

Succesfully over-used it to learn Kubernetes, again very basic, but get the job done !

## Build

Under linux just :

 	make

or osx :

 	make osx

and of course :

  	docker build --rm -t jbaptiste/envspitter .

You can also pull the working docker at the same address : 

  	docker pull jbaptiste/envspitter

## Example

The server is listening on 127.0.0.1:8081 by default, override it with -l or --listen

  	$ docker run -t -i jbaptiste/envspitter --help
  	Usage of /usr/local/bin/envspitter:
      	-l, --listen="127.0.0.1:8081"  IPaddress to listen from 
 
Running as docker
 
  	$ docker run -t -p 9090:9090 -i jbaptiste/envspitter --listen 0.0.0.0:9090
  	2014/07/04 12:48:01 Listening on 0.0.0.0:9090 ...