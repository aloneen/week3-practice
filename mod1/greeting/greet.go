// mod1/greeting/greet.go

package greeting

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return "Hello, " + name + "!"
}
