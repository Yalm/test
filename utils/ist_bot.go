package utils

import (
	"regexp"
	"strings"
)

var listBots = []string{
	" daum[/|\\s]",
	" deusu/",
	"(?:^|\\s)site",
	"@[a-z]",
	"\\(at\\)[a-z]",
	"\\(github\\.com/",
	"\\[at\\][a-z]",
	"^12345",
	"^<",
	"^ace explorer",
	"^acoon",
	"^active",
	"^ad muncher",
	"^ahc/",
	"^amiga$",
	"^anglesharp/",
	"^anonymous",
	"^apache",
	"^apple-pubsub/",
	"^applicationhealthservice",
	"^arachni/",
	"^astute srm",
	"^avsdevicesdk/",
	"^axios/",
	"^bidtellect/",
	"^biglotron",
	"^bitdiscovery",
	"^blackboard safeassign",
	"^blocknote.net",
	"^blogtrottr/",
	"^braze sender",
	"^browsershots",
	"^cakephp",
	"^camo asset proxy",
	"^captivenetworksupport",
	"^castro",
	"^clamav[\\s/]",
	"^cobweb/",
	"^coccoc",
	"^coldfusion",
	"^collectd/",
	"^custom$",
	"^dap ",
	"^datadog agent/",
	"^davclnt",
	"^ddg[_-]android",
	"^deluge",
	"^discourse",
	"^dispatch/\\d",
	"^disqus/",
	"^downcast/",
	"^duckduckgo",
	"^enigma browser",
	"^evernote clip resolver",
	"^facebook",
	"^faraday",
	"^fdm[/\\s]\\d",
	"^flashget",
	"^friendica",
	"^getright/",
	"^gigablastopensource",
	"^gobuster/",
	"^googal",
	"^goose",
	"^gozilla/",
	"^greenbrowser",
	"^hatena",
	"^hexometer",
	"^hobbit",
	"^hotzonu",
	"^hwcdn/",
	"^infox-wisg",
	"^ingrid/\\d",
	"^integrity/",
	"^invision",
	"^jeode/",
	"^jetbrains",
	"^jetty/",
	"^jigsaw",
	"^libwww",
	"^linkdex",
	"^lwp-",
	"^lwp::simple",
	"^mailchimp\\.com$",
	"^metauri",
	"^microsoft bits",
	"^microsoft data",
	"^microsoft office existence",
	"^microsoft office protocol discovery",
	"^microsoft windows network diagnostics",
	"^microsoft-cryptoapi",
	"^microsoft-webdav-miniredir",
	"^monit",
	"^movabletype",
	"^mozilla/\\d\\.\\d \\(compatible;?\\)$",
	"^my browser$",
	"^navermailapp",
	"^netsurf",
	"^nibbler",
	"^ning",
	"^node-superagent",
	"^notetextview",
	"^nuzzel",
	"^octopus",
	"^offline explorer",
	"^ossproxy",
	"^pagething",
	"^panscient",
	"^pearltrees",
	"^perimeterx",
	"^photon/",
	"^php",
	"^postman",
	"^postrank",
	"^prometheus/",
	"^python",
	"^ramblermail",
	"^read",
	"^reed",
	"^restsharp/",
	"^robozilla/",
	"^ruby$",
	"^scrapy",
	"^seo",
	"^shareaza",
	"^shockwave flash",
	"^shortlinktranslate",
	"^sistrix",
	"^sixy.ch/",
	"^smallproxy",
	"^snap$",
	"^snapchat",
	"^space bison",
	"^spotify/",
	"^sprinklr",
	"^svn",
	"^swcd ",
	"^t-online browser",
	"^taringa",
	"^test certificate info",
	"^the knowledge ai",
	"^thinklab",
	"^thumb",
	"^traackr.com",
	"^transmission",
	"^tumblr/",
	"^ucmore",
	"^upflow/",
	"^user_agent",
	"^vbulletin",
	"^venus/fedoraplanet",
	"^viber$",
	"^w3c",
	"^webbandit/",
	"^webcollage/",
	"^webcopier",
	"^wget",
	"^whatsapp",
	"^whatweb",
	"^www-mechanize",
	"^xenu link sleuth",
	"^xymon",
	"^yahoo",
	"^yandex",
	"^zabbix",
	"^zdm/\\d",
	"^zeushdthree",
	"^zmeu$",
	"adbeat\\.com",
	"appinsights",
	"archive",
	"ask jeeves/teoma",
	"bit.ly/",
	"bluecoat drtr",
	"bot",
	"browsex",
	"burpcollaborator",
	"capture",
	"catch",
	"check",
	"chrome-lighthouse",
	"chromeframe",
	"client",
	"cloud",
	"crawl",
	"cron",
	"daemon",
	"dareboost",
	"datanyze",
	"dataprovider",
	"dejaclick",
	"dmbrowser",
	"download",
	"email",
	"evc-batch/",
	"feed",
	"fetch",
	"finder",
	"firephp",
	"freesafeip",
	"ghost",
	"gomezagent",
	"google",
	"headlesschrome/",
	"http",
	"httrack",
	"hubspot marketing grader",
	"hydra",
	"ibisbrowser",
	"images",
	"index",
	"ips-agent",
	"java",
	"jorgee",
	"library",
	"mail\\.ru/",
	"manager",
	"monitor",
	"neustar wpm",
	"news",
	"nutch",
	"offbyone",
	"optimize",
	"pagespeed",
	"parse",
	"perl",
	"phantom",
	"pingdom",
	"powermarks",
	"preview",
	"probe",
	"ptst[/ ]\\d",
	"reader",
	"rigor",
	"rss",
	"scan",
	"scrape",
	"search",
	"server",
	"sogou",
	"sparkler/",
	"spider",
	"statuscake",
	"stumbleupon\\.com",
	"supercleaner",
	"synapse",
	"synthetic",
	"toolbar",
	"torrent",
	"tracemyfile",
	"transcoder",
	"trendsmapresolver",
	"twingly recon",
	"url",
	"valid",
	"virtuoso",
	"wappalyzer",
	"webglance",
	"webkit2png",
	"websitemetadataretriever",
	"whatcms/",
	"wordpress",
	"zgrab"}

func IsBot(userAgent string) bool {
	matched, _ := regexp.MatchString("(?i)"+strings.Join(listBots, "|"), userAgent)
	return matched
}