package config

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

// ListManagedRules invokes the config.ListManagedRules API synchronously
func (client *Client) ListManagedRules(request *ListManagedRulesRequest) (response *ListManagedRulesResponse, err error) {
	response = CreateListManagedRulesResponse()
	err = client.DoAction(request, response)
	return
}

// ListManagedRulesWithChan invokes the config.ListManagedRules API asynchronously
func (client *Client) ListManagedRulesWithChan(request *ListManagedRulesRequest) (<-chan *ListManagedRulesResponse, <-chan error) {
	responseChan := make(chan *ListManagedRulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListManagedRules(request)
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

// ListManagedRulesWithCallback invokes the config.ListManagedRules API asynchronously
func (client *Client) ListManagedRulesWithCallback(request *ListManagedRulesRequest, callback func(response *ListManagedRulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListManagedRulesResponse
		var err error
		defer close(result)
		response, err = client.ListManagedRules(request)
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

// ListManagedRulesRequest is the request struct for api ListManagedRules
type ListManagedRulesRequest struct {
	*requests.RpcRequest
	RiskLevel  requests.Integer `position:"Query" name:"RiskLevel"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	Keyword    string           `position:"Query" name:"Keyword"`
}

// ListManagedRulesResponse is the response struct for api ListManagedRules
type ListManagedRulesResponse struct {
	*responses.BaseResponse
	RequestId    string       `json:"RequestId" xml:"RequestId"`
	ManagedRules ManagedRules `json:"ManagedRules" xml:"ManagedRules"`
}

// CreateListManagedRulesRequest creates a request to invoke ListManagedRules API
func CreateListManagedRulesRequest() (request *ListManagedRulesRequest) {
	request = &ListManagedRulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Config", "2020-09-07", "ListManagedRules", "", "")
	request.Method = requests.POST
	return
}

// CreateListManagedRulesResponse creates a response to parse from ListManagedRules response
func CreateListManagedRulesResponse() (response *ListManagedRulesResponse) {
	response = &ListManagedRulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
