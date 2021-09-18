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

// DeleteDynamicImage invokes the vod.DeleteDynamicImage API synchronously
func (client *Client) DeleteDynamicImage(request *DeleteDynamicImageRequest) (response *DeleteDynamicImageResponse, err error) {
	response = CreateDeleteDynamicImageResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteDynamicImageWithChan invokes the vod.DeleteDynamicImage API asynchronously
func (client *Client) DeleteDynamicImageWithChan(request *DeleteDynamicImageRequest) (<-chan *DeleteDynamicImageResponse, <-chan error) {
	responseChan := make(chan *DeleteDynamicImageResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteDynamicImage(request)
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

// DeleteDynamicImageWithCallback invokes the vod.DeleteDynamicImage API asynchronously
func (client *Client) DeleteDynamicImageWithCallback(request *DeleteDynamicImageRequest, callback func(response *DeleteDynamicImageResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteDynamicImageResponse
		var err error
		defer close(result)
		response, err = client.DeleteDynamicImage(request)
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

// DeleteDynamicImageRequest is the request struct for api DeleteDynamicImage
type DeleteDynamicImageRequest struct {
	*requests.RpcRequest
	VideoId         string `position:"Query" name:"VideoId"`
	DynamicImageIds string `position:"Query" name:"DynamicImageIds"`
}

// DeleteDynamicImageResponse is the response struct for api DeleteDynamicImage
type DeleteDynamicImageResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDeleteDynamicImageRequest creates a request to invoke DeleteDynamicImage API
func CreateDeleteDynamicImageRequest() (request *DeleteDynamicImageRequest) {
	request = &DeleteDynamicImageRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DeleteDynamicImage", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteDynamicImageResponse creates a response to parse from DeleteDynamicImage response
func CreateDeleteDynamicImageResponse() (response *DeleteDynamicImageResponse) {
	response = &DeleteDynamicImageResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
