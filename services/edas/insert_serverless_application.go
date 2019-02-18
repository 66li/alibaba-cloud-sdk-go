package edas

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

// InsertServerlessApplication invokes the edas.InsertServerlessApplication API synchronously
// api document: https://help.aliyun.com/api/edas/insertserverlessapplication.html
func (client *Client) InsertServerlessApplication(request *InsertServerlessApplicationRequest) (response *InsertServerlessApplicationResponse, err error) {
	response = CreateInsertServerlessApplicationResponse()
	err = client.DoAction(request, response)
	return
}

// InsertServerlessApplicationWithChan invokes the edas.InsertServerlessApplication API asynchronously
// api document: https://help.aliyun.com/api/edas/insertserverlessapplication.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) InsertServerlessApplicationWithChan(request *InsertServerlessApplicationRequest) (<-chan *InsertServerlessApplicationResponse, <-chan error) {
	responseChan := make(chan *InsertServerlessApplicationResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.InsertServerlessApplication(request)
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

// InsertServerlessApplicationWithCallback invokes the edas.InsertServerlessApplication API asynchronously
// api document: https://help.aliyun.com/api/edas/insertserverlessapplication.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) InsertServerlessApplicationWithCallback(request *InsertServerlessApplicationRequest, callback func(response *InsertServerlessApplicationResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *InsertServerlessApplicationResponse
		var err error
		defer close(result)
		response, err = client.InsertServerlessApplication(request)
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

// InsertServerlessApplicationRequest is the request struct for api InsertServerlessApplication
type InsertServerlessApplicationRequest struct {
	*requests.RoaRequest
	WebContainer    string           `position:"Query" name:"WebContainer"`
	JarStartArgs    string           `position:"Query" name:"JarStartArgs"`
	Memory          requests.Integer `position:"Query" name:"Memory"`
	CommandArgs     string           `position:"Query" name:"CommandArgs"`
	Replicas        requests.Integer `position:"Query" name:"Replicas"`
	Readiness       string           `position:"Query" name:"Readiness"`
	Liveness        string           `position:"Query" name:"Liveness"`
	Cpu             requests.Integer `position:"Query" name:"Cpu"`
	Envs            string           `position:"Query" name:"Envs"`
	PackageVersion  string           `position:"Query" name:"PackageVersion"`
	Command         string           `position:"Query" name:"Command"`
	CustomHostAlias string           `position:"Query" name:"CustomHostAlias"`
	Deploy          requests.Boolean `position:"Query" name:"Deploy"`
	VSwitchId       string           `position:"Query" name:"VSwitchId"`
	Jdk             string           `position:"Query" name:"Jdk"`
	AppDescription  string           `position:"Query" name:"AppDescription"`
	JarStartOptions string           `position:"Query" name:"JarStartOptions"`
	AppName         string           `position:"Query" name:"AppName"`
	NamespaceId     string           `position:"Query" name:"NamespaceId"`
	PackageUrl      string           `position:"Query" name:"PackageUrl"`
	VpcId           string           `position:"Query" name:"VpcId"`
	ImageUrl        string           `position:"Query" name:"ImageUrl"`
	PackageType     string           `position:"Query" name:"PackageType"`
}

// InsertServerlessApplicationResponse is the response struct for api InsertServerlessApplication
type InsertServerlessApplicationResponse struct {
	*responses.BaseResponse
	Code    int    `json:"Code" xml:"Code"`
	Message string `json:"Message" xml:"Message"`
	Data    Data   `json:"Data" xml:"Data"`
}

// CreateInsertServerlessApplicationRequest creates a request to invoke InsertServerlessApplication API
func CreateInsertServerlessApplicationRequest() (request *InsertServerlessApplicationRequest) {
	request = &InsertServerlessApplicationRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("Edas", "2017-08-01", "InsertServerlessApplication", "/pop/v5/k8s/pop/pop_serverless_app_create_without_deploy", "", "")
	request.Method = requests.POST
	return
}

// CreateInsertServerlessApplicationResponse creates a response to parse from InsertServerlessApplication response
func CreateInsertServerlessApplicationResponse() (response *InsertServerlessApplicationResponse) {
	response = &InsertServerlessApplicationResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
