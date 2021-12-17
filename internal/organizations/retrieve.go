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

package organizations

import (
	"github.com/packethost/packngo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Retrieve() *cobra.Command {
	var organizationID string
	// retrieveOrganizationCmd represents the retrieveOrganization command
	retrieveOrganizationCmd := &cobra.Command{
		Use:     "get",
		Aliases: []string{"list"},
		Short:   "Retrieves an organization or list of organizations",
		Long: `Example:
	
To retrieve list of all available organizations:
metal organization get

To retrieve a single organization:
metal organization get -i [organization-id]

	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			listOpts := c.Servicer.ListOptions(nil, nil)
			if organizationID == "" {
				orgs, _, err := c.Service.List(listOpts)
				if err != nil {
					return errors.Wrap(err, "Could not list Organizations")
				}

				data := make([][]string, len(orgs))

				for i, p := range orgs {
					data[i] = []string{p.ID, p.Name, p.Created}
				}
				header := []string{"ID", "Name", "Created"}

				return c.Out.Output(orgs, header, &data)
			} else {
				getOpts := &packngo.GetOptions{Includes: listOpts.Includes, Excludes: listOpts.Excludes}
				org, _, err := c.Service.Get(organizationID, getOpts)
				if err != nil {
					return errors.Wrap(err, "Could not get Organization")
				}

				data := make([][]string, 1)

				data[0] = []string{org.ID, org.Name, org.Created}
				header := []string{"ID", "Name", "Created"}

				return c.Out.Output(org, header, &data)
			}
		},
	}

	retrieveOrganizationCmd.Flags().StringVarP(&organizationID, "organization-id", "i", "", "UUID of the organization")
	return retrieveOrganizationCmd
}
