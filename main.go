package main

import (
	"fmt"
	"os"

	"github.com/nav9v/code2txt/cmd"
)

// version is set during build time via ldflags
var version = "dev"

func main() {
	banner := `
                                                                	  
   ██████╗ ██████╗ ██████╗ ███████╗██████╗ ████████╗██╗  ██╗████████╗  
  ██╔════╝██╔═══██╗██╔══██╗██╔════╝╚════██╗╚══██╔══╝╚██╗██╔╝╚══██╔══╝  
  ██║     ██║   ██║██║  ██║█████╗   █████╔╝   ██║    ╚███╔╝    ██║     
  ██║     ██║   ██║██║  ██║██╔══╝  ██╔═══╝    ██║    ██╔██╗    ██║     
  ╚██████╗╚██████╔╝██████╔╝███████╗███████╗   ██║   ██╔╝ ██╗   ██║     
   ╚═════╝ ╚═════╝ ╚═════╝ ╚══════╝╚══════╝   ╚═╝   ╚═╝  ╚═╝   ╚═╝     
                                                                      `

	fmt.Println(banner)
	fmt.Printf("version %s\n", version)
	fmt.Println()

	if err := cmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
