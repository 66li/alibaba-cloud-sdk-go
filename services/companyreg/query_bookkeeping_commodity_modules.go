package companyreg

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

// QueryBookkeepingCommodityModules invokes the companyreg.QueryBookkeepingCommodityModules API synchronously
func (client *Client) QueryBookkeepingCommodityModules(request *QueryBookkeepingCommodityModulesRequest) (response *QueryBookkeepingCommodityModulesResponse, err error) {
	response = CreateQueryBookkeepingCommodityModulesResponse()
	err = client.DoAction(request, response)
	return
}

// QueryBookkeepingCommodityModulesWithChan invokes the companyreg.QueryBookkeepingCommodityModules API asynchronously
func (client *Client) QueryBookkeepingCommodityModulesWithChan(request *QueryBookkeepingCommodityModulesRequest) (<-chan *QueryBookkeepingCommodityModulesResponse, <-chan error) {
	responseChan := make(chan *QueryBookkeepingCommodityModulesResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryBookkeepingCommodityModules(request)
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

// QueryBookkeepingCommodityModulesWithCallback invokes the companyreg.QueryBookkeepingCommodityModules API asynchronously
func (client *Client) QueryBookkeepingCommodityModulesWithCallback(request *QueryBookkeepingCommodityModulesRequest, callback func(response *QueryBookkeepingCommodityModulesResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryBookkeepingCommodityModulesResponse
		var err error
		defer close(result)
		response, err = client.QueryBookkeepingCommodityModules(request)
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

// QueryBookkeepingCommodityModulesRequest is the request struct for api QueryBookkeepingCommodityModules
type QueryBookkeepingCommodityModulesRequest struct {
	*requests.RpcRequest
}

// QueryBookkeepingCommodityModulesResponse is the response struct for api QueryBookkeepingCommodityModules
type QueryBookkeepingCommodityModulesResponse struct {
	*responses.BaseResponse
	RequestId string                 `json:"RequestId" xml:"RequestId"`
	Data      map[string]interface{} `json:"Data" xml:"Data"`
}

// CreateQueryBookkeepingCommodityModulesRequest creates a request to invoke QueryBookkeepingCommodityModules API
func CreateQueryBookkeepingCommodityModulesRequest() (request *QueryBookkeepingCommodityModulesRequest) {
	request = &QueryBookkeepingCommodityModulesRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2019-05-08", "QueryBookkeepingCommodityModules", "companyreg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateQueryBookkeepingCommodityModulesResponse creates a response to parse from QueryBookkeepingCommodityModules response
func CreateQueryBookkeepingCommodityModulesResponse() (response *QueryBookkeepingCommodityModulesResponse) {
	response = &QueryBookkeepingCommodityModulesResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
