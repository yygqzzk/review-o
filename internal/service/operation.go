package service

import (
	"context"

	pb "github.com/yygqzhu/review-o/api/operation/v1"
	"github.com/yygqzhu/review-o/internal/biz"
)

type OperationService struct {
	pb.UnimplementedOperationServer
	uc *biz.OperationUsecase
}

func NewOperationService(uc *biz.OperationUsecase) *OperationService {
	return &OperationService{uc: uc}
}

func (s *OperationService) AuditReview(ctx context.Context, req *pb.AuditReviewReq) (*pb.AuditReviewRsp, error) {
	err := s.uc.AuditReview(ctx, &biz.AuditReviewParam{
		ReviewID:  req.GetReviewID(),
		Status:    int(req.GetStatus()),
		OpReason:  req.GetOpReason(),
		OpRemarks: req.GetOpRemarks(),
		OpUser:    req.GetOpUser(),
	})
	return &pb.AuditReviewRsp{}, err
}
func (s *OperationService) AuditAppeal(ctx context.Context, req *pb.AuditAppealReq) (*pb.AuditAppealRsp, error) {
	err := s.uc.AuditAppeal(ctx, &biz.AuditAppealParam{
		AppealID:  req.GetAppealID(),
		ReviewID:  req.GetReviewID(),
		Status:    int(req.GetStatus()),
		OpReason:  req.GetOpReason(),
		OpRemarks: req.GetOpRemarks(),
		OpUser:    req.GetOpUser(),
	})
	return &pb.AuditAppealRsp{}, err
}
