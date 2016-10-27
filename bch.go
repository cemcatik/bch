package main

import (
	"fmt"
	"os"

	"github.com/urfave/cli"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/crypto/ssh/terminal"
)

func hash(c *cli.Context) error {
	fmt.Println("Enter password: ")
	password, _ := terminal.ReadPassword(0)

	factor := c.Int("factor")

	hash, err := bcrypt.GenerateFromPassword(password, factor)
	if err == nil {
		fmt.Println(string(hash))
	} else {
		fmt.Printf("Error: %s\n", err)
	}

	return err
}

func verify(c *cli.Context) error {
	if !c.Args().Present() {
		fmt.Println("Error: a hash must be provided in order to verify")
		os.Exit(1)
	}

	hashed := c.Args().First()
	hashedBytes := []byte(hashed)

	fmt.Println("Enter password: ")
	password, _ := terminal.ReadPassword(0)

	err := bcrypt.CompareHashAndPassword(hashedBytes, password)
	fmt.Println(err == nil)

	return err
}

func main() {
	app := cli.NewApp()
	app.Name = "bch"
	app.Usage = "Generate and verify bcrypt password hashes"
	app.Commands = []cli.Command{
		{
			Name:  "hash",
			Usage: "Generate bcrypt password hash",
			Flags: []cli.Flag{
				cli.IntFlag{
					Name:  "factor, f",
					Value: 12,
					Usage: "work factor",
				},
			},
			Action: hash,
		},
		{
			Name:      "verify",
			Usage:     "Verify given hash against entered password",
			ArgsUsage: "[hash]",
			Action:    verify,
		},
	}

	app.Run(os.Args)
}
