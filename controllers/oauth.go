package controllers

import (
	"github.com/gogather/oauth"
)

type OAuthController struct {
	BaseController
}

func (this *OAuthController) Post() {
	oauthGithub := &oauth.GithubOAuth{}
	oauthGithub.NewGithubOAuth(this.GetString("access_token"))
	json, err := oauthGithub.GetData()
	if err != nil {
		this.Ctx.WriteString("Response Error!")
	}
	data := json.(map[string]interface{})

	this.Data["login"] = data["login"].(string)
	this.Data["avatar_url"] = data["avatar_url"].(string)
	this.Data["name"] = data["name"].(string)
}

func (this *OAuthController) Get() {
	this.Ctx.WriteString("content")
}