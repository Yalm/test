package utils

import (
	"net/url"

	"github.com/yalm/cloud-messaging/models"
)

func GenerateSocialTag(dynamicLink *models.DynamicLink) string {
	values := url.Values{}
	values.Add("socialTitle", dynamicLink.DynamicLinkInfo.SocialMetaTagInfo.SocialTitle)
	values.Add("socialImageUrl", dynamicLink.DynamicLinkInfo.SocialMetaTagInfo.SocialImageLink)
	values.Add("socialDescription", dynamicLink.DynamicLinkInfo.SocialMetaTagInfo.SocialDescription)
	values.Add("androidPackageName", dynamicLink.DynamicLinkInfo.AndroidInfo.AndroidPackageName)
	values.Add("androidUrl", dynamicLink.DynamicLinkInfo.AndroidInfo.AndroidUrl)
	values.Add("link", dynamicLink.DynamicLinkInfo.Link)
	values.Add("iosUrl", dynamicLink.DynamicLinkInfo.IosInfo.IosAppStoreId)
	query := values.Encode()
	return renderParam + "?" + query
}
