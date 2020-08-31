package vcs

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

// Data is a nested struct in vcs response
type Data struct {
	PhoneNo      string                          `json:"PhoneNo" xml:"PhoneNo"`
	Name         string                          `json:"Name" xml:"Name"`
	PageSize     int64                           `json:"PageSize" xml:"PageSize"`
	MaxId        string                          `json:"MaxId" xml:"MaxId"`
	Attachment   string                          `json:"Attachment" xml:"Attachment"`
	CatalogId    int                             `json:"CatalogId" xml:"CatalogId"`
	UserId       int                             `json:"UserId" xml:"UserId"`
	FaceUrl      string                          `json:"FaceUrl" xml:"FaceUrl"`
	IsvSubId     string                          `json:"IsvSubId" xml:"IsvSubId"`
	Age          string                          `json:"Age" xml:"Age"`
	OssPath      string                          `json:"OssPath" xml:"OssPath"`
	BizId        string                          `json:"BizId" xml:"BizId"`
	Success      int64                           `json:"Success" xml:"Success"`
	TaskId       string                          `json:"TaskId" xml:"TaskId"`
	Description  string                          `json:"Description" xml:"Description"`
	LiveAddress  string                          `json:"LiveAddress" xml:"LiveAddress"`
	PageNumber   int64                           `json:"PageNumber" xml:"PageNumber"`
	DataSourceId string                          `json:"DataSourceId" xml:"DataSourceId"`
	SceneType    string                          `json:"SceneType" xml:"SceneType"`
	TotalCount   int64                           `json:"TotalCount" xml:"TotalCount"`
	UserGroupId  int                             `json:"UserGroupId" xml:"UserGroupId"`
	PersonId     string                          `json:"PersonId" xml:"PersonId"`
	KafkaTopic   string                          `json:"KafkaTopic" xml:"KafkaTopic"`
	StructList   string                          `json:"StructList" xml:"StructList"`
	Address      string                          `json:"Address" xml:"Address"`
	FaceImageUrl string                          `json:"FaceImageUrl" xml:"FaceImageUrl"`
	IdNumber     string                          `json:"IdNumber" xml:"IdNumber"`
	Gender       string                          `json:"Gender" xml:"Gender"`
	QualityScore string                          `json:"QualityScore" xml:"QualityScore"`
	UserName     string                          `json:"UserName" xml:"UserName"`
	PicUrl       string                          `json:"PicUrl" xml:"PicUrl"`
	Total        int64                           `json:"Total" xml:"Total"`
	PlateNo      string                          `json:"PlateNo" xml:"PlateNo"`
	ProfileId    int                             `json:"ProfileId" xml:"ProfileId"`
	Attributes   Attributes                      `json:"Attributes" xml:"Attributes"`
	ResultObject []ResultObjectItem              `json:"ResultObject" xml:"ResultObject"`
	Records      []RecordsItemInGetMonitorResult `json:"Records" xml:"Records"`
	TagList      []TagListItem                   `json:"TagList" xml:"TagList"`
	FaceList     []Face                          `json:"FaceList" xml:"FaceList"`
	BodyList     []Body                          `json:"BodyList" xml:"BodyList"`
}
