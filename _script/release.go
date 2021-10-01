package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

type platformInfo struct {
	goos   string
	goarch string
}

func (p platformInfo) String() string { return p.goos + "-" + p.goarch }

func main() {
	log.SetFlags(0)

	flag.Parse()

	platforms := []platformInfo{
		{"linux", "amd64"},
		{"linux", "arm64"},
		{"darwin", "amd64"},
		{"darwin", "arm64"},
		{"windows", "amd64"},
	}

	for _, platform := range platforms {
		if err := prepareArchive(platform); err != nil {
			log.Printf("error: build %s: %v", platform, err)
		}
	}
}

func prepareArchive(platform platformInfo) error {
	log.Printf("building %s", platform)
	binaryExt := ""
	if platform.goos == "windows" {
		binaryExt = ".exe"
	}

	binaryName := filepath.Join("build", "CodeQuery"+binaryExt)
	buildCmd := exec.Command("go", "build",
		"-o", binaryName,
		".")
	buildCmd.Env = append([]string{}, os.Environ()...) // Copy env slice
	buildCmd.Env = append(buildCmd.Env, "GOOS="+platform.goos)
	buildCmd.Env = append(buildCmd.Env, "GOARCH="+platform.goarch)
	out, err := buildCmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("run %s: %v: %s", buildCmd, err, out)
	}

	// Pack it into the archive.
	archiveName := "CodeQuery-" + platform.String() + ".zip"
	zipCmd := exec.Command("zip", archiveName, "CodeQuery"+binaryExt)
	zipCmd.Dir = "build"
	log.Printf("creating %s archive", archiveName)
	if out, err := zipCmd.CombinedOutput(); err != nil {
		return fmt.Errorf("make archive: %v: %s", err, out)
	}

	return nil
}
