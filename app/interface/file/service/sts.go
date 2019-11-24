package service

import (
	"context"
	"valerian/app/interface/file/model"
	"valerian/library/conf/env"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
)

const uatPolicy = `
{
    "Version": "1",
    "Statement": [
     {
           "Effect": "Allow",
           "Action": [
             "oss:*",
           ],
           "Resource": [
             "acs:oss:*:*:flywiki",
           ]
     }
    ]
}
`

const prodPolicy = `
{
    "Version": "1",
    "Statement": [
     {
           "Effect": "Allow",
           "Action": [
             "oss:*",
           ],
           "Resource": [
             "acs:oss:*:*:stonote",
           ]
     }
    ]
}
`

func (p *Service) AssumeRole(c context.Context) (resp *model.STSResp, err error) {
	assumeRoleReq := sts.CreateAssumeRoleRequest()
	assumeRoleReq.RoleArn = p.c.Aliyun.RoleArn
	assumeRoleReq.RoleSessionName = p.c.Aliyun.RoleSessionName

	if env.DeployEnv == env.DeployEnvProd {
		assumeRoleReq.Policy = prodPolicy
	} else {
		assumeRoleReq.Policy = uatPolicy
	}

	p.stsClient.SetHTTPSInsecure(true)
	var ret *sts.AssumeRoleResponse
	if ret, err = p.stsClient.AssumeRole(assumeRoleReq); err != nil {
		return
	}
	resp = &model.STSResp{
		AccessKeySecret: ret.Credentials.AccessKeySecret,
		Expiration:      ret.Credentials.Expiration,
		AccessKeyId:     ret.Credentials.AccessKeyId,
		SecurityToken:   ret.Credentials.SecurityToken,
	}

	return
}
