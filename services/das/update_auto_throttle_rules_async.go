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

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// UpdateAutoThrottleRulesAsync invokes the das.UpdateAutoThrottleRulesAsync API synchronously
func (client *Client) UpdateAutoThrottleRulesAsync(request *UpdateAutoThrottleRulesAsyncRequest) (response *UpdateAutoThrottleRulesAsyncResponse, err error) {
	response = CreateUpdateAutoThrottleRulesAsyncResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateAutoThrottleRulesAsyncWithChan invokes the das.UpdateAutoThrottleRulesAsync API asynchronously
func (client *Client) UpdateAutoThrottleRulesAsyncWithChan(request *UpdateAutoThrottleRulesAsyncRequest) (<-chan *UpdateAutoThrottleRulesAsyncResponse, <-chan error) {
	responseChan := make(chan *UpdateAutoThrottleRulesAsyncResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateAutoThrottleRulesAsync(request)
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

// UpdateAutoThrottleRulesAsyncWithCallback invokes the das.UpdateAutoThrottleRulesAsync API asynchronously
func (client *Client) UpdateAutoThrottleRulesAsyncWithCallback(request *UpdateAutoThrottleRulesAsyncRequest, callback func(response *UpdateAutoThrottleRulesAsyncResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateAutoThrottleRulesAsyncResponse
		var err error
		defer close(result)
		response, err = client.UpdateAutoThrottleRulesAsync(request)
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

// UpdateAutoThrottleRulesAsyncRequest is the request struct for api UpdateAutoThrottleRulesAsync
type UpdateAutoThrottleRulesAsyncRequest struct {
	*requests.RpcRequest
	ResultId               string           `position:"Query" name:"ResultId"`
	CpuSessionRelation     string           `position:"Query" name:"CpuSessionRelation"`
	AbnormalDuration       string           `position:"Query" name:"AbnormalDuration"`
	MaxThrottleTime        string           `position:"Query" name:"MaxThrottleTime"`
	ConsoleContext         string           `position:"Query" name:"ConsoleContext"`
	CpuUsage               string           `position:"Query" name:"CpuUsage"`
	AutoKillSession        requests.Boolean `position:"Query" name:"AutoKillSession"`
	AllowThrottleStartTime string           `position:"Query" name:"AllowThrottleStartTime"`
	AllowThrottleEndTime   string           `position:"Query" name:"AllowThrottleEndTime"`
	InstanceIds            string           `position:"Query" name:"InstanceIds"`
	ActiveSessions         requests.Integer `position:"Query" name:"ActiveSessions"`
}

// UpdateAutoThrottleRulesAsyncResponse is the response struct for api UpdateAutoThrottleRulesAsync
type UpdateAutoThrottleRulesAsyncResponse struct {
	*responses.BaseResponse
	Code      int64                              `json:"Code" xml:"Code"`
	Message   string                             `json:"Message" xml:"Message"`
	RequestId string                             `json:"RequestId" xml:"RequestId"`
	Success   bool                               `json:"Success" xml:"Success"`
	Data      DataInUpdateAutoThrottleRulesAsync `json:"Data" xml:"Data"`
}

// CreateUpdateAutoThrottleRulesAsyncRequest creates a request to invoke UpdateAutoThrottleRulesAsync API
func CreateUpdateAutoThrottleRulesAsyncRequest() (request *UpdateAutoThrottleRulesAsyncRequest) {
	request = &UpdateAutoThrottleRulesAsyncRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("DAS", "2020-01-16", "UpdateAutoThrottleRulesAsync", "", "")
	request.Method = requests.POST
	return
}

// CreateUpdateAutoThrottleRulesAsyncResponse creates a response to parse from UpdateAutoThrottleRulesAsync response
func CreateUpdateAutoThrottleRulesAsyncResponse() (response *UpdateAutoThrottleRulesAsyncResponse) {
	response = &UpdateAutoThrottleRulesAsyncResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
