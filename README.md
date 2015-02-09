# SSHClientTest
Golang SSH Client Test

## How to build ?
 - `$git clone <this repo>`
 - `$go get golang.org/x/crypto/ssh`
 - `$go build src/SSHClient.go`
 
## How to run ?
 - `$./SSHClient <user> <host:port> <command>`
 
 It will prompt for a password, leave it empty if you want to use your private keys located in $HOME/.ssh/id_rsa
