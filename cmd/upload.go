package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func uploadRunner(_ *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("provide single directory argument")
	}
	dir := args[0]
	err := os.Chdir(dir)
	if err != nil {
		return err
	}
	files, err := ioutil.ReadDir(".")
	if err != nil {
		return err
	}
	var newest os.FileInfo
	for _, f := range files {
		if f.Mode().IsRegular() && f.ModTime().After(newest.ModTime()) {
			newest = f
		}
	}
	if newest.Name() == "" {
		return fmt.Errorf("no files found")
	}
	fmt.Printf("found %s", newest)
	return nil
}

var uploadCmd = &cobra.Command{
	Use:   "upload DIR",
	Short: "Upload newest file in DIR",
	RunE:  uploadRunner,
}

func init() {
	rootCmd.AddCommand(uploadCmd)
}
