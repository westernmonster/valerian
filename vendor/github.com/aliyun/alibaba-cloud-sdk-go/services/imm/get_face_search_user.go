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

// GetFaceSearchUser invokes the imm.GetFaceSearchUser API synchronously
// api document: https://help.aliyun.com/api/imm/getfacesearchuser.html
func (client *Client) GetFaceSearchUser(request *GetFaceSearchUserRequest) (response *GetFaceSearchUserResponse, err error) {
	response = CreateGetFaceSearchUserResponse()
	err = client.DoAction(request, response)
	return
}

// GetFaceSearchUserWithChan invokes the imm.GetFaceSearchUser API asynchronously
// api document: https://help.aliyun.com/api/imm/getfacesearchuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetFaceSearchUserWithChan(request *GetFaceSearchUserRequest) (<-chan *GetFaceSearchUserResponse, <-chan error) {
	responseChan := make(chan *GetFaceSearchUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetFaceSearchUser(request)
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

// GetFaceSearchUserWithCallback invokes the imm.GetFaceSearchUser API asynchronously
// api document: https://help.aliyun.com/api/imm/getfacesearchuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetFaceSearchUserWithCallback(request *GetFaceSearchUserRequest, callback func(response *GetFaceSearchUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetFaceSearchUserResponse
		var err error
		defer close(result)
		response, err = client.GetFaceSearchUser(request)
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

// GetFaceSearchUserRequest is the request struct for api GetFaceSearchUser
type GetFaceSearchUserRequest struct {
	*requests.RpcRequest
	Project   string `position:"Query" name:"Project"`
	GroupName string `position:"Query" name:"GroupName"`
	User      string `position:"Query" name:"User"`
}

// GetFaceSearchUserResponse is the response struct for api GetFaceSearchUser
type GetFaceSearchUserResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	GroupName  string `json:"GroupName" xml:"GroupName"`
	Count      int    `json:"Count" xml:"Count"`
	Status     string `json:"Status" xml:"Status"`
	CreateTime string `json:"CreateTime" xml:"CreateTime"`
	ModifyTime string `json:"ModifyTime" xml:"ModifyTime"`
	GroupId    string `json:"GroupId" xml:"GroupId"`
	User       string `json:"User" xml:"User"`
}

// CreateGetFaceSearchUserRequest creates a request to invoke GetFaceSearchUser API
func CreateGetFaceSearchUserRequest() (request *GetFaceSearchUserRequest) {
	request = &GetFaceSearchUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("imm", "2017-09-06", "GetFaceSearchUser", "imm", "openAPI")
	return
}

// CreateGetFaceSearchUserResponse creates a response to parse from GetFaceSearchUser response
func CreateGetFaceSearchUserResponse() (response *GetFaceSearchUserResponse) {
	response = &GetFaceSearchUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
