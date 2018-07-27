/*
* Copyright 2018 - Present Okta, Inc.
*
* Licensed under the Apache License, Version 2.0 (the "License");
* you may not use this file except in compliance with the License.
* You may obtain a copy of the License at
*
*      http://www.apache.org/licenses/LICENSE-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an "AS IS" BASIS,
* WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
* See the License for the specific language governing permissions and
* limitations under the License.
*/

// AUTO-GENERATED!  DO NOT EDIT FILE DIRECTLY

package okta

import (
)

type ApplicationSettingsNotificationsVpn struct {
	HelpUrl string `json:"helpUrl,omitempty"`
	Message string `json:"message,omitempty"`
	Network *ApplicationSettingsNotificationsVpnNetwork `json:"network,omitempty"`
}

func (m *ApplicationSettingsNotificationsVpn) WithHelpUrl(v string) *ApplicationSettingsNotificationsVpn {
	m.HelpUrl = v
	return m
}

func (m *ApplicationSettingsNotificationsVpn) WithMessage(v string) *ApplicationSettingsNotificationsVpn {
	m.Message = v
	return m
}

func (m *ApplicationSettingsNotificationsVpn) WithNetwork(v *ApplicationSettingsNotificationsVpnNetwork) *ApplicationSettingsNotificationsVpn {
	m.Network = v
	return m
}


