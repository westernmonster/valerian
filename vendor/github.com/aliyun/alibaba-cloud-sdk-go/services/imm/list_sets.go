package imm

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

// ListSets invokes the imm.ListSets API synchronously
// api document: https://help.aliyun.com/api/imm/listsets.html
func (client *Client) ListSets(request *ListSetsRequest) (response *ListSetsResponse, err error) {
	response = CreateListSetsResponse()
	err = client.DoAction(request, response)
	return
}

// ListSetsWithChan invokes the imm.ListSets API asynchronously
// api document: https://help.aliyun.com/api/imm/listsets.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListSetsWithChan(request *ListSetsRequest) (<-chan *ListSetsResponse, <-chan error) {
	responseChan := make(chan *ListSetsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.ListSets(request)
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

// ListSetsWithCallback invokes the imm.ListSets API asynchronously
// api document: https://help.aliyun.com/api/imm/listsets.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) ListSetsWithCallback(request *ListSetsRequest, callback func(response *ListSetsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *ListSetsResponse
		var err error
		defer close(result)
		response, err = client.ListSets(request)
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

// ListSetsRequest is the request struct for api ListSets
type ListSetsRequest struct {
	*requests.RpcRequest
	Marker  string `position:"Query" name:"Marker"`
	Project string `position:"Query" name:"Project"`
}

// ListSetsResponse is the response struct for api ListSets
type ListSetsResponse struct {
	*responses.BaseResponse
	RequestId  string     `json:"RequestId" xml:"RequestId"`
	NextMarker string     `json:"NextMarker" xml:"NextMarker"`
	Sets       []SetsItem `json:"Sets" xml:"Sets"`
}

// CreateListSetsRequest creates a request to invoke ListSets API
func CreateListSetsRequest() (request *ListSetsRequest) {
	request = &ListSetsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "ListSets", "imm", "openAPI")
	return
}

// CreateListSetsResponse creates a response to parse from ListSets response
func CreateListSetsResponse() (response *ListSetsResponse) {
	response = &ListSetsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
