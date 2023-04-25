package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type context struct {
	directory           string
	spreadsheetOfTotals string
	statements          []string
}

type TransactionReader struct {
	transaction string
}

func (c *context) getDirectory() {
	fmt.Println("What folder are your statements in?")
	input := bufio.NewReader(os.Stdin)
	dir, err := input.ReadString('\n')
	directory := strings.TrimSuffix(dir, "\n")
	if err != nil {
		fmt.Println(err)
	}
	c.directory = directory
}

func (c *context) getNameForTotalsSpreadsheet() {
	fmt.Println("What do you want the final spreadsheet to be called?")
	input := bufio.NewReader(os.Stdin)
	name, err := input.ReadString('\n')
	if err != nil {
		fmt.Println(err)
	}
	c.spreadsheetOfTotals = name + ".csv"
}

func (c *context) getStatements(directory string) {
	statements, _ := os.ReadDir(directory)
	s := []string{}
	for _, v := range statements {
		s = append(s, v.Name())
	}
	c.statements = s
}

func main() {
	var context context
	context.getDirectory()
	context.getNameForTotalsSpreadsheet()
	context.getStatements(context.directory)
}
