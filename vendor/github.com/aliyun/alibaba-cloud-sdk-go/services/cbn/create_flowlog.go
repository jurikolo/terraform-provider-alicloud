package cbn

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

// CreateFlowlog invokes the cbn.CreateFlowlog API synchronously
func (client *Client) CreateFlowlog(request *CreateFlowlogRequest) (response *CreateFlowlogResponse, err error) {
	response = CreateCreateFlowlogResponse()
	err = client.DoAction(request, response)
	return
}

// CreateFlowlogWithChan invokes the cbn.CreateFlowlog API asynchronously
func (client *Client) CreateFlowlogWithChan(request *CreateFlowlogRequest) (<-chan *CreateFlowlogResponse, <-chan error) {
	responseChan := make(chan *CreateFlowlogResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateFlowlog(request)
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

// CreateFlowlogWithCallback invokes the cbn.CreateFlowlog API asynchronously
func (client *Client) CreateFlowlogWithCallback(request *CreateFlowlogRequest, callback func(response *CreateFlowlogResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateFlowlogResponse
		var err error
		defer close(result)
		response, err = client.CreateFlowlog(request)
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

// CreateFlowlogRequest is the request struct for api CreateFlowlog
type CreateFlowlogRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	ClientToken          string           `position:"Query" name:"ClientToken"`
	CenId                string           `position:"Query" name:"CenId"`
	Description          string           `position:"Query" name:"Description"`
	ProjectName          string           `position:"Query" name:"ProjectName"`
	LogStoreName         string           `position:"Query" name:"LogStoreName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string           `position:"Query" name:"OwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	FlowLogName          string           `position:"Query" name:"FlowLogName"`
}

// CreateFlowlogResponse is the response struct for api CreateFlowlog
type CreateFlowlogResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
	Success   string `json:"Success" xml:"Success"`
	FlowLogId string `json:"FlowLogId" xml:"FlowLogId"`
}

// CreateCreateFlowlogRequest creates a request to invoke CreateFlowlog API
func CreateCreateFlowlogRequest() (request *CreateFlowlogRequest) {
	request = &CreateFlowlogRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Cbn", "2017-09-12", "CreateFlowlog", "cbn", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateFlowlogResponse creates a response to parse from CreateFlowlog response
func CreateCreateFlowlogResponse() (response *CreateFlowlogResponse) {
	response = &CreateFlowlogResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
