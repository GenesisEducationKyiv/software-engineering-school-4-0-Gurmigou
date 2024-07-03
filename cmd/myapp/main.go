package main

import "se-school-case/initializer"

func main() {
	apiService := initializer.NewApi()
	apiService.HandleRequests()
}
