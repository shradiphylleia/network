# network
# tcp
echo server implemented in go

tcp is a communication protocol allowing two computers to exchange data over a network

a tcp server opens a port, waits for connection , accpets clients, reads bytes, sends byte back

echo server: receives data and sends the same data back

servers use flags as different env need different configs such as ports for dev and prod, flags make them configurable

ports are number: InVar
host string StringVar

&config.Host: config is struct then Host is the variable inside it & gives the memory address of this variable

host+port would be where the server listens for connections
host: which ip address to bind to: 127.0.0.1 , 0.0.0.0 would listen on all available network interfaces

port:
ports below 1024 would need admin/root permissions
we can have multiple services run, each service listens on a different port [how is the choice made?]
localhost:3000 etc

http: 80
https: 443
psql: 5432
redis: 6379

for my implementation:
0.0.0.0:8080 := accept TCP connections on port 8080 from any interface

flag.Parse()
read the args and apply them to flags, without parse it would not read even if i would have set them up, fallbacks
 define the flags, parse and store values


Go exports functions with CAPITAL letters.
Folders == packages
module-name/folder-name :: import path
