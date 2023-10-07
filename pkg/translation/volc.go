// 火山翻译接口
package translation

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"time"

	"github.com/volcengine/volc-sdk-golang/base"
)

type (
	// 火山翻译返回结果
	VolcTranslate struct {
		ResponseMetaData struct {
			RequestId, Action, Version, Service, Region string
		}
		TranslationList []TranslationList
	}
	// 翻译文本列表
	TranslationList struct {
		Translation, DetectedSourceLanguage string
		Extra                               any
	}
)

var client *base.Client

// 初始化
func InitVolc(accessKey, secretKey string) {
	client = base.NewClient(&base.ServiceInfo{
		Timeout: 5 * time.Second,
		Host:    "translate.volcengineapi.com",
		Header: http.Header{
			"Accept": []string{"application/json"},
		},
		Credentials: base.Credentials{Region: base.RegionCnNorth1, Service: "translate"},
	}, map[string]*base.ApiInfo{
		"TranslateText": {
			Method: http.MethodPost,
			Path:   "/",
			Query: url.Values{
				"Action":  []string{"TranslateText"},
				"Version": []string{"2020-06-01"},
			},
		},
	})
	client.SetAccessKey(accessKey)
	client.SetSecretKey(secretKey)
}

// 翻译
//
//	Text => 文本
//	source => 原语种，不传为自动检测语种
//	target => 目标语种
//
// 语种代号 语种中文名 语种英文名: zh 中文(简体) Chinese (simplified), zh-Hant 中文(繁体) Chinese (traditional), zh-Hant-hk 中文(香港繁体) Chinese (Hongkong traditional), zh-Hant-tw 中文(台湾繁体) Chinese (Taiwan traditional), tn 札那语 Tswana, vi 越南语 Vietnamese, iu 伊努克提图特语 Inuktitut, it 意大利语 Italian, id 印尼语 Indonesian, hi 印地语 Hindi, en 英语 English, ho 希里莫图语 Hiri Motu, he 希伯来语 Hebrew, es 西班牙语 Spanish, el 现代希腊语 Modern Greek, uk 乌克兰语 Ukrainian, ur 乌尔都语 Urdu, tk 土库曼语 Turkmen, tr 土耳其语 Turkish, ti 提格里尼亚语 Tigrinya, ty 塔希提语 Tahitian, tl 他加禄语 Tagalog, to 汤加语 Tongan, th 泰语 Thai, ta 泰米尔语 Tamil, te 泰卢固语 Telugu, sl 斯洛文尼亚语 Slovenian, sk 斯洛伐克语 Slovak, ss 史瓦帝语 Swati, eo 世界语 Esperanto, sm 萨摩亚语 Samoan, sg 桑戈语 Sango, st 塞索托语 Southern Sotho, sv 瑞典语 Swedish, ja 日语 Japanese, tw 契维语 Twi, qu 奇楚瓦语 Quechua, pt 葡萄牙语 Portuguese, pa 旁遮普语 Punjabi, no 挪威语 Norwegian, nb 挪威布克莫尔语 Norwegian Bokmål, nr 南恩德贝勒语 South Ndebele, my 缅甸语 Burmese, bn 孟加拉语 Bengali, mn 蒙古语 Mongolian, mh 马绍尔语 Marshallese, mk 马其顿语 Macedonian, ml 马拉亚拉姆语 Malayalam, mr 马拉提语 Marathi, ms 马来语 Malay, lu 卢巴卡丹加语 Luba-Katanga, ro 罗马尼亚语 Romanian, lt 立陶宛语 Lithuanian, lv 拉脱维亚语 Latvian, lo 老挝语 Lao lao, kj 宽亚玛语 Kwanyama, hr 克罗地亚语 Croatian, kn 坎纳达语 Kannada, ki 基库尤语 Kikuyu, cs 捷克语 Czech, ca 加泰隆语 Catalan, nl 荷兰语 Dutch, ko 韩语 Korean, ht 海地克里奥尔语 Haitian Creole, gu 古吉拉特语 Gujarati, ka 格鲁吉亚语 Georgian, kl 格陵兰语 Greenlandic, km 高棉语 Khmer, lg 干达语 Ganda, kg 刚果语 Kongo, fi 芬兰语 Finnish, fj 斐济语 Fijian, fr 法语 French, ru 俄语 Russian, ng 恩敦加语 Ndonga, de 德语 German, tt 鞑靼语 Tatar, da 丹麦语 Danish, ts 聪加语 Tsonga, cv 楚瓦什语 Chuvash, fa 波斯语 Persian, bs 波斯尼亚语 Bosnian, pl 波兰语 Polish, bi 比斯拉玛语 Bislama, nd 北恩德贝勒语 North Ndebele, ba 巴什基尔语 Bashkir, bg 保加利亚语 Bulgarian, az 阿塞拜疆语 Azerbaijani, ar 阿拉伯语 Arabic, af 阿非利堪斯语 Afrikaans, sq 阿尔巴尼亚语 Albanian, ab 阿布哈兹语 Abkhazian, os 奥塞梯语 Ossetian, ee 埃维语 Ewe, et 爱沙尼亚语 Estonian, ay 艾马拉语 Aymara, lzh 中文（文言文） Chinese (classical), am 阿姆哈拉语 Amharic, ckb 中库尔德语 Central Kurdish, cy 威尔士语 Welsh, gl 加利西亚语 Galician, ha 豪萨语 Hausa, hy 亚美尼亚语 Armenian, ig 伊博语 Igbo, kmr 北库尔德语 Northern Kurdish, ln 林加拉语 Lingala, nso 北索托语 Northern Sotho, ny 齐切瓦语 Chewa, om 奥洛莫语 Oromo, sn 绍纳语 Shona, so 索马里语 Somali, sr 塞尔维亚语 Serbian, sw 斯瓦希里语 Swahili, xh 科萨语 Xhosa, yo 约鲁巴语 Yoruba, zu 祖鲁语 Zulu
func Translate(text, source, target string) (t VolcTranslate, err error) {
	body, _ := json.Marshal(struct {
		SourceLanguage string   `json:"SourceLanguage"`
		TargetLanguage string   `json:"TargetLanguage"`
		TextList       []string `json:"TextList"`
	}{
		SourceLanguage: source,
		TargetLanguage: target,
		TextList:       []string{text},
	})

	t = VolcTranslate{}
	resp, code, err := client.Json("TranslateText", nil, string(body))
	if err != nil {
		return
	}
	if code != 200 {
		return t, errors.New(string(resp))
	}

	err = json.Unmarshal(resp, &t)
	if err != nil {
		return
	}
	return
}
