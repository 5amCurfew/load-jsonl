package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "xtkt [config.json] | load-jsonl [config.json]",
		Short: "Appends records to a jsonl file",
		Run:   load,
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func load(cmd *cobra.Command, args []string) {
	reader := bufio.NewReader(os.Stdin)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		var record map[string]interface{}
		if err := json.Unmarshal([]byte(line), &record); err != nil {
			fmt.Println("Error unmarshaling JSON:", err)
			continue
		}

		if recordType, ok := record["type"].(string); ok && recordType == "RECORD" {
			file, _ := os.OpenFile(record["stream"].(string)+".jsonl", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

			recordBytes, err := json.Marshal(record["record"].(map[string]interface{}))
			if err != nil {
				fmt.Println("Error marshaling JSON:", err)
				continue
			}

			recordLine := fmt.Sprintf("%s\n", recordBytes)
			if _, err := file.WriteString(recordLine); err != nil {
				fmt.Println("Error writing to file:", err)
				continue
			}
		}
	}

	fmt.Println("Records appended to file.")
}
