package arms

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

// UpdateIntegration invokes the arms.UpdateIntegration API synchronously
func (client *Client) UpdateIntegration(request *UpdateIntegrationRequest) (response *UpdateIntegrationResponse, err error) {
	response = CreateUpdateIntegrationResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateIntegrationWithChan invokes the arms.UpdateIntegration API asynchronously
func (client *Client) UpdateIntegrationWithChan(request *UpdateIntegrationRequest) (<-chan *UpdateIntegrationResponse, <-chan error) {
	responseChan := make(chan *UpdateIntegrationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateIntegration(request)
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

// UpdateIntegrationWithCallback invokes the arms.UpdateIntegration API asynchronously
func (client *Client) UpdateIntegrationWithCallback(request *UpdateIntegrationRequest, callback func(response *UpdateIntegrationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateIntegrationResponse
		var err error
		defer close(result)
		response, err = client.UpdateIntegration(request)
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

// UpdateIntegrationRequest is the request struct for api UpdateIntegration
type UpdateIntegrationRequest struct {
	*requests.RpcRequest
	ShortToken                 string           `position:"Body" name:"ShortToken"`
	FieldRedefineRules         string           `position:"Body" name:"FieldRedefineRules"`
	Stat                       string           `position:"Body" name:"Stat"`
	Liveness                   string           `position:"Body" name:"Liveness"`
	IntegrationId              requests.Integer `position:"Body" name:"IntegrationId"`
	Description                string           `position:"Body" name:"Description"`
	ApiEndpoint                string           `position:"Body" name:"ApiEndpoint"`
	AutoRecover                requests.Boolean `position:"Body" name:"AutoRecover"`
	RecoverTime                requests.Integer `position:"Body" name:"RecoverTime"`
	DuplicateKey               string           `position:"Body" name:"DuplicateKey"`
	IntegrationName            string           `position:"Body" name:"IntegrationName"`
	State                      requests.Boolean `position:"Body" name:"State"`
	ExtendedFieldRedefineRules string           `position:"Body" name:"ExtendedFieldRedefineRules"`
	IntegrationProductType     string           `position:"Body" name:"IntegrationProductType"`
}

// UpdateIntegrationResponse is the response struct for api UpdateIntegration
type UpdateIntegrationResponse struct {
	*responses.BaseResponse
	RequestId   string      `json:"RequestId" xml:"RequestId"`
	Integration Integration `json:"Integration" xml:"Integration"`
}

// CreateUpdateIntegrationRequest creates a request to invoke UpdateIntegration API
func CreateUpdateIntegrationRequest() (request *UpdateIntegrationRequest) {
	request = &UpdateIntegrationRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("ARMS", "2019-08-08", "UpdateIntegration", "arms", "openAPI")
	request.Method = requests.POST
	return
}

// CreateUpdateIntegrationResponse creates a response to parse from UpdateIntegration response
func CreateUpdateIntegrationResponse() (response *UpdateIntegrationResponse) {
	response = &UpdateIntegrationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
