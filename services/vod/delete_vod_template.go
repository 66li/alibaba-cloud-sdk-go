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

// DeleteVodTemplate invokes the vod.DeleteVodTemplate API synchronously
// api document: https://help.aliyun.com/api/vod/deletevodtemplate.html
func (client *Client) DeleteVodTemplate(request *DeleteVodTemplateRequest) (response *DeleteVodTemplateResponse, err error) {
	response = CreateDeleteVodTemplateResponse()
	err = client.DoAction(request, response)
	return
}

// DeleteVodTemplateWithChan invokes the vod.DeleteVodTemplate API asynchronously
// api document: https://help.aliyun.com/api/vod/deletevodtemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteVodTemplateWithChan(request *DeleteVodTemplateRequest) (<-chan *DeleteVodTemplateResponse, <-chan error) {
	responseChan := make(chan *DeleteVodTemplateResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DeleteVodTemplate(request)
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

// DeleteVodTemplateWithCallback invokes the vod.DeleteVodTemplate API asynchronously
// api document: https://help.aliyun.com/api/vod/deletevodtemplate.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DeleteVodTemplateWithCallback(request *DeleteVodTemplateRequest, callback func(response *DeleteVodTemplateResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DeleteVodTemplateResponse
		var err error
		defer close(result)
		response, err = client.DeleteVodTemplate(request)
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

// DeleteVodTemplateRequest is the request struct for api DeleteVodTemplate
type DeleteVodTemplateRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	VodTemplateId        string           `position:"Query" name:"VodTemplateId"`
}

// DeleteVodTemplateResponse is the response struct for api DeleteVodTemplate
type DeleteVodTemplateResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	VodTemplateId string `json:"VodTemplateId" xml:"VodTemplateId"`
}

// CreateDeleteVodTemplateRequest creates a request to invoke DeleteVodTemplate API
func CreateDeleteVodTemplateRequest() (request *DeleteVodTemplateRequest) {
	request = &DeleteVodTemplateRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("vod", "2017-03-21", "DeleteVodTemplate", "vod", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDeleteVodTemplateResponse creates a response to parse from DeleteVodTemplate response
func CreateDeleteVodTemplateResponse() (response *DeleteVodTemplateResponse) {
	response = &DeleteVodTemplateResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
