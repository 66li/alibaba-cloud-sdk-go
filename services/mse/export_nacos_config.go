package mse

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

// ExportNacosConfig invokes the mse.ExportNacosConfig API synchronously
func (client *Client) ExportNacosConfig(request *ExportNacosConfigRequest) (response *ExportNacosConfigResponse, err error) {
	response = CreateExportNacosConfigResponse()
	err = client.DoAction(request, response)
	return
}

// ExportNacosConfigWithChan invokes the mse.ExportNacosConfig API asynchronously
func (client *Client) ExportNacosConfigWithChan(request *ExportNacosConfigRequest) (<-chan *ExportNacosConfigResponse, <-chan error) {
	responseChan := make(chan *ExportNacosConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ExportNacosConfig(request)
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

// ExportNacosConfigWithCallback invokes the mse.ExportNacosConfig API asynchronously
func (client *Client) ExportNacosConfigWithCallback(request *ExportNacosConfigRequest, callback func(response *ExportNacosConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ExportNacosConfigResponse
		var err error
		defer close(result)
		response, err = client.ExportNacosConfig(request)
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

// ExportNacosConfigRequest is the request struct for api ExportNacosConfig
type ExportNacosConfigRequest struct {
	*requests.RpcRequest
	InstanceId  string `position:"Query" name:"InstanceId"`
	DataId      string `position:"Query" name:"DataId"`
	AppName     string `position:"Query" name:"AppName"`
	NamespaceId string `position:"Query" name:"NamespaceId"`
	Ids         string `position:"Query" name:"Ids"`
	Group       string `position:"Query" name:"Group"`
}

// ExportNacosConfigResponse is the response struct for api ExportNacosConfig
type ExportNacosConfigResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	Success        bool   `json:"Success" xml:"Success"`
	Message        string `json:"Message" xml:"Message"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Code           int    `json:"Code" xml:"Code"`
	DynamicMessage string `json:"DynamicMessage" xml:"DynamicMessage"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateExportNacosConfigRequest creates a request to invoke ExportNacosConfig API
func CreateExportNacosConfigRequest() (request *ExportNacosConfigRequest) {
	request = &ExportNacosConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "ExportNacosConfig", "", "")
	request.Method = requests.POST
	return
}

// CreateExportNacosConfigResponse creates a response to parse from ExportNacosConfig response
func CreateExportNacosConfigResponse() (response *ExportNacosConfigResponse) {
	response = &ExportNacosConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
