package main

var emotes_dictionary = []string {
	"4Head",
	"8-)",
	":(",
	":(",
	":)",
	":-(",
	":-)",
	":-/",
	":-D",
	":-O",
	":-P",
	":-Z",
	":-\\",
	":-o",
	":-p",
	":-z",
	":-|",
	":/",
	":/",
	":D",
	":D",
	":O",
	":O",
	":P",
	":P",
	":Z",
	":\\",
	":o",
	":p",
	":z",
	":|",
	":|",
	";)",
	";)",
	";-)",
	";-P",
	";-p",
	";P",
	";P",
	";p",
	"<3",
	"<3",
	">(",
	">(",
	"ANELE",
	"ArgieB8",
	"ArsonNoSexy",
	"AsexualPride",
	"AsianGlow",
	"B)",
	"B)",
	"B-)",
	"BCWarrior",
	"BOP",
	"BabyRage",
	"BatChest",
	"BegWan",
	"BibleThump",
	"BigBrother",
	"BigPhish",
	"BisexualPride",
	"BlackLivesMatter",
	"BlargNaut",
	"BloodTrail",
	"BrainSlug",
	"BrokeBack",
	"BuddhaBar",
	"CaitlynS",
	"CarlSmile",
	"ChefFrank",
	"ChewyYAY",
	"CoolCat",
	"CoolStoryBob",
	"CorgiDerp",
	"CrreamAwk",
	"CurseLit",
	"DAESuppy",
	"DBstyle",
	"DansGame",
	"DarkKnight",
	"DarkMode",
	"DatSheffy",
	"DendiFace",
	"DinoDance",
	"DogFace",
	"DoritosChip",
	"DxCat",
	"EarthDay",
	"EleGiggle",
	"EntropyWins",
	"ExtraLife",
	"FBBlock",
	"FBCatch",
	"FBChallenge",
	"FBPass",
	"FBPenalty",
	"FBRun",
	"FBSpiral",
	"FBtouchdown",
	"FUNgineer",
	"FailFish",
	"FallCry",
	"FallHalp",
	"FallWinning",
	"FamilyMan",
	"FlawlessVictory",
	"FootBall",
	"FootGoal",
	"FootYellow",
	"ForSigmar",
	"FrankerZ",
	"FreakinStinkin",
	"FutureMan",
	"GayPride",
	"GenderFluidPride",
	"Getcamped",
	"GingerPower",
	"GivePLZ",
	"GlitchCat",
	"GlitchLit",
	"GlitchNRG",
	"GoatEmotey",
	"GoldPLZ",
	"GrammarKing",
	"GunRun",
	"HSCheers",
	"HSWP",
	"HarleyWink",
	"HassaanChop",
	"HeyGuys",
	"HolidayCookie",
	"HolidayLog",
	"HolidayPresent",
	"HolidaySanta",
	"HolidayTree",
	"HotPokket",
	"HungryPaimon",
	"ImTyping",
	"IntersexPride",
	"InuyoFace",
	"ItsBoshyTime",
	"JKanStyle",
	"Jebaited",
	"Jebasted",
	"JonCarnage",
	"KAPOW",
	"KEKHeim",
	"Kappa",
	"Kappa",
	"KappaClaus",
	"KappaPride",
	"KappaRoss",
	"KappaWealth",
	"Kappu",
	"Keepo",
	"KevinTurtle",
	"Kippa",
	"KomodoHype",
	"KonCha",
	"Kreygasm",
	"LUL",
	"LaundryBasket",
	"Lechonk",
	"LesbianPride",
	"LionOfYara",
	"MVGame",
	"Mau5",
	"MaxLOL",
	"MechaRobot",
	"MercyWing1",
	"MercyWing2",
	"MikeHogu",
	"MingLee",
	"ModLove",
	"MorphinTime",
	"MrDestructoid",
	"MyAvatar",
	"NewRecord",
	"NiceTry",
	"NinjaGrumpy",
	"NomNom",
	"NonbinaryPride",
	"NotATK",
	"NotLikeThis",
	"O.O",
	"O.o",
	"OSFrog",
	"O_O",
	"O_o",
	"O_o",
	"OhMyDog",
	"OneHand",
	"OpieOP",
	"OptimizePrime",
	"PJSalt",
	"PJSugar",
	"PMSTwin",
	"PRChase",
	"PanicVis",
	"PansexualPride",
	"PartyHat",
	"PartyTime",
	"PeoplesChamp",
	"PermaSmug",
	"PicoMause",
	"PikaRamen",
	"PinkMercy",
	"PipeHype",
	"PixelBob",
	"PizzaTime",
	"PogBones",
	"PogChamp",
	"Poooound",
	"PopCorn",
	"PopNemo",
	"PoroSad",
	"PotFriend",
	"PowerUpL",
	"PowerUpR",
	"PraiseIt",
	"PrimeMe",
	"PunOko",
	"PunchTrees",
	"R)",
	"R)",
	"R-)",
	"RaccAttack",
	"RalpherZ",
	"RedCoat",
	"ResidentSleeper",
	"RitzMitz",
	"RlyTho",
	"RuleFive",
	"RyuChamp",
	"SMOrc",
	"SSSsss",
	"SUBprise",
	"SabaPing",
	"SeemsGood",
	"SeriousSloth",
	"ShadyLulu",
	"ShazBotstix",
	"Shush",
	"SingsMic",
	"SingsNote",
	"SmoocherZ",
	"SoBayed",
	"SoonerLater",
	"Squid1",
	"Squid2",
	"Squid3",
	"Squid4",
	"StinkyCheese",
	"StinkyGlitch",
	"StoneLightning",
	"StrawBeary",
	"SuperVinlin",
	"SwiftRage",
	"TBAngel",
	"TF2John",
	"TPFufun",
	"TPcrunchyroll",
	"TTours",
	"TakeNRG",
	"TearGlove",
	"TehePelo",
	"ThankEgg",
	"TheIlluminati",
	"TheRinger",
	"TheTarFu",
	"TheThing",
	"ThunBeast",
	"TinyFace",
	"TombRaid",
	"TooSpicy",
	"TransgenderPride",
	"TriHard",
	"TwitchConHYPE",
	"TwitchLit",
	"TwitchRPG",
	"TwitchSings",
	"TwitchUnity",
	"TwitchVotes",
	"UWot",
	"UnSane",
	"UncleNox",
	"VirtualHug",
	"VoHiYo",
	"VoteNay",
	"VoteYea",
	"WTRuck",
	"WholeWheat",
	"WhySoSerious",
	"WutFace",
	"YouDontSay",
	"YouWHY",
	"bleedPurple",
	"cmonBruh",
	"copyThis",
	"duDudu",
	"imGlitch",
	"mcaT",
	"o.O",
	"o.o",
	"o_O",
	"o_o",
	"panicBasket",
	"pastaThat",
	"riPepperonis",
	"twitchRaid",

	// RUBIUS

	"rubb2IQ",
	"rubb600IQ",
	"rubbAFK",
	"rubbAYAYA",
	"rubbAngry",
	"rubbAsco",
	"rubbBed",
	"rubbBless",
	"rubbBobo",
	"rubbC",
	"rubbCARRIED",
	"rubbCam",
	"rubbCi",
	"rubbCiego",
	"rubbComfy",
	"rubbDehecho",
	"rubbDrive",
	"rubbDross",
	"rubbDu",
	"rubbExe",
	"rubbFBI",
	"rubbFL",
	"rubbFOCUS",
	"rubbFeels",
	"rubbGasm",
	"rubbGun",
	"rubbH",
	"rubbHart",
	"rubbHeCrazy",
	"rubbJoy",
	"rubbKEKW",
	"rubbKhe",
	"rubbL",
	"rubbLeep",
	"rubbLof",
	"rubbMIEDO",
	"rubbMalo",
	"rubbN",
	"rubbNice",
	"rubbNo",
	"rubbNotStnks",
	"rubbOld",
	"rubbOso",
	"rubbPaw",
	"rubbPhone",
	"rubbPlebs",
	"rubbPls",
	"rubbPog",
	"rubbPrime",
	"rubbRAGE",
	"rubbRare",
	"rubbRaspy",
	"rubbRatatopo",
	"rubbReee",
	"rubbRich",
	"rubbSad",
	"rubbSchizo",
	"rubbSeh",
	"rubbSmash",
	"rubbTacos",
	"rubbTexto",
	"rubbTonks",
	"rubbToxic",
	"rubbVape",
	"rubbW",
	"rubbWc",
	"rubbWeird",
	"rubbWilson",
	"rubbSilver",
	"rubbGold",
	"rubbXddddddddd",
	"rubbComo",
	"rubbOof",

	// POPULARES

	"OMEGALUL",
	"monkaMEGA",
	"Pog",
	"peepoSad",
	"PepeHands",
	"Prayge ",
	"widepeepoFeliz",
	"HasMuerto",
	"LULW",
	"pikachuS",
	"Pepega",
	"peepoClown",
	"KEKW",
	"COPIUM",
	"monkaW",
	"WICKED",
	"5cabeza",
	"AnchoDuro",
	"POGGERS",
	"Madge",
	"monkaS",
	"PepegaTeléfono",
	"AYAYA",
	"peepoWTF",
	"TooLewd",
	"WeirdChamp",
	"widepeepoSad",
	"POGGIES",
	"Sadge",
	"peepoGrasa",
	"FeelsBadMan",
	"MEGALUL",
	"HYPERS",
	"PulseF",
	"YEP",
	"Pepepains",
	"FeelsOkayMan",
	"PauseChamp",
	"Kappa",
	"Okayge",
	"EZY",
	"widepeepoFeliz",
	"HandsUp",
	"HeavyBreathing",
	"PepoPiensa en",
	"4Raro",
	"3cabeza",
	"Thonk",
	"gachiGASM",
	"TRIGGERED",
	"FeelsWeirdMan",
	"widepeepoManta",
	"monkaHmm",
	"Klappa",
	"forsenCD",
	"WaitWhat",
	"KKomrade",
	"PeepoGlad ",
	"PepoG ",
	"Pogey",
	"monkaOMEGA ",
	"FeelsWowMan",
	"monkaGun",
	"4Casa",
	"peepoManta",
	"SmileW",
	"FeelsStrongMan ",
	"monkaCristo",
	"monkaEyes",
	"pepeJAM",
	"HYPERBRUH ",
	"¡GG!",
	"peepoFeliz",
	"peepoPoo",
	"monkaTOS",
	"POGGERS",
	"peepoLove",
	"Piedras",
	"weSmart",
	"PagMan",
	"FeelsGoodMan",
	"Hmm",
	"KKonaW",
	"Hahaa",
	"FeelsDankMan",
	"gachiHYPER",
	"REEeee",
	"SadCatW",
	"monkaH",
	"HYPERDANSGAME",
	"4Punta",
	"sadKEK",
	"FeelsSpecialMan",
	"ThisIsFine",
	"KEKWait",
	"2cabeza",
	"PepeLaugh",
	"monkaGIGA",
	"WideHardo",
	"SuchMeme",
	"4Encogerse de hombros",
	"widepeepoAbrazo",
	"pepePoint",
	"JUSTDOIT",
	"MaN",
	"gachiBASS",
	"EHEHE",
	"VWasted",
	"peepoAbrazo",
	"eShrug",
	"Clap",
	"Bruh",
	"PepeLmao",
	"OkayChamp",
	"stopS ",
	"MikuStare",
	"POOGERS",
	"FeelsBirthdayMan",
	"Bedge",
	"KannaNom",
	"KKona",
	"RAGEY",
	"StrimPlz",
	"peepoShy",
	"monkaVelocidad",
	"PLEASENO",
	"FacePalm",
	"4RaroW",
	"CatWhat",
	"WTFF",
	"okiRIP ",
	"PepeClown",
	"FeelsAmazingMan",
	"FeelsWeirdManW ",
	"KEKL",
	"AlmohadaNo",
	"pepeGun",
	"AlmohadaSí",
	"Susge",
	"peepoFinger",
	"HYPERS",
	"Tuturu",
	"nessiepls",
	"alienpls",
	"catmunch",
	"popnemo",
	"pokcroagunk",
	"peepogift",
	"roierdance",
	"peepocomfy",
	"huh",
	"tokyodchipichipi",
	"uuh",
	"euh",
	"booba",
	"boobapeeking",
	"hmmm",
	"boobest",
	"peepocheer",
	"hmm",
	"enough",
	"tole",
	"peepostripper",
	"chinacat",
	"obsesstriste",
	"nizagouwu",
	"fullwineo",
	"nyanpls",
	"pls",
	"ppoverheat",
	"weebrun",
	"blobhypers",
	"clueless",
	"rubbstare",
	"wideduckass",
	"galaxyunpacked",
	"neondance",
	"squid",
	"dinkdonk",
	"yuzz",
	"maryblogdance",
	"wolfracelebrate",
	"miouvtunn",
	"winttewiggle",
	"saxtime",
	"yuzz",
	"catdisco",
	"vivirtop",
}