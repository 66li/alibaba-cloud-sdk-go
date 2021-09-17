package dyplsapi

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

// QueryPhoneNoAByTrackNo invokes the dyplsapi.QueryPhoneNoAByTrackNo API synchronously
func (client *Client) QueryPhoneNoAByTrackNo(request *QueryPhoneNoAByTrackNoRequest) (response *QueryPhoneNoAByTrackNoResponse, err error) {
	response = CreateQueryPhoneNoAByTrackNoResponse()
	err = client.DoAction(request, response)
	return
}

// QueryPhoneNoAByTrackNoWithChan invokes the dyplsapi.QueryPhoneNoAByTrackNo API asynchronously
func (client *Client) QueryPhoneNoAByTrackNoWithChan(request *QueryPhoneNoAByTrackNoRequest) (<-chan *QueryPhoneNoAByTrackNoResponse, <-chan error) {
	responseChan := make(chan *QueryPhoneNoAByTrackNoResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.QueryPhoneNoAByTrackNo(request)
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

// QueryPhoneNoAByTrackNoWithCallback invokes the dyplsapi.QueryPhoneNoAByTrackNo API asynchronously
func (client *Client) QueryPhoneNoAByTrackNoWithCallback(request *QueryPhoneNoAByTrackNoRequest, callback func(response *QueryPhoneNoAByTrackNoResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *QueryPhoneNoAByTrackNoResponse
		var err error
		defer close(result)
		response, err = client.QueryPhoneNoAByTrackNo(request)
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

// QueryPhoneNoAByTrackNoRequest is the request struct for api QueryPhoneNoAByTrackNo
type QueryPhoneNoAByTrackNoRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	CabinetNo            string           `position:"Query" name:"CabinetNo"`
	PhoneNoX             string           `position:"Query" name:"PhoneNoX"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	TrackNo              string           `position:"Query" name:"trackNo"`
}

// QueryPhoneNoAByTrackNoResponse is the response struct for api QueryPhoneNoAByTrackNo
type QueryPhoneNoAByTrackNoResponse struct {
	*responses.BaseResponse
	Code      string         `json:"Code" xml:"Code"`
	Message   string         `json:"Message" xml:"Message"`
	RequestId string         `json:"RequestId" xml:"RequestId"`
	Module    []PhoneNoAInfo `json:"Module" xml:"Module"`
}

// CreateQueryPhoneNoAByTrackNoRequest creates a request to invoke QueryPhoneNoAByTrackNo API
func CreateQueryPhoneNoAByTrackNoRequest() (request *QueryPhoneNoAByTrackNoRequest) {
	request = &QueryPhoneNoAByTrackNoRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dyplsapi", "2017-05-25", "QueryPhoneNoAByTrackNo", "", "")
	request.Method = requests.POST
	return
}

// CreateQueryPhoneNoAByTrackNoResponse creates a response to parse from QueryPhoneNoAByTrackNo response
func CreateQueryPhoneNoAByTrackNoResponse() (response *QueryPhoneNoAByTrackNoResponse) {
	response = &QueryPhoneNoAByTrackNoResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
