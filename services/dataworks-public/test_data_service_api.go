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

// TestDataServiceApi invokes the dataworks_public.TestDataServiceApi API synchronously
func (client *Client) TestDataServiceApi(request *TestDataServiceApiRequest) (response *TestDataServiceApiResponse, err error) {
	response = CreateTestDataServiceApiResponse()
	err = client.DoAction(request, response)
	return
}

// TestDataServiceApiWithChan invokes the dataworks_public.TestDataServiceApi API asynchronously
func (client *Client) TestDataServiceApiWithChan(request *TestDataServiceApiRequest) (<-chan *TestDataServiceApiResponse, <-chan error) {
	responseChan := make(chan *TestDataServiceApiResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.TestDataServiceApi(request)
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

// TestDataServiceApiWithCallback invokes the dataworks_public.TestDataServiceApi API asynchronously
func (client *Client) TestDataServiceApiWithCallback(request *TestDataServiceApiRequest, callback func(response *TestDataServiceApiResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *TestDataServiceApiResponse
		var err error
		defer close(result)
		response, err = client.TestDataServiceApi(request)
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

// TestDataServiceApiRequest is the request struct for api TestDataServiceApi
type TestDataServiceApiRequest struct {
	*requests.RpcRequest
	PathParams  *[]TestDataServiceApiPathParams `position:"Body" name:"PathParams"  type:"Repeated"`
	BodyContent string                          `position:"Body" name:"BodyContent"`
	BodyParams  *[]TestDataServiceApiBodyParams `position:"Body" name:"BodyParams"  type:"Repeated"`
	QueryParam  *[]TestDataServiceApiQueryParam `position:"Body" name:"QueryParam"  type:"Repeated"`
	HeadParams  *[]TestDataServiceApiHeadParams `position:"Body" name:"HeadParams"  type:"Repeated"`
	ApiId       requests.Integer                `position:"Query" name:"ApiId"`
}

// TestDataServiceApiPathParams is a repeated param struct in TestDataServiceApiRequest
type TestDataServiceApiPathParams struct {
	ParamKey   string `name:"ParamKey"`
	ParamValue string `name:"ParamValue"`
}

// TestDataServiceApiBodyParams is a repeated param struct in TestDataServiceApiRequest
type TestDataServiceApiBodyParams struct {
	ParamKey   string `name:"ParamKey"`
	ParamValue string `name:"ParamValue"`
}

// TestDataServiceApiQueryParam is a repeated param struct in TestDataServiceApiRequest
type TestDataServiceApiQueryParam struct {
	ParamKey   string `name:"ParamKey"`
	ParamValue string `name:"ParamValue"`
}

// TestDataServiceApiHeadParams is a repeated param struct in TestDataServiceApiRequest
type TestDataServiceApiHeadParams struct {
	ParamKey   string `name:"ParamKey"`
	ParamValue string `name:"ParamValue"`
}

// TestDataServiceApiResponse is the response struct for api TestDataServiceApi
type TestDataServiceApiResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Data      Data   `json:"Data" xml:"Data"`
}

// CreateTestDataServiceApiRequest creates a request to invoke TestDataServiceApi API
func CreateTestDataServiceApiRequest() (request *TestDataServiceApiRequest) {
	request = &TestDataServiceApiRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("dataworks-public", "2020-05-18", "TestDataServiceApi", "", "")
	request.Method = requests.POST
	return
}

// CreateTestDataServiceApiResponse creates a response to parse from TestDataServiceApi response
func CreateTestDataServiceApiResponse() (response *TestDataServiceApiResponse) {
	response = &TestDataServiceApiResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
