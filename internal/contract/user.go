package contract

type (
	UserRequest struct {
		Id        uint   `json:"id"`
		FullName  string `json:"full_name"`
		Email     string `json:"email"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
	DeleteUserReq struct {
		Ids []uint `json:"ids"`
	}
	ChangeStatusRes struct {
		Id           uint `json:"id" valid:"required~please enter id"`
		StatusUserId uint `json:"status_id" valid:"required~please enter status_id"`
	}
	UserResponse struct {
		Id        uint   `json:"id"`
		FullName  string `json:"full_name"`
		Email     string `json:"email"`
		CreatedAt string `json:"created_at"`
		UpdatedAt string `json:"updated_at"`
	}
)

func (u UserRequest) GetPage() uint {
	//TODO implement me
	panic("implement me")
}

func (u UserRequest) GetLimit() uint {
	//TODO implement me
	panic("implement me")
}

type ListUserRequest struct {
	ListRequest
}
