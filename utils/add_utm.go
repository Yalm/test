package utils

import (
	"net/url"

	"github.com/yalm/cloud-messaging/models"
)

func AddUtm(dynamicLink *models.DynamicLink) (string, error) {
	url, err := url.Parse(dynamicLink.DynamicLinkInfo.Link)
	if err != nil {
		return "", err
	}

	query := url.Query()
	if len(dynamicLink.DynamicLinkInfo.AnalyticsInfo.GooglePlayAnalytics.UtmCampaign) > 0 {
		query.Set("utm_campaign", dynamicLink.DynamicLinkInfo.AnalyticsInfo.GooglePlayAnalytics.UtmCampaign)
	}

	if len(dynamicLink.DynamicLinkInfo.AnalyticsInfo.GooglePlayAnalytics.UtmMedium) > 0 {
		query.Set("utm_medium", dynamicLink.DynamicLinkInfo.AnalyticsInfo.GooglePlayAnalytics.UtmMedium)
	}

	if len(dynamicLink.DynamicLinkInfo.AnalyticsInfo.GooglePlayAnalytics.UtmSource) > 0 {
		query.Set("utm_source", dynamicLink.DynamicLinkInfo.AnalyticsInfo.GooglePlayAnalytics.UtmSource)
	}
	url.RawQuery = query.Encode()

	return url.String(), nil
}
