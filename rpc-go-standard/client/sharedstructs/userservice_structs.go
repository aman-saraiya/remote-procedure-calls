package sharedstructs

// User represents a user with an ID, Name, and Email.
type User struct {
	ID int
	Name string
	Email string
}

// Request argument for AddUser RPC
type AddUserArg struct {
	Name  string
	Email string
}

// Response returned by AddUser RPC
type AddUserRet struct {
	ID int
}

// Request argument for GetUser RPC
type GetUserArg struct {
	ID int
}

// Response returned by GetUser RPC
type GetUserRet struct {
	User User
}
