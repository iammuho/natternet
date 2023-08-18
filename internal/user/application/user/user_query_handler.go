package user

import (
	"github.com/iammuho/natternet/cmd/app/context"
	"github.com/iammuho/natternet/internal/user/application/user/dto"
	"github.com/iammuho/natternet/internal/user/domain/services"
	"github.com/iammuho/natternet/internal/user/domain/values"
	"github.com/iammuho/natternet/pkg/errorhandler"
)

type UserQueryHandler struct {
	ctx                     context.AppContext
	userQueryDomainServices services.UserQueryDomainServices
}

func NewUserQueryHandler(ctx context.AppContext, userQueryDomainServices services.UserQueryDomainServices) *UserQueryHandler {
	return &UserQueryHandler{
		ctx:                     ctx,
		userQueryDomainServices: userQueryDomainServices,
	}
}

func (s *UserQueryHandler) QueryUserByID(req *dto.QueryUserByIDReqDTO) (*values.UserValue, *errorhandler.Response) {
	return s.userQueryDomainServices.FindByID(req.UserID)
}
