package service

import (
	"context"
	"valerian/app/interface/common/model"
	"valerian/library/conf/env"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/sts"
)

func (p *Service) AssumeRole(c context.Context) (resp *model.STSResp, err error) {
	assumeRoleReq := sts.CreateAssumeRoleRequest()
	assumeRoleReq.RoleArn = p.c.Aliyun.RoleArn
	assumeRoleReq.RoleSessionName = p.c.Aliyun.RoleSessionName
	assumeRoleReq.SetScheme("https")

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

	if env.DeployEnv == env.DeployEnvProd {
		resp.BucketName = "stonote"
	} else {
		resp.BucketName = "flywiki"
	}

	return
}
