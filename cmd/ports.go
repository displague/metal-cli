// Copyright © 2021 Packet CLI Developers
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package cmd

import (
	"fmt"
	"strconv"

	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var (
	vnID   string
	portID string
)

// portDisbondCmd represents the portDisbond command
var disbondPortCmd = &cobra.Command{
	Use:   "disbond",
	Short: "Disbonds a port",
	Long: `Example:

packet port disbond -i [port_id]
  
  `,
	RunE: func(cmd *cobra.Command, args []string) error {
		p, _, err := apiClient.Ports.Disbond(portID, true)
		if err != nil {
			return errors.Wrap(err, "Could not disbond Port")
		}

		data := [][]string{
			{p.ID, p.Name, strconv.FormatBool(p.Data.Bonded)},
		}
		header := []string{"ID", "Name", "Bonded"}
		return output(p, header, &data)
	},
}

// portBondCmd represents the portBond command
var bondPortCmd = &cobra.Command{
	Use:   "bond",
	Short: "Bonds a port",
	Long: `Example:

packet port bond --id [port_id]
  
  `,
	RunE: func(cmd *cobra.Command, args []string) error {
		p, _, err := apiClient.Ports.Bond(portID, true)
		if err != nil {
			return errors.Wrap(err, "Could not bond Port")
		}

		data := make([][]string, 1)

		data[0] = []string{p.ID, p.Name, strconv.FormatBool(p.Data.Bonded)}
		header := []string{"ID", "Name", "Bonded"}
		return output(p, header, &data)
	},
}

// deleteNativeVLANCmd represents the deleteNativeVLAN command
var deleteNativeVLANCmd = &cobra.Command{
	Use:   "delete-native-vlan",
	Short: "Deletes a port's native VLAN",
	Long: `Example:

packet port delete-native-vlan --id [port_UUID]

`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if !force {
			prompt := promptui.Prompt{
				Label:     fmt.Sprintf("Are you sure you want to delete the native VLAN on port %s: ", portID),
				IsConfirm: true,
			}

			_, err := prompt.Run()
			if err != nil {
				return nil
			}
		}
		return errors.Wrap(deleteNativeVLAN(portID), "Could not delete port's native VLAN")
	},
}

func deleteNativeVLAN(id string) error {
	_, _, err := apiClient.Ports.UnassignNative(id)
	if err != nil {
		return err
	}
	fmt.Println("Port native VLAN", id, "successfully deleted.")
	return nil
}

// portCmd represents the port command
var portCmd = &cobra.Command{
	Use:     "port",
	Aliases: []string{"ports"},
	Short:   "Port operations",
	Long:    `Port operations: bond, disbond, get`,
}

// retrivePortCmd represents the retrivePort command
var retrivePortCmd = &cobra.Command{
	Use:     "get",
	Aliases: []string{"list"},
	Short:   "Retrieves all available ports or a single port",
	Long: `Example:

Retrieve all ports:
packet port get
  
Retrieve a specific port:
packet port get -i [port_UUID]

When using "--json" or "--yaml", "--include=members" is implied.
	`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if portID == "" {
			return fmt.Errorf("Must specifythe port id")
		}

		inc := []string{"members"}

		// don't fetch extra details that won't be rendered
		if !isYaml && !isJSON {
			inc = nil
		}

		listOpts := listOptions(inc, nil)

		p, _, err := apiClient.Ports.Get(portID, listOpts)
		if err != nil {
			return errors.Wrap(err, "Could not get Port")
		}

		data := [][]string{
			{p.ID, p.Name, p.Type, p.Data.MAC},
		}

		header := []string{"ID", "Name", "Type", "MAC"}
		return output(p, header, &data)
	},
}

func init() {
	deleteNativeVLANCmd.Flags().StringVarP(&portID, "id", "i", "", "UUID of the port")
	_ = deleteNativeVLANCmd.MarkFlagRequired("id")

	deleteNativeVLANCmd.Flags().BoolVarP(&force, "force", "f", false, "Force removal of the port")

	disbondPortCmd.Flags().StringVarP(&portID, "id", "i", "", "ID of the port")

	_ = disbondPortCmd.MarkFlagRequired("id")

	bondPortCmd.Flags().StringVarP(&portID, "id", "i", "", "ID of the port")

	_ = bondPortCmd.MarkFlagRequired("id")

	rootCmd.AddCommand(portCmd)
	portCmd.AddCommand(bondPortCmd, disbondPortCmd, deleteNativeVLANCmd, retrivePortCmd)

	retrivePortCmd.Flags().StringVarP(&portID, "port-id", "i", "", "UUID of the port")

	// assignVLANCmd.Flags().StringVarP(&vnID, "virtual-network-id", "v", "", "ID of the VLAN")

}
