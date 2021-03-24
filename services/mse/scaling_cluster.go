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

// ScalingCluster invokes the mse.ScalingCluster API synchronously
func (client *Client) ScalingCluster(request *ScalingClusterRequest) (response *ScalingClusterResponse, err error) {
	response = CreateScalingClusterResponse()
	err = client.DoAction(request, response)
	return
}

// ScalingClusterWithChan invokes the mse.ScalingCluster API asynchronously
func (client *Client) ScalingClusterWithChan(request *ScalingClusterRequest) (<-chan *ScalingClusterResponse, <-chan error) {
	responseChan := make(chan *ScalingClusterResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ScalingCluster(request)
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

// ScalingClusterWithCallback invokes the mse.ScalingCluster API asynchronously
func (client *Client) ScalingClusterWithCallback(request *ScalingClusterRequest, callback func(response *ScalingClusterResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ScalingClusterResponse
		var err error
		defer close(result)
		response, err = client.ScalingCluster(request)
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

// ScalingClusterRequest is the request struct for api ScalingCluster
type ScalingClusterRequest struct {
	*requests.RpcRequest
	ClusterSpecification string           `position:"Query" name:"ClusterSpecification"`
	Cpu                  requests.Integer `position:"Query" name:"Cpu"`
	ClusterId            string           `position:"Query" name:"ClusterId"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	MemoryCapacity       requests.Integer `position:"Query" name:"MemoryCapacity"`
	InstanceCount        requests.Integer `position:"Query" name:"InstanceCount"`
}

// ScalingClusterResponse is the response struct for api ScalingCluster
type ScalingClusterResponse struct {
	*responses.BaseResponse
	Success   bool   `json:"Success" xml:"Success"`
	Message   string `json:"Message" xml:"Message"`
	ErrorCode string `json:"ErrorCode" xml:"ErrorCode"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateScalingClusterRequest creates a request to invoke ScalingCluster API
func CreateScalingClusterRequest() (request *ScalingClusterRequest) {
	request = &ScalingClusterRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("mse", "2019-05-31", "ScalingCluster", "", "")
	request.Method = requests.POST
	return
}

// CreateScalingClusterResponse creates a response to parse from ScalingCluster response
func CreateScalingClusterResponse() (response *ScalingClusterResponse) {
	response = &ScalingClusterResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
