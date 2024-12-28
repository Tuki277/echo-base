package contract

type (
	LoginRequest struct {
		Email    string `json:"email" valid:"email~Email is not valid"`
		Password string `json:"password" valid:"stringlength(6|50)~Password is at least 6 characters"`
	}
	RegisterRequest struct {
		FullName string `json:"full_name" valid:"required~please enter full_name"`
		Email    string `json:"email" valid:"email~Email is not valid"`
		Password string `json:"password" valid:"stringlength(6|50)~Password is at least 6 characters"`
	}
	AuthResponse struct {
		IsSuccessfull bool   `json:"is_successfull"`
		Token         string `json:"token"`
	}
)
