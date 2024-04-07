package app

import (
	"fmt"

	"github.com/toastsandwich/LCP/config"
	"github.com/toastsandwich/LCP/routes"
)

func Start(port string) {
	echo := config.Echo
	routes.Init(echo)
	fmt.Println("endpoints : ")
	fmt.Println("[+] Doctor")
	fmt.Println("/getAllDoctors")
	fmt.Println("/getDoctorByID/:id")
	fmt.Println("/addDoctor/")
	fmt.Println("/deleteDoctor/:id")
	fmt.Println("[+] Lab Assistant")
	fmt.Println("/getAllLabAssistants")
	fmt.Println("getLabAssistantByID/:id")
	fmt.Println("/addLabAssistant/")
	fmt.Println("/deleteLabAssistant/:id")
	echo.Logger.Fatal(echo.Start(port))
}
