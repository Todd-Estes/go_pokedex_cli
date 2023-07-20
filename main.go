package main

import (
    "bufio"
    "fmt"
		"os"
		"time"
		// "github.com/Todd-Estes/internal/pokecache"
		"github.com/Todd-Estes/go_pokedex_cli/internal/pokecache"
)



// return map[string]cliCommand{
//     "help": {
//         name:        "help",
//         description: "Displays a help message",
//         callback:    commandHelp,
//     },
//     "exit": {
//         name:        "exit",
//         description: "Exit the Pokedex",
//         callback:    commandExit,
//     },
// }


// StringPrompt asks for a string value using the label
// func StringPrompt(label string) string {
//     var s string
//     r := bufio.NewReader(os.Stdin)
//     for {
//         fmt.Fprint(os.Stderr, label+" ")
//         s, _ = r.ReadString('\n')
//         if s != "" {
//             break
//         }
//     }
//     return strings.TrimSpace(s)
// }

func main() {
	reader := bufio.NewScanner(os.Stdin)
	location := Locations{Next: "https://pokeapi.co/api/v2/location"}
	duration := time.Duration(1 * time.Minute)
	// requestCache is a pointer
	requestCache := pokecache.NewCache(duration)

		for {
			fmt.Print("Pokedex > ")
			reader.Scan()
			input := reader.Text()

		command, ok := getCommands()[input]
		if ok {
			err := command.callback()
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else if input == "map" {
			getPokeLocations(location.Next, &location, requestCache)
			continue
		} else if input == "mapb" {
			getPrevPokeLocations(&location, requestCache)
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
	}
}

func commandHelp() error {
	fmt.Printf("%s\n%s\n\n", "Welcome to the Pokedex!", "Usage:")

	for _, cmd := range getCommands() {
		fmt.Printf("%s: %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandExit() error {
	os.Exit(0)
	return nil
}