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

package hardware

import (
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Move() *cobra.Command {
	var projectID, hardwareReservationID string

	moveHardwareReservationCmd := &cobra.Command{
		Use:   "move",
		Short: "Move hardware reservation to another project",
		Long: `Example:

metal hardware_reservation move -i [hardware_reservation_UUID] -p [project_UUID]
`,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			header := []string{"ID", "Facility", "Plan", "Created"}
			r, _, err := c.Service.Move(hardwareReservationID, projectID)
			if err != nil {
				return errors.Wrap(err, "Could not move Hardware Reservation")
			}

			data := make([][]string, 1)

			data[0] = []string{r.ID, r.Facility.Code, r.Plan.Name, r.CreatedAt.String()}

			return c.Out.Output(r, header, &data)
		},
	}

	moveHardwareReservationCmd.Flags().StringVarP(&hardwareReservationID, "id", "i", "", "UUID of the hardware reservation")
	moveHardwareReservationCmd.Flags().StringVarP(&projectID, "project-id", "p", "", "Project ID (METAL_PROJECT_ID)")
	_ = moveHardwareReservationCmd.MarkFlagRequired("project-id")
	_ = moveHardwareReservationCmd.MarkFlagRequired("id")

	return moveHardwareReservationCmd
}
