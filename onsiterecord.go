package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	DefaultSleepTime = 100
)

type UserData struct {
	DtRegister           string  `json:"DtRegister"`
	GroupCNPJ            string  `json:"GroupCNPJ"`
	CNPJ                 string  `json:"CNPJ"`
	EmpID                int64   `json:"EmpID"`
	EmpName              string  `json:"EmpName"`
	EmpPhone             string  `json:"EmpPhone"`
	EmpLoginID           int64   `json:"EmpLoginID"`
	EmpLogin             string  `json:"EmpLogin"`
	EmpPass              string  `json:"EmpPass"`
	EmpLoginPhone        string  `json:"EmpLoginPhone"`
	EmpSID               string  `json:"EmpSID"`
	ActiveLogin          bool    `json:"ActiveLogin"`
	LoginMsg             string  `json:"LoginMsg"`
	TypeLogin            string  `json:"TypeLogin"`
	EmpCPF               string  `json:"EmpCPF"`
	EmpRG                string  `json:"EmpRG"`
	EmpEmissorRG         string  `json:"EmpEmissorRG"`
	EmpDtEmissaoRG       string  `json:"EmpDtEmissaoRG"`
	EmpStateRG           string  `json:"EmpStateRG"`
	EmpDtBirth           string  `json:"EmpDtBirth"`
	Gender               string  `json:"Gender"`
	Card                 int     `json:"Card"`
	FormattedCard        string  `json:"FormattedCard"`
	Chip                 string  `json:"Chip"`
	BalanceID            int64   `json:"BalanceID"`
	Value                float64 `json:"Value"`
	EstimatedBalanceDays int64   `json:"EstimatedBalanceDays"`
	DtUpdateBalance      string  `json:"DtUpdateBalance"`
	PendencyBalance      float64 `json:"PendencyBalance"`
	PendencyRecharge     int64   `json:"PendencyRecharge"`
	DtOldRecharge        string  `json:"DtOldRecharge"`
	Org                  int64   `json:"Org"`
	OrgName              string  `json:"OrgName"`
	OrgSite              string  `json:"OrgSite"`
	OrgEmail             string  `json:"OrgEmail"`
	OrgPhone             string  `json:"OrgPhone"`
	PathBalance          string  `json:"PathBalance"`
	PathRegister         string  `json:"PathRegister"`
	Instructions         string  `json:"Instructions"`
	Product              string  `json:"Product"`
	NotFound             bool    `json:"NotFound"`
	Status               int     `json:"Status"`
	MsgError             string  `json:"MsgError"`
	FileScheduled        string  `json:"FileScheduled"`
	FileRecovered        string  `json:"FileRecovered"`
	Registered           int64   `json:"Registered"`
	EmpFornecedorID      int64   `json:"EmpFornecedorID"`
}

var (
	userList   []UserData
	filepath   string
	screenData string
)

func init() {
	flag.StringVar(&filepath, "f", "", "Sets the file path for the json data")
}

func main() {
	flag.Parse()
	readjson(filepath)
	for _, userData := range userList {
		cmd("adb shell am force-stop org.mozilla.firefox", 10)
		onSiteRegistry(userData)
	}
}

func onSiteRegistry(userData UserData) error {

	if !validatePW(userData.EmpPass) {
		return fmt.Errorf("Invalid password")
	}

	var err error
	log.Println("Opening firefox in private")
	log.Println("Opening riocards registration page")
	cmd("adb shell am start -a android.activity.MAIN -n org.mozilla.firefox/org.mozilla.gecko.BrowserApp -d https://minhaconta.riocardmais.com.br/cadastro/home --ez private_tab true", 50)
	waitEnter("captcha")
	log.Printf("Filling form ---> CPF: %s", userData.EmpCPF)
	err = inputByRegexp("bounds=\"(\\[\\d+,\\d+\\]\\[\\d+,\\d+\\])\"><node NAF=\"true\" index=\"0\" text=\"\" resource-id=\"cpf\"", userData.EmpCPF, true)
	if err != nil {
		log.Printf("inputByRegexp err:\n%v", err)
		log.Printf("Failed to reach %v field coords", userData.EmpCPF)
		return err
	}

	log.Printf("Filling form ---> Card: %s", userData.FormattedCard)
	err = inputByRegexp("bounds=\"(\\[\\d+,\\d+\\]\\[\\d+,\\d+\\])\" /><node index=\"2\" text=\"Número do cartão\"", userData.FormattedCard, true)
	if err != nil {
		log.Printf("inputByRegexp err:\n%v", err)
		log.Printf("Failed to reach %v field coords", userData.FormattedCard)
		return err
	}

	log.Printf("Filling form ---> Name: %s", userData.EmpName)
	err = inputByRegexp("bounds=\"(\\[\\d+,\\d+\\]\\[\\d+,\\d+\\])\"><node NAF=\"true\" index=\"0\" text=\"\" resource-id=\"nome\"", userData.EmpName, false)
	if err != nil {
		log.Printf("inputByRegexp err:\n%v", err)
		log.Printf("Failed to reach %v field coords", userData.EmpName)
		return err
	}

	log.Printf("Filling form ---> Birthday: %s", userData.EmpDtBirth)
	err = inputByRegexp("bounds=\"(\\[\\d+,\\d+\\]\\[\\d+,\\d+\\])\" /><node index=\"2\" text=\"Data de nascimento\"", userData.EmpDtBirth, true)
	if err != nil {
		log.Printf("inputByRegexp err:\n%v", err)
		log.Printf("Failed to reach %v field coords", userData.EmpDtBirth)
		return err
	}

	sexradio := ""
	if userData.Gender == "F" {
		log.Printf("Filling form ---> Sex: %s", userData.Gender)
		sexradio = "bounds=\"(\\[\\d+,\\d+\\]\\[\\d+,\\d+\\])\" /><node index=\"1\" text=\"Feminino\""
	} else {
		log.Printf("Filling form ---> Sex: M or empty")
		sexradio = "bounds=\"(\\[\\d+,\\d+\\]\\[\\d+,\\d+\\])\" /><node index=\"1\" text=\"Masculino\""
	}
	err = inputByRegexp(sexradio, "", false)
	if err != nil {
		log.Printf("inputByRegexp err:\n%v", err)
		log.Println("Failed to fetch sexradio field coords")
		return err
	}

	log.Printf("Filling form ---> Email: %s", userData.EmpLogin)
	err = inputByRegexp("bounds=\"(\\[\\d+,\\d+\\]\\[\\d+,\\d+\\])\"><node NAF=\"true\" index=\"0\" text=\"\" resource-id=\"email\"", userData.EmpLogin, false)
	if err != nil {
		log.Printf("inputByRegexp err:\n%v", err)
		log.Printf("Failed to fetch %v field coords", userData.EmpLogin)
		return err
	}

	log.Printf("Filling form ---> Confirma Email: %s", userData.EmpLogin)
	err = inputByRegexp("bounds=\"(\\[\\d+,\\d+\\]\\[\\d+,\\d+\\])\"><node NAF=\"true\" index=\"0\" text=\"\" resource-id=\"cEmail\"", userData.EmpLogin, false)
	if err != nil {
		log.Printf("inputByRegexp err:\n%v", err)
		log.Printf("Failed to fetch %v field coords", userData.EmpLogin)
		return err
	}
	userData.EmpLoginPhone = randPhoneNumber()
	log.Printf("Filling form ---> Email again: %s", userData.EmpLoginPhone)
	err = inputByRegexp("bounds=\"(\\[\\d+,\\d+\\]\\[\\d+,\\d+\\])\" /><node index=\"2\" text=\"DDD \\+ Telefone\"", userData.EmpLoginPhone, true)
	if err != nil {
		log.Printf("inputByRegexp err:\n%v", err)
		log.Printf("Failed to fetch %v field coords", userData.EmpLoginPhone)
		return err
	}

	log.Printf("Filling form ---> Password: %s", userData.EmpPass)
	err = inputByRegexp("bounds=\"(\\[\\d+,\\d+\\]\\[\\d+,\\d+\\])\"><node NAF=\"true\" index=\"0\" text=\"\" resource-id=\"keySenha\"", userData.EmpPass, false)
	if err != nil {
		log.Printf("inputByRegexp err:\n%v", err)
		log.Printf("Failed to fetch %v field coords", userData.EmpPass)
		return err
	}

	log.Printf("Filling form ---> Password again: %s", userData.EmpPass)
	err = inputByRegexp("bounds=\"(\\[\\d+,\\d+\\]\\[\\d+,\\d+\\])\" /><node index=\"2\" text=\"Confirmação de senha\"", userData.EmpPass, false)
	if err != nil {
		log.Printf("inputByRegexp err:\n%v", err)
		log.Printf("Failed to fetch %v field coords", userData.EmpPass)
		return err
	}

	log.Printf("Filling form ---> Accepting agreement terms")
	err = inputByRegexp("bounds=\"(\\[\\d+,\\d+\\]\\[\\d+,\\d+\\])\" /><node NAF=\"true\" index=\"1\" text=\"\"", "", false)
	if err != nil {
		log.Printf("inputByRegexp err:\n%v", err)
		log.Println("Failed to fetch agreement terms field coords")
		return err
	}

	waitEnter("cadastrar")
	log.Printf("Press CADASTRAR button")
	err = inputByRegexp("<node index=\"\\d\" text=\"CADASTRAR\".*bounds=\"(\\[\\d+,\\d+\\]\\[\\d+,\\d+\\])\\\" /></node></node></node></node><node index=\"1\".*text=\"©Copyright Riocard 2019", "", false)
	if err != nil {
		log.Printf("inputByRegexp err:\n%v", err)
		log.Println("Failed to fetch record button field coords")
		return err
	}

	return nil
}

func xmlDumpScreen() string {
	waitIdle()
	return cmd("adb shell cat /sdcard/window_dump.xml", 1)
}

func xml2tap(xmlcoords string) (coords, error) {
	coords := coords{}
	openbracket := "["
	closebracket := "]"
	joinedbracket := "]["
	if string(xmlcoords[0]) == openbracket && string(xmlcoords[len(xmlcoords)-1]) == closebracket && strings.Contains(xmlcoords, joinedbracket) {
		stringcoords := strings.Split(xmlcoords, "][")
		leftcoords := strings.Split(string(stringcoords[0][1:]), ",")
		rightcoords := strings.Split(string(stringcoords[1][:len(stringcoords[1])-1]), ",")
		x1, err := strconv.Atoi(leftcoords[0])
		if err != nil {
			return coords, err
		}
		y1, err := strconv.Atoi(leftcoords[1])
		if err != nil {
			return coords, err
		}
		x2, err := strconv.Atoi(rightcoords[0])
		if err != nil {
			return coords, err
		}
		y2, err := strconv.Atoi(rightcoords[1])
		if err != nil {
			return coords, err
		}
		coords.x = (x1 + x2) / 2
		coords.y = (y1 + y2) / 2
		return coords, nil
	}
	return coords, fmt.Errorf("input data must start with '[', end with ']' and contain ']['")
}

func inputByRegexp(exp, inputData string, splitted bool) error {
	regexpcoords, err := applyRegexp(exp, xmlDumpScreen(), false)
	if err != nil {
		log.Printf("applyregexp err:\n%v", err)
		return err
	}
	if len(regexpcoords) > 1 {
		coords, err := xml2tap(regexpcoords[1])
		if err != nil {
			log.Printf("xml2tap err:\n%v", err)
			return err
		}
		tapScreen(coords.x, coords.y, 1)
		if splitted {
			splittedInput(inputData)
			return nil
		}
		if strings.Contains(inputData, " ") {
			log.Println("Filling field with spaces")
			fieldWithSpace(inputData)
			return nil
		}
		log.Println("Filling field without spaces")
		inputText(inputData)
		return nil
	}
	return fmt.Errorf("Unable to find regular expression")
}

func tapEnter(delay int) {
	log.Printf("Tapping enter and waiting %ds\n", delay)
	out := cmd("adb shell input keyevent 66", delay)
	if len(out) > 0 {
		log.Printf("cmd output:\n%v", out)
	}
}

func fieldWithSpace(value string) {
	cmd("adb shell input text $(echo '"+value+"' | sed 's/ /\\%s/g')", 1)
}

func readjson(path string) ([]UserData, error) {
	jsonFile, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(jsonFile), &userList)
	if err != nil {
		return nil, err
	}
	return userList, nil

}

func shell(arg string) string {
	args := strings.Split(arg, " ")
	if len(args) == 1 {
		args = append(args, "")
	} else if len(args) < 1 {
		log.Println("Invalid command")
		return ""
	}
	out, err := exec.Command(args[0], args[1:]...).CombinedOutput()
	if err != nil {
		log.Printf("Command: '%v'; Output: %v; Error: %v\n", arg, string(out), err)
		return err.Error()
	}
	return string(out)
}

func tapScreen(x, y, delay int) {
	xstring := strconv.Itoa(x)
	ystring := strconv.Itoa(y)
	cmd("adb shell input tap "+xstring+" "+ystring, delay)
}

func sleeper(delay int) {
	if delay > 15 {
		log.Printf("Waiting a longer sleep: %vs", delay*DefaultSleepTime/1000)
	}
	for i := 0; i < delay; i++ {
		time.Sleep(time.Duration(DefaultSleepTime) * time.Millisecond)
	}
}

func cmd(arg string, delay int) string {
	out := shell(arg)
	sleeper(delay)
	return out
}

func inputText(text string) {
	cmd("adb shell input text "+text, 0)
}

func splittedInput(value string) {
	for _, char := range strings.Split(value, "") {
		inputText(char)
		sleeper(1)
	}
}

func waitEnter(text string) {
	log.Printf("Waiting for '%v'", text)
	log.Printf("Press <enter> to continue or <ctrl+c> to interrupt")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	log.Printf("Now, where was I?")
	log.Printf("Oh yes...\n\n")
}

func applyRegexp(exp, readtext string, mustLog bool) ([]string, error) {
	if mustLog {
		log.Println("Reading screen in search of regexp")
	}

	re := regexp.MustCompile(exp)
	match := re.FindStringSubmatch(readtext)
	if len(match) < 1 {
		log.Println("Unable to find match for regexp")
		return []string{""}, fmt.Errorf("No match found (<1) for expression:\n%s", exp)
	}
	if len(match) > 0 {
		if len(match[1]) > 0 {
			if mustLog {
				log.Printf("match[1]: %v", match[1])
			}
			return match, nil
		} else if len(match) > 2 {
			if len(match[2]) > 0 {
				if mustLog {
					log.Printf("match[2]: %v", match[2])
				}
				return []string{match[0], match[2]}, nil
			}
		}
	}
	return []string{""}, errors.New("No match found (=1)")
}

func waitIdle() {
	for i := 0; strings.Contains(cmd("adb shell uiautomator dump", 0), "idle state"); i++ {
		sleeper(1)
		if i >= 5 {
			log.Println("App is not reaching idle state.")
			return
		}
	}
}

func randPhoneNumber() string {
	ddd := []string{"11", "12", "13", "14", "15", "16", "17", "18", "19", "21", "22", "24", "27", "28", "31", "32", "33", "34", "35", "37", "38", "41", "42", "43", "44", "45", "46", "47", "48", "49", "51", "53", "54", "55", "61", "62", "63", "64", "65", "66", "67", "68", "69", "71", "73", "74", "75", "77", "79", "81", "82", "83", "84", "85", "86", "87", "88", "89", "91", "92", "93", "94", "95", "96", "97", "98", "99"}
	newnumber := ddd[randNumber(len(ddd))] + "9"
	for i := 0; i < 8; i++ {
		newnumber += strconv.Itoa(randNumber(10))
	}
	log.Printf("New number generated: %v", newnumber)
	return newnumber
}

func validatePW(pass string) bool {
	log.Printf("Validating password %s", pass)
	if len(pass) < 10 || len(pass) > 12 || strings.Contains(pass, "*") || strings.Contains(pass, "#") {
		return false
	}
	expression := []string{"([A-Z]+)", "([a-z]+)", "(\\d+)", "([@|%]+)"}
	for i := range expression {
		value, err := applyRegexp(expression[i], pass, false)
		if err != nil {
			fmt.Printf("applyRegexp err: %v", err)
			return false
		}
		if len(value) == 0 {
			fmt.Printf("Expression %s not satisfied", expression[i])
			return false
		}
	}
	return true
}

func randNumber(n int) int {
	time.Sleep(time.Duration(1) * time.Millisecond)
	return rand.New(rand.NewSource(time.Now().UnixNano())).Intn(n)
}
