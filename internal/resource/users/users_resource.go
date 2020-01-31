package usersresource

type GetUserInfoResponse struct {
	Status   string
	Username string
	Email	 string
	Message	 string
}

func SetLoginResponse(status string,
	username string,
	email string,
	message string) (GetUserInfoResponse) {
		var response GetUserInfoResponse
		response.Status = status
		response.Username = username
		response.Email = email
		response.Message = message
		return response
}