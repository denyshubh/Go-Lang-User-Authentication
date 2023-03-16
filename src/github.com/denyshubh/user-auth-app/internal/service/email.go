package service

func SendWelcomeEmail(user *model.User) {
	// Simulate a time-consuming task like sending an email
	time.Sleep(3 * time.Second)
	fmt.Printf("Welcome email sent to: %s\n", user.Email)
}
