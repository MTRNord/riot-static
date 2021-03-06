// Copyright 2017 Michael Telatynski <7t3chguy@gmail.com>
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

package mxclient

type PowerLevel int

func (powerLevel PowerLevel) String() string {
	switch int(powerLevel) {
	case 100:
		return "Admin"
	case 50:
		return "Moderator"
	case 0:
		return "User"
	case -1:
		return "Muted"
	default:
		return "Custom"
	}
}

func (powerLevel PowerLevel) Int() int {
	return int(powerLevel)
}

type MemberInfo struct {
	MXID        string
	Membership  string
	DisplayName string
	AvatarURL   MXCURL
	PowerLevel  PowerLevel
}

func NewMemberInfo(mxid string) *MemberInfo {
	return &MemberInfo{
		MXID:       mxid,
		Membership: "leave",
	}
}

func (memberInfo MemberInfo) GetName() string {
	if memberInfo.DisplayName != "" {
		return memberInfo.DisplayName
	} else {
		return memberInfo.MXID
	}
}
