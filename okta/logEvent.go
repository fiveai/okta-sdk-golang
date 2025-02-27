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
	"fmt"
	"github.com/fiveai/okta-sdk-golang/okta/query"
	"time"
)

type LogEventResource resource

type LogEvent struct {
	Actor                 *LogActor                 `json:"actor,omitempty"`
	AuthenticationContext *LogAuthenticationContext `json:"authenticationContext,omitempty"`
	Client                *LogClient                `json:"client,omitempty"`
	DebugContext          *LogDebugContext          `json:"debugContext,omitempty"`
	DisplayMessage        string                    `json:"displayMessage,omitempty"`
	EventType             string                    `json:"eventType,omitempty"`
	LegacyEventType       string                    `json:"legacyEventType,omitempty"`
	Outcome               *LogOutcome               `json:"outcome,omitempty"`
	Published             *time.Time                `json:"published,omitempty"`
	Request               *LogRequest               `json:"request,omitempty"`
	SecurityContext       *LogSecurityContext       `json:"securityContext,omitempty"`
	Severity              string                    `json:"severity,omitempty"`
	Target                []*LogTarget              `json:"target,omitempty"`
	Transaction           *LogTransaction           `json:"transaction,omitempty"`
	Uuid                  string                    `json:"uuid,omitempty"`
	Version               string                    `json:"version,omitempty"`
}

func (m *LogEventResource) GetLogs(qp *query.Params) ([]*LogEvent, *Response, error) {
	url := fmt.Sprintf("/api/v1/logs")
	if qp != nil {
		url = url + qp.String()
	}
	req, err := m.client.requestExecutor.NewRequest("GET", url, nil)
	if err != nil {
		return nil, nil, err
	}

	var logEvent []*LogEvent
	resp, err := m.client.requestExecutor.Do(req, &logEvent)
	if err != nil {
		return nil, resp, err
	}
	return logEvent, resp, nil
}
