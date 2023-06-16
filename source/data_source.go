package source

const StringInfo = `
{"nameEN":"china","nameZH":"中国","code":"CN"},
{"nameEN":"afghanistan","nameZH":"阿富汗","code":"AF"},
{"nameEN":"åland islands","nameZH":"奥兰群岛","code":"AX"},
{"nameEN":"albania","nameZH":"阿尔巴尼亚","code":"AL"},
{"nameEN":"algeria","nameZH":"阿尔及利亚","code":"DZ"},
{"nameEN":"american samoa","nameZH":"美国属萨摩亚","code":"AS"},
{"nameEN":"andorra","nameZH":"安道尔","code":"AD"},
{"nameEN":"angola","nameZH":"安哥拉","code":"AO"},
{"nameEN":"anguilla","nameZH":"安圭拉","code":"AI"},
{"nameEN":"antarctica","nameZH":"南极洲","code":"AQ"},
{"nameEN":"antigua and barbuda","nameZH":"安提瓜和巴布达","code":"AG"},
{"nameEN":"argentina","nameZH":"阿根廷","code":"AR"},
{"nameEN":"armenia","nameZH":"亚美尼亚","code":"AM"},
{"nameEN":"aruba","nameZH":"阿鲁巴","code":"AW"},
{"nameEN":"australia","nameZH":"澳大利亚","code":"AU"},
{"nameEN":"austria","nameZH":"奥地利","code":"AT"},
{"nameEN":"azerbaijan","nameZH":"阿塞拜疆","code":"AZ"},
{"nameEN":"bahamas","nameZH":"巴哈马","code":"BS"},
{"nameEN":"bahrain","nameZH":"巴林","code":"BH"},
{"nameEN":"bangladesh","nameZH":"孟加拉国","code":"BD"},
{"nameEN":"barbados","nameZH":"巴巴多斯","code":"BB"},
{"nameEN":"belarus","nameZH":"白俄罗斯","code":"BY"},
{"nameEN":"belgium","nameZH":"比利时","code":"BE"},
{"nameEN":"belize","nameZH":"伯利兹","code":"BZ"},
{"nameEN":"benin","nameZH":"贝宁","code":"BJ"},
{"nameEN":"bermuda","nameZH":"百慕大","code":"BM"},
{"nameEN":"bhutan","nameZH":"不丹","code":"BT"},
{"nameEN":"bolivia","nameZH":"玻利维亚","code":"BO"},
{"nameEN":"bosnia and herzegovina","nameZH":"波斯尼亚和黑塞哥维那","code":"BA"},
{"nameEN":"botswana","nameZH":"博茨瓦纳","code":"BW"},
{"nameEN":"bouvet island","nameZH":"布维岛","code":"BV"},
{"nameEN":"brazil","nameZH":"巴西","code":"BR"},
{"nameEN":"british indian ocean territory","nameZH":"英属印度洋领地","code":"IO"},
{"nameEN":"brunei darussalam","nameZH":"文莱达鲁萨兰国","code":"BN"},
{"nameEN":"bulgaria","nameZH":"保加利亚","code":"BG"},
{"nameEN":"burkina faso","nameZH":"布基纳法索","code":"BF"},
{"nameEN":"burundi","nameZH":"布隆迪","code":"BI"},
{"nameEN":"cabo verde","nameZH":"佛得角","code":"CV"},
{"nameEN":"cambodia","nameZH":"柬埔寨","code":"KH"},
{"nameEN":"cameroon","nameZH":"喀麦隆","code":"CM"},
{"nameEN":"canada","nameZH":"加拿大","code":"CA"},
{"nameEN":"cayman islands","nameZH":"开曼群岛","code":"KY"},
{"nameEN":"central african republic","nameZH":"中非共和国","code":"CF"},
{"nameEN":"chad","nameZH":"乍得","code":"TD"},
{"nameEN":"chile","nameZH":"智利","code":"CL"},
{"nameEN":"christmas island","nameZH":"圣诞岛","code":"CX"},
{"nameEN":"cocos (keeling) islands","nameZH":"科科斯（基林）群岛","code":"CC"},
{"nameEN":"colombia","nameZH":"哥伦比亚","code":"CO"},
{"nameEN":"comoros","nameZH":"科摩罗","code":"KM"},
{"nameEN":"congo","nameZH":"刚果","code":"CG"},
{"nameEN":"congo, the democratic republic of the","nameZH":"刚果民主共和国","code":"CD"},
{"nameEN":"cook islands","nameZH":"库克群岛","code":"CK"},
{"nameEN":"costa rica","nameZH":"哥斯达黎加","code":"CR"},
{"nameEN":"côte d'ivoire","nameZH":"科特迪瓦","code":"CI"},
{"nameEN":"croatia","nameZH":"克罗地亚","code":"HR"},
{"nameEN":"cuba","nameZH":"古巴","code":"CU"},
{"nameEN":"cyprus","nameZH":"塞浦路斯","code":"CY"},
{"nameEN":"czech republic","nameZH":"捷克共和国","code":"CZ"},
{"nameEN":"denmark","nameZH":"丹麦","code":"DK"},
{"nameEN":"djibouti","nameZH":"吉布提","code":"DJ"},
{"nameEN":"dominica","nameZH":"多米尼克","code":"DM"},
{"nameEN":"dominican republic","nameZH":"多米尼加共和国","code":"DO"},
{"nameEN":"ecuador","nameZH":"厄瓜多尔","code":"EC"},
{"nameEN":"egypt","nameZH":"埃及","code":"EG"},
{"nameEN":"el salvador","nameZH":"萨尔瓦多","code":"SV"},
{"nameEN":"equatorial guinea","nameZH":"赤道几内亚","code":"GQ"},
{"nameEN":"eritrea","nameZH":"厄立特里亚","code":"ER"},
{"nameEN":"estonia","nameZH":"爱沙尼亚","code":"EE"},
{"nameEN":"eswatini","nameZH":"斯威士兰","code":"SZ"},
{"nameEN":"ethiopia","nameZH":"埃塞俄比亚","code":"ET"},
{"nameEN":"falkland islands (malvinas)","nameZH":"福克兰群岛（马尔维纳斯）","code":"FK"},
{"nameEN":"faroe islands","nameZH":"法罗群岛","code":"FO"},
{"nameEN":"fiji","nameZH":"斐济","code":"FJ"},
{"nameEN":"finland","nameZH":"芬兰","code":"FI"},
{"nameEN":"france","nameZH":"法国","code":"FR"},
{"nameEN":"french guiana","nameZH":"法属圭亚那","code":"GF"},
{"nameEN":"french polynesia","nameZH":"法属波利尼西亚","code":"PF"},
{"nameEN":"french southern territories","nameZH":"法属南部领地","code":"TF"},
{"nameEN":"gabon","nameZH":"加蓬","code":"GA"},
{"nameEN":"gambia","nameZH":"冈比亚","code":"GM"},
{"nameEN":"georgia","nameZH":"格鲁吉亚","code":"GE"},
{"nameEN":"germany","nameZH":"德国","code":"DE"},
{"nameEN":"ghana","nameZH":"加纳","code":"GH"},
{"nameEN":"gibraltar","nameZH":"直布罗陀","code":"GI"},
{"nameEN":"greece","nameZH":"希腊","code":"GR"},
{"nameEN":"greenland","nameZH":"格陵兰","code":"GL"},
{"nameEN":"grenada","nameZH":"格林纳达","code":"GD"},
{"nameEN":"guadeloupe","nameZH":"瓜德罗普","code":"GP"},
{"nameEN":"guam","nameZH":"关岛","code":"GU"},
{"nameEN":"guatemala","nameZH":"危地马拉","code":"GT"},
{"nameEN":"guernsey","nameZH":"根西岛","code":"GG"},
{"nameEN":"guinea","nameZH":"几内亚","code":"GN"},
{"nameEN":"guinea-bissau","nameZH":"几内亚比绍","code":"GW"},
{"nameEN":"guyana","nameZH":"圭亚那","code":"GY"},
{"nameEN":"haiti","nameZH":"海地","code":"HT"},
{"nameEN":"heard island and mcdonald islands","nameZH":"赫德岛和麦克唐纳群岛","code":"HM"},
{"nameEN":"holy see (vatican city state)","nameZH":"梵蒂冈城国","code":"VA"},
{"nameEN":"honduras","nameZH":"洪都拉斯","code":"HN"},
{"nameEN":"hong kong","nameZH":"香港","code":"HK"},
{"nameEN":"hungary","nameZH":"匈牙利","code":"HU"},
{"nameEN":"iceland","nameZH":"冰岛","code":"IS"},
{"nameEN":"india","nameZH":"印度","code":"IN"},
{"nameEN":"indonesia","nameZH":"印度尼西亚","code":"ID"},
{"nameEN":"iran, islamic republic of","nameZH":"伊朗伊斯兰共和国","code":"IR"},
{"nameEN":"iraq","nameZH":"伊拉克","code":"IQ"},
{"nameEN":"ireland","nameZH":"爱尔兰","code":"IE"},
{"nameEN":"isle of man","nameZH":"马恩岛","code":"IM"},
{"nameEN":"israel","nameZH":"以色列","code":"IL"},
{"nameEN":"italy","nameZH":"意大利","code":"IT"},
{"nameEN":"jamaica","nameZH":"牙买加","code":"JM"},
{"nameEN":"japan","nameZH":"日本","code":"JP"},
{"nameEN":"jersey","nameZH":"泽西岛","code":"JE"},
{"nameEN":"jordan","nameZH":"约旦","code":"JO"},
{"nameEN":"kazakhstan","nameZH":"哈萨克斯坦","code":"KZ"},
{"nameEN":"kenya","nameZH":"肯尼亚","code":"KE"},
{"nameEN":"kiribati","nameZH":"基里巴斯","code":"KI"},
{"nameEN":"korea, democratic people's republic of","nameZH":"朝鲜民主主义人民共和国","code":"KP"},
{"nameEN":"korea, republic of","nameZH":"韩国","code":"KR"},
{"nameEN":"kuwait","nameZH":"科威特","code":"KW"},
{"nameEN":"kyrgyzstan","nameZH":"吉尔吉斯斯坦","code":"KG"},
{"nameEN":"lao people's democratic republic","nameZH":"老挝人民民主共和国","code":"LA"},
{"nameEN":"latvia","nameZH":"拉脱维亚","code":"LV"},
{"nameEN":"lebanon","nameZH":"黎巴嫩","code":"LB"},
{"nameEN":"lesotho","nameZH":"莱索托","code":"LS"},
{"nameEN":"liberia","nameZH":"利比里亚","code":"LR"},
{"nameEN":"libya","nameZH":"利比亚","code":"LY"},
{"nameEN":"liechtenstein","nameZH":"列支敦士登","code":"LI"},
{"nameEN":"lithuania","nameZH":"立陶宛","code":"LT"},
{"nameEN":"luxembourg","nameZH":"卢森堡","code":"LU"},
{"nameEN":"macao","nameZH":"澳门","code":"MO"},
{"nameEN":"madagascar","nameZH":"马达加斯加","code":"MG"},
{"nameEN":"malawi","nameZH":"马拉维","code":"MW"},
{"nameEN":"malaysia","nameZH":"马来西亚","code":"MY"},
{"nameEN":"maldives","nameZH":"马尔代夫","code":"MV"},
{"nameEN":"mali","nameZH":"马里","code":"ML"},
{"nameEN":"malta","nameZH":"马耳他","code":"MT"},
{"nameEN":"marshall islands","nameZH":"马绍尔群岛","code":"MH"},
{"nameEN":"martinique","nameZH":"马提尼克","code":"MQ"},
{"nameEN":"mauritania","nameZH":"毛里塔尼亚","code":"MR"},
{"nameEN":"mauritius","nameZH":"毛里求斯","code":"MU"},
{"nameEN":"mayotte","nameZH":"马约特","code":"YT"},
{"nameEN":"mexico","nameZH":"墨西哥","code":"MX"},
{"nameEN":"micronesia, federated states of","nameZH":"密克罗尼西亚联邦","code":"FM"},
{"nameEN":"moldova, republic of","nameZH":"摩尔多瓦共和国","code":"MD"},
{"nameEN":"monaco","nameZH":"摩纳哥","code":"MC"},
{"nameEN":"mongolia","nameZH":"蒙古","code":"MN"},
{"nameEN":"montenegro","nameZH":"黑山","code":"ME"},
{"nameEN":"montserrat","nameZH":"蒙特塞拉特","code":"MS"},
{"nameEN":"morocco","nameZH":"摩洛哥","code":"MA"},
{"nameEN":"mozambique","nameZH":"莫桑比克","code":"MZ"},
{"nameEN":"myanmar","nameZH":"缅甸","code":"MM"},
{"nameEN":"namibia","nameZH":"纳米比亚","code":"NA"},
{"nameEN":"nauru","nameZH":"瑙鲁","code":"NR"},
{"nameEN":"nepal","nameZH":"尼泊尔","code":"NP"},
{"nameEN":"netherlands","nameZH":"荷兰","code":"NL"},
{"nameEN":"new caledonia","nameZH":"新喀里多尼亚","code":"NC"},
{"nameEN":"new zealand","nameZH":"新西兰","code":"NZ"},
{"nameEN":"nicaragua","nameZH":"尼加拉瓜","code":"NI"},
{"nameEN":"niger","nameZH":"尼日尔","code":"NE"},
{"nameEN":"nigeria","nameZH":"尼日利亚","code":"NG"},
{"nameEN":"niue","nameZH":"纽埃","code":"NU"},
{"nameEN":"norfolk island","nameZH":"诺福克岛","code":"NF"},
{"nameEN":"north macedonia","nameZH":"北马其顿","code":"MK"},
{"nameEN":"northern mariana islands","nameZH":"北马里亚纳群岛","code":"MP"},
{"nameEN":"norway","nameZH":"挪威","code":"NO"},
{"nameEN":"oman","nameZH":"阿曼","code":"OM"},
{"nameEN":"pakistan","nameZH":"巴基斯坦","code":"PK"},
{"nameEN":"palau","nameZH":"帕劳","code":"PW"},
{"nameEN":"palestine, state of","nameZH":"巴勒斯坦国","code":"PS"},
{"nameEN":"panama","nameZH":"巴拿马","code":"PA"},
{"nameEN":"papua new guinea","nameZH":"巴布亚新几内亚","code":"PG"},
{"nameEN":"paraguay","nameZH":"巴拉圭","code":"PY"},
{"nameEN":"peru","nameZH":"秘鲁","code":"PE"},
{"nameEN":"philippines","nameZH":"菲律宾","code":"PH"},
{"nameEN":"pitcairn","nameZH":"皮特凯恩","code":"PN"},
{"nameEN":"poland","nameZH":"波兰","code":"PL"},
{"nameEN":"portugal","nameZH":"葡萄牙","code":"PT"},
{"nameEN":"puerto rico","nameZH":"波多黎各","code":"PR"},
{"nameEN":"qatar","nameZH":"卡塔尔","code":"QA"},
{"nameEN":"réunion","nameZH":"留尼汪","code":"RE"},
{"nameEN":"romania","nameZH":"罗马尼亚","code":"RO"},
{"nameEN":"russian federation","nameZH":"俄罗斯联邦","code":"RU"},
{"nameEN":"rwanda","nameZH":"卢旺达","code":"RW"},
{"nameEN":"saint barthélemy","nameZH":"圣巴泰勒米","code":"BL"},
{"nameEN":"saint helena, ascension and tristan da cunha","nameZH":"圣赫勒拿、阿森松岛和特里斯坦-达库尼亚岛","code":"SH"},
{"nameEN":"saint kitts and nevis","nameZH":"圣基茨和尼维斯","code":"KN"},
{"nameEN":"saint lucia","nameZH":"圣卢西亚","code":"LC"},
{"nameEN":"saint martin (french part)","nameZH":"法属圣马丁","code":"MF"},
{"nameEN":"saint pierre and miquelon","nameZH":"圣皮埃尔和密克隆群岛","code":"PM"},
{"nameEN":"saint vincent and the grenadines","nameZH":"圣文森特和格林纳丁斯","code":"VC"},
{"nameEN":"samoa","nameZH":"萨摩亚","code":"WS"},
{"nameEN":"san marino","nameZH":"圣马力诺","code":"SM"},
{"nameEN":"sao tome and principe","nameZH":"圣多美和普林西比","code":"ST"},
{"nameEN":"saudi arabia","nameZH":"沙特阿拉伯","code":"SA"},
{"nameEN":"senegal","nameZH":"塞内加尔","code":"SN"},
{"nameEN":"serbia","nameZH":"塞尔维亚","code":"RS"},
{"nameEN":"seychelles","nameZH":"塞舌尔","code":"SC"},
{"nameEN":"sierra leone","nameZH":"塞拉利昂","code":"SL"},
{"nameEN":"singapore","nameZH":"新加坡","code":"SG"},
{"nameEN":"sint maarten (dutch part)","nameZH":"荷属圣马丁","code":"SX"},
{"nameEN":"slovakia","nameZH":"斯洛伐克","code":"SK"},
{"nameEN":"slovenia","nameZH":"斯洛文尼亚","code":"SI"},
{"nameEN":"solomon islands","nameZH":"所罗门群岛","code":"SB"},
{"nameEN":"somalia","nameZH":"索马里","code":"SO"},
{"nameEN":"south africa","nameZH":"南非","code":"ZA"},
{"nameEN":"south georgia and the south sandwich islands","nameZH":"南乔治亚和南桑威奇群岛","code":"GS"},
{"nameEN":"south sudan","nameZH":"南苏丹","code":"SS"},
{"nameEN":"spain","nameZH":"西班牙","code":"ES"},
{"nameEN":"sri lanka","nameZH":"斯里兰卡","code":"LK"},
{"nameEN":"sudan","nameZH":"苏丹","code":"SD"},
{"nameEN":"suriname","nameZH":"苏里南","code":"SR"},
{"nameEN":"svalbard and jan mayen","nameZH":"斯瓦尔巴群岛和扬马延岛","code":"SJ"},
{"nameEN":"sweden","nameZH":"瑞典","code":"SE"},
{"nameEN":"switzerland","nameZH":"瑞士","code":"CH"},
{"nameEN":"syrian arab republic","nameZH":"叙利亚阿拉伯共和国","code":"SY"},
{"nameEN":"taiwan, province of china","nameZH":"中国台湾省","code":"TW"},
{"nameEN":"tajikistan","nameZH":"塔吉克斯坦","code":"TJ"},
{"nameEN":"tanzania, united republic of","nameZH":"坦桑尼亚联合共和国","code":"TZ"},
{"nameEN":"thailand","nameZH":"泰国","code":"TH"},
{"nameEN":"timor-leste","nameZH":"东帝汶","code":"TL"},
{"nameEN":"togo","nameZH":"多哥","code":"TG"},
{"nameEN":"tokelau","nameZH":"托克劳","code":"TK"},
{"nameEN":"tonga","nameZH":"汤加","code":"TO"},
{"nameEN":"trinidad and tobago","nameZH":"特立尼达和多巴哥","code":"TT"},
{"nameEN":"tunisia","nameZH":"突尼斯","code":"TN"},
{"nameEN":"turkey","nameZH":"土耳其","code":"TR"},
{"nameEN":"turkmenistan","nameZH":"土库曼斯坦","code":"TM"},
{"nameEN":"turks and caicos islands","nameZH":"特克斯和凯科斯群岛","code":"TC"},
{"nameEN":"tuvalu","nameZH":"图瓦卢","code":"TV"},
{"nameEN":"uganda","nameZH":"乌干达","code":"UG"},
{"nameEN":"ukraine","nameZH":"乌克兰","code":"UA"},
{"nameEN":"united arab emirates","nameZH":"阿拉伯联合酋长国","code":"AE"},
{"nameEN":"united kingdom","nameZH":"英国","code":"GB"},
{"nameEN":"united states","nameZH":"美国","code":"US"},
{"nameEN":"united states minor outlying islands","nameZH":"美国本土外小岛屿","code":"UM"},
{"nameEN":"uruguay","nameZH":"乌拉圭","code":"UY"},
{"nameEN":"uzbekistan","nameZH":"乌兹别克斯坦","code":"UZ"},
{"nameEN":"vanuatu","nameZH":"瓦努阿图","code":"VU"},
{"nameEN":"venezuela, bolivarian republic of","nameZH":"委内瑞拉玻利瓦尔共和国","code":"VE"},
{"nameEN":"viet nam","nameZH":"越南","code":"VN"},
{"nameEN":"virgin islands, british","nameZH":"英属维尔京群岛","code":"VG"},
{"nameEN":"virgin islands, u.s.","nameZH":"美属维尔京群岛","code":"VI"},
{"nameEN":"wallis and futuna","nameZH":"瓦利斯和富图纳","code":"WF"},
{"nameEN":"western sahara","nameZH":"西撒哈拉","code":"EH"},
{"nameEN":"yemen","nameZH":"也门","code":"YE"},
{"nameEN":"zambia","nameZH":"赞比亚","code":"ZM"},
{"nameEN":"zimbabwe","nameZH":"津巴布韦","code":"ZW"}
`