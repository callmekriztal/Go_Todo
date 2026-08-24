package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	done := Done{}

	path, err := todoFilePath()
	if err != nil {
		panic(err)
	}

	storage := NewStorage[Done](path)

	if err := storage.Load(&done); err != nil {
		if !errors.Is(err, os.ErrNotExist) {
			panic(err)
		}
	}

	cmdFlags := NewCmdFlags()

	if err := cmdFlags.Execute(&done); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	if err := storage.Save(done); err != nil {
		panic(err)
	}
}
