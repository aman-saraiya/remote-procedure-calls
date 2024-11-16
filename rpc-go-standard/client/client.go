/*
 * Client to send RPC requests to the server.
 */
package main

import (
	"fmt"
	"log"
	"net/rpc"

	ss "github.com/aman-saraiya/remote-procedure-calls/rpc-go-standard/client/sharedstructs"
)

func invokeAddUserRPC(client *rpc.Client, email, name string) int {
	// Initialize the AddUser RPC request argument
	addUserArg := &ss.AddUserArg{
		Email: email,
		Name:  name,
	}

	// Declare the RPC response variable to store the return value
	var addUserRet ss.AddUserRet
	
	// Calling the UserService's AddUser RPC
	err := client.Call("UserService.AddUser", addUserArg, &addUserRet)
	if err != nil {
		fmt.Errorf("Error calling AddUser: %v\n", err)
	}
	fmt.Printf("Added user with ID: %d\n", addUserRet.ID)
	
	return addUserRet.ID 
}

func invokeGetUserRPC(client *rpc.Client, id int) {
	// Initialize the RPC request argument
	getUserArgs := &ss.GetUserArg{ID: id}

	// Declaring the RPC response variable to store the return value 
	var getUserRet ss.GetUserRet
	
	// Calling the UserService's GetUser RPC
	err := client.Call("UserService.GetUser", getUserArgs, &getUserRet)
	if err != nil {
		log.Fatalf("Error calling GetUser: %v", err)
	}
	fmt.Printf("Retrieved user: ID = %d, Name = %s, Email = %s\n",
		getUserRet.User.ID, getUserRet.User.Name, getUserRet.User.Email)
}

func main() {
	// Attempts to connect to the RPC server
	client, err := rpc.Dial("tcp", "localhost:2024")
	if err != nil {
		fmt.Errorf("Error connecting to server: %v", err)
		return
	}
	defer client.Close()

	// Adding new User
	idOfCreatedUser := invokeAddUserRPC(client, "amansaraiya937@gmail.com", "Aman Saraiya")

	// Retrieving the newly created User
	invokeGetUserRPC(client, idOfCreatedUser)
}
