package main

import (
	".."
	"../filters"
	"fmt"
	"os"
)

func main() {
	sha, err := gitmedia.Decode(os.Stdin)
	if err != nil {
		gitmedia.Panic(err, "Error reading git-media meta data from stdin:")
	}

	var path string
	if len(os.Args) > 1 {
		path = os.Args[1]
	} else {
		path = sha
	}

	fmt.Fprintf(os.Stderr, "Downloading media: %s\n", path)

	err = gitmediafilters.Smudge(os.Stdout, sha)
	if err != nil {
		smudgerr := err.(*gitmediafilters.SmudgeError)
		gitmedia.Panic(err, "Error reading file from local media dir: %s", smudgerr.Filename)
	}
}
