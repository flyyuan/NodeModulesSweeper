package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"os"
	"path/filepath"
)

func main() {
	var dir string
	var dirsToDelete []string

	fmt.Print("Please enter the directory to scan (default is current directory): ")
	fmt.Scanln(&dir)

	if dir == "" {
		dir = "."
	}

	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() && info.Name() == "node_modules" {
			dirsToDelete = append(dirsToDelete, path)
			return filepath.SkipDir // Skip the subdirectories of this node_modules
		}
		return nil
	})

	if err != nil {
		fmt.Println("An error occurred while scanning:", err)
		return
	}

	if len(dirsToDelete) == 0 {
		fmt.Println("No node_modules directories found.")
		return
	}

	for {
		options := append([]string{"\033[31mDelete All\033[0m", "Exit"}, dirsToDelete...) // ANSI escape codes for red color
		prompt := &survey.Select{
			Message: "Please select a directory to delete:",
			Options: options,
		}

		var choice string
		if err := survey.AskOne(prompt, &choice); err != nil {
			fmt.Println("An error occurred:", err)
			return
		}

		if choice == "Exit" {
			return
		}

		if choice == "\033[31mDelete All\033[0m" {
			for _, d := range dirsToDelete {
				if err := os.RemoveAll(d); err != nil {
					fmt.Printf("Error deleting %s: %s\n", d, err)
				} else {
					fmt.Printf("Successfully deleted %s\n", d)
				}
			}
			fmt.Println("No node_modules directories found.")
			return
		} else {
			if err := os.RemoveAll(choice); err != nil {
				fmt.Printf("Error deleting %s: %s\n", choice, err)
			} else {
				fmt.Printf("Successfully deleted %s\n", choice)
				// Remove the deleted directory from the list
				dirsToDelete = remove(dirsToDelete, choice)
			}
		}
	}
}

func remove(slice []string, item string) []string {
	for i, a := range slice {
		if a == item {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
