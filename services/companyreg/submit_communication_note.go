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

// SubmitCommunicationNote invokes the companyreg.SubmitCommunicationNote API synchronously
func (client *Client) SubmitCommunicationNote(request *SubmitCommunicationNoteRequest) (response *SubmitCommunicationNoteResponse, err error) {
	response = CreateSubmitCommunicationNoteResponse()
	err = client.DoAction(request, response)
	return
}

// SubmitCommunicationNoteWithChan invokes the companyreg.SubmitCommunicationNote API asynchronously
func (client *Client) SubmitCommunicationNoteWithChan(request *SubmitCommunicationNoteRequest) (<-chan *SubmitCommunicationNoteResponse, <-chan error) {
	responseChan := make(chan *SubmitCommunicationNoteResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.SubmitCommunicationNote(request)
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

// SubmitCommunicationNoteWithCallback invokes the companyreg.SubmitCommunicationNote API asynchronously
func (client *Client) SubmitCommunicationNoteWithCallback(request *SubmitCommunicationNoteRequest, callback func(response *SubmitCommunicationNoteResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *SubmitCommunicationNoteResponse
		var err error
		defer close(result)
		response, err = client.SubmitCommunicationNote(request)
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

// SubmitCommunicationNoteRequest is the request struct for api SubmitCommunicationNote
type SubmitCommunicationNoteRequest struct {
	*requests.RpcRequest
	Note            string           `position:"Query" name:"Note"`
	Type            requests.Integer `position:"Query" name:"Type"`
	ActionRequestId string           `position:"Query" name:"ActionRequestId"`
	OperatorType    requests.Integer `position:"Query" name:"OperatorType"`
	BizCode         string           `position:"Query" name:"BizCode"`
	BizId           string           `position:"Query" name:"BizId"`
}

// SubmitCommunicationNoteResponse is the response struct for api SubmitCommunicationNote
type SubmitCommunicationNoteResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateSubmitCommunicationNoteRequest creates a request to invoke SubmitCommunicationNote API
func CreateSubmitCommunicationNoteRequest() (request *SubmitCommunicationNoteRequest) {
	request = &SubmitCommunicationNoteRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("companyreg", "2019-05-08", "SubmitCommunicationNote", "companyreg", "openAPI")
	request.Method = requests.POST
	return
}

// CreateSubmitCommunicationNoteResponse creates a response to parse from SubmitCommunicationNote response
func CreateSubmitCommunicationNoteResponse() (response *SubmitCommunicationNoteResponse) {
	response = &SubmitCommunicationNoteResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
