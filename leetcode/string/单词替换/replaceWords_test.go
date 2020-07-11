package main

import (
	"fmt"
	"testing"
)

func Test_ReplaceWords(T *testing.T) {
	fmt.Println(replaceWords([]string{"cat", "bat", "rat"}, "the cattle was rattled by the battery"))

	dict := []string{"e", "k", "c", "harqp", "h", "gsafc", "vn", "lqp", "soy", "mr", "x", "iitgm", "sb", "oo", "spj", "gwmly", "iu", "z", "f", "ha", "vds", "v", "vpx", "fir", "t", "xo", "apifm", "tlznm", "kkv", "nxyud", "j", "qp", "omn", "zoxp", "mutu", "i", "nxth", "dwuer", "sadl", "pv", "w", "mding", "mubem", "xsmwc", "vl", "farov", "twfmq", "ljhmr", "q", "bbzs", "kd", "kwc", "a", "buq", "sm", "yi", "nypa", "xwz", "si", "amqx", "iy", "eb", "qvgt", "twy", "rf", "dc", "utt", "mxjfu", "hm", "trz", "lzh", "lref", "qbx", "fmemr", "gil", "go", "qggh", "uud", "trnhf", "gels", "dfdq", "qzkx", "qxw"}
	sentence := "ikkbp miszkays wqjferqoxjwvbieyk gvcfldkiavww vhokchxz dvypwyb bxahfzcfanteibiltins ueebf lqhflvwxksi dco kddxmckhvqifbuzkhstp wc ytzzlm gximjuhzfdjuamhsu gdkbmhpnvy ifvifheoxqlbosfww mengfdydekwttkhbzenk wjhmmyltmeufqvcpcxg hthcuovils ldipovluo aiprogn nusquzpmnogtjkklfhta klxvvlvyh nxzgnrveghc mpppfhzjkbucv cqcft uwmahhqradjtf iaaasabqqzmbcig zcpvpyypsmodtoiif qjuiqtfhzcpnmtk yzfragcextvx ivnvgkaqs iplazv jurtsyh gzixfeugj rnukjgtjpim hscyhgoru aledyrmzwhsz xbahcwfwm hzd ygelddphxnbh rvjxtlqfnlmwdoezh zawfkko iwhkcddxgpqtdrjrcv bbfj mhs nenrqfkbf spfpazr wrkjiwyf cw dtd cqibzmuuhukwylrnld dtaxhddidfwqs bgnnoxgyynol hg dijhrrpnwjlju muzzrrsypzgwvblf zbugltrnyzbg hktdviastoireyiqf qvufxgcixvhrjqtna ipfzhuvgo daee r nlipyfszvxlwqw yoq dewpgtcrzausqwhh qzsaobsghgm ichlpsjlsrwzhbyfhm ksenb bqprarpgnyemzwifqzz oai pnqottd nygesjtlpala qmxixtooxtbrzyorn gyvukjpc s mxhlkdaycskj uvwmerplaibeknltuvd ocnn frotscysdyclrc ckcttaceuuxzcghw pxbd oklwhcppuziixpvihihp"
	fmt.Println(replaceWords(dict, sentence))

	dict = []string{"a", "aa", "aaa", "aaaa"}
	sentence = "a aa a aaaa aaa aaa aaa aaaaaa bbb baba ababa"
	fmt.Println(replaceWords(dict, sentence))

	dict = []string{"vmfydkuo", "ymaw", "cecqzhk", "qtm", "snrvddn", "leddmqvl", "fvylup", "fno", "efxwj", "svbjksnd", "iweiryjx", "coxb", "zvqvp", "iozicai", "fkmdeeh", "bvzox", "xweubr", "vyt", "vrroaeec", "pwz", "ggbbgm", "vxo", "bmdgvut", "bzlz", "bsjnqck", "lno", "qbduddyu", "yfyxad", "uatc", "gkuk", "luowi", "yit", "rsxmse", "svdwvbpm", "rulbk", "osaakwc", "kzw", "krycw", "ptf", "nle", "zdbod", "pzu", "lmxe", "ewv", "hvjins", "yttl", "vke", "ynlhab", "bwnz", "cyekvnr", "fdbj", "pijgjp", "rzph", "onye", "synyce", "ujtlwgg", "qptbkdzl", "wjxbotrl", "etgmofr", "sak", "elr", "yvnra"}
	sentence = "ggkmrfmejugv webfxzezusd tqgvsetkxti ebeozlplvhlq nvnfelhumodikkhosg q glpxtztajbyyocl ct hqttgnkrcb mpeol krmkltmie ymufnfoenqejemigcum wcsarrgfpwwrbbso avymhdentscbirdiimkv ocpurbsupqvmpwsroc ajlgjpkijtxv azvb okthrxhqecd yffvdpkiqjhyba mabqzkiirj arscyscvdicanjmcotf x fqelimvwuyr qpek xbrelcyzqo durpihk peupgqntivpojlr ipihjybpfm bjhlielhxxxpxblyn niuvhshhdseljeoqtbza manlezwinkkaxp xueypijsgtppae lcvtxhxxlkdkdjeegv pqzotlgb nrivllttlseb qkgsscuyfmutoaepy ypzgbozsv"
	fmt.Println(replaceWords(dict, sentence))
}
