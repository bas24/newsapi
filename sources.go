package newsapi

var sources = map[string]string{
	"abc-news-au":             "general",       // www.abc.net.au | general
	"ars-technica":            "technology",    // www.arstechnica.com | technology
	"associated-press":        "general",       // www.bigstory.ap.org | general
	"bbc-news":                "general",       // www.bbc.co.uk | general
	"bbc-sport":               "sport",         // www.bbc.co.uk/sport | sport
	"bloomberg":               "business",      // www.bloomberg.com | business
	"business-insider":        "business",      // www.businessinsider.com | business
	"business-insider-uk":     "business",      // www.uk.businessinsider.com | business
	"cnbc":                    "business",      // www.cnbc.com | business
	"cnn":                     "general",       // www.us.cnn.com | general
	"daily-mail":              "entertainment", // www.i.dailymail.co.uk | entertainment
	"engadget":                "technology",    // www.engadget.com | technology
	"entertainment-weekly":    "entertainment", // www.ew.com | entertainment
	"espn":                    "sport",         // www.espn.go.com | sport
	"espn-cric-info":          "sport",         // www.espncricinfo.com | sport
	"football-italia":         "sport",         // www.football-italia.net | sport
	"fortune":                 "business",      // www.fortune.com | business
	"four-four-two":           "sport",         // www.fourfourtwo.com | sport
	"fox-sports":              "sport",         // www.foxsports.com | sport
	"google-news":             "general",       // aggregator | general
	"hacker-news":             "technology",    // aggregator | technology
	"ign":                     "gaming",        // www.ign.com | gaming
	"independent":             "general",       // www.independent.co.uk | general
	"mashable":                "entertainment", // www.mashable.com | entertainment
	"metro":                   "general",       // www.metro.co.uk | general
	"mirror":                  "general",       // www.mirror.co.uk | general
	"mtv-news":                "music",         // www.mtv.com | music
	"mtv-news-uk":             "music",         // www.mtv.co.uk | music
	"national-geographic":     "science",       // www.news.nationalgeographic.com | science-and-nature
	"new-scientist":           "science",       // www.newscientist.com | science-and-nature
	"newsweek":                "general",       // www.newsweek.com | general
	"new-york-magazine":       "general",       // aggregator | general
	"nfl-news":                "sport",         // www.nfl.com | sport
	"polygon":                 "gaming",        // www.polygon.com | gaming
	"recode":                  "technology",    // www.recode.net | technology
	"reddit-r-all":            "general",       // aggregator | general
	"reuters":                 "general",       // www.feeds.reuters.com | general
	"sky-news":                "general",       // www.news.sky.com | general
	"sky-sports-news":         "sport",         // www.skysports.com | sport
	"talksport":               "sport",         // www.talksport.com | sport
	"techcrunch":              "technology",    // www.techcrunch.com | technology
	"techradar":               "technology",    // www.techradar.com | technology
	"the-economist":           "business",      // www.economist.com | business
	"the-guardian-au":         "general",       // www.theguardian.com/australia-news | general
	"the-guardian-uk":         "general",       // www.theguardian.com/us-news | general
	"the-hindu":               "general",       // www.thehindu.com | general
	"the-huffington-post":     "general",       // www.huffingtonpost.com | general
	"the-lad-bible":           "entertainment", // www.theladbible.com | entertainment
	"the-new-york-times":      "general",       // www.nytimes.com | general
	"the-next-web":            "technology",    // www.thenextweb.com | technology
	"the-sport-bible":         "sport",         // www.thesportbible.com | sport
	"the-telegraph":           "general",       // www.telegraph.co.uk | general
	"the-times-of-india":      "general",       // www.timesofindia.indiatimes.com | general
	"the-verge":               "technology",    // www.theverge.com | technology
	"the-wall-street-journal": "business",      // www.wsj.com | business
	"the-washington-post":     "general",       // www.washingtonpost.com | general
	"time":                    "general",       // www.time.com | general
	"usa-today":               "general",       // www.rssfeeds.usatoday.com | general
}