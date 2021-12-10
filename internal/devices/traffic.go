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

package devices

import (
	"fmt"
	"time"

	"github.com/packethost/packngo"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

func (c *Client) Traffic() *cobra.Command {
	var (
		startTime, endTime string
	)
	var retrieveDeviceCmd = &cobra.Command{
		Use:   "traffic",
		Short: "Retrieves device list or device details",

		Long: `Example:
	
metal device get --id [device_UUID]

	`,
		RunE: func(cmd *cobra.Command, args []string) error {
			deviceID, _ := cmd.Flags().GetString("id")
			listOpts := c.Servicer.ListOptions(nil, nil)
			opts := &packngo.TrafficRequest{}
			if startTime != "" {
				parsedTime, err := time.Parse(time.RFC3339, startTime)
				if err != nil {
					return errors.Wrap(err, fmt.Sprintf("Could not parse time %q", startTime))
				}
				opts.StartedAt = &packngo.Timestamp{Time: parsedTime}
			}

			if endTime != "" {
				parsedTime, err := time.Parse(time.RFC3339, endTime)
				if err != nil {
					return errors.Wrap(err, fmt.Sprintf("Could not parse time %q", endTime))
				}
				opts.EndedAt = &packngo.Timestamp{Time: parsedTime}
			}

			device, _, err := c.Service.GetTraffic(deviceID, opts, listOpts)
			if err != nil {
				return errors.Wrap(err, "Could not get Devices")
			}
			header := []string{} // "ID", "Hostname", "OS", "State", "Created"}

			data := make([][]string, 1)
			data[0] = []string{} // device.ID, device.Hostname, device.OS.Name, device.State, device.Created}

			return c.Out.Output(device, header, &data)
		},
	}

	retrieveDeviceCmd.Flags().StringP("id", "i", "", "UUID of the device")
	retrieveDeviceCmd.Flags().StringVar(&startTime, "started-at", "", `Device termination time: --termination-time="15:04:05"`)
	retrieveDeviceCmd.Flags().StringVar(&endTime, "ended-at", "", `Device termination time: --termination-time="15:04:05"`)
	_ = retrieveDeviceCmd.MarkFlagRequired("id")
	return retrieveDeviceCmd
}
