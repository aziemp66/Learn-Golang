package main

type Blacklist func(string) bool

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUsername("eko", blacklist)
	registerUsername("admin", func(s string) bool {
		return s == "admin"
	})
}

func registerUsername(username string, blacklist Blacklist) {
	if blacklist(username) {
		println("You are blocked", username)
	} else {
		println("Welcome", username)
	}
}

// func blaclistAdmin(name string) bool {
// 	return name == "admin"
// }

// func blacklistRoot(name string) bool {
// 	return name == "root"
// }
