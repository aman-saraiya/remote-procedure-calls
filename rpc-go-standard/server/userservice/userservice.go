/* 
 * This package implements the UserService RPC service.
 */

package userservice
import (
	"fmt"
	"log"

	ss "github.com/aman-saraiya/remote-procedure-calls/rpc-go-standard/client/sharedstructs"
)

// Defining the RPC Service struct
type UserService struct {
	Users map[int]ss.User 
}

// Function to create RPC service instance
func NewUserService() *UserService {
	return &UserService{
		Users: make(map[int]ss.User),
	}
}

// Handler for AddUser RPC
func (svc *UserService) AddUser(arg *ss.AddUserArg, ret *ss.AddUserRet) error {
	log.Println("Handling AddUser RPC with argument ", arg)
id := len(svc.Users) + 1
	newUser := ss.User{
		ID:    id,
		Email: arg.Email,
		Name:  arg.Name,
	}
	svc.Users[id] = newUser
	ret.ID = id
	return nil
}

// Handler for GetUser RPC
func (svc *UserService) GetUser(arg *ss.GetUserArg, ret *ss.GetUserRet) error {
	log.Println("Handling GetUser RPC with argument ", arg)
	user, exists := svc.Users[arg.ID]
	if !exists {
		return fmt.Errorf("User with ID %d not found", arg.ID)
	}
	ret.User = user
	return nil
}
