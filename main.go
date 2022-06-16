package main

import (
	"bufio"
	_ "embed"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/anaskhan96/soup"
	"github.com/mattn/go-tty"
)

// Holds a collection of types from Phenom dropfile
type PhenomDrop struct {
	Node          int
	BbsName       string
	UserName      string
	SysopName     string
	SecLevel      int
	TimeLeft      int
	Cols          int
	Rows          int
	OsType        string
	BbsDir        string
	BbsDomain     string
	LoadableFonts bool
	XtendPalette  bool
}

const (
	Esc         = "\u001B["
	Osc         = "\u001B]"
	Bel         = "\u0007"
	EraseScreen = Esc + "2J"
	Idle        = 120

	Reset     = Esc + "0m"
	Black     = Esc + "30m"
	Red       = Esc + "31m"
	Green     = Esc + "32m"
	Yellow    = Esc + "33m"
	Blue      = Esc + "34m"
	Magenta   = Esc + "35m"
	Cyan      = Esc + "36m"
	White     = Esc + "37m"
	BlackHi   = Esc + "30;1m"
	RedHi     = Esc + "31;1m"
	GreenHi   = Esc + "32;1m"
	YellowHi  = Esc + "33;1m"
	BlueHi    = Esc + "34;1m"
	MagentaHi = Esc + "35;1m"
	CyanHi    = Esc + "36;1m"
	WhiteHi   = Esc + "37;1m"

	BgBlack     = Esc + "40m"
	BgRed       = Esc + "41m"
	BgGreen     = Esc + "42m"
	BgYellow    = Esc + "43m"
	BgBlue      = Esc + "44m"
	BgMagenta   = Esc + "45m"
	BgCyan      = Esc + "46m"
	BgWhite     = Esc + "47m"
	BgBlackHi   = Esc + "40;1m"
	BgRedHi     = Esc + "41;1m"
	BgGreenHi   = Esc + "42;1m"
	BgYellowHi  = Esc + "43;1m"
	BgBlueHi    = Esc + "44;1m"
	BgMagentaHi = Esc + "45;1m"
	BgCyanHi    = Esc + "46;1m"
	BgWhiteHi   = Esc + "47;1m"
)

var (
	Pd       PhenomDrop
	DropPath string
)

// NewTimer boots a user after being idle too long
func NewTimer(seconds int, action func()) *time.Timer {
	timer := time.NewTimer(time.Second * time.Duration(seconds))

	go func() {
		<-timer.C
		action()
	}()
	return timer
}

// Move cursor to X, Y location
func MoveCursor(x int, y int) {
	fmt.Printf(Esc+"%d;%df", y, x)
}

// Erase the screen
func ClearScreen() {
	fmt.Print(EraseScreen)
	MoveCursor(0, 0)
}

// Returns all values as strings
func DropFileData(path string) (string, string, string, string, string, string, string, string, string, string, string, string, string) {
	var node string
	var bbsname string
	var username string
	var sysopname string
	var seclevel string
	var timeleft string
	var cols string
	var rows string
	var ostype string
	var bbsdir string
	var bbsdomain string
	var loadablefonts string
	var xtendpalette string

	file, err := os.Open(strings.ToLower(path + "/phenomdrop.txt"))
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var text []string

	for scanner.Scan() {
		text = append(text, scanner.Text())
	}

	count := 1
	for _, line := range text {
		if count == 1 {
			node = line
		}
		if count == 2 {
			bbsname = line
		}
		if count == 3 {
			sysopname = line
		}
		if count == 4 {
			username = line
		}
		if count == 5 {
			seclevel = line
		}
		if count == 6 {
			timeleft = line
		}
		if count == 7 {
			cols = line
		}
		if count == 8 {
			rows = line
		}
		if count == 9 {
			ostype = line
		}
		if count == 10 {
			bbsdir = line
		}
		if count == 11 {
			bbsdomain = line
		}
		if count == 12 {
			loadablefonts = line
		}
		if count == 13 {
			xtendpalette = line
		}
		if count == 14 {
			break
		}
		count++
		continue
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return node, bbsname, username, sysopname, seclevel, timeleft, cols, rows, ostype, bbsdir, bbsdomain, loadablefonts, xtendpalette
}

// Print text at an X, Y location
func PrintStringLoc(text string, x int, y int) {
	yLoc := y
	s := bufio.NewScanner(strings.NewReader(text))
	for s.Scan() {
		fmt.Fprintf(os.Stdout, Esc+strconv.Itoa(yLoc)+";"+strconv.Itoa(x)+"f"+s.Text())
		yLoc++
	}
}

func getNumEnding() string {

	dayStr := time.Now().Day()

	if dayStr-1 == 1 && len(fmt.Sprint(dayStr)) == 1 {
		return "st"
	} else if dayStr-1 == 2 && dayStr-2 != 12 {
		return "nd"
	} else if dayStr-1 == 3 {
		return "rd"
	} else {
		return "th"
	}
}

func generateEventList() {

	day := time.Now().Day()
	month := time.Now().Month()
	year := time.Now().Year()
	current_time := time.Now()

	ClearScreen()

	resp, err := soup.Get("https://www.timeanddate.com/on-this-day")
	if err != nil {
		os.Exit(1)
	}

	fmt.Print("\r\n " + BlackHi + Reset + "-" + Cyan + "---" + GreenHi + "-" + Reset + Cyan + "--" + GreenHi + "-" + Reset + Cyan + "-" + GreenHi + "--------- ------------------------------------ ------ -- -  " + Reset)

	fmt.Print("\r\n " + BgGreen + WhiteHi + ">> " + GreenHi + "Glimpse In Time v1.1  " + Reset + BgGreen + Black + ">>" + BgBlack + Green + ">>  " + Reset + WhiteHi + "by " + CyanHi + "Smooth " + Reset + Cyan + "<" + WhiteHi + "PHEN0M" + Reset + Cyan + ">" + Reset)
	fmt.Print("\r\n " + BlackHi + "-" + Reset + Cyan + "--" + GreenHi + "--" + Reset + Cyan + "---" + GreenHi + "-" + Reset + Cyan + "-" + GreenHi + "----- --- -------------------------------- ------ -- -  " + Reset)

	fmt.Printf("\r\n "+BgRed+Black+">>"+BgBlack+" "+MagentaHi+"On "+Reset+YellowHi+"THIS DAY"+MagentaHi+", These "+YellowHi+"EVENTS "+MagentaHi+"Happened... "+Reset+Red+":: "+Yellow+" %v %v%v "+Red+" ::"+Reset, month, day, getNumEnding())
	fmt.Print("\r\n " + BlackHi + "-" + Reset + Cyan + "--" + GreenHi + "--" + Reset + Cyan + "---" + GreenHi + "-" + Reset + Cyan + "-" + GreenHi + "--" + Reset + Cyan + "--- " + GreenHi + "--- ---------------------------- ------ -- -  " + Reset)

	MoveCursor(1, 8)
	index := 1
	doc := soup.HTMLParse(resp)
	events := doc.FindAllStrict("h3", "class", "otd-title")
	for _, event := range events {
		fmt.Print(" " + CyanHi + event.Find("strong").Text() + Reset + Cyan + " <" + BlackHi + ":" + Reset + Cyan + "> " + WhiteHi + strings.TrimSpace(event.Text()) + Reset + "\r\n\r\n")

		if index == 5 {
			break
		}
		index++
	}

	fmt.Print(" " + BlackHi + "-" + Reset + Cyan + "---" + GreenHi + "-" + Reset + Cyan + "--" + GreenHi + "-" + Reset + Cyan + "-" + GreenHi + "-----" + Reset + Cyan + "-" + GreenHi + "--------------------------------------- ---  --- -- -  " + Reset)
	fmt.Printf("\r\n "+BgRed+Black+">>"+BgBlack+" "+WhiteHi+"Generated on %v %v, %v at %v", month, day, year, current_time.Format("3:4 PM"))
	fmt.Print("\r\n " + BlackHi + "-" + Reset + Cyan + "---" + GreenHi + "-" + Reset + Cyan + "--" + GreenHi + "-" + Reset + Cyan + "-" + GreenHi + "-----" + Reset + Cyan + "-" + GreenHi + "--------------------------------------- ---  --- -- -  " + Reset)

	MoveCursor(0, 23)
	fmt.Print("                   " + BgBlueHi + WhiteHi + "<" + Reset + Cyan + "<  " + BlackHi + "... " + Reset + White + "press " + WhiteHi + "ANY KEY " + Reset + White + "to " + WhiteHi + "CONTINUE " + Reset + BlackHi + "... " + Reset + Cyan + ">" + BgBlue + WhiteHi + ">" + Reset)

}

func init() {
	// Use FLAG to get command line paramenters
	pathPtr := flag.String("path", "", "path to node directory")
	required := []string{"path"}

	flag.Parse()

	seen := make(map[string]bool)
	flag.Visit(func(f *flag.Flag) { seen[f.Name] = true })
	for _, req := range required {
		if !seen[req] {
			// or possibly use `log.Fatalf` instead of:
			fmt.Fprintf(os.Stderr, "missing path to node directory, e.g.: ./phenomdroptest -%s /bbs/temp/1 \n", req)
			os.Exit(2) // the same exit code flag.Parse uses
		}
	}

	// read the drop file and save to struct
	DropPath = *pathPtr
	node, bbsname, username, sysopname, seclevel, timeleft, cols, rows, ostype, bbsdir, bbsdomain, loadablefonts, xtendpalette := DropFileData(DropPath)

	// convert some values to int or bool
	intnode, _ := strconv.Atoi(node)
	intcols, _ := strconv.Atoi(cols)
	introws, _ := strconv.Atoi(rows)
	intseclevel, _ := strconv.Atoi(seclevel)
	inttimeleft, _ := strconv.Atoi(timeleft)
	boolloadablefonts, _ := strconv.ParseBool(loadablefonts)
	boolxtendpalette, _ := strconv.ParseBool(xtendpalette)

	// asign to struct
	Pd = PhenomDrop{
		Node:          intnode,
		BbsName:       bbsname,
		UserName:      username,
		SysopName:     sysopname,
		SecLevel:      intseclevel,
		TimeLeft:      inttimeleft,
		Cols:          intcols,
		Rows:          introws,
		OsType:        ostype,
		BbsDir:        bbsdir,
		BbsDomain:     bbsdomain,
		LoadableFonts: boolloadablefonts,
		XtendPalette:  boolxtendpalette,
	}
}

func main() {
	// Start the idle timer
	shortTimer := NewTimer(Idle, func() {
		fmt.Println("\r\nYou've been idle for too long... exiting!")
		time.Sleep(1 * time.Second)
		os.Exit(0)
	})
	defer shortTimer.Stop()

	ClearScreen()
	MoveCursor(0, 0)

	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	for {

		generateEventList()
		_, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}
		os.Exit(0)

	}
}
