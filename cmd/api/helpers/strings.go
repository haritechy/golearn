package helpers

func Stringshelp(Name string, Message string, Email string, PhoneNumber string) (userWelcome string, useEmailBody string, adminEmailBody string, sheduled string) {
	userWelcome = "Welcome to [Kar ai]! You are now an official user, and here are some resources to help you get started."
	useEmailBody = "Dear " + Name + ",\n\n" +
		"Thank you for approaching us for our services. We have received your message and will get back to you shortly. Here's a copy of your message:\n\n" +
		"\"" + Message + "\"\n\n" +
		"Best regards,\n" +
		"Kar ai Team"

	adminEmailBody = "You have received a new customer message from:\n\n" +
		"Name: " + Name + "\n" +
		"Phone Number: " + PhoneNumber + "\n" +
		"Email: " + Email + "\n\n" +
		"Message:\n" +
		"\"" + Message + "\"\n\n" +
		"Please follow up with the customer as soon as possible."
	sheduled = "meeting sheduled"

	return
}
