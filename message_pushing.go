// Copyright 2017 The go-oceanconnect authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package oceanconnect

type Notification string

const (
	// NotificationDeviceAdded is used to notify initial device logins.
	// When a device registers with the OceanConnect (the device creates messages on
	// the OceanConnect and obtains the password), the OceanConnect sends a
	// notification to the application or a new device is added to the gateway, the
	// OceanConnect invokes this interface to send a notification to the application.
	NotificationDeviceAdded Notification = "deviceAdded"
	// NotificationDeviceInfoChanged is used after receiving device information changes
	// (changes of static information such as the device name and manufacturer ID).
	NotificationDeviceInfoChanged Notification = "deviceInfoChanged"
	// NotificationDeviceDataChanged is used after receiving device data changes
	// (dynamic changes such as changes of service attribute values).
	NotificationDeviceDataChanged Notification = "deviceDataChanged"
	// NotificationDeviceDeleted is used when learning that a
	// non-directly-connected device is deleted
	NotificationDeviceDeleted Notification = "deviceDeleted"
	// NotificationMessageConfirm is used after receiving an acknowledgment
	// message from the gateway, for example, the OceanConnect sends a command
	// to the gateway and the gateway acknowledges the message.
	NotificationMessageConfirm Notification = "messageConfirm"
	// NotificationCommandResponse is used after receiving a response command
	// from a device (gateway or common device), for example, the OceanConnect sends a
	// command to the device and the device returns a response command after running the
	// command, such as video call, video recording, and screenshot
	NotificationCommandResponse Notification = "commandRsp"
	// NotificationDeviceEvent after receiving an event (for example, insufficient
	// UDS storage space) from a device
	NotificationDeviceEvent Notification = "deviceEvent"
	// NotificationServiceInfoChanged is sent when learning device service information
	// changes, the OceanConnect invokes this interface to send a notification to the
	// application.
	NotificationServiceInfoChanged Notification = "serviceInfoChanged"
	// NotificationRuleEvent is used when generates the corresponding rule event
	// notification to NA when the rule is triggered
	NotificationRuleEvent Notification = "ruleEvent"
)

//DeviceDataChanged struct with device data
type DeviceDataChanged struct {
	DeviceID  string
	GatewayID string
	RequestID string
	Service   `json:"service"`
}
