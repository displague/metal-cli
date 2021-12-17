// Copyright © 2018 Jasmin Gacic <jasmin@stackpointcloud.com>
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

package vlan

import (
	"strconv"

	"github.com/packethost/packngo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Create() *cobra.Command {
	var vxlan int
	var projectID, metro, facility, description string

	// createVirtualNetworkCmd represents the createVirtualNetwork command
	createVirtualNetworkCmd := &cobra.Command{
		Use:   "create",
		Short: "Creates a virtual network",
		Long: `Example:

metal virtual-network create --project-id [project_UUID] { --metro [metro_code] --vlan [vlan] | --facility [facility_code] }

`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			req := &packngo.VirtualNetworkCreateRequest{
				ProjectID: projectID,
				Metro:     metro,
				Facility:  facility,
				VXLAN:     vxlan,
			}
			if description != "" {
				req.Description = description
			}

			n, _, err := c.Service.Create(req)
			if err != nil {
				return errors.Wrap(err, "Could not create ProjectVirtualNetwork")
			}

			data := make([][]string, 1)

			// TODO(displague) metro is not in the response
			data[0] = []string{n.ID, n.Description, strconv.Itoa(n.VXLAN), n.MetroCode, n.FacilityCode, n.CreatedAt}

			header := []string{"ID", "Description", "VXLAN", "Metro", "Facility", "Created"}

			return c.Out.Output(n, header, &data)
		},
	}

	createVirtualNetworkCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "Project ID (METAL_PROJECT_ID)")
	createVirtualNetworkCmd.Flags().StringVarP(&facility, "facility", "f", "", "Code of the facility")
	createVirtualNetworkCmd.Flags().StringVarP(&metro, "metro", "m", "", "Code of the metro")
	createVirtualNetworkCmd.Flags().StringVarP(&description, "description", "d", "", "Description of the virtual network")
	createVirtualNetworkCmd.Flags().IntVarP(&vxlan, "vxlan", "", 0, "VXLAN id to use (can only be used with --metro)")

	_ = createVirtualNetworkCmd.MarkFlagRequired("project-id")
	return createVirtualNetworkCmd
}
