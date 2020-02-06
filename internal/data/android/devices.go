package android

import (
	"encoding/json"
	"github.com/deletescape/noise/pkg/entities"
)

const devicesJson = `[
  {
    "manufacturer": "Google",
    "market_name": "Pixel 2",
    "codename": "walleye",
    "model": "Pixel 2"
  },
  {
    "manufacturer": "Google",
    "market_name": "Pixel 2 XL",
    "codename": "taimen",
    "model": "Pixel 2 XL"
  },
  {
    "manufacturer": "Google",
    "market_name": "Pixel 3",
    "codename": "blueline",
    "model": "Pixel 3"
  },
  {
    "manufacturer": "Google",
    "market_name": "Pixel 3a",
    "codename": "sargo",
    "model": "Pixel 3a"
  },
  {
    "manufacturer": "Google",
    "market_name": "Pixel 3a XL",
    "codename": "bonito",
    "model": "Pixel 3a XL"
  },
  {
    "manufacturer": "Google",
    "market_name": "Pixel 3 XL",
    "codename": "crosshatch",
    "model": "Pixel 3 XL"
  },
  {
    "manufacturer": "Google",
    "market_name": "Pixel 4",
    "codename": "flame",
    "model": "Pixel 4"
  },
  {
    "manufacturer": "Google",
    "market_name": "Pixel 4 XL",
    "codename": "coral",
    "model": "Pixel 4 XL"
  },
  {
    "manufacturer": "Huawei",
    "market_name": "Mate 9",
    "codename": "HWMHA",
    "model": "MHA-AL00"
  },
  {
    "manufacturer": "Huawei",
    "market_name": "Mate 9",
    "codename": "HWMHA",
    "model": "MHA-L09"
  },
  {
    "manufacturer": "Huawei",
    "market_name": "Mate 9",
    "codename": "HWMHA",
    "model": "MHA-L29"
  },
  {
    "manufacturer": "Huawei",
    "market_name": "Mate 9",
    "codename": "HWMHA",
    "model": "MHA-TL00"
  },
  {
    "manufacturer": "Huawei",
    "market_name": "Mate 10",
    "codename": "HWALP",
    "model": "ALP-AL00"
  },
  {
    "manufacturer": "Huawei",
    "market_name": "Mate 10",
    "codename": "HWALP",
    "model": "ALP-L09"
  },
  {
    "manufacturer": "Huawei",
    "market_name": "Mate 10",
    "codename": "HWALP",
    "model": "ALP-L29"
  },
  {
    "manufacturer": "Huawei",
    "market_name": "Mate 10",
    "codename": "HWALP",
    "model": "ALP-TL00"
  },
  {
    "manufacturer": "Huawei",
    "market_name": "Nexus 6P",
    "codename": "angler",
    "model": "Nexus 6P"
  },
  {
    "manufacturer": "Huawei",
    "market_name": "Honor View 10",
    "codename": "HWBKL",
    "model": "BKL-L04"
  },
  {
    "manufacturer": "Huawei",
    "market_name": "Honor View 10",
    "codename": "HWBKL",
    "model": "BKL-L09"
  },
  {
    "manufacturer": "Huawei",
    "market_name": "Honor 7X",
    "codename": "HWBND-H",
    "model": "BND-L21"
  },
  {
    "manufacturer": "Huawei",
    "market_name": "Honor 7X",
    "codename": "HWBND-H",
    "model": "BND-L24"
  },
  {
    "manufacturer": "Huawei",
    "market_name": "Honor 7X",
    "codename": "HWBND-H",
    "model": "BND-L31"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G5",
    "codename": "h1",
    "model": "LG-F700K"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G5",
    "codename": "h1",
    "model": "LG-F700L"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G5",
    "codename": "h1",
    "model": "LG-F700S"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G5",
    "codename": "h1",
    "model": "LG-H820"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G5",
    "codename": "h1",
    "model": "LG-H820PR"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G5",
    "codename": "h1",
    "model": "LG-H830"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G5",
    "codename": "h1",
    "model": "LG-H831"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G5",
    "codename": "h1",
    "model": "LG-H850"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G5",
    "codename": "h1",
    "model": "LG-H858"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G5",
    "codename": "h1",
    "model": "LG-H860"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G5",
    "codename": "h1",
    "model": "LG-H868"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G5",
    "codename": "h1",
    "model": "LGAS992"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G5",
    "codename": "h1",
    "model": "LGLS992"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G5",
    "codename": "h1",
    "model": "LGUS992"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G5",
    "codename": "h1",
    "model": "RS988"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G5",
    "codename": "h1",
    "model": "VS987"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G6",
    "codename": "lucye",
    "model": "LG-AS993"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G6",
    "codename": "lucye",
    "model": "LG-H870"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G6",
    "codename": "lucye",
    "model": "LG-H870AR"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G6",
    "codename": "lucye",
    "model": "LG-H870DS"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G6",
    "codename": "lucye",
    "model": "LG-H870I"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G6",
    "codename": "lucye",
    "model": "LG-H870S"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G6",
    "codename": "lucye",
    "model": "LG-H871"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G6",
    "codename": "lucye",
    "model": "LG-H871S"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G6",
    "codename": "lucye",
    "model": "LG-H872"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G6",
    "codename": "lucye",
    "model": "LG-H872PR"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G6",
    "codename": "lucye",
    "model": "LG-H873"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G6",
    "codename": "lucye",
    "model": "LG-LS993"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G6",
    "codename": "lucye",
    "model": "LGM-G600K"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G6",
    "codename": "lucye",
    "model": "LGM-G600L"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G6",
    "codename": "lucye",
    "model": "LGM-G600S"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G6",
    "codename": "lucye",
    "model": "LGUS997"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG G6",
    "codename": "lucye",
    "model": "VS988"
  },
  {
    "manufacturer": "LGE",
    "market_name": "Nexus 4",
    "codename": "mako",
    "model": "Nexus 4"
  },
  {
    "manufacturer": "LGE",
    "market_name": "Nexus 5",
    "codename": "hammerhead",
    "model": "Nexus 5"
  },
  {
    "manufacturer": "LGE",
    "market_name": "Nexus 5X",
    "codename": "bullhead",
    "model": "Nexus 5X"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG V50 ThinQ",
    "codename": "flashlmdd",
    "model": "LM-V500"
  },
  {
    "manufacturer": "LGE",
    "market_name": "LG V50 ThinQ",
    "codename": "flashlmdd",
    "model": "LM-V500N"
  },
  {
    "manufacturer": "Motorola",
    "market_name": "Nexus 6",
    "codename": "shamu",
    "model": "Nexus 6"
  },
  {
    "manufacturer": "Motorola",
    "market_name": "Moto Z",
    "codename": "griffin",
    "model": "XT1650"
  },
  {
    "manufacturer": "Motorola",
    "market_name": "Moto Z",
    "codename": "griffin",
    "model": "XT1650-05"
  },
  {
    "manufacturer": "OnePlus",
    "market_name": "OnePlus 6",
    "codename": "OnePlus6",
    "model": "ONEPLUS A6003"
  },
  {
    "manufacturer": "OnePlus",
    "market_name": "OnePlus 6T",
    "codename": "OnePlus6T",
    "model": "ONEPLUS A6013"
  },
  {
    "manufacturer": "OnePlus",
    "market_name": "OnePlus 6T",
    "codename": "OnePlus6TSingle",
    "model": "ONEPLUS A6013"
  },
  {
    "manufacturer": "OnePlus",
    "market_name": "OnePlus 7",
    "codename": "OnePlus7",
    "model": "GM1905"
  },
  {
    "manufacturer": "OnePlus",
    "market_name": "OnePlus 7 Pro",
    "codename": "OP7ProNRSpr",
    "model": "GM1925"
  },
  {
    "manufacturer": "OnePlus",
    "market_name": "OnePlus 7 Pro",
    "codename": "OnePlus7Pro",
    "model": "GM1917"
  },
  {
    "manufacturer": "OnePlus",
    "market_name": "OnePlus 7 Pro",
    "codename": "OnePlus7ProTMO",
    "model": "GM1915"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy A50",
    "codename": "a50",
    "model": "SM-A505F"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy A50",
    "codename": "a50",
    "model": "SM-A505FM"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy A50",
    "codename": "a50",
    "model": "SM-A505FN"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy A50",
    "codename": "a50",
    "model": "SM-A505G"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy A50",
    "codename": "a50",
    "model": "SM-A505GN"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy A50",
    "codename": "a50",
    "model": "SM-A505GT"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy A50",
    "codename": "a50",
    "model": "SM-A505N"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy A50",
    "codename": "a50",
    "model": "SM-A505U"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy A50",
    "codename": "a50",
    "model": "SM-A505U1"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy A50",
    "codename": "a50",
    "model": "SM-A505W"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy A50",
    "codename": "a50",
    "model": "SM-A505YN"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy A50",
    "codename": "a50",
    "model": "SM-S506DL"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy A6",
    "codename": "a6elteaio",
    "model": "SM-A600AZ"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy A6",
    "codename": "a6elteatt",
    "model": "SM-A600A"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy A6",
    "codename": "a6eltemtr",
    "model": "SM-A600T1"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy A6",
    "codename": "a6eltespr",
    "model": "SM-A600P"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy A6",
    "codename": "a6eltetmo",
    "model": "SM-A600T"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy A6",
    "codename": "a6elteue",
    "model": "SM-A600U"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy A6",
    "codename": "a6lte",
    "model": "SM-A600F"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy A6",
    "codename": "a6lte",
    "model": "SM-A600FN"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy A6",
    "codename": "a6lte",
    "model": "SM-A600G"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy A6",
    "codename": "a6lte",
    "model": "SM-A600GN"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy A6",
    "codename": "a6lteks",
    "model": "SM-A600N"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note7",
    "codename": "SC-01J",
    "model": "SC-01J"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note7",
    "codename": "SCV34",
    "model": "SCV34"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note7",
    "codename": "gracelte",
    "model": "SM-N930F"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note7",
    "codename": "gracelte",
    "model": "SM-N930X"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note7",
    "codename": "graceltektt",
    "model": "SM-N930K"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note7",
    "codename": "graceltelgt",
    "model": "SM-N930L"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note7",
    "codename": "gracelteskt",
    "model": "SM-N930S"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note7",
    "codename": "graceqlteacg",
    "model": "SM-N930R7"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note7",
    "codename": "graceqlteatt",
    "model": "SAMSUNG-SM-N930A"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note7",
    "codename": "graceqltebmc",
    "model": "SM-N930W8"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note7",
    "codename": "graceqltechn",
    "model": "SM-N9300"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note7",
    "codename": "graceqltedcm",
    "model": "SGH-N037"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note7",
    "codename": "graceqltelra",
    "model": "SM-N930R6"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note7",
    "codename": "graceqltespr",
    "model": "SM-N930P"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note7",
    "codename": "graceqltetfnvzw",
    "model": "SM-N930VL"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note7",
    "codename": "graceqltetmo",
    "model": "SM-N930T"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note7",
    "codename": "graceqlteue",
    "model": "SM-N930U"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note7",
    "codename": "graceqlteusc",
    "model": "SM-N930R4"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note7",
    "codename": "graceqltevzw",
    "model": "SM-N930V"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note8",
    "codename": "SC-01K",
    "model": "SC-01K"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note8",
    "codename": "SCV37",
    "model": "SCV37"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note8",
    "codename": "greatlte",
    "model": "SM-N950F"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note8",
    "codename": "greatlteks",
    "model": "SM-N950N"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note8",
    "codename": "greatlteks",
    "model": "SM-N950XN"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note8",
    "codename": "greatqlte",
    "model": "SM-N950U"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note8",
    "codename": "greatqltechn",
    "model": "SM-N9500"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note8",
    "codename": "greatqltecmcc",
    "model": "SM-N9508"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note8",
    "codename": "greatqltecs",
    "model": "SM-N950W"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note8",
    "codename": "greatqlteue",
    "model": "SM-N950U1"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note9",
    "codename": "SC-01L",
    "model": "SC-01L"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note9",
    "codename": "SCV40",
    "model": "SCV40"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note9",
    "codename": "crownlte",
    "model": "SM-N960F"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note9",
    "codename": "crownlteks",
    "model": "SM-N960N"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note9",
    "codename": "crownqltechn",
    "model": "SM-N9600"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note9",
    "codename": "crownqltecs",
    "model": "SM-N960W"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note9",
    "codename": "crownqltesq",
    "model": "SM-N960U"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy Note9",
    "codename": "crownqlteue",
    "model": "SM-N960U1"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "SC-04F",
    "model": "SC-04F"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "SCL23",
    "model": "SCL23"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "k3g",
    "model": "SM-G900H"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "klte",
    "model": "SM-G9008W"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "klte",
    "model": "SM-G9009W"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "klte",
    "model": "SM-G900F"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "klte",
    "model": "SM-G900FQ"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "klte",
    "model": "SM-G900I"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "klte",
    "model": "SM-G900M"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "klte",
    "model": "SM-G900MD"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "klteMetroPCS",
    "model": "SM-G900T1"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "klteMetroPCS",
    "model": "SM-G900T4"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "klteacg",
    "model": "SM-G900R7"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "klteaio",
    "model": "SAMSUNG-SM-G900AZ"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "klteatt",
    "model": "SAMSUNG-SM-G900A"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "kltecan",
    "model": "SM-G900W8"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "klteduoszn",
    "model": "SM-G9006W"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "kltektt",
    "model": "SM-G900K"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "kltelgt",
    "model": "SM-G900L"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "kltelra",
    "model": "SM-G900R6"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "klteskt",
    "model": "SM-G900S"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "kltespr",
    "model": "SM-G900P"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "kltetfnvzw",
    "model": "SM-S903VL"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "kltetmo",
    "model": "SM-G900T"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "kltetmo",
    "model": "SM-G900T3"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "klteusc",
    "model": "SM-G900R4"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "kltevzw",
    "model": "SM-G900V"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "kwifi",
    "model": "SM-G900X"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "lentisltektt",
    "model": "SM-G906K"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "lentisltelgt",
    "model": "SM-G906L"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5",
    "codename": "lentislteskt",
    "model": "SM-G906S"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5 Neo",
    "codename": "s5neolte",
    "model": "SM-G903F"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5 Neo",
    "codename": "s5neolte",
    "model": "SM-G903M"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S5 Neo",
    "codename": "s5neoltecan",
    "model": "SM-G903W"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6",
    "codename": "SC-05G",
    "model": "SC-05G"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6",
    "codename": "zeroflte",
    "model": "SM-G920F"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6",
    "codename": "zeroflte",
    "model": "SM-G920I"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6",
    "codename": "zeroflte",
    "model": "SM-G920X"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6",
    "codename": "zeroflteacg",
    "model": "SM-G920R7"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6",
    "codename": "zeroflteaio",
    "model": "SAMSUNG-SM-G920AZ"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6",
    "codename": "zeroflteatt",
    "model": "SAMSUNG-SM-G920A"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6",
    "codename": "zerofltebmc",
    "model": "SM-G920W8"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6",
    "codename": "zerofltechn",
    "model": "SM-G9200"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6",
    "codename": "zerofltechn",
    "model": "SM-G9208"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6",
    "codename": "zerofltectc",
    "model": "SM-G9209"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6",
    "codename": "zerofltektt",
    "model": "SM-G920K"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6",
    "codename": "zerofltelgt",
    "model": "SM-G920L"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6",
    "codename": "zerofltelra",
    "model": "SM-G920R6"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6",
    "codename": "zerofltemtr",
    "model": "SM-G920T1"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6",
    "codename": "zeroflteskt",
    "model": "SM-G920S"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6",
    "codename": "zerofltespr",
    "model": "SM-G920P"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6",
    "codename": "zerofltetfnvzw",
    "model": "SM-S906L"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6",
    "codename": "zerofltetfnvzw",
    "model": "SM-S907VL"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6",
    "codename": "zerofltetmo",
    "model": "SM-G920T"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6",
    "codename": "zeroflteusc",
    "model": "SM-G920R4"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6",
    "codename": "zerofltevzw",
    "model": "SM-G920V"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge",
    "codename": "404SC",
    "model": "404SC"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge",
    "codename": "SC-04G",
    "model": "SC-04G"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge",
    "codename": "SCV31",
    "model": "SCV31"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge",
    "codename": "zerolte",
    "model": "SM-G925I"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge",
    "codename": "zerolte",
    "model": "SM-G925X"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge",
    "codename": "zerolteacg",
    "model": "SM-G925R7"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge",
    "codename": "zerolteatt",
    "model": "SAMSUNG-SM-G925A"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge",
    "codename": "zeroltebmc",
    "model": "SM-G925W8"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge",
    "codename": "zeroltechn",
    "model": "SM-G9250"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge",
    "codename": "zeroltektt",
    "model": "SM-G925K"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge",
    "codename": "zeroltelra",
    "model": "SM-G925R6"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge",
    "codename": "zerolteskt",
    "model": "SM-G925S"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge",
    "codename": "zeroltespr",
    "model": "SM-G925P"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge",
    "codename": "zeroltetmo",
    "model": "SM-G925T"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge",
    "codename": "zerolteusc",
    "model": "SM-G925R4"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge",
    "codename": "zeroltevzw",
    "model": "SM-G925V"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 edge",
    "codename": "zerolte",
    "model": "SM-G925F"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 edge",
    "codename": "zeroltelgt",
    "model": "SM-G925L"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge+",
    "codename": "zenlte",
    "model": "SM-G9287C"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge+",
    "codename": "zenlte",
    "model": "SM-G928C"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge+",
    "codename": "zenlte",
    "model": "SM-G928G"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge+",
    "codename": "zenlte",
    "model": "SM-G928I"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge+",
    "codename": "zenlte",
    "model": "SM-G928X"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge+",
    "codename": "zenlteatt",
    "model": "SAMSUNG-SM-G928A"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge+",
    "codename": "zenltebmc",
    "model": "SM-G928W8"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge+",
    "codename": "zenltechn",
    "model": "SM-G9280"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge+",
    "codename": "zenltektt",
    "model": "SM-G928K"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge+",
    "codename": "zenltekx",
    "model": "SM-G928N0"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge+",
    "codename": "zenltelgt",
    "model": "SM-G928L"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge+",
    "codename": "zenlteskt",
    "model": "SM-G928S"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge+",
    "codename": "zenltespr",
    "model": "SM-G928P"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge+",
    "codename": "zenltetmo",
    "model": "SM-G928T"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge+",
    "codename": "zenlteusc",
    "model": "SM-G928R4"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 Edge+",
    "codename": "zenltevzw",
    "model": "SM-G928V"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 edge+",
    "codename": "zenlte",
    "model": "SM-G9287"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 edge+",
    "codename": "zenlte",
    "model": "SM-G928F"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S6 edge+",
    "codename": "zenlte",
    "model": "SM-G928G"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7",
    "codename": "herolte",
    "model": "SM-G930F"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7",
    "codename": "herolte",
    "model": "SM-G930X"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7",
    "codename": "heroltebmc",
    "model": "SM-G930W8"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7",
    "codename": "heroltektt",
    "model": "SM-G930K"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7",
    "codename": "heroltelgt",
    "model": "SM-G930L"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7",
    "codename": "herolteskt",
    "model": "SM-G930S"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7",
    "codename": "heroqlteacg",
    "model": "SM-G930R7"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7",
    "codename": "heroqlteaio",
    "model": "SAMSUNG-SM-G930AZ"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7",
    "codename": "heroqlteatt",
    "model": "SAMSUNG-SM-G930A"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7",
    "codename": "heroqltecctvzw",
    "model": "SM-G930VC"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7",
    "codename": "heroqltechn",
    "model": "SM-G9300"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7",
    "codename": "heroqltechn",
    "model": "SM-G9308"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7",
    "codename": "heroqltelra",
    "model": "SM-G930R6"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7",
    "codename": "heroqltemtr",
    "model": "SM-G930T1"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7",
    "codename": "heroqltespr",
    "model": "SM-G930P"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7",
    "codename": "heroqltetfnvzw",
    "model": "SM-G930VL"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7",
    "codename": "heroqltetmo",
    "model": "SM-G930T"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7",
    "codename": "heroqlteue",
    "model": "SM-G930U"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7",
    "codename": "heroqlteusc",
    "model": "SM-G930R4"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7",
    "codename": "heroqltevzw",
    "model": "SM-G930V"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7 Edge",
    "codename": "SC-02H",
    "model": "SC-02H"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7 Edge",
    "codename": "SCV33",
    "model": "SCV33"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7 Edge",
    "codename": "hero2lte",
    "model": "SM-G935X"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7 Edge",
    "codename": "hero2ltebmc",
    "model": "SM-G935W8"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7 Edge",
    "codename": "hero2ltektt",
    "model": "SM-G935K"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7 Edge",
    "codename": "hero2lteskt",
    "model": "SM-G935S"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7 Edge",
    "codename": "hero2qlteatt",
    "model": "SAMSUNG-SM-G935A"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7 Edge",
    "codename": "hero2qltecctvzw",
    "model": "SM-G935VC"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7 Edge",
    "codename": "hero2qltespr",
    "model": "SM-G935P"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7 Edge",
    "codename": "hero2qltetmo",
    "model": "SM-G935T"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7 Edge",
    "codename": "hero2qlteusc",
    "model": "SM-G935R4"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7 Edge",
    "codename": "hero2qltevzw",
    "model": "SM-G935V"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7 edge",
    "codename": "hero2lte",
    "model": "SM-G935F"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7 edge",
    "codename": "hero2ltelgt",
    "model": "SM-G935L"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7 edge",
    "codename": "hero2qltechn",
    "model": "SM-G9350"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S7 edge",
    "codename": "hero2qlteue",
    "model": "SM-G935U"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S8",
    "codename": "SC-02J",
    "model": "SC-02J"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S8",
    "codename": "SCV36",
    "model": "SCV36"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S8",
    "codename": "dreamlte",
    "model": "SM-G950F"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S8",
    "codename": "dreamlteks",
    "model": "SM-G950N"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S8",
    "codename": "dreamqltecan",
    "model": "SM-G950W"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S8",
    "codename": "dreamqltechn",
    "model": "SM-G9500"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S8",
    "codename": "dreamqltecmcc",
    "model": "SM-G9508"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S8",
    "codename": "dreamqltesq",
    "model": "SM-G950U"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S8",
    "codename": "dreamqlteue",
    "model": "SM-G950U1"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S8+",
    "codename": "SC-03J",
    "model": "SC-03J"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S8+",
    "codename": "SCV35",
    "model": "SCV35"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S8+",
    "codename": "dream2lte",
    "model": "SM-G955F"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S8+",
    "codename": "dream2lteks",
    "model": "SM-G955N"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S8+",
    "codename": "dream2qltecan",
    "model": "SM-G955W"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S8+",
    "codename": "dream2qltechn",
    "model": "SM-G9550"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S8+",
    "codename": "dream2qltesq",
    "model": "SM-G955U"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S8+",
    "codename": "dream2qlteue",
    "model": "SM-G955U1"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S9",
    "codename": "SC-02K",
    "model": "SC-02K"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S9",
    "codename": "SCV38",
    "model": "SCV38"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S9",
    "codename": "starlte",
    "model": "SM-G960F"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S9",
    "codename": "starlteks",
    "model": "SM-G960N"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S9",
    "codename": "starqltechn",
    "model": "SM-G9600"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S9",
    "codename": "starqltecmcc",
    "model": "SM-G9608"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S9",
    "codename": "starqltecs",
    "model": "SM-G960W"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S9",
    "codename": "starqltesq",
    "model": "SM-G960U"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S9",
    "codename": "starqlteue",
    "model": "SM-G960U1"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S9+",
    "codename": "SC-03K",
    "model": "SC-03K"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S9+",
    "codename": "SCV39",
    "model": "SCV39"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S9+",
    "codename": "star2lte",
    "model": "SM-G965F"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S9+",
    "codename": "star2lteks",
    "model": "SM-G965N"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S9+",
    "codename": "star2qltechn",
    "model": "SM-G9650"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S9+",
    "codename": "star2qltecs",
    "model": "SM-G965W"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S9+",
    "codename": "star2qltesq",
    "model": "SM-G965U"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S9+",
    "codename": "star2qlteue",
    "model": "SM-G965U1"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S10",
    "codename": "SC-03L",
    "model": "SC-03L"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S10",
    "codename": "SCV41",
    "model": "SCV41"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S10",
    "codename": "beyond1",
    "model": "SM-G973F"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S10",
    "codename": "beyond1",
    "model": "SM-G973N"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S10",
    "codename": "beyond1q",
    "model": "SM-G9730"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S10",
    "codename": "beyond1q",
    "model": "SM-G9738"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S10",
    "codename": "beyond1q",
    "model": "SM-G973C"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S10",
    "codename": "beyond1q",
    "model": "SM-G973U"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S10",
    "codename": "beyond1q",
    "model": "SM-G973U1"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S10",
    "codename": "beyond1q",
    "model": "SM-G973W"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S10e",
    "codename": "beyond0",
    "model": "SM-G970F"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S10e",
    "codename": "beyond0",
    "model": "SM-G970N"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S10e",
    "codename": "beyond0q",
    "model": "SM-G9700"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S10e",
    "codename": "beyond0q",
    "model": "SM-G9708"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S10e",
    "codename": "beyond0q",
    "model": "SM-G970U"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S10e",
    "codename": "beyond0q",
    "model": "SM-G970U1"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S10e",
    "codename": "beyond0q",
    "model": "SM-G970W"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S10+",
    "codename": "SC-04L",
    "model": "SC-04L"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S10+",
    "codename": "SCV42",
    "model": "SCV42"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S10+",
    "codename": "beyond2",
    "model": "SM-G975F"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S10+",
    "codename": "beyond2",
    "model": "SM-G975N"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S10+",
    "codename": "beyond2q",
    "model": "SM-G9750"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S10+",
    "codename": "beyond2q",
    "model": "SM-G9758"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S10+",
    "codename": "beyond2q",
    "model": "SM-G975U"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S10+",
    "codename": "beyond2q",
    "model": "SM-G975U1"
  },
  {
    "manufacturer": "Samsung",
    "market_name": "Galaxy S10+",
    "codename": "beyond2q",
    "model": "SM-G975W"
  },
  {
    "manufacturer": "Nokia",
    "market_name": "Nokia 3.1 Plus",
    "codename": "RHD",
    "model": "Nokia 3.1 Plus"
  },
  {
    "manufacturer": "Nokia",
    "market_name": "Nokia 3.1 Plus",
    "codename": "ROO",
    "model": "Nokia 3.1 Plus"
  },
  {
    "manufacturer": "Nokia",
    "market_name": "Nokia 3.1 Plus",
    "codename": "ROON_sprout",
    "model": "Nokia 3.1 Plus"
  },
  {
    "manufacturer": "Nokia",
    "market_name": "Nokia 3.1 Plus",
    "codename": "ROO_sprout",
    "model": "Nokia 3.1 Plus"
  },
  {
    "manufacturer": "Nokia",
    "market_name": "Nokia 7.1",
    "codename": "CTL_sprout",
    "model": "Nokia 7.1"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia 1",
    "codename": "802SO",
    "model": "802SO"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia 1",
    "codename": "J8110",
    "model": "J8110"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia 1",
    "codename": "J8170",
    "model": "J8170"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia 1",
    "codename": "J9110",
    "model": "J9110"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia 1",
    "codename": "J9180",
    "model": "J9180"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia 1",
    "codename": "SO-03L",
    "model": "SO-03L"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia 1",
    "codename": "SOV40",
    "model": "SOV40"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia 10",
    "codename": "I3113",
    "model": "I3113"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia 10",
    "codename": "I3123",
    "model": "I3123"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia 10",
    "codename": "I4113",
    "model": "I4113"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia 10",
    "codename": "I4193",
    "model": "I4193"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia 10 Plus",
    "codename": "I3213",
    "model": "I3213"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia 10 Plus",
    "codename": "I3223",
    "model": "I3223"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia 10 Plus",
    "codename": "I4213",
    "model": "I4213"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia 10 Plus",
    "codename": "I4293",
    "model": "I4293"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia XZ3",
    "codename": "801SO",
    "model": "801SO"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia XZ3",
    "codename": "H8416",
    "model": "H8416"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia XZ3",
    "codename": "H9436",
    "model": "H9436"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia XZ3",
    "codename": "H9493",
    "model": "H9493"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia XZ3",
    "codename": "SO-01L",
    "model": "SO-01L"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia XZ3",
    "codename": "SOV39",
    "model": "SOV39"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia XZ2",
    "codename": "702SO",
    "model": "702SO"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia XZ2",
    "codename": "H8216",
    "model": "H8216"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia XZ2",
    "codename": "H8266",
    "model": "H8266"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia XZ2",
    "codename": "H8276",
    "model": "H8276"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia XZ2",
    "codename": "H8296",
    "model": "H8296"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia XZ2",
    "codename": "SO-03K",
    "model": "SO-03K"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia XZ2",
    "codename": "SOV37",
    "model": "SOV37"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia XZ2 Premium",
    "codename": "H8116",
    "model": "H8116"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia XZ2 Premium",
    "codename": "H8166",
    "model": "H8166"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia XZ2 Premium",
    "codename": "SO-04K",
    "model": "SO-04K"
  },
  {
    "manufacturer": "Sony",
    "market_name": "Xperia XZ2 Premium",
    "codename": "SOV38",
    "model": "SOV38"
  }
]
`

var Devices []entities.AndroidDevice

func init() {
	_ = json.Unmarshal([]byte(devicesJson), &Devices)
}

