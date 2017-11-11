//  Copyright 2015 Walter Schulze
//
//  Licensed under the Apache License, Version 2.0 (the "License");
//  you may not use this file except in compliance with the License.
//  You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
//  Unless required by applicable law or agreed to in writing, software
//  distributed under the License is distributed on an "AS IS" BASIS,
//  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//  See the License for the specific language governing permissions and
//  limitations under the License.

package main

import (
	"github.com/gogo/protobuf/proto"
	. "github.com/katydid/katydid/relapse/combinator"
	. "github.com/katydid/katydid/relapse/funcs"
)

func init() {
	var scarBusStop = G{
		"main": InOrder(
			Any(),
			In("PocketRoses",
				Any(),
				In("ScarBusStop", Value(Contains(StringVar(), StringConst("a")))),
				Any(),
			),
			Any(),
		),
	}
	BenchValidateProtoNumEtc("TypewriterPrisonScarBusStop", scarBusStop, RandomTypewriterPrison)
	ValidateProtoNum("TypewriterPrisonScarBusStop", scarBusStop, &TypewriterPrison{PocketRoses: &PocketRoses{ScarBusStop: proto.String("a")}}, true)

	var daisySled = G{
		"main": InOrder(
			Any(),
			In("PocketRoses",
				Any(),
				In("DaisySled", Value(IntEq(IntVar(), IntConst(1)))),
				Any(),
			),
			Any(),
		),
	}
	BenchValidateProtoNumEtc("TypewriterPrisonDaisySled", daisySled, RandomTypewriterPrison)
	ValidateProtoNum("TypewriterPrisonDaisySled", daisySled, &TypewriterPrison{PocketRoses: &PocketRoses{DaisySled: proto.Int64(1)}}, true)

	var smileLetter = G{
		"main": InOrder(
			Any(),
			In("PocketRoses",
				Any(),
				In("SmileLetter", Value(BoolVar())),
				Any(),
			),
			Any(),
		),
	}
	BenchValidateProtoNumEtc("TypewriterPrisonSmileLetter", smileLetter, RandomTypewriterPrison)
	ValidateProtoNum("TypewriterPrisonSmileLetter", smileLetter, &TypewriterPrison{PocketRoses: &PocketRoses{SmileLetter: proto.Bool(true)}}, true)

	var menuPaperClip = G{
		"main": InOrder(
			Any(),
			In("PocketRoses",
				Any(),
				In("MenuPaperclip",
					InAny(Value(Contains(StringVar(), StringConst("a")))),
					Any(),
				),
				Any(),
			),
			Any(),
		),
	}
	BenchValidateProtoNumEtc("TypewriterPrisonMenuPaperclip", menuPaperClip, RandomTypewriterPrison)
	ValidateProtoNum("TypewriterPrisonMenuPaperclip", menuPaperClip, &TypewriterPrison{PocketRoses: &PocketRoses{MenuPaperclip: []string{"a"}}}, true)

	var mapShark = G{
		"main": InOrder(
			Any(),
			In("PocketRoses",
				Any(),
				In("MapShark", Value(Contains(StringVar(), StringConst("a")))),
				Any(),
			),
			Any(),
		),
	}
	BenchValidateProtoNumEtc("TypewriterPrisonMapShark", mapShark, RandomTypewriterPrison)
	ValidateProtoNum("TypewriterPrisonMapShark", mapShark, &TypewriterPrison{PocketRoses: &PocketRoses{MapShark: proto.String("a")}}, true)

}
