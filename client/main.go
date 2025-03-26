package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/hetu-project/hetu-checkpoint/config"
	"github.com/hetu-project/hetu-checkpoint/proto/types"
)

func main() {
	url := "http://localhost:8080/reqblssign"
	request := config.Request{
		ValidatorAddress: "hetu1...",
		Checkpoint: types.RawCheckpoint{
			EpochNum:    100,
			BlockHash:   &types.BlockHash{}, // Initialize with appropriate hash
			Bitmap:      []byte("0x1234567890abcdef..."),
			BlsMultiSig: nil,
		},
	}

	message, err := json.Marshal(request)
	if err != nil {
		log.Fatalf("Error marshaling request: %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(message))
	if err != nil {
		log.Fatalf("Error creating request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response: %v", err)
	}

	fmt.Printf("%s\n", body)
}
