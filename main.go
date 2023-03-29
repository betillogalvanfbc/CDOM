package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/common-nighthawk/go-figure"
	"github.com/fatih/color"
	"github.com/likexian/whois"
)

func main() {
	myFigure := figure.NewFigure("CHEDOM", "", true)
	myFigure.Print()
	if len(os.Args) < 2 {
		fmt.Println("No name provided")
		return
	}
	domain := os.Args[1]
	result, err := whois.Whois(domain)
	if err != nil {
		fmt.Println("Error:", err)
	}
	if strings.Contains(result, "Domain Name:") {
		red := color.New(color.FgRed).SprintFunc()
		fmt.Println(red(domain, " Registrado"))
	} else {
		green := color.New(color.FgGreen).SprintFunc()
		fmt.Println(green(domain, " Venta"))
	}
}
