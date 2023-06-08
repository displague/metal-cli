package vrf

import (
	"fmt"
	"strconv"
	"github.com/packethost/packngo"
	"github.com/spf13/cobra"
)


func (c *Client) Create() *cobra.Command {
	var (
	    metro string
	    name string
	    description string
	    ipRanges []string
	    localASN int
        )

	// createVRFCmd represents the creatVRF command
	createVRFCmd := &cobra.Command{
		Use:   `create vrf <my_vrf> [-m <metro_code>] [-AS int] [-I str] [-d <description>]`,
		Short: "Creates a virtual network.",
		Long:  "Creates a VRF",
		Example: `# Creates a VRF, metal vrf create `,

		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true

			req := &packngo.VRFCreateRequest{
				Metro:     metro,
				Name:	   name,
				IPranges,  IPranges,
				LocalASN,  LocalASN,
			}
			if description != "" {
				req.Description = description
			}
			// Retreived these params from packetngo based on the VRF function there, assuming vrfs need these to work
			// From Packetngo >>>>
			// IPRanges is a list of all IPv4 and IPv6 Ranges that will be available to
			// BGP Peers. IPv4 addresses must be /8 or smaller with a minimum size of
			// /29. IPv6 must be /56 or smaller with a minimum size of /64. Ranges must
			// not overlap other ranges within the VRF.
			// >>>>>>>>>>>>>>>>>>>
			// Not quite sure if a CIDR can be specified here using a "/" or how it works exactly in tandem with VRFs
			// may also need some logic for processing errors on local asn, also not sure about this
			n, _, err := c.Service.Create(req)
			if err != nil {
				return fmt.Errorf("Could not create VRF: %w", err)
			}

			data := make([][]string, 1)

			// This output block below is probably incorrect but leaving it for now for testing later.
			data[0] = []string{n.Name, n.Description, strconv.Itoa(n.LocalASN), strconv.Itoa(n.IPrange) n.MetroCode, n.CreatedAt}
			header := []string{"Name", "Metro", "Description", "LocalASN", "IPrange" "Created"}

			return c.Out.Output(n, header, &data)
		},
	}

	createVRFCmd.Flags().StringVarP(&Name)
	createVRFCmd.Flags().StringVarP(&metro)
	createVRFCmd.Flags().IntVarP(&LocalASN)
	createVRFCmd.Flags().IntVarP(&ip_ranges)
	_ = createDeviceCmd.MarkFlagRequired("name")
	_ = createDeviceCmd.MarkFlagRequired("nmetro")
	_ = createDeviceCmd.MarkFlagRequired("LocalASN")
	_ = createDeviceCmd.MarkFlagRequired("IPrange")

	// making them all required here 

	return createDeviceCmd
}





