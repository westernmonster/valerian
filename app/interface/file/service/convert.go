package service

import (
	"context"
	"fmt"
	"valerian/app/interface/file/model"
	"valerian/library/log"

	"github.com/aliyun/alibaba-cloud-sdk-go/services/imm"
)

func (p *Service) CreateOfficeConversionTask(c context.Context, arg *model.ArgConvertOfficeTask) (err error) {
	req := imm.CreateCreateOfficeConversionTaskRequest()
	req.Project = "stonote"
	req.SrcUri = arg.SrcUri
	req.TgtType = "pdf"
	req.TgtUri = arg.TgtUri
	req.SetScheme("https")

	var ret *imm.CreateOfficeConversionTaskResponse
	if ret, err = p.immClient.CreateOfficeConversionTask(req); err != nil {
		log.For(c).Error(fmt.Sprintf("service.CreateOfficeConversionTask() error(%+v)", err))
		return
	}

	log.For(c).Info(fmt.Sprintf("service.CreateOfficeConversionTask() resp(%+v)", ret))

	return
}
