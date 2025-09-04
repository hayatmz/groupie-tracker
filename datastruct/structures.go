package datastruct

var Videos = []string{
// Queen - Bohemian Rhapsody
"https://www.youtube.com/embed/fJ9rUzIMcZQ?si=nwATmj8KRXRcMP2i",
// SOJa - Rest of my life
"https://www.youtube.com/embed/X572Mp_r46E?si=7V5WIOu97aWF2rwK",
// Pink Floyd - another brick in the wall
"https://www.youtube.com/embed/5IpYOF4Hi6Q?si=JbDlfgt5BsyjDIh6",
// scorpions - Wind of change
"https://www.youtube.com/embed/n4RjJKxsamQ?si=TAOLILyMe1CGLmSi",
// xxxtentacion - changes
"https://www.youtube.com/embed/f0bbDFRYD_A?si=jkLYNPhubm-vN3GO",
// mac miller - yeah
"https://www.youtube.com/embed/qPRKc5BDe8A?si=nfSsXUfAN1lqYz_E",
// joyner lucas - 24 hours to live
"https://www.youtube.com/embed/xFdfTN-Wffo?si=I9o7O50KuxMFwu12",
// kendreick lamar - rich spirit
"https://www.youtube.com/embed/toBTPGfurLc?si=bCKUO6dUfOi41hHR",
// ACDC - Thunderstruck
"https://www.youtube.com/embed/v2AC41dglnM?si=vB3nCqCh9WX2ADFr",
// Pearl Jam - Alive
"https://www.youtube.com/embed/qM0zINtulhM?si=FXf64VfC2yef6eTn",
// Katy Perry - Firework
"https://www.youtube.com/embed/QGJuMBdaqIw?si=iRoy3SC6_ku-p0Nr",
// Rihanna - Man Down
"https://www.youtube.com/embed/sEhy-RXkNo0?si=LF9aNoyh_nKUzXkO",
// Genesis - Jesus He Knows Me
"https://www.youtube.com/embed/35K6vQRt67g?si=CxS-8AiZ6kMRwUHh",
// Phil Collins - In the air tonight
"https://www.youtube.com/embed/YkADj0TPrJA?si=ds9nbJ2XFvoy9zy_",
// Led Zeppelin - whole lotta love
"https://www.youtube.com/embed/HQmmM_qwG4k?si=TEhGek17r3njuRHk",
// The Jimi Hendrix Experience - hey joe
"https://www.youtube.com/embed/gUPifXX0foU?si=Bebar2NaDKSm6_Nj",
// bee gee - stayin' alive
"https://www.youtube.com/embed/fNFzfwLM72c?si=_YrHL3p0DZliL2yq",
// deep purple - smoke on the water
"https://www.youtube.com/embed/eu5lv2Umn3M?si=N4avzAed0wWfGz62",
// aerosmith - cryin'
"https://www.youtube.com/embed/qfNmyxV2Ncw?si=M_ijakHKfU-Nr-rj",
// dire straits - sultans of swing
"https://www.youtube.com/embed/h0ffIJ7ZO4U?si=I-y_WBWtpcw1YHU1",
// mamonas assassinas - vira-vira
"https://www.youtube.com/embed/1WjI3DLOk4c?si=N46PA_WuhV5_ECPK",
// Thirty seconds to mars - Stuck
"https://www.youtube.com/embed/di-VTrW7Kr0?si=wVE3gwTpLS7vHyJH",
// Imagine Dragons - Children of the sky
"https://www.youtube.com/embed/zHfIkyLh-Ew?si=yZ7D_m1AQFfqKce1",
// Juice WRLD - Glo’d UpJuice WRLD - Glo’d Up
"https://www.youtube.com/embed/EYygYkYldHw?si=D0Aogl4M3xf1Q3Jo",
// Logic - Fear
"https://www.youtube.com/embed/KMZ-JbvRzFs?si=96OFkBv6FBN1hLx3",
// Alec Benjamin - Pick Me
"https://www.youtube.com/embed/9pkZDoRyNLY?si=M_r6SbC17Ay45NsN",
// Bobby McFerrin - Don't Worry Be Happy
"https://www.youtube.com/embed/d-diB65scQU?si=taFDJF096VnJIQn0",
// R3HAB, INNA, Sash! – Rock My Body
"https://www.youtube.com/embed/aYCuEONIIDM?si=RH5c1n16i4SgeaS7",
// Post Malone ft. 21 Savage - rockstar
"https://www.youtube.com/embed/UceaB4D0jpo?si=keSTIRrRdrDeZPbs",
// Travis Scott - I KNOW ?
"https://www.youtube.com/embed/X7aF3nZOS98?si=k1QSy5a_7qRuxUZX",
// J. Cole - MIDDLE CHILD
"https://www.youtube.com/embed/WILNIXZr2oc?si=5I9piAaCHWUSJgPE",
// Nickelback - Savin' Me
"https://www.youtube.com/embed/_JQiEs32SqQ?si=B2NBllwSKvzMMSF4",
// Mobb Deep - Survival of the Fittest
"https://www.youtube.com/embed/Dz5VzLz67WA?si=tLBIdrXNAH-rn4_6",
// Guns N' Roses - November Rain
"https://www.youtube.com/embed/8SbUC-UaAxE?si=mOfPgLPKP-oPPgkR",
// N.W.A. - Straight Outta Compton
"https://www.youtube.com/embed/TMZi25Pq3T8?si=5oIOGZVQjLe23Gcc",
// U2 - With Or Without You
"https://www.youtube.com/embed/ujNeHIo7oTE?si=aMQ6M1mTQtqPB7Eo",
// Arctic Monkeys - Do I Wanna Know?
"https://www.youtube.com/embed/bpOSxM0rNPM?si=E4lKvwZa8Y3c05uN",
// Fall Out Boy - Thnks fr th Mmrs
"https://www.youtube.com/embed/onzL0EM1pKY?si=vKukogKFEelNMrVO",
// Gorillaz - She's My Collar
"https://www.youtube.com/embed/f8NwLXYIHS4?si=CJkIpRoLgkBI1HQX",
// Eagles - Tequila Sunrise
"https://www.youtube.com/embed/bZxhQJC9hWk?si=VOojqGkwqf5unXIF",
// Linkin Park - In The End
"https://www.youtube.com/embed/eVTXPUF4Oz4?si=Xrd0fEXDAjF3ibIz",
// Red Hot Chili Peppers - The Shape I'm Takin'
"https://www.youtube.com/embed/Lqf406eSorw?si=rzphOrqNRCxEwlHt",
// Eminem - The Real Slim Shady
"https://www.youtube.com/embed/eJO5HU_7_1w?si=-F1ysMkB7KH4blT8",
// Green Day - Fancy Sauce
"https://www.youtube.com/embed/yqPVq9DZYvM?si=clfYkrOZ_pbeCAZb",
// Metallica: Enter Sandman
"https://www.youtube.com/embed/CD-E-LDc384?si=FHR_Ce1FuyZxe4th",
// Coldplay - Adventure Of A Lifetime
"https://www.youtube.com/embed/QtXby3twMmI?si=t_PLzZB9JCUJxUbn",
// Maroon 5 - Girls Like You ft. Cardi B
"https://www.youtube.com/embed/aJOTlE1K90k?si=kkN4f_ErYCqKkaBl",
// twenty one pilots: Stressed Out
"https://www.youtube.com/embed/pXRviuL6vMY?si=CtYC_R0QZzbgFrhO",
// The Rolling Stones - Angry
"https://www.youtube.com/embed/_mEC54eTuGw?si=qJujsUT_eoubecUo",
// Muse - Starlight
"https://www.youtube.com/embed/Pgum6OT_VH8?si=lgXEFZ8PXeQdGhCY",
// Foo Fighters - The Pretender
"https://www.youtube.com/embed/SBjQ9tuuTJQ?si=cfUd5SNrTAuB3_Cc",
// The Chainsmokers - Don't Let Me Down
"https://www.youtube.com/embed/Io0fBr1XBUA?si=QZsyEqbNzlErM-vk",
}

type Artists struct {
	ID					int					`json:"id"`
	Image				string				`json:"image"`
	Name				string				`json:"name"`
	Members				[]string			`json:"members"`
	Creation			int					`json:"creationDate"`
	FirstAlbum			string				`json:"firstAlbum"`
	LocationsRecup 		[]string
	ConcertDatesRecup 	[]string
	RelationsRecup  	map[string][]string
	Link				string
}

type IndexLocations struct {
	ID			int			`json:"id"`
	Locations 	[]string 	`json:"locations"`
}

type Locations struct {
	Index []IndexLocations 
}

type IndexDates struct {
	ID		int 		`json:"id"`
	Dates	[]string 	`json:"dates"`
}

type ConcertDates struct {
	Index []IndexDates
}

type IndexRelations struct {
	ID 				int 					`json:"id"`
	DatesLocations 	map[string][]string 	`json:"datesLocations"`
}

type Relations struct {
	Index []IndexRelations
}

// Pour renvoyer à la page HTML.
type PageData struct {
    ArtistHTML    []Artists
}