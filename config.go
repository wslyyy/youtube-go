package youtube_go

const (
	REFERER_YOUTUBE           = "https://www.youtube.com/"
	REFERER_YOUTUBE_MOBILE    = "https://m.youtube.com/"
	REFERER_YOUTUBE_MUSIC     = "https://music.youtube.com/"
	REFERER_YOUTUBE_KIDS      = "https://www.youtubekids.com/"
	REFERER_YOUTUBE_STUDIO    = "https://studio.youtube.com/"
	REFERER_YOUTUBE_ANALYTICS = "https://analytics.youtube.com/"
	REFERER_GOOGLE_ASSISTANT  = "https://assistant.google.com/"

	USER_AGENT_WEB              = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.157 Safari/537.36"
	USER_AGENT_ANDROID          = "Mozilla/5.0 (Linux; Android 6.0; Nexus 5 Build/MRA58N) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/65.0.3325.181 Mobile Safari/537.36"
	USER_AGENT_IOS              = "Mozilla/5.0 (iPhone; CPU iPhone OS 15_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/98.2  Mobile/15E148 Safari/605.1.15"
	USER_AGENT_TV_HTML5         = "Mozilla/5.0 (PlayStation 4 5.55) AppleWebKit/601.2 (KHTML, like Gecko)"
	USER_AGENT_TV_APPLE         = "AppleCoreMedia/1.0.0.12B466 (Apple TV; U; CPU OS 8_1_3 like Mac OS X; en_us)"
	USER_AGENT_TV_ANDROID       = "Mozilla/5.0 (Linux; Android 5.1.1; AFTT Build/LVY48F; wv) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/49.0.2623.10"
	USER_AGENT_XBOX_ONE         = "Mozilla/5.0 (Windows NT 10.0; Win64; x64; Xbox; Xbox One) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/46.0.2486.0 Safari/537.36 Edge/13.10553"
	USER_AGENT_GOOGLE_ASSISTANT = "Mozilla/5.0 (Linux; Android 11; Pixel 2; DuplexWeb-Google/1.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.193 Mobile Safari/537.36"
)

var config = Config{
	Host:    "youtubei.googleapis.com",
	BaseURL: "https://youtubei.googleapis.com/youtubei/v1/",
	Clients: []ClientContext{
		{ClientID: 1, ClientName: "WEB", ClientVersion: "2.20230728.00.00", UserAgent: USER_AGENT_WEB, Referer: REFERER_YOUTUBE, APIKey: "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"},
		{ClientID: 2, ClientName: "MWEB", ClientVersion: "2.20211214.00.00", UserAgent: USER_AGENT_ANDROID, Referer: REFERER_YOUTUBE_MOBILE, APIKey: "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"},
		{ClientID: 3, ClientName: "ANDROID", ClientVersion: "17.13.3", UserAgent: USER_AGENT_ANDROID, APIKey: "AIzaSyA8eiZmM1FaDVjRy-df2KTyQ_vz_yYM39w"},
		{ClientID: 5, ClientName: "IOS", ClientVersion: "17.14.2", UserAgent: USER_AGENT_IOS, APIKey: "AIzaSyB-63vPrdThhKuerbB2N_l7Kwwcxj6yUAc"},
		{ClientID: 7, ClientName: "TVHTML5", ClientVersion: "7.20210224.00.00", UserAgent: USER_AGENT_TV_HTML5, APIKey: "AIzaSyDCU8hByM-4DrUqRUYnGn-3llEO78bcxq8"},
		{ClientID: 8, ClientName: "TVLITE", ClientVersion: "2", UserAgent: USER_AGENT_TV_HTML5},
		{ClientID: 10, ClientName: "TVANDROID", ClientVersion: "1.0", UserAgent: USER_AGENT_TV_ANDROID},
		{ClientID: 13, ClientName: "XBOXONEGUIDE", ClientVersion: "1.0", UserAgent: USER_AGENT_XBOX_ONE},
		{ClientID: 14, ClientName: "ANDROID_CREATOR", ClientVersion: "21.06.103", UserAgent: USER_AGENT_ANDROID, APIKey: "AIzaSyD_qjV8zaaUMehtLkrKFgVeSX_Iqbtyws8"},
		{ClientID: 15, ClientName: "IOS_CREATOR", ClientVersion: "20.47.100", UserAgent: USER_AGENT_IOS, APIKey: "AIzaSyAPyF5GfQI-kOa6nZwO8EsNrGdEx9bioNs"},
		{ClientID: 16, ClientName: "TVAPPLE", ClientVersion: "1.0", UserAgent: USER_AGENT_TV_APPLE},
		{ClientID: 18, ClientName: "ANDROID_KIDS", ClientVersion: "7.12.3", UserAgent: USER_AGENT_ANDROID, APIKey: "AIzaSyAxxQKWYcEX8jHlflLt2Qcbb-rlolzBhhk"},
		{ClientID: 19, ClientName: "IOS_KIDS", ClientVersion: "5.42.2", UserAgent: USER_AGENT_IOS, APIKey: "AIzaSyA6_JWXwHaVBQnoutCv1-GvV97-rJ949Bc"},
		{ClientID: 21, ClientName: "ANDROID_MUSIC", ClientVersion: "5.01", UserAgent: USER_AGENT_ANDROID, APIKey: "AIzaSyAOghZGza2MQSZkY_zfZ370N-PUdXEo8AI"},
		{ClientID: 23, ClientName: "ANDROID_TV", ClientVersion: "2.16.032", UserAgent: USER_AGENT_TV_ANDROID},
		{ClientID: 26, ClientName: "IOS_MUSIC", ClientVersion: "4.16.1", UserAgent: USER_AGENT_IOS, APIKey: "AIzaSyBAETezhkwP0ZWA02RsqT1zu78Fpt0bC_s"},
		{ClientID: 27, ClientName: "MWEB_TIER_2", ClientVersion: "9.20220325", UserAgent: USER_AGENT_ANDROID, Referer: REFERER_YOUTUBE_MOBILE},
		{ClientID: 28, ClientName: "ANDROID_VR", ClientVersion: "1.28.63", UserAgent: USER_AGENT_ANDROID},
		{ClientID: 29, ClientName: "ANDROID_UNPLUGGED", ClientVersion: "6.13", UserAgent: USER_AGENT_ANDROID},
		{ClientID: 30, ClientName: "ANDROID_TESTSUITE", ClientVersion: "1.9", UserAgent: USER_AGENT_ANDROID},
		{ClientID: 31, ClientName: "WEB_MUSIC_ANALYTICS", ClientVersion: "0.2", UserAgent: USER_AGENT_WEB, Referer: REFERER_YOUTUBE_ANALYTICS},
		{ClientID: 33, ClientName: "IOS_UNPLUGGED", ClientVersion: "6.13", UserAgent: USER_AGENT_IOS},
		{ClientID: 38, ClientName: "ANDROID_LITE", ClientVersion: "3.26.1", UserAgent: USER_AGENT_ANDROID},
		{ClientID: 39, ClientName: "IOS_EMBEDDED_PLAYER", ClientVersion: "2.3", UserAgent: USER_AGENT_IOS},
		{ClientID: 41, ClientName: "WEB_UNPLUGGED", ClientVersion: "1.20220403", UserAgent: USER_AGENT_WEB, Referer: REFERER_YOUTUBE},
		{ClientID: 42, ClientName: "WEB_EXPERIMENTS", ClientVersion: "1", UserAgent: USER_AGENT_WEB, Referer: REFERER_YOUTUBE},
		{ClientID: 43, ClientName: "TVHTML5_CAST", ClientVersion: "1.1", UserAgent: USER_AGENT_TV_HTML5},
		{ClientID: 55, ClientName: "ANDROID_EMBEDDED_PLAYER", ClientVersion: "17.13.3", UserAgent: USER_AGENT_ANDROID, APIKey: "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"},
		{ClientID: 56, ClientName: "WEB_EMBEDDED_PLAYER", ClientVersion: "1.20220413.01.00", UserAgent: USER_AGENT_WEB, Referer: REFERER_YOUTUBE, APIKey: "AIzaSyAO_FJ2SlqU8Q4STEHLGCilw_Y9_11qcW8"},
		{ClientID: 57, ClientName: "TVHTML5_AUDIO", ClientVersion: "2.0", UserAgent: USER_AGENT_TV_HTML5},
		{ClientID: 58, ClientName: "TV_UNPLUGGED_CAST", ClientVersion: "0.1", UserAgent: USER_AGENT_TV_HTML5},
		{ClientID: 59, ClientName: "TVHTML5_KIDS", ClientVersion: "3.20220325", UserAgent: USER_AGENT_TV_HTML5},
		{ClientID: 60, ClientName: "WEB_HEROES", ClientVersion: "0.1", UserAgent: USER_AGENT_WEB, Referer: REFERER_YOUTUBE},
		{ClientID: 61, ClientName: "WEB_MUSIC", ClientVersion: "1.0", UserAgent: USER_AGENT_WEB, Referer: REFERER_YOUTUBE_MUSIC},
		{ClientID: 62, ClientName: "WEB_CREATOR", ClientVersion: "1.20210223.01.00", UserAgent: USER_AGENT_WEB, Referer: REFERER_YOUTUBE_STUDIO, APIKey: "AIzaSyBUPetSUmoZL-OhlxA7wSac5XinrygCqMo"},
		{ClientID: 63, ClientName: "TV_UNPLUGGED_ANDROID", ClientVersion: "1.22.062.06.90", UserAgent: USER_AGENT_TV_ANDROID},
		{ClientID: 64, ClientName: "IOS_LIVE_CREATION_EXTENSION", ClientVersion: "17.13.3", UserAgent: USER_AGENT_IOS},
		{ClientID: 65, ClientName: "TVHTML5_UNPLUGGED", ClientVersion: "6.13", UserAgent: USER_AGENT_TV_HTML5},
		{ClientID: 66, ClientName: "IOS_MESSAGES_EXTENSION", ClientVersion: "16.20", UserAgent: USER_AGENT_IOS, APIKey: "AIzaSyDCU8hByM-4DrUqRUYnGn-3llEO78bcxq8"},
		{ClientID: 67, ClientName: "WEB_REMIX", ClientVersion: "1.20230724.00.00", UserAgent: USER_AGENT_WEB, Referer: REFERER_YOUTUBE_MUSIC, APIKey: "AIzaSyC9XL3ZjWddXya6X74dJoCTL-WEYFDNX30"},
		{ClientID: 68, ClientName: "IOS_UPTIME", ClientVersion: "1.0", UserAgent: USER_AGENT_IOS},
		{ClientID: 69, ClientName: "WEB_UNPLUGGED_ONBOARDING", ClientVersion: "0.1", UserAgent: USER_AGENT_WEB},
		{ClientID: 70, ClientName: "WEB_UNPLUGGED_OPS", ClientVersion: "0.1", UserAgent: USER_AGENT_WEB},
		{ClientID: 71, ClientName: "WEB_UNPLUGGED_PUBLIC", ClientVersion: "0.1", UserAgent: USER_AGENT_WEB},
		{ClientID: 72, ClientName: "TVHTML5_VR", ClientVersion: "0.1", UserAgent: USER_AGENT_TV_HTML5},
		{ClientID: 74, ClientName: "ANDROID_TV_KIDS", ClientVersion: "1.16.80", UserAgent: USER_AGENT_TV_ANDROID},
		{ClientID: 75, ClientName: "TVHTML5_SIMPLY", ClientVersion: "1.0", UserAgent: USER_AGENT_TV_HTML5},
		{ClientID: 76, ClientName: "WEB_KIDS", ClientVersion: "2.20220414.00.00", Referer: REFERER_YOUTUBE_KIDS, UserAgent: USER_AGENT_WEB, APIKey: "AIzaSyBbZV_fZ3an51sF-mvs5w37OqqbsTOzwtU"},
		{ClientID: 77, ClientName: "MUSIC_INTEGRATIONS", ClientVersion: "0.1"},
		{ClientID: 80, ClientName: "TVHTML5_YONGLE", ClientVersion: "0.1", UserAgent: USER_AGENT_TV_HTML5},
		{ClientID: 84, ClientName: "GOOGLE_ASSISTANT", ClientVersion: "0.1", UserAgent: USER_AGENT_GOOGLE_ASSISTANT},
		{ClientID: 85, ClientName: "TVHTML5_SIMPLY_EMBEDDED_PLAYER", ClientVersion: "2.0", UserAgent: USER_AGENT_TV_HTML5},
		{ClientID: 87, ClientName: "WEB_INTERNAL_ANALYTICS", ClientVersion: "0.1", UserAgent: USER_AGENT_WEB, Referer: REFERER_YOUTUBE_ANALYTICS},
		{ClientID: 88, ClientName: "WEB_PARENT_TOOLS", ClientVersion: "1.20220403", UserAgent: USER_AGENT_WEB},
		{ClientID: 89, ClientName: "GOOGLE_MEDIA_ACTIONS", ClientVersion: "0.1", UserAgent: USER_AGENT_GOOGLE_ASSISTANT},
		{ClientID: 90, ClientName: "WEB_PHONE_VERIFICATION", ClientVersion: "1.0.0", UserAgent: USER_AGENT_WEB},
		{ClientID: 92, ClientName: "IOS_PRODUCER", ClientVersion: "0.1", UserAgent: USER_AGENT_IOS},
		{ClientID: 93, ClientName: "TVHTML5_FOR_KIDS", ClientVersion: "7.20220325", UserAgent: USER_AGENT_TV_HTML5},
	},
}
