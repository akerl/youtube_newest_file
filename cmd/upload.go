package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"google.golang.org/api/youtube/v3"
)

func uploadRunner(_ *cobra.Command, args []string) error {
	if len(args) != 1 {
		return fmt.Errorf("provide single directory argument")
	}
	file, err := getFilePath(args[1])
	if err != nil {
		return err
	}
	return uploadFile(file)
}

func getFilePath(dir string) (string, error) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		return "", err
	}
	var newest os.FileInfo
	for _, f := range files {
		if f.Mode().IsRegular() && (newest == nil || f.ModTime().After(newest.ModTime())) {
			newest = f
		}
	}
	if newest == nil {
		return "", fmt.Errorf("no files found")
	}
	return filepath.Join(dir, newest), nil
}

func uploadFile(file string) error {
	client := getClient(youtube.YoutubeUploadScope)

	service, err := youtube.New(client)
	if err != nil {
		return err
	}

	upload := &youtube.Video{
		Snippet: &youtube.VideoSnippet{
			Title:      getTitle(),
			CategoryId: getCategory(),
		},
		Status: &youtube.VideoStatus{PrivacyStatus: getPrivacy()},
	}
	call := service.Videos.Insert("snippet,status", upload)

	file, err := os.Open(file)
	defer file.Close()
	if err != nil {
		return err
	}

	_, err := call.Media(file).Do()
	if err != nil {
		return err
	}
	return nil
}

func getTitle() string {
	return time.Now().Format("2006/01/02 15:04 Recording")
}

func getCategory() string {
	return "20"
}

func getPrivacy() string {
	return "unlisted"
}

var uploadCmd = &cobra.Command{
	Use:   "upload DIR",
	Short: "Upload newest file in DIR",
	RunE:  uploadRunner,
}

func init() {
	rootCmd.AddCommand(uploadCmd)
}
