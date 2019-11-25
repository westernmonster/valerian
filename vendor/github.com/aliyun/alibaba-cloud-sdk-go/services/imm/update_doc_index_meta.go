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

// UpdateDocIndexMeta invokes the imm.UpdateDocIndexMeta API synchronously
// api document: https://help.aliyun.com/api/imm/updatedocindexmeta.html
func (client *Client) UpdateDocIndexMeta(request *UpdateDocIndexMetaRequest) (response *UpdateDocIndexMetaResponse, err error) {
	response = CreateUpdateDocIndexMetaResponse()
	err = client.DoAction(request, response)
	return
}

// UpdateDocIndexMetaWithChan invokes the imm.UpdateDocIndexMeta API asynchronously
// api document: https://help.aliyun.com/api/imm/updatedocindexmeta.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateDocIndexMetaWithChan(request *UpdateDocIndexMetaRequest) (<-chan *UpdateDocIndexMetaResponse, <-chan error) {
	responseChan := make(chan *UpdateDocIndexMetaResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.UpdateDocIndexMeta(request)
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

// UpdateDocIndexMetaWithCallback invokes the imm.UpdateDocIndexMeta API asynchronously
// api document: https://help.aliyun.com/api/imm/updatedocindexmeta.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) UpdateDocIndexMetaWithCallback(request *UpdateDocIndexMetaRequest, callback func(response *UpdateDocIndexMetaResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *UpdateDocIndexMetaResponse
		var err error
		defer close(result)
		response, err = client.UpdateDocIndexMeta(request)
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

// UpdateDocIndexMetaRequest is the request struct for api UpdateDocIndexMeta
type UpdateDocIndexMetaRequest struct {
	*requests.RpcRequest
	CustomKey1 string `position:"Query" name:"CustomKey1"`
	Set        string `position:"Query" name:"Set"`
	CustomKey5 string `position:"Query" name:"CustomKey5"`
	CustomKey4 string `position:"Query" name:"CustomKey4"`
	CustomKey3 string `position:"Query" name:"CustomKey3"`
	CustomKey2 string `position:"Query" name:"CustomKey2"`
	Project    string `position:"Query" name:"Project"`
	CustomKey6 string `position:"Query" name:"CustomKey6"`
	Name       string `position:"Query" name:"Name"`
	UniqueId   string `position:"Query" name:"UniqueId"`
}

// UpdateDocIndexMetaResponse is the response struct for api UpdateDocIndexMeta
type UpdateDocIndexMetaResponse struct {
	*responses.BaseResponse
	RequestId         string `json:"RequestId" xml:"RequestId"`
	IndexCreatedTime  string `json:"IndexCreatedTime" xml:"IndexCreatedTime"`
	IndexModifiedTime string `json:"IndexModifiedTime" xml:"IndexModifiedTime"`
}

// CreateUpdateDocIndexMetaRequest creates a request to invoke UpdateDocIndexMeta API
func CreateUpdateDocIndexMetaRequest() (request *UpdateDocIndexMetaRequest) {
	request = &UpdateDocIndexMetaRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "UpdateDocIndexMeta", "imm", "openAPI")
	return
}

// CreateUpdateDocIndexMetaResponse creates a response to parse from UpdateDocIndexMeta response
func CreateUpdateDocIndexMetaResponse() (response *UpdateDocIndexMetaResponse) {
	response = &UpdateDocIndexMetaResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
