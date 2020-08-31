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

// GetDeviceLiveUrl invokes the vcs.GetDeviceLiveUrl API synchronously
// api document: https://help.aliyun.com/api/vcs/getdeviceliveurl.html
func (client *Client) GetDeviceLiveUrl(request *GetDeviceLiveUrlRequest) (response *GetDeviceLiveUrlResponse, err error) {
	response = CreateGetDeviceLiveUrlResponse()
	err = client.DoAction(request, response)
	return
}

// GetDeviceLiveUrlWithChan invokes the vcs.GetDeviceLiveUrl API asynchronously
// api document: https://help.aliyun.com/api/vcs/getdeviceliveurl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDeviceLiveUrlWithChan(request *GetDeviceLiveUrlRequest) (<-chan *GetDeviceLiveUrlResponse, <-chan error) {
	responseChan := make(chan *GetDeviceLiveUrlResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetDeviceLiveUrl(request)
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

// GetDeviceLiveUrlWithCallback invokes the vcs.GetDeviceLiveUrl API asynchronously
// api document: https://help.aliyun.com/api/vcs/getdeviceliveurl.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetDeviceLiveUrlWithCallback(request *GetDeviceLiveUrlRequest, callback func(response *GetDeviceLiveUrlResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetDeviceLiveUrlResponse
		var err error
		defer close(result)
		response, err = client.GetDeviceLiveUrl(request)
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

// GetDeviceLiveUrlRequest is the request struct for api GetDeviceLiveUrl
type GetDeviceLiveUrlRequest struct {
	*requests.RpcRequest
	CorpId string `position:"Body" name:"CorpId"`
	GbId   string `position:"Body" name:"GbId"`
}

// GetDeviceLiveUrlResponse is the response struct for api GetDeviceLiveUrl
type GetDeviceLiveUrlResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Url       string `json:"Url" xml:"Url"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateGetDeviceLiveUrlRequest creates a request to invoke GetDeviceLiveUrl API
func CreateGetDeviceLiveUrlRequest() (request *GetDeviceLiveUrlRequest) {
	request = &GetDeviceLiveUrlRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "GetDeviceLiveUrl", "vcs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetDeviceLiveUrlResponse creates a response to parse from GetDeviceLiveUrl response
func CreateGetDeviceLiveUrlResponse() (response *GetDeviceLiveUrlResponse) {
	response = &GetDeviceLiveUrlResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
