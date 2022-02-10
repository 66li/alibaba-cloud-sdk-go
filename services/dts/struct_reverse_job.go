package dts

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

// ReverseJob is a nested struct in dts response
type ReverseJob struct {
	OriginType                    string                          `json:"OriginType" xml:"OriginType"`
	IsDemoJob                     bool                            `json:"IsDemoJob" xml:"IsDemoJob"`
	BeginTimestamp                string                          `json:"BeginTimestamp" xml:"BeginTimestamp"`
	Reserved                      string                          `json:"Reserved" xml:"Reserved"`
	DtsJobId                      string                          `json:"DtsJobId" xml:"DtsJobId"`
	GroupId                       string                          `json:"GroupId" xml:"GroupId"`
	JobType                       string                          `json:"JobType" xml:"JobType"`
	Checkpoint                    string                          `json:"Checkpoint" xml:"Checkpoint"`
	PayType                       string                          `json:"PayType" xml:"PayType"`
	DtsJobClass                   string                          `json:"DtsJobClass" xml:"DtsJobClass"`
	FinishTime                    string                          `json:"FinishTime" xml:"FinishTime"`
	Status                        string                          `json:"Status" xml:"Status"`
	ConsumptionCheckpoint         string                          `json:"ConsumptionCheckpoint" xml:"ConsumptionCheckpoint"`
	TaskType                      string                          `json:"TaskType" xml:"TaskType"`
	DtsJobDirection               string                          `json:"DtsJobDirection" xml:"DtsJobDirection"`
	ErrorMessage                  string                          `json:"ErrorMessage" xml:"ErrorMessage"`
	CreateTime                    string                          `json:"CreateTime" xml:"CreateTime"`
	EndTimestamp                  string                          `json:"EndTimestamp" xml:"EndTimestamp"`
	DatabaseCount                 int                             `json:"DatabaseCount" xml:"DatabaseCount"`
	ConsumptionClient             string                          `json:"ConsumptionClient" xml:"ConsumptionClient"`
	DbObject                      string                          `json:"DbObject" xml:"DbObject"`
	SynchronizationDirection      string                          `json:"SynchronizationDirection" xml:"SynchronizationDirection"`
	Delay                         int64                           `json:"Delay" xml:"Delay"`
	ExpireTime                    string                          `json:"ExpireTime" xml:"ExpireTime"`
	DtsJobName                    string                          `json:"DtsJobName" xml:"DtsJobName"`
	EtlCalculator                 string                          `json:"EtlCalculator" xml:"EtlCalculator"`
	ReverseJob                    string                          `json:"ReverseJob" xml:"ReverseJob"`
	DestNetType                   string                          `json:"DestNetType" xml:"DestNetType"`
	AppName                       string                          `json:"AppName" xml:"AppName"`
	DtsInstanceID                 string                          `json:"DtsInstanceID" xml:"DtsInstanceID"`
	SubscribeTopic                string                          `json:"SubscribeTopic" xml:"SubscribeTopic"`
	MigrationMode                 MigrationMode                   `json:"MigrationMode" xml:"MigrationMode"`
	DataSynchronizationStatus     DataSynchronizationStatus       `json:"DataSynchronizationStatus" xml:"DataSynchronizationStatus"`
	SourceEndpoint                SourceEndpoint                  `json:"SourceEndpoint" xml:"SourceEndpoint"`
	DataInitializationStatus      DataInitializationStatus        `json:"DataInitializationStatus" xml:"DataInitializationStatus"`
	DestinationEndpoint           DestinationEndpoint             `json:"DestinationEndpoint" xml:"DestinationEndpoint"`
	RetryState                    RetryState                      `json:"RetryState" xml:"RetryState"`
	DataEtlStatus                 DataEtlStatus                   `json:"DataEtlStatus" xml:"DataEtlStatus"`
	SubscriptionHost              SubscriptionHost                `json:"SubscriptionHost" xml:"SubscriptionHost"`
	StructureInitializationStatus StructureInitializationStatus   `json:"StructureInitializationStatus" xml:"StructureInitializationStatus"`
	PrecheckStatus                PrecheckStatusInDescribeDtsJobs `json:"PrecheckStatus" xml:"PrecheckStatus"`
	Performance                   Performance                     `json:"Performance" xml:"Performance"`
	SubscriptionDataType          SubscriptionDataType            `json:"SubscriptionDataType" xml:"SubscriptionDataType"`
	TagList                       []TagListItem                   `json:"TagList" xml:"TagList"`
}
