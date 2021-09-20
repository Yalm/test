package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type NavigationInfo struct {
	EnableForcedRedirect bool `bson:"enableForcedRedirect" json:"enableForcedRedirect"`
}

type Suffix struct {
	Option       string `bson:"option" json:"option"`
	CustomSuffix string `bson:"customSuffix" json:"customSuffix"`
}

type SocialMetaTagInfo struct {
	SocialTitle       string `bson:"socialTitle" json:"socialTitle"`
	SocialDescription string `bson:"socialDescription" json:"socialDescription"`
	SocialImageLink   string `bson:"socialImageLink" json:"socialImageLink"`
}

type AndroidInfo struct {
	AndroidPackageName           string `bson:"androidPackageName" json:"androidPackageName"`
	AndroidAppName               string `bson:"androidAppName" json:"androidAppName"`
	AndroidUrl                   string `bson:"androidUrl" json:"androidUrl"`
	AndroidFallbackLink          string `bson:"androidFallbackLink" json:"androidFallbackLink"`
	AndroidMinPackageVersionCode string `bson:"androidMinPackageVersionCode" json:"androidMinPackageVersionCode"`
}

type IosInfo struct {
	IosAppStoreId   string `bson:"iosAppStoreId" json:"iosAppStoreId"`
	IosCustomScheme string `bson:"iosCustomScheme" json:"iosCustomScheme"`
}

type GooglePlayAnalytics struct {
	UtmSource   string `bson:"utmSource" json:"utmSource"`
	UtmMedium   string `bson:"utmMedium" json:"utmMedium"`
	UtmCampaign string `bson:"utmCampaign" json:"utmCampaign"`
	UtmTerm     string `bson:"utmTerm" json:"utmTerm"`
	UtmContent  string `bson:"utmContent" json:"utmContent"`
}

type AnalyticsInfo struct {
	GooglePlayAnalytics GooglePlayAnalytics
}

type DynamicLinkInfo struct {
	DomainUriPrefix   string `bson:"domainUriPrefix" json:"domainUriPrefix"`
	Link              string `bson:"link" json:"link"`
	AndroidInfo       AndroidInfo
	IosInfo           IosInfo
	NavigationInfo    NavigationInfo
	Hostname          string
	SocialMetaTagInfo SocialMetaTagInfo
	AnalyticsInfo     AnalyticsInfo
}

type DynamicLink struct {
	ID              primitive.ObjectID `bson:"_id" json:"id"`
	Status          int                `bson:"status" json:"status"`
	DynamicLinkInfo DynamicLinkInfo
	Suffix          Suffix
}
