package das

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// DataInGetAutoResourceOptimizeRules is a nested struct in das response
type DataInGetAutoResourceOptimizeRules struct {
	TotalAutoResourceOptimizeRulesCount                     int64                             `json:"TotalAutoResourceOptimizeRulesCount" xml:"TotalAutoResourceOptimizeRulesCount"`
	EnableAutoResourceOptimizeCount                         int64                             `json:"EnableAutoResourceOptimizeCount" xml:"EnableAutoResourceOptimizeCount"`
	TurnOffAutoResourceOptimizeCount                        int64                             `json:"TurnOffAutoResourceOptimizeCount" xml:"TurnOffAutoResourceOptimizeCount"`
	HasEnableRuleButNotDasProCount                          int64                             `json:"HasEnableRuleButNotDasProCount" xml:"HasEnableRuleButNotDasProCount"`
	NeverEnableAutoResourceOptimizeOrReleasedInstanceCount  int64                             `json:"NeverEnableAutoResourceOptimizeOrReleasedInstanceCount" xml:"NeverEnableAutoResourceOptimizeOrReleasedInstanceCount"`
	NeverEnableAutoResourceOptimizeOrReleasedInstanceIdList []string                          `json:"NeverEnableAutoResourceOptimizeOrReleasedInstanceIdList" xml:"NeverEnableAutoResourceOptimizeOrReleasedInstanceIdList"`
	EnableAutoResourceOptimizeList                          []EnableAutoResourceOptimizeList  `json:"EnableAutoResourceOptimizeList" xml:"EnableAutoResourceOptimizeList"`
	TurnOffAutoResourceOptimizeList                         []TurnOffAutoResourceOptimizeList `json:"TurnOffAutoResourceOptimizeList" xml:"TurnOffAutoResourceOptimizeList"`
	HasEnableRuleButNotDasProList                           []HasEnableRuleButNotDasProList   `json:"HasEnableRuleButNotDasProList" xml:"HasEnableRuleButNotDasProList"`
}
