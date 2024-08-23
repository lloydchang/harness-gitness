//  Copyright 2023 Harness, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package types

import (
	"time"

	"github.com/harness/gitness/registry/types/enum"
)

// CleanupPolicy DTO object.
type CleanupPolicy struct {
	ID            int64
	RegistryID    int64
	Name          string
	VersionPrefix []string
	PackagePrefix []string
	ExpiryTime    int64
	CreatedAt     time.Time
	UpdatedAt     time.Time
	CreatedBy     int64
	UpdatedBy     int64
}

// CleanupPolicyPrefix DTO object.
type CleanupPolicyPrefix struct {
	ID              int64
	CleanupPolicyID int64
	Prefix          string
	PrefixType      enum.PrefixType
}