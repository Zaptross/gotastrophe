package main

import (
	"fmt"
	"os"
	"strings"
	"syscall"
	"unicode"
)

// string template for the main.go file
const mainTemplate = `package main

import (
	"github.com/gin-gonic/gin"

	routes "github.com/Zaptross/gotastrophy/routes"
)

func main() {
	r := gin.Default()

	r.%s("%s", routes.%s)

	r.Run()
}`

func main() {
	// read contents of ./routes/routes.go
	println(os.Getwd())
	contents, err := os.ReadFile("./routes/routes.go")

	if err != nil {
		println("Error reading routes.go")
		panic(err)
	}

	// convert contents to string
	routesContents := string(contents)

	// split routes.go into lines
	lines := strings.Split(routesContents, "\n")

	println(len(mainTemplate))

	// ensure ./pico directory exists
	if _, err := os.Stat("./pico"); os.IsNotExist(err) {
		os.Mkdir("./pico", 0755)
	}

	// open build.sh file
	buildFile, err := os.Create("./pico/build.sh")

	if err != nil {
		println("Error creating build.sh")
		panic(err)
	}

	// find each line with GET, POST, PUT, DELETE and Handler
	for _, line := range lines {
		for _, method := range []string{"GET", "POST", "PUT", "DELETE"} {
			// if line contains method, Handler, and func we know it's not the single process function
			if strings.Contains(line, method) && strings.Contains(line, "Handler") && strings.Contains(line, "func") {
				// get handlerFuncName from method to Handler
				handlerFuncName := line[strings.Index(line, method) : strings.Index(line, "Handler")+7]

				// get substring between method and Handler
				routeName := handlerFuncName[strings.Index(handlerFuncName, method)+len(method) : strings.Index(handlerFuncName, "Handler")]

				// if dir ./pico/routeName doesn't exist, create it
				if _, err := os.Stat("./pico/" + routeName); os.IsNotExist(err) {
					os.Mkdir("./pico/"+routeName, 0755)
				}

				// create main.go file in ./pico/routeName
				mainFile, err := os.Create("./pico/" + routeName + "/main.go")

				if err != nil {
					println(fmt.Sprintf("Error creating main.go in %s", routeName))
					panic(err)
				}

				// write mainTemplate to main.go
				_, err = mainFile.WriteString(fmt.Sprintf(mainTemplate, method, toKebabCase(routeName), handlerFuncName))

				if err != nil {
					println(fmt.Sprintf("Error writing main.go in %s", routeName))
					panic(err)
				}

				println(fmt.Sprintf("Created main.go for /%s", toKebabCase(routeName)))

				// write build command to build.sh
				_, err = buildFile.WriteString(fmt.Sprintf("go build -o ./pico/%s/%s ./pico/%s/main.go\n", routeName, routeName, routeName))

				if err != nil {
					println(fmt.Sprintf("Error writing build command for %s to build.sh", routeName))
					panic(err)
				}

				// close main.go file
				mainFile.Close()
			}
		}
	}

	// close build.sh file
	buildFile.Close()
	println("Created ./pico/build.sh")

	println("Building all picoservices...")
	// run build.sh
	buildErr := syscall.Exec("/bin/sh", []string{"sh", "./pico/build.sh"}, os.Environ())

	if buildErr != nil {
		println("Error running build.sh")
		println(err)
	}

	println("Done!")
}

// replace uppercase letters with lowercase letters and add a dash before each uppercase letter
func toKebabCase(s string) string {
	var sb strings.Builder
	for i, c := range s {
		if i > 0 && unicode.IsUpper(c) {
			sb.WriteRune('-')
		}
		sb.WriteRune(unicode.ToLower(c))
	}
	return sb.String()
}
