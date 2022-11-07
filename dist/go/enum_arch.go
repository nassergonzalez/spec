// Code generated by scripts/generate.js; DO NOT EDIT.

// Copyright 2022 Harness, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package yaml

import "encoding/json"

// List of cpu architectures.
type Arch int

// Arch enumeration.
const (
	ArchNone Arch = iota
	ArchAmd64
	ArchArm
	ArchArm64
	Arch386
	ArchPpc
	ArchPpc64
	ArchPpc64le
	ArchRiscv
	ArchRiscv64
	ArchS390
	ArchS390x
	ArchSparc
	ArchSparc64
)

// String returns the Arch as a string.
func (e Arch) String() string {
	switch e {
	case ArchAmd64:
		return "amd64"
	case ArchArm:
		return "arm"
	case ArchArm64:
		return "arm64"
	case Arch386:
		return "386"
	case ArchPpc:
		return "ppc"
	case ArchPpc64:
		return "ppc64"
	case ArchPpc64le:
		return "ppc64le"
	case ArchRiscv:
		return "riscv"
	case ArchRiscv64:
		return "riscv64"
	case ArchS390:
		return "s390"
	case ArchS390x:
		return "s390x"
	case ArchSparc:
		return "sparc"
	case ArchSparc64:
		return "sparc64"
	default:
		return ""
	}
}

// MarshalJSON marshals the type as a JSON string.
func (e Arch) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

// UnmarshalJSON unmashals a quoted json string to the enum value.
func (e *Arch) UnmarshalJSON(b []byte) error {
	var v string
	json.Unmarshal(b, &v)
	switch v {
	case "amd64":
		*e = ArchAmd64
	case "arm":
		*e = ArchArm
	case "arm64":
		*e = ArchArm64
	case "386":
		*e = Arch386
	case "ppc":
		*e = ArchPpc
	case "ppc64":
		*e = ArchPpc64
	case "ppc64le":
		*e = ArchPpc64le
	case "riscv":
		*e = ArchRiscv
	case "riscv64":
		*e = ArchRiscv64
	case "s390":
		*e = ArchS390
	case "s390x":
		*e = ArchS390x
	case "sparc":
		*e = ArchSparc
	case "sparc64":
		*e = ArchSparc64
	default:
		*e = ArchNone
	}
	return nil
}