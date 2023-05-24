package vrf

import (
	"github.com/equinix/metal-cli/internal/outputs"
	"github.com/packethost/packngo"
	"github.com/spf13/cobra"
)

type Client struct {
	Servicer Servicer
	Service  packngo.VRFService
	Out      outputs.Outputer
}

func (c *Client) NewCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:     `vrf`,
		Aliases: []string{"vrf"},
		Short:   "VRF operations : create, TODO: make other commands.",
		Long:    "Experimental VRF function"

		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			if root := cmd.Root(); root != nil {
				if root.PersistentPreRun != nil {
					root.PersistentPreRun(cmd, args)
				}
			}
			c.Service = c.Servicer.API(cmd).ProjectVRFs //Not sure about this endpoint
		},
	}

	cmd.AddCommand(
		c.List(),
		c.Create(),
		c.Update(),
		c.Get(),
		c.ListIPs(),
		c.Delete(),
	)
	return cmd
}

type Servicer interface {
	API(*cobra.Command) *packngo.Client
	ListOptions(defaultIncludes, defaultExcludes []string) *packngo.ListOptions
}

func NewClient(s Servicer, out outputs.Outputer) *Client {
	return &Client{
		Servicer: s,
		Out:      out,
	}
}
