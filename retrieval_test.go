package subsonic

import "testing"

func runRetrievalTests(client Client, t *testing.T) {
	sampleSong := getSampleSong(client)

	t.Run("Stream", func(t *testing.T) {
		// Purposefully choose an ID that returns an error
		_, err := client.Stream("1", nil)
		if err == nil {
			t.Error("An error was not returned on ID 1")
		}
		contents, err := client.Stream(sampleSong.ID, nil)
		if err != nil {
			t.Error(err)
		}
		if contents == nil {
			t.Error("No content returned")
		}
	})
}