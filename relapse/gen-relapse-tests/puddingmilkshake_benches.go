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
	. "github.com/katydid/katydid/relapse/combinator"
	. "github.com/katydid/katydid/relapse/funcs"
)

func init() {
	var bridgePepper = G{
		"main": InOrder(
			Any(),
			In("FinanceJudo",
				Any(),
				In("SaladWorry",
					Any(),
					In("SpyCarpenter",
						Any(),
						In("BridgePepper", InAny(Value(Contains(StringVar(), StringConst("a"))))),
						Any(),
					),
					Any(),
				),
				Any(),
			),
			Any(),
		),
	}
	BenchValidateProtoNumEtc("BridgePepper", bridgePepper, RandomPuddingMilkshake)
	ValidateProtoNum("BridgePepper", bridgePepper, &PuddingMilkshake{FinanceJudo: &FinanceJudo{SaladWorry: &SaladWorry{SpyCarpenter: &SpyCarpenter{BridgePepper: []string{"a"}}}}}, true)

	var bridgePepperAndFountainTarget = G{
		"main": InOrder(
			Any(),
			In("FinanceJudo",
				Any(),
				In("SaladWorry",
					Any(),
					In("SpyCarpenter",
						AllOf(
							InPath("BridgePepper", InAny(Value(Contains(StringVar(), StringConst("a"))))),
							InPath("FountainTarget", InAny(Value(Contains(StringVar(), StringConst("a"))))),
						),
					),
					Any(),
				),
				Any(),
			),
			Any(),
		),
	}
	BenchValidateProtoNumEtc("BridgePepperAndFountainTarget", bridgePepperAndFountainTarget, RandomPuddingMilkshake)
	ValidateProtoNum("BenchBridgePepperAndFountainTarget", bridgePepperAndFountainTarget, &PuddingMilkshake{FinanceJudo: &FinanceJudo{SaladWorry: &SaladWorry{SpyCarpenter: &SpyCarpenter{BridgePepper: []string{"a"}, FountainTarget: []string{"a"}}}}}, true)
}
