package alikafka

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

// CreatePostPayOrder invokes the alikafka.CreatePostPayOrder API synchronously
func (client *Client) CreatePostPayOrder(request *CreatePostPayOrderRequest) (response *CreatePostPayOrderResponse, err error) {
	response = CreateCreatePostPayOrderResponse()
	err = client.DoAction(request, response)
	return
}

// CreatePostPayOrderWithChan invokes the alikafka.CreatePostPayOrder API asynchronously
func (client *Client) CreatePostPayOrderWithChan(request *CreatePostPayOrderRequest) (<-chan *CreatePostPayOrderResponse, <-chan error) {
	responseChan := make(chan *CreatePostPayOrderResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreatePostPayOrder(request)
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

// CreatePostPayOrderWithCallback invokes the alikafka.CreatePostPayOrder API asynchronously
func (client *Client) CreatePostPayOrderWithCallback(request *CreatePostPayOrderRequest, callback func(response *CreatePostPayOrderResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreatePostPayOrderResponse
		var err error
		defer close(result)
		response, err = client.CreatePostPayOrder(request)
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

// CreatePostPayOrderRequest is the request struct for api CreatePostPayOrder
type CreatePostPayOrderRequest struct {
	*requests.RpcRequest
	PaidType        requests.Integer `position:"Query" name:"PaidType"`
	DiskSize        requests.Integer `position:"Query" name:"DiskSize"`
	IoMax           requests.Integer `position:"Query" name:"IoMax"`
	IoMaxSpec       string           `position:"Query" name:"IoMaxSpec"`
	DiskType        string           `position:"Query" name:"DiskType"`
	TopicQuota      requests.Integer `position:"Query" name:"TopicQuota"`
	EipMax          requests.Integer `position:"Query" name:"EipMax"`
	SpecType        string           `position:"Query" name:"SpecType"`
	ResourceGroupId string           `position:"Query" name:"ResourceGroupId"`
	DeployType      requests.Integer `position:"Query" name:"DeployType"`
}

// CreatePostPayOrderResponse is the response struct for api CreatePostPayOrder
type CreatePostPayOrderResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
	OrderId   string `json:"OrderId" xml:"OrderId"`
}

// CreateCreatePostPayOrderRequest creates a request to invoke CreatePostPayOrder API
func CreateCreatePostPayOrderRequest() (request *CreatePostPayOrderRequest) {
	request = &CreatePostPayOrderRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alikafka", "2019-09-16", "CreatePostPayOrder", "alikafka", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreatePostPayOrderResponse creates a response to parse from CreatePostPayOrder response
func CreateCreatePostPayOrderResponse() (response *CreatePostPayOrderResponse) {
	response = &CreatePostPayOrderResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
