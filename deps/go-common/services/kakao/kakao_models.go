package kakao

type KakaoOAuthTokenRequest struct {
	GrantType   string `form:"grant_type"`
	ClientId    string `form:"client_id"`
	RedirectUri string `form:"redirect_uri"`
	Code        string `form:"code"`
}

type KakaoOAuthTokenResponse struct {
	TokenType             string `json:"token_type"`
	AccessToken           string `json:"access_token"`
	IDToken               string `json:"id_token"`
	ExpiresIn             int64  `json:"expires_in"`
	RefreshToken          string `json:"refresh_token"`
	RefreshTokenExpiresIn int64  `json:"refresh_token_expires_in"`
	Scope                 string `json:"scope"`
}

type KakaoOpenIDTokenInfoRequest struct {
	IDToken string `form:"id_token"`
}

type KakaoOpenIDTokenInfoResponse struct {
	Issuer    string `json:"iss"`
	Subject   string `json:"sub"`
	Audience  string `json:"aud"`
	ExpiresIn int64  `json:"exp"`
	IssuedAt  int64  `json:"iat"`
}
