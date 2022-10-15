package access

var ApplicationVersion float32 = 1.82
var apiKeys string = "OIHIO0*W798sO89Bboufyo^yodoYOyaiiausda^"

// Jika diawali Huruf BESAR maka bisa diakses dari luar packages
func SayHello(name string) string {
	return "Hello " + name
}

// Jika diawali Huruf kecil maka tidak bisa diakses dari luar packages
func sayGoodbye(name string) string {
	return "Goodbye " + name
}
