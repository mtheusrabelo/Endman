package prompt

import (
	"bufio"
	"fmt"
	"os"

	"strings"

	Config "github.com/matheusrabelo/Endman/configuration"
	"github.com/matheusrabelo/Endman/routes"
)

var addedRoutes []string

//Run starts the prompt
func Run() {
	fmt.Println(Config.Application.Message["endman"])

	routes.Listen(Config.Application.Port)
	fmt.Println(Config.Application.Message["listening"], Config.Application.Port)

	scanner := bufio.NewScanner(os.Stdin)
	var path string
	var method string

	for {
		fmt.Print(Config.Application.Message["addRoute"])
		scanner.Scan()
		path = scanner.Text()

		if path == "end" {
			end()
			break
		}

		fmt.Print(Config.Application.Message["addMethod"])
		scanner.Scan()
		method = scanner.Text()

		if method != "end" {
			routes.AddRoute(method, path)
			addedRoutes = append(addedRoutes, strings.ToUpper(method)+" "+path)
			fmt.Println(Config.Application.Message["added"], addedRoutes[len(addedRoutes)-1])
		} else {
			end()
			break
		}
	}

}

func end() {
	fmt.Println(Config.Application.Message["line"])
	for _, r := range addedRoutes {
		fmt.Println(r)
	}
	fmt.Println(Config.Application.Message["line"])
}
