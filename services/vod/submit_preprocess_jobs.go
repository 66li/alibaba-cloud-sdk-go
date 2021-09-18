package vod

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

// SubmitPreprocessJobs invokes the vod.SubmitPreprocessJobs API synchronously
func (client *Client) SubmitPreprocessJobs(request *SubmitPreprocessJobsRequest) (response *SubmitPreprocessJobsResponse, err error) {
	response = CreateSubmitPreprocessJobsResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitPreprocessJobsWithChan invokes the vod.SubmitPreprocessJobs API asynchronously
func (client *Client) SubmitPreprocessJobsWithChan(request *SubmitPreprocessJobsRequest) (<-chan *SubmitPreprocessJobsResponse, <-chan error) {
	responseChan := make(chan *SubmitPreprocessJobsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitPreprocessJobs(request)
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

// SubmitPreprocessJobsWithCallback invokes the vod.SubmitPreprocessJobs API asynchronously
func (client *Client) SubmitPreprocessJobsWithCallback(request *SubmitPreprocessJobsRequest, callback func(response *SubmitPreprocessJobsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitPreprocessJobsResponse
		var err error
		defer close(result)
		response, err = client.SubmitPreprocessJobs(request)
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

// SubmitPreprocessJobsRequest is the request struct for api SubmitPreprocessJobs
type SubmitPreprocessJobsRequest struct {
	*requests.RpcRequest
	VideoId        string `position:"Query" name:"VideoId"`
	PreprocessType string `position:"Query" name:"PreprocessType"`
}

// SubmitPreprocessJobsResponse is the response struct for api SubmitPreprocessJobs
type SubmitPreprocessJobsResponse struct {
	*responses.BaseResponse
	RequestId      string         `json:"RequestId" xml:"RequestId"`
	PreprocessJobs PreprocessJobs `json:"PreprocessJobs" xml:"PreprocessJobs"`
}

// CreateSubmitPreprocessJobsRequest creates a request to invoke SubmitPreprocessJobs API
func CreateSubmitPreprocessJobsRequest() (request *SubmitPreprocessJobsRequest) {
	request = &SubmitPreprocessJobsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "SubmitPreprocessJobs", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitPreprocessJobsResponse creates a response to parse from SubmitPreprocessJobs response
func CreateSubmitPreprocessJobsResponse() (response *SubmitPreprocessJobsResponse) {
	response = &SubmitPreprocessJobsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
