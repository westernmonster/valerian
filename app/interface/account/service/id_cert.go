package service

import (
	"context"
	"valerian/app/interface/account/model"
	account "valerian/app/service/account/api"
	"valerian/library/cloudauth"
	"valerian/library/ecode"
	"valerian/library/net/metadata"
)

// RequestIDCertification 发起认证请求，获取 Token
func (p *Service) RequestIDCertification(c context.Context) (token cloudauth.VerifyTokenData, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var t *account.RequestIDCertResp
	if t, err = p.d.RequestIDCert(c, aid); err != nil {
		return
	}

	token = cloudauth.VerifyTokenData{
		CloudauthPageUrl: t.CloudauthPageUrl,
		STSToken: cloudauth.STSToken{
			AccessKeyId:     t.STSToken.AccessKeyId,
			AccessKeySecret: t.STSToken.AccessKeySecret,
			Expiration:      t.STSToken.Expiration,
			EndPoint:        t.STSToken.EndPoint,
			BucketName:      t.STSToken.BucketName,
			Path:            t.STSToken.Path,
			Token:           t.STSToken.Token,
		},
		VerifyToken: cloudauth.VerifyToken{
			Token:           t.VerifyToken.Token,
			DurationSeconds: int(t.VerifyToken.DurationSeconds),
		},
	}

	return

}

// 查询认证状态
func (p *Service) GetIDCertificationStatus(c context.Context) (status int, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var resp *account.IDCertStatus
	if resp, err = p.d.RefreshIDCertStatus(c, aid); err != nil {
		return
	}

	status = int(resp.Status)

	return
}

func (p *Service) GetIDCert(c context.Context) (resp *model.IDCertResp, err error) {
	aid, ok := metadata.Value(c, metadata.Aid).(int64)
	if !ok {
		err = ecode.AcquireAccountIDFailed
		return
	}

	var v *account.IDCertInfo
	if v, err = p.d.GetIDCert(c, aid); err != nil {
		return
	}

	resp = &model.IDCertResp{
		Status:               (v.Status),
		AuditConclusions:     v.AuditConclusions,
		Name:                 v.Name,
		IdentificationNumber: v.IdentificationNumber,
		IDCardType:           v.IDCardType,
		IDCardStartDate:      v.IDCardStartDate,
		IDCardExpiry:         v.IDCardExpiry,
		Sex:                  v.Sex,
		IDCardFrontPic:       v.IDCardFrontPic,
		IDCardBackPic:        v.IDCardBackPic,
		CreatedAt:            v.CreatedAt,
		UpdatedAt:            v.UpdatedAt,
	}

	return

}
