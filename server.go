package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/schema"
	"github.com/kelseyhightower/envconfig"
)

type envVars struct {
	Debug       bool
	SlackTokens string
	TriggerWord string
}

type slackRequest struct {
	Token       string `schema:"token"`
	TeamID      string `schema:"team_id"`
	TeamDomain  string `schema:"team_domain"`
	ChannelID   string `schema:"channel_id"`
	ServiceID   string `schema:"service_id"`
	ChannelName string `schema:"channel_name"`
	Timestamp   string `schema:"timestamp"`
	UserID      string `schema:"user_id"`
	UserName    string `schema:"user_name"`
	Text        string `schema:"text"`
	TriggerWord string `schema:"trigger_word"`
}

type slackResponse struct {
	Text string `json:"text"`
}

var env envVars
var flips map[string]string

func main() {
	// get env vars
	err := envconfig.Process("flippy", &env)
	if err != nil {
		log.Fatal(err.Error())
	}

	// if we don't have a TriggerWord set in env, default to 'flip'
	if env.TriggerWord == "" {
		env.TriggerWord = "flip"
	}

	flips = getFlipMap()

	r := mux.NewRouter()

	r.HandleFunc("/slack", slackPostHandler).Methods("POST")

	http.Handle("/", r)

	log.Printf("Listening on http://0.0.0.0:%v...\n", os.Getenv("PORT"))
	if err = http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("PORT")), nil); err != nil {
		log.Fatal(err)
	}

}

func slackPostHandler(w http.ResponseWriter, r *http.Request) {

	hookRequest := slackRequest{}
	table := "┻━┻"
	var reversed string

	start := time.Now()
	log.Printf(
		"%s\t%s\t%s",
		r.Method,
		r.RequestURI,
		start,
	)

	err := r.ParseForm()

	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}

	defer r.Body.Close()

	err = validateSlackRequest(r, &hookRequest)

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// the text passed in is prepended with the hook trigger, so remove it to get the text to flip, if any
	triggerText := strings.Replace(strings.Trim(hookRequest.Text, " "), env.TriggerWord, "", 1)

	if triggerText != "" {
		reversed = flipText(triggerText)
	}

	table = reverseString(table + reversed)

	resp := slackResponse{
		Text: "(╯°□°）╯︵ " + table,
	}

	log.Println(resp)

	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.NotFound(w, r)
		return
	}

}

func flipText(input string) string {
	var flipped string
	for _, rune := range input {
		letter := string(rune)
		// get matches
		if flips[letter] != "" {
			flipped += flips[letter]
		} else {
			flipped += letter
		}
	}

	return flipped
}

func validateSlackRequest(r *http.Request, hookRequest *slackRequest) error {

	decoder := schema.NewDecoder()
	err := decoder.Decode(hookRequest, r.PostForm)

	if err != nil {
		return err
	}

	log.Printf("%#v\n", hookRequest)

	if !strings.Contains(env.SlackTokens, hookRequest.Token) || hookRequest.TriggerWord != env.TriggerWord {
		return errors.New("invalid token or trigger")
	}

	return nil
}

func reverseString(input string) string {
	// Get Unicode code points.
	n := 0
	rune := make([]rune, len(input))
	for _, r := range input {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	output := string(rune)

	return output
}

func getFlipMap() map[string]string {
	var flips = make(map[string]string)

	flips["a"] = "ɐ"
	flips["b"] = "q"
	flips["c"] = "ɔ"
	flips["d"] = "p"
	flips["e"] = "ǝ"
	flips["f"] = "ɟ"
	flips["g"] = "ƃ"
	flips["h"] = "ɥ"
	flips["i"] = "ᴉ"
	flips["j"] = "ɾ"
	flips["k"] = "ʞ"
	flips["l"] = "l"
	flips["m"] = "ɯ"
	flips["n"] = "u"
	flips["o"] = "o"
	flips["p"] = "d"
	flips["q"] = "b"
	flips["r"] = "ɹ"
	flips["s"] = "s"
	flips["t"] = "ʇ"
	flips["u"] = "n"
	flips["v"] = "ʌ"
	flips["w"] = "ʍ"
	flips["x"] = "x"
	flips["y"] = "ʎ"
	flips["z"] = "z"

	flips["A"] = "∀"
	flips["B"] = "B"
	flips["C"] = "Ɔ"
	flips["D"] = "D"
	flips["E"] = "Ǝ"
	flips["F"] = "Ⅎ"
	flips["G"] = "פ"
	flips["H"] = "H"
	flips["I"] = "I"
	flips["J"] = "ſ"
	flips["K"] = "K"
	flips["L"] = "˥"
	flips["M"] = "W"
	flips["N"] = "N"
	flips["O"] = "O"
	flips["P"] = "Ԁ"
	flips["Q"] = "Q"
	flips["R"] = "R"
	flips["S"] = "S"
	flips["T"] = "┴"
	flips["U"] = "∩"
	flips["V"] = "Λ"
	flips["W"] = "M"
	flips["X"] = "X"
	flips["Y"] = "⅄"
	flips["Z"] = "z"

	flips["0"] = "0"
	flips["1"] = "Ɩ"
	flips["2"] = "ᄅ"
	flips["3"] = "Ɛ"
	flips["4"] = "ㄣ"
	flips["5"] = "ϛ"
	flips["6"] = "9"
	flips["7"] = "ㄥ"
	flips["8"] = "8"
	flips["9"] = "6"

	flips[","] = "'"
	flips["."] = "˙"
	flips["?"] = "¿"
	flips["!"] = "¡"
	flips["\""] = ",,"
	flips["'"] = ","
	flips["`"] = ","
	flips["("] = ")"
	flips[")"] = "("
	flips["["] = "]"
	flips["]"] = "["
	flips["{"] = "}"
	flips["}"] = "{"
	flips["<"] = ">"
	flips[">"] = "<"
	flips["&"] = "⅋"
	flips["_"] = "‾"

	return flips
}
