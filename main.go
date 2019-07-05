package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func wipeOutputDirectory() error {
	dir, err := os.Open("/out")
	if err != nil {
		return err
	}
	defer dir.Close()
	names, err := dir.Readdirnames(-1)
	if err != nil {
		return err
	}
	for _, name := range names {
		err = os.RemoveAll("/out"+name)
		if err != nil {
			return err
		}
	}
	return nil
}

func main() {
	err := wipeOutputDirectory()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Could not delete the contents of the output directory. Is it mounted read/write? %s\n", err.Error())
		os.Exit(1)
	}

	env := os.Getenv("SANIC_ENV")
	if env == "" {
		fmt.Fprintf(os.Stderr, "SANIC_ENV was not set?\n")
		os.Exit(1)
	}
	if _, err = os.Open("/in/overlays/"+env+"/kustomization.yaml"); err != nil {
		fmt.Fprintf(os.Stderr, "Could not read the kustomization file for environment %s. Expected at deploy/in/overlays/%s/kustomization.yaml to exist.", env, env)
		os.Exit(1)
	}

	cmd := exec.Command("/app/kustomize", "build", "/in/overlays/"+env, "--output", "/out")
	out := bytes.Buffer{}
	cmd.Stderr = &out
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		fmt.Fprint(os.Stderr, out.String())
		fmt.Fprintf(os.Stderr, "Could not run templater for environment %s: %s\n", env, err.Error())
		os.Exit(1)
	}
}
