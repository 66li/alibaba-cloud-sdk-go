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

// GetURLUploadInfos invokes the vod.GetURLUploadInfos API synchronously
// api document: https://help.aliyun.com/api/vod/geturluploadinfos.html
func (client *Client) GetURLUploadInfos(request *GetURLUploadInfosRequest) (response *GetURLUploadInfosResponse, err error) {
	response = CreateGetURLUploadInfosResponse()
	err = client.DoAction(request, response)
	return
}

// GetURLUploadInfosWithChan invokes the vod.GetURLUploadInfos API asynchronously
// api document: https://help.aliyun.com/api/vod/geturluploadinfos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetURLUploadInfosWithChan(request *GetURLUploadInfosRequest) (<-chan *GetURLUploadInfosResponse, <-chan error) {
	responseChan := make(chan *GetURLUploadInfosResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetURLUploadInfos(request)
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

// GetURLUploadInfosWithCallback invokes the vod.GetURLUploadInfos API asynchronously
// api document: https://help.aliyun.com/api/vod/geturluploadinfos.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetURLUploadInfosWithCallback(request *GetURLUploadInfosRequest, callback func(response *GetURLUploadInfosResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetURLUploadInfosResponse
		var err error
		defer close(result)
		response, err = client.GetURLUploadInfos(request)
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

// GetURLUploadInfosRequest is the request struct for api GetURLUploadInfos
type GetURLUploadInfosRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	JobIds               string           `position:"Query" name:"JobIds"`
	UploadURLs           string           `position:"Query" name:"UploadURLs"`
}

// GetURLUploadInfosResponse is the response struct for api GetURLUploadInfos
type GetURLUploadInfosResponse struct {
	*responses.BaseResponse
	RequestId         string                `json:"RequestId" xml:"RequestId"`
	NonExists         []string              `json:"NonExists" xml:"NonExists"`
	URLUploadInfoList []UrlUploadJobInfoDTO `json:"URLUploadInfoList" xml:"URLUploadInfoList"`
}

// CreateGetURLUploadInfosRequest creates a request to invoke GetURLUploadInfos API
func CreateGetURLUploadInfosRequest() (request *GetURLUploadInfosRequest) {
	request = &GetURLUploadInfosRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "GetURLUploadInfos", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateGetURLUploadInfosResponse creates a response to parse from GetURLUploadInfos response
func CreateGetURLUploadInfosResponse() (response *GetURLUploadInfosResponse) {
	response = &GetURLUploadInfosResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
