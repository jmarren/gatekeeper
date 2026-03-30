package other

import "fmt"

func HandleUser(user IUser) {
	fmt.Printf("user.GetFirstName() = %s\n", user.GetFirstName())
}
