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

type ApplicationCredentialsUsernameTemplate struct {
	Suffix string `json:"suffix,omitempty"`
	Template string `json:"template,omitempty"`
	Type string `json:"type,omitempty"`
}

func (m *ApplicationCredentialsUsernameTemplate) WithSuffix(v string) *ApplicationCredentialsUsernameTemplate {
	m.Suffix = v
	return m
}

func (m *ApplicationCredentialsUsernameTemplate) WithTemplate(v string) *ApplicationCredentialsUsernameTemplate {
	m.Template = v
	return m
}

func (m *ApplicationCredentialsUsernameTemplate) WithType(v string) *ApplicationCredentialsUsernameTemplate {
	m.Type = v
	return m
}


