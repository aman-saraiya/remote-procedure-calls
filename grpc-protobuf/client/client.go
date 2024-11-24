package main

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	gen_usersvc "github.com/aman-saraiya/remote-procedure-calls/grpc-protobuf/client/generated/userservice"
	gen_statssvc "github.com/aman-saraiya/remote-procedure-calls/grpc-protobuf/client/generated/statsservice"
)

func invokeAddUserGRPC(userSvcClient gen_usersvc.UserServiceClient, name, email string) int32 {
	// Initialize the AddUser RPC request argument
	addUserArg := &gen_usersvc.AddUserArg{
		Email: email,
		Name:  name,
	}

	// Declaring the RPC response variable to store the return value
	var addUserRet *gen_usersvc.AddUserRet

	// Calling the User Service's AddUser RPC
	addUserRet, err := userSvcClient.AddUser(context.Background(), addUserArg)
	if err != nil {
		fmt.Errorf("Error calling AddUser: %v\n", err)
	}
	fmt.Printf("Added user with ID: %d\n\n", addUserRet.GetId())
	
	return addUserRet.GetId()
}

func invokeGetUserGRPC(userSvcClient gen_usersvc.UserServiceClient, id int32){
	// Initialize the RPC request argument
	getUserArg := &gen_usersvc.GetUserArg{Id: id}

	// Declaring the RPC response variable to store the return value 
	var getUserRet *gen_usersvc.GetUserRet
	
	// Calling the UserService's GetUser RPC
	getUserRet, err := userSvcClient.GetUser(context.Background(), getUserArg)
	if err != nil {
		fmt.Errorf("Error calling GetUser: %v", err)
	}
	retrievedUser := getUserRet.GetUser()

	fmt.Printf("Retrieved user: Id = %d, Name = %s, Email = %s\n\n",
		retrievedUser.GetId(), retrievedUser.GetName(), retrievedUser.GetEmail())
}

func invokeGetStatsGRPC(client gen_statssvc.StatsServiceClient, numbers []float64) {
	//Initialize the RPC request argument
	getStatsArg := &gen_statssvc.GetStatsArg{Numbers: numbers}

	// Declaring the RPC response variable to store the response
	var getStatsRet *gen_statssvc.GetStatsRet

	// Calling the StatsService's GetStats RPC
	getStatsRet, err := client.GetStats(context.Background(), getStatsArg)
	if err != nil {
		fmt.Errorf("Error calling GetStats: %v", err)
	}
	fmt.Printf("Mean: %.2f, Median: %.2f\n", getStatsRet.Mean, getStatsRet.Median)
}

func invokeRunningStatsGRPC(client gen_statssvc.StatsServiceClient, numbers []float64) {
	// Calling the RunningStats RPC to start the stream
	stream, err := client.RunningStats(context.Background())
	if err != nil {
		fmt.Errorf("Could not start RunningStats stream: %v", err)
	}

	// Starting a separate go routine to send the numbers over the stream
	go func() {
		for idx, num := range numbers {
			// Initializing RunningStatsArg
			runningStatsArg := &gen_statssvc.RunningStatsArg{Number: num}
			fmt.Println("Sending ", runningStatsArg, " to server over stream.")
			fmt.Println("Overall Data Streamed ", numbers[:idx+1])
			if err := stream.Send(runningStatsArg); err != nil {
				fmt.Errorf("Could not send number: %v", err)
			}
			time.Sleep(2 * time.Second)
		}
		stream.CloseSend()
	}()

	// Receive the running mean and median from the server
	for {
		resp, err := stream.Recv()
		if err != nil {
			if err.Error() == "EOF" {
				fmt.Println("Server has closed the stream")
				break
			}
			fmt.Errorf("Could not receive running stats: %v", err)
		}
		fmt.Printf("Running Mean: %.2f, Running Median: %.2f\n\n", resp.Mean, resp.Median)
	}
}

func main() {
	// Connect to the server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Errorf("could not connect: %v", err)
	}
	defer conn.Close()

	// Initializing clients for UserService and StatsService
	// Notice here that both the service clients use the same
	// underlying connection to send the requests.
	userSvcClient := gen_usersvc.NewUserServiceClient(conn)
	statsSvcClient := gen_statssvc.NewStatsServiceClient(conn)

	fmt.Println("UserService RPC Calls\n")
	// Adding new User
	fmt.Println("Calling AddUser RPC")
	idOfCreatedUser := invokeAddUserGRPC(userSvcClient, "Aman Saraiya", "amansaraiya937@gmail.com")

	fmt.Println("Retrieving the newly created User")
	// Retrieving the newly created User
	invokeGetUserGRPC(userSvcClient, idOfCreatedUser)
	
	fmt.Println("\nStatsService RPC Calls\n")
	// Get Mean and Median for numbers
	numbers := []float64{6.5, 4.0, 0.5, 4.5, 9.5}
	fmt.Println("Calling GetStats RPC with numbers: ", numbers)
	invokeGetStatsGRPC(statsSvcClient, numbers)

	// Demonstrating the RunningStats bidirectional streaming RPC
	fmt.Println("\nCalling RunningStats RPC")
	invokeRunningStatsGRPC(statsSvcClient, numbers)
}
