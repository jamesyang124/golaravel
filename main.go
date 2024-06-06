package main

func main() {
	c := initApplication()

	c.App.ListenAndServe()
}
