package mycode

// PageModeInfo 页面模块信息
type PageModeInfo struct {
	PageID   int64  `json:"page_id"`
	Name     string `json:"name"`
	Version  int64  `json:"version"`
	LastTime int64  `json:"last_time"`
	Models   struct {
		ID   int64 `json:"id"`
		Type int64 `json:"type"`
	} `json:"models"`
}

// ModelDisplayInfo 模块详情信息
type ModelDisplayInfo struct {
	ID    int64  `json:"id"`
	Title string `json:"title"`
	/*
		不同模块类型数据不同
	*/
	GameIDList []int64 `json:"game_id_list"`
}

// ModelGameInfo 游戏详情信息
type ModelGameInfo struct {
	ID           int64          `json:"id"`
	Title        string         `json:"title"`
	ActUrl       string         `json:"act_url"`
	RankTitleV2  string         `json:"rank_title_v2"`
	DownloadDesc string         `json:"download_desc"`
	Summary      string         `json:"summary"`
	Score        float64        `json:"score"`
	ScoreV2      float64        `json:"score_v2"`
	Animation    int64          `json:"animation"`
	BCnt         int64          `json:"b_cnt"`
	PublishType  int64          `json:"publish_type"`
	IsSubscribe  bool           `json:"is_subscribe"`
	ClickPlayTag bool           `json:"click_play_tag"`
	Banners      []*GameBanner  `json:"banners"`
	DownloadInfo *DownloadInfo  `json:"download_info"`
	ServerInfo   *ServerInfo    `json:"server_info"`
	GameTags     []*GameTagInfo `json:"game_tags"`
}

type GameBanner struct {
	URL     string `json:"url"`
	URLType int64  `json:"url_type"`
}

type DownloadInfo struct {
	DownloadAble    int64          `json:"download_able"`
	Icon            string         `json:"icon"`
	GameCloudName   string         `json:"game_cloud_name"`
	GameCloudIcon   string         `json:"game_cloud_icon"`
	GameType        int64          `json:"game_type"`
	PreDownloadAble bool           `json:"pre_download_able"`
	ViewCommunity   int64          `json:"view_community"`
	CircleID        int64          `json:"circle_id"`
	APK             *APKInfo       `json:"apk"`
	Subscribe       *SubscribeInfo `json:"subscribe"`
}

type APKInfo struct {
	GameID      int64  `json:"game_id"`
	PackageName string `json:"package_name"`
	GameApk     string `json:"game_apk"`
	GameApkSsi  string `json:"game_apk_ssi"`
	ApkHash     string `json:"apk_hash"`
	ApkSize     int64  `json:"apk_size"`
	VersionCode int64  `json:"version_code"`
}

type SubscribeInfo struct {
	Type            int64  `json:"type"`
	T               int64  `json:"t"`
	S               int64  `json:"s"`
	Text            string `json:"text"`
	Auto            int64  `json:"auto"`
	WelfareTitle1   string `json:"welfare_title_1"`
	WelfareContent1 string `json:"welfare_content_1"`
	WelfareTitle2   string `json:"welfare_title_2"`
	WelfareContent2 string `json:"welfare_content_2"`
	GiftIcon        string `json:"gift_icon"`
	ActivityBanner  string `json:"activity_banner"`
	ActivityUrl     string `json:"activity_url"`
	ActivityH5Url   string `json:"activity_h5_url"`
	ActID           string `json:"act_id"`
	Count           int64  `json:"count"`
	CancelSubscribe int64  `json:"cancel_subscribe"`
}

type ServerInfo struct {
	ContentID int64  `json:"content_id"`
	TraceID   string `json:"trace_id"`
	Channel   string `json:"channel"`
}

type GameTagInfo struct {
	TagID   int64  `json:"tag_id"`
	TagType int64  `json:"tag_type"`
	V       int64  `json:"v"`
	Name    string `json:"name"`
	ActUrl  string `json:"act_url"`
}
