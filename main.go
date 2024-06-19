package main

import "se-school-case/api"

func main() {
	apiService := api.NewApi()
	apiService.HandleRequests()
}
