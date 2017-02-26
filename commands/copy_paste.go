package commands

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/jutkko/cli"
	"github.com/jutkko/copy-pasta/runcommands"
	"github.com/jutkko/copy-pasta/store"
	minio "github.com/minio/minio-go"
)

type InvalidConfig struct {
	error  string
	status int
}

func (ic *InvalidConfig) Error() string {
	return ic.error
}

type CopyPasteCommand struct {
	Ui cli.Ui
}

func (c *CopyPasteCommand) Help() string {
	return `Usage to paste: copy-pasta [--paste]
Usage to copy: <some command with output> | copy-pasta

    Copy or paste using copy-pasta. Use --paste to force copy-pasta to
		ignore its stdin and output from the current target.
`
}

func (c *CopyPasteCommand) Run(args []string) int {
	config, invalidConfig := loadRunCommands()
	if invalidConfig != nil {
		c.Ui.Error(fmt.Sprintf("Failed to load the runcommands file: %s", invalidConfig.Error()))
		os.Exit(invalidConfig.status)
	}

	copyPasteCommand := flag.NewFlagSet("", flag.ExitOnError)
	copyPastePasteOption := copyPasteCommand.Bool("paste", false, "")

	// not tested, may be too hard
	err := copyPasteCommand.Parse(args)
	if err != nil {
		return 10
	}

	if config != nil {
		content, err := copyPaste(config.CurrentTarget, *copyPastePasteOption)
		if err != nil {
			c.Ui.Error(fmt.Sprintf("Failed to load the runcommands file: %s", err.Error()))
			os.Exit(-15)
		}

		// cannot use c.Ui sicne it prints a newline
		fmt.Print(content)
	}

	return 0
}

func (c *CopyPasteCommand) Synopsis() string {
	return "Copy or paste using copy-pasta"
}

func copyPaste(target *runcommands.Target, paste bool) (string, error) {
	client, err := minioClient(target)
	if err != nil {
		return "", errors.New(fmt.Sprintf("Failed initializing client: %s\n", err.Error()))
	}

	if isFromAPipe() && !paste {
		if err = store.S3Write(client, target.BucketName, "default-object-name", target.Location, os.Stdin); err != nil {
			return "", errors.New(fmt.Sprintf("Failed writing to the bucket: %s\n", err.Error()))
		}

		return "", nil
	} else {
		content, err := store.S3Read(client, target.BucketName, "default-object-name")
		if err != nil {
			return "", errors.New(fmt.Sprintf("Have you copied yet? Failed reading the bucket: %s\n", err.Error()))
		}

		return content, nil
	}
}

func isFromAPipe() bool {
	stat, err := os.Stdin.Stat()
	if err != nil {
		log.Fatal(err)
	}

	return (stat.Mode() & os.ModeCharDevice) == 0
}

func minioClient(t *runcommands.Target) (*minio.Client, error) {
	endpoint := t.Endpoint
	accessKeyID := t.AccessKey
	secretAccessKey := t.SecretAccessKey
	useSSL := true

	// Initialize minio client object
	return minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
}

func getOrElse(key, defaultValue string) string {
	result := os.Getenv(key)
	if result == "" {
		return defaultValue
	}

	return result
}

func loadRunCommands() (*runcommands.Config, *InvalidConfig) {
	loadedConfig, err := runcommands.Load()
	if err != nil {
		return nil, &InvalidConfig{
			error:  "Please log in",
			status: 1,
		}
	}

	return loadedConfig, nil
}
