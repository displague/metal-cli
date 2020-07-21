package cmd

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/ghodss/yaml"
	"github.com/olekukonko/tablewriter"
	"github.com/packethost/packngo"
	"github.com/spf13/cobra"
)

// Cli struct
type Cli struct {
	Client  *packngo.Client
	MainCmd *cobra.Command
}

// VERSION build
var (
	Version string = "0.0.7"
)

// NewCli struct
func NewCli() *Cli {
	var err error
	cli := &Cli{}
	cli.Client, err = packngo.NewClientWithBaseURL("Packet CLI", os.Getenv("PACKET_TOKEN"), nil, "https://api.packet.net/")
	cli.Client.UserAgent = fmt.Sprintf("packet-cli/%s %s", Version, cli.Client.UserAgent)
	if err != nil {
		fmt.Println("Client error:", err)
		return nil
	}
	rootCmd.Execute()
	rootCmd.DisableSuggestions = false
	cli.MainCmd = rootCmd
	return cli
}

func output(in interface{}, header []string, data *[][]string) {
	if !isJSON && !isYaml {

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoWrapText(false)
		table.SetAlignment(tablewriter.ALIGN_LEFT)
		table.SetHeader(header)
		table.AppendBulk(*data)
		table.Render()
	} else if isJSON {
		output, err := json.MarshalIndent(in, "", "  ")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(output))
	} else if isYaml {
		output, err := yaml.Marshal(in)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(output))
	}
}

func outputMergingCells(in interface{}, header []string, data *[][]string) {
	if !isJSON && !isYaml {

		table := tablewriter.NewWriter(os.Stdout)
		table.SetAutoMergeCells(true)
		table.SetRowLine(true)
		table.SetHeader(header)
		table.AppendBulk(*data)
		table.Render()
	} else if isJSON {
		output, err := json.MarshalIndent(in, "", "  ")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(output))
	} else if isYaml {
		output, err := yaml.Marshal(in)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(output))
	}
}
