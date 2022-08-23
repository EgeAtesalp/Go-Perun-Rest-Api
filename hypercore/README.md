# Hypercore implementation of User Session Management

This implementation has three core functionalities.
## Starting a User Session
User inputs their relevant information (as model.UserSession containing IP and Port) and starts their hypercore (nodejs) server.

## Retrieving User Information
User inputs alias of user they pretend to connect to. Information is retrieved from the hypercore server and correctly parsed as model.UserSession

## Automatic syncing between nodes
Through hyperswarm, a module of the hypercore protocol, different user servers are automatically synced with current information.

## Usage
Usage of this implementation is through three different functions:

StartUserSession() starts hypercore server and inputs User information into append-only log.
GetUsers() retrieves information of all users and parses it into a Golang map (map[string]model.UserSession).
FindUserSession() retrieves desired entry of map containing all user information.


## Demo
The screenshots below show Alice creating her server successfully, Alice's Server connecting to Bob's Server succesfully and Alice retrieving Bob's information successfully. The code of the full test in in the Demo Code Screenshot and in api_paymentchannel.go.

![Demo Code](/screenshots/hypercode.png)
![Terminal Output](/screenshots/hypertest.png "Optional Title")
![Hypercore Server Output](/screenshots/hyperserver.png "Optional Title")

## Known Issues

This hypercore implementation was not successfully packaged into a Docker image. As there are many dependencies, both in the Golang and NodeJS part, it is cumbersome to individually install all dependencies on a new system. In addition the current implementation utilizes a system call that open a gnome-terminal. While this is useful for demonstration purposes (see Screenshots in this README), it is not ideal for a final implementation. Even so, the system call should work in Ubuntu and other gnome systems by default.

This implementation was mainly tested in a local network. 