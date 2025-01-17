package alikafka

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

// StartInstance invokes the alikafka.StartInstance API synchronously
func (client *Client) StartInstance(request *StartInstanceRequest) (response *StartInstanceResponse, err error) {
	response = CreateStartInstanceResponse()
	err = client.DoAction(request, response)
	return
}

// StartInstanceWithChan invokes the alikafka.StartInstance API asynchronously
func (client *Client) StartInstanceWithChan(request *StartInstanceRequest) (<-chan *StartInstanceResponse, <-chan error) {
	responseChan := make(chan *StartInstanceResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.StartInstance(request)
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

// StartInstanceWithCallback invokes the alikafka.StartInstance API asynchronously
func (client *Client) StartInstanceWithCallback(request *StartInstanceRequest, callback func(response *StartInstanceResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *StartInstanceResponse
		var err error
		defer close(result)
		response, err = client.StartInstance(request)
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

// StartInstanceRequest is the request struct for api StartInstance
type StartInstanceRequest struct {
	*requests.RpcRequest
	SelectedZones        string           `position:"Query" name:"SelectedZones"`
	IsEipInner           requests.Boolean `position:"Query" name:"IsEipInner"`
	SecurityGroup        string           `position:"Query" name:"SecurityGroup"`
	DeployModule         string           `position:"Query" name:"DeployModule"`
	IsSetUserAndPassword requests.Boolean `position:"Query" name:"IsSetUserAndPassword"`
	Password             string           `position:"Query" name:"Password"`
	Notifier             string           `position:"Query" name:"Notifier"`
	IsForceSelectedZones requests.Boolean `position:"Query" name:"IsForceSelectedZones"`
	VSwitchId            string           `position:"Query" name:"VSwitchId"`
	UserPhoneNum         string           `position:"Query" name:"UserPhoneNum"`
	InstanceId           string           `position:"Query" name:"InstanceId"`
	VpcId                string           `position:"Query" name:"VpcId"`
	CrossZone            requests.Boolean `position:"Query" name:"CrossZone"`
	Name                 string           `position:"Query" name:"Name"`
	ServiceVersion       string           `position:"Query" name:"ServiceVersion"`
	ZoneId               string           `position:"Query" name:"ZoneId"`
	KMSKeyId             string           `position:"Query" name:"KMSKeyId"`
	Config               string           `position:"Query" name:"Config"`
	Username             string           `position:"Query" name:"Username"`
}

// StartInstanceResponse is the response struct for api StartInstance
type StartInstanceResponse struct {
	*responses.BaseResponse
	Code      int    `json:"Code" xml:"Code"`
	Message   string `json:"Message" xml:"Message"`
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   bool   `json:"Success" xml:"Success"`
}

// CreateStartInstanceRequest creates a request to invoke StartInstance API
func CreateStartInstanceRequest() (request *StartInstanceRequest) {
	request = &StartInstanceRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("alikafka", "2019-09-16", "StartInstance", "alikafka", "openAPI")
	request.Method = requests.POST
	return
}

// CreateStartInstanceResponse creates a response to parse from StartInstance response
func CreateStartInstanceResponse() (response *StartInstanceResponse) {
	response = &StartInstanceResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
