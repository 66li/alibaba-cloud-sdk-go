package dataworks_public

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

// GetMetaTableListByCategory invokes the dataworks_public.GetMetaTableListByCategory API synchronously
// api document: https://help.aliyun.com/api/dataworks-public/getmetatablelistbycategory.html
func (client *Client) GetMetaTableListByCategory(request *GetMetaTableListByCategoryRequest) (response *GetMetaTableListByCategoryResponse, err error) {
	response = CreateGetMetaTableListByCategoryResponse()
	err = client.DoAction(request, response)
	return
}

// GetMetaTableListByCategoryWithChan invokes the dataworks_public.GetMetaTableListByCategory API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/getmetatablelistbycategory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetMetaTableListByCategoryWithChan(request *GetMetaTableListByCategoryRequest) (<-chan *GetMetaTableListByCategoryResponse, <-chan error) {
	responseChan := make(chan *GetMetaTableListByCategoryResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetMetaTableListByCategory(request)
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

// GetMetaTableListByCategoryWithCallback invokes the dataworks_public.GetMetaTableListByCategory API asynchronously
// api document: https://help.aliyun.com/api/dataworks-public/getmetatablelistbycategory.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetMetaTableListByCategoryWithCallback(request *GetMetaTableListByCategoryRequest, callback func(response *GetMetaTableListByCategoryResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetMetaTableListByCategoryResponse
		var err error
		defer close(result)
		response, err = client.GetMetaTableListByCategory(request)
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

// GetMetaTableListByCategoryRequest is the request struct for api GetMetaTableListByCategory
type GetMetaTableListByCategoryRequest struct {
	*requests.RpcRequest
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	CategoryId requests.Integer `position:"Query" name:"CategoryId"`
}

// GetMetaTableListByCategoryResponse is the response struct for api GetMetaTableListByCategory
type GetMetaTableListByCategoryResponse struct {
	*responses.BaseResponse
	RequestId      string `json:"RequestId" xml:"RequestId"`
	ErrorCode      string `json:"ErrorCode" xml:"ErrorCode"`
	ErrorMessage   string `json:"ErrorMessage" xml:"ErrorMessage"`
	HttpStatusCode int    `json:"HttpStatusCode" xml:"HttpStatusCode"`
	Success        bool   `json:"Success" xml:"Success"`
	Data           Data   `json:"Data" xml:"Data"`
}

// CreateGetMetaTableListByCategoryRequest creates a request to invoke GetMetaTableListByCategory API
func CreateGetMetaTableListByCategoryRequest() (request *GetMetaTableListByCategoryRequest) {
	request = &GetMetaTableListByCategoryRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "GetMetaTableListByCategory", "dide", "openAPI")
	request.Method = requests.GET
	return
}

// CreateGetMetaTableListByCategoryResponse creates a response to parse from GetMetaTableListByCategory response
func CreateGetMetaTableListByCategoryResponse() (response *GetMetaTableListByCategoryResponse) {
	response = &GetMetaTableListByCategoryResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}