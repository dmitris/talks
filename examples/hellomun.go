package hellomun

// Greeting returns a greeting.
func Greeting(who string) string {
	if who == "" {
		who = "world"
	}
	return "Servus, " + who + "!"
}
