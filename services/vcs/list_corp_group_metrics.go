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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// ListCorpGroupMetrics invokes the vcs.ListCorpGroupMetrics API synchronously
// api document: https://help.aliyun.com/api/vcs/listcorpgroupmetrics.html
func (client *Client) ListCorpGroupMetrics(request *ListCorpGroupMetricsRequest) (response *ListCorpGroupMetricsResponse, err error) {
	response = CreateListCorpGroupMetricsResponse()
	err = client.DoAction(request, response)
	return
}

// ListCorpGroupMetricsWithChan invokes the vcs.ListCorpGroupMetrics API asynchronously
// api document: https://help.aliyun.com/api/vcs/listcorpgroupmetrics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListCorpGroupMetricsWithChan(request *ListCorpGroupMetricsRequest) (<-chan *ListCorpGroupMetricsResponse, <-chan error) {
	responseChan := make(chan *ListCorpGroupMetricsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListCorpGroupMetrics(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// ListCorpGroupMetricsWithCallback invokes the vcs.ListCorpGroupMetrics API asynchronously
// api document: https://help.aliyun.com/api/vcs/listcorpgroupmetrics.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListCorpGroupMetricsWithCallback(request *ListCorpGroupMetricsRequest, callback func(response *ListCorpGroupMetricsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListCorpGroupMetricsResponse
		var err error
		defer close(result)
		response, err = client.ListCorpGroupMetrics(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// ListCorpGroupMetricsRequest is the request struct for api ListCorpGroupMetrics
type ListCorpGroupMetricsRequest struct {
	*requests.RpcRequest
	CorpId      string `position:"Body" name:"CorpId"`
	GroupId     string `position:"Body" name:"GroupId"`
	EndTime     string `position:"Body" name:"EndTime"`
	StartTime   string `position:"Body" name:"StartTime"`
	DeviceId    string `position:"Body" name:"DeviceId"`
	PageNumber  string `position:"Body" name:"PageNumber"`
	DeviceGroup string `position:"Body" name:"DeviceGroup"`
	TagCode     string `position:"Body" name:"TagCode"`
	PageSize    string `position:"Body" name:"PageSize"`
	UserGroup   string `position:"Body" name:"UserGroup"`
}

// ListCorpGroupMetricsResponse is the response struct for api ListCorpGroupMetrics
type ListCorpGroupMetricsResponse struct {
	*responses.BaseResponse
	Code       string     `json:"Code" xml:"Code"`
	Message    string     `json:"Message" xml:"Message"`
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	PageNumber int        `json:"PageNumber" xml:"PageNumber"`
	PageSize   int        `json:"PageSize" xml:"PageSize"`
	TotalCount int        `json:"TotalCount" xml:"TotalCount"`
	Success    string     `json:"Success" xml:"Success"`
	Data       []DataItem `json:"Data" xml:"Data"`
}

// CreateListCorpGroupMetricsRequest creates a request to invoke ListCorpGroupMetrics API
func CreateListCorpGroupMetricsRequest() (request *ListCorpGroupMetricsRequest) {
	request = &ListCorpGroupMetricsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "ListCorpGroupMetrics", "vcs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateListCorpGroupMetricsResponse creates a response to parse from ListCorpGroupMetrics response
func CreateListCorpGroupMetricsResponse() (response *ListCorpGroupMetricsResponse) {
	response = &ListCorpGroupMetricsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
