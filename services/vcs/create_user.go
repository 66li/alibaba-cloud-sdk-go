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

// CreateUser invokes the vcs.CreateUser API synchronously
// api document: https://help.aliyun.com/api/vcs/createuser.html
func (client *Client) CreateUser(request *CreateUserRequest) (response *CreateUserResponse, err error) {
	response = CreateCreateUserResponse()
	err = client.DoAction(request, response)
	return
}

// CreateUserWithChan invokes the vcs.CreateUser API asynchronously
// api document: https://help.aliyun.com/api/vcs/createuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateUserWithChan(request *CreateUserRequest) (<-chan *CreateUserResponse, <-chan error) {
	responseChan := make(chan *CreateUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateUser(request)
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

// CreateUserWithCallback invokes the vcs.CreateUser API asynchronously
// api document: https://help.aliyun.com/api/vcs/createuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateUserWithCallback(request *CreateUserRequest, callback func(response *CreateUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateUserResponse
		var err error
		defer close(result)
		response, err = client.CreateUser(request)
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

// CreateUserRequest is the request struct for api CreateUser
type CreateUserRequest struct {
	*requests.RpcRequest
	CorpId       string           `position:"Body" name:"CorpId"`
	Gender       requests.Integer `position:"Body" name:"Gender"`
	PlateNo      string           `position:"Body" name:"PlateNo"`
	IdNumber     string           `position:"Body" name:"IdNumber"`
	FaceImageUrl string           `position:"Body" name:"FaceImageUrl"`
	Attachment   string           `position:"Body" name:"Attachment"`
	IsvSubId     string           `position:"Body" name:"IsvSubId"`
	Address      string           `position:"Body" name:"Address"`
	UserGroupId  requests.Integer `position:"Body" name:"UserGroupId"`
	PhoneNo      string           `position:"Body" name:"PhoneNo"`
	BizId        string           `position:"Body" name:"BizId"`
	Age          requests.Integer `position:"Body" name:"Age"`
	UserName     string           `position:"Body" name:"UserName"`
}

// CreateUserResponse is the response struct for api CreateUser
type CreateUserResponse struct {
	*responses.BaseResponse
	Code      string `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateCreateUserRequest creates a request to invoke CreateUser API
func CreateCreateUserRequest() (request *CreateUserRequest) {
	request = &CreateUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vcs", "2020-05-15", "CreateUser", "vcs", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateUserResponse creates a response to parse from CreateUser response
func CreateCreateUserResponse() (response *CreateUserResponse) {
	response = &CreateUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
