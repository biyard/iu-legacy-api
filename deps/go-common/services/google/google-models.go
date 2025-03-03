package google

type GetTokenInfoRequest struct {
	IDToken string `form:"id_token"`
}

type GetTokenInfoResponse struct {
	Issuer        string `json:"iss"`
	Subject       string `json:"sub"`
	Audience      string `json:"aud"`
	IssuedAt      string `json:"iat"`
	Expiration    string `json:"exp"`
	Email         string `json:"email"`
	EmailVerified string `json:"email_verified"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Locale        string `json:"locale"`
}
