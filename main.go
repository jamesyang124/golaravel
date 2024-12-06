package main

func main() {
	c := initApplication()

	// should close external resource connection
	// before this function returned by defer
	c.App.ListenAndServe()
}
