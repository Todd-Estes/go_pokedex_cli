package main

import (
    "bufio"
    "fmt"
    "os"
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
	// res, err := getPokeLocations()
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Println(res)
	// }
	
	reader := bufio.NewScanner(os.Stdin)
		location := Locations{Next: "https://pokeapi.co/api/v2/location"}
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
			getPokeLocations(&location)
			continue
		} else if input == "mapb" {
			getPrevPokeLocations(&location)
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