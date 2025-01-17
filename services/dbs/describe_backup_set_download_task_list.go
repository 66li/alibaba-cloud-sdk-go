package dbs

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

// DescribeBackupSetDownloadTaskList invokes the dbs.DescribeBackupSetDownloadTaskList API synchronously
func (client *Client) DescribeBackupSetDownloadTaskList(request *DescribeBackupSetDownloadTaskListRequest) (response *DescribeBackupSetDownloadTaskListResponse, err error) {
	response = CreateDescribeBackupSetDownloadTaskListResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeBackupSetDownloadTaskListWithChan invokes the dbs.DescribeBackupSetDownloadTaskList API asynchronously
func (client *Client) DescribeBackupSetDownloadTaskListWithChan(request *DescribeBackupSetDownloadTaskListRequest) (<-chan *DescribeBackupSetDownloadTaskListResponse, <-chan error) {
	responseChan := make(chan *DescribeBackupSetDownloadTaskListResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeBackupSetDownloadTaskList(request)
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

// DescribeBackupSetDownloadTaskListWithCallback invokes the dbs.DescribeBackupSetDownloadTaskList API asynchronously
func (client *Client) DescribeBackupSetDownloadTaskListWithCallback(request *DescribeBackupSetDownloadTaskListRequest, callback func(response *DescribeBackupSetDownloadTaskListResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeBackupSetDownloadTaskListResponse
		var err error
		defer close(result)
		response, err = client.DescribeBackupSetDownloadTaskList(request)
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

// DescribeBackupSetDownloadTaskListRequest is the request struct for api DescribeBackupSetDownloadTaskList
type DescribeBackupSetDownloadTaskListRequest struct {
	*requests.RpcRequest
	ClientToken             string           `position:"Query" name:"ClientToken"`
	BackupSetDownloadTaskId string           `position:"Query" name:"BackupSetDownloadTaskId"`
	BackupPlanId            string           `position:"Query" name:"BackupPlanId"`
	PageNum                 requests.Integer `position:"Query" name:"PageNum"`
	OwnerId                 string           `position:"Query" name:"OwnerId"`
	PageSize                requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeBackupSetDownloadTaskListResponse is the response struct for api DescribeBackupSetDownloadTaskList
type DescribeBackupSetDownloadTaskListResponse struct {
	*responses.BaseResponse
	HttpStatusCode int                                      `json:"HttpStatusCode" xml:"HttpStatusCode"`
	PageNum        int                                      `json:"PageNum" xml:"PageNum"`
	RequestId      string                                   `json:"RequestId" xml:"RequestId"`
	ErrCode        string                                   `json:"ErrCode" xml:"ErrCode"`
	Success        bool                                     `json:"Success" xml:"Success"`
	ErrMessage     string                                   `json:"ErrMessage" xml:"ErrMessage"`
	TotalPages     int                                      `json:"TotalPages" xml:"TotalPages"`
	TotalElements  int                                      `json:"TotalElements" xml:"TotalElements"`
	PageSize       int                                      `json:"PageSize" xml:"PageSize"`
	Items          ItemsInDescribeBackupSetDownloadTaskList `json:"Items" xml:"Items"`
}

// CreateDescribeBackupSetDownloadTaskListRequest creates a request to invoke DescribeBackupSetDownloadTaskList API
func CreateDescribeBackupSetDownloadTaskListRequest() (request *DescribeBackupSetDownloadTaskListRequest) {
	request = &DescribeBackupSetDownloadTaskListRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Dbs", "2019-03-06", "DescribeBackupSetDownloadTaskList", "", "")
	request.Method = requests.POST
	return
}

// CreateDescribeBackupSetDownloadTaskListResponse creates a response to parse from DescribeBackupSetDownloadTaskList response
func CreateDescribeBackupSetDownloadTaskListResponse() (response *DescribeBackupSetDownloadTaskListResponse) {
	response = &DescribeBackupSetDownloadTaskListResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
