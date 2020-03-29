package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	var mode string
	flag.StringVar(&mode, "mode", "", "The parser mode ('teams' or 'players')")
	flag.Parse()

	if mode == "" {
		log.Fatal("mode flag not provided")
	}

	switch mode {
	case "teams":
		parseTeams()
		break
	case "players":
		parsePlayers()
		break
	default:
		log.Fatal("error mode " + mode + "not supported")
	}
}

const (
	agentesLibres = "Agentes Libres"
	sudamericana  = "Sudamericana"
)

func parseTeams() {
	type team struct {
		Name      string `json:"name"`
		ShieldSrc string `json:"img"`
		League    string `json:"league"`
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("error getting current directory: ", err)
	}
	file, err := os.Open(dir + "/cmd/parser/teams.json")
	defer file.Close()
	if err != nil {
		log.Fatal("error openning teams.json: ", err)
	}
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("error reading teams.json file: ", err)
	}

	var teams []team
	if err := json.Unmarshal(byteValue, &teams); err != nil {
		log.Fatal("error unmarshalling data to teams", err)
	}

	f, err := os.Create(dir + "/cmd/parser/init-teams.sql")
	if err != nil {
		log.Fatal("error creating init-teams.sql: ", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, t := range teams {
		if t.League != agentesLibres && t.League != sudamericana {
			_, err := w.WriteString("INSERT INTO generator.fifa_team (name, shield_src,  league_id) values('" + sanitize(t.Name) + "',  '" + sanitize(t.ShieldSrc) + "'," + "( SELECT id FROM generator.fifa_league WHERE name = '" + sanitize(t.League) + "'));\n")
			if err != nil {
				log.Fatal("error writting string: ", err)
			}
		}
	}

	w.Flush()

}

func parsePlayers() {
	type player struct {
		Name       string `json:"name"`
		Position   string `json:"position"`
		Number     string `json:"number"`
		PictureSrc string `json:"picture"`
		Team       string `json:"team"`
		League     string `json:"league"`
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatal("error getting current directory: ", err)
	}
	file, err := os.Open(dir + "/cmd/parser/players.json")
	defer file.Close()
	if err != nil {
		log.Fatal("error openning players.json: ", err)
	}
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal("error reading players.json file: ", err)
	}

	var players []player
	if err := json.Unmarshal(byteValue, &players); err != nil {
		log.Fatal("error unmarshalling data to players", err)
	}

	f, err := os.Create(dir + "/cmd/parser/init-players.sql")
	if err != nil {
		log.Fatal("error creating init-players.sql: ", err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	for _, p := range players {
		if p.League != agentesLibres && p.League != sudamericana && p.Team != agentesLibres {
			_, err := w.WriteString("INSERT INTO generator.fifa_player(name, picture_src, number, position, team_id) values('" + sanitize(p.Name) + "',  '" + sanitize(p.PictureSrc) + "',  '" + sanitize(p.Number) + "',  '" + sanitize(p.Position) + "'," + "( SELECT ft.id FROM generator.fifa_team ft JOIN generator.fifa_league fl ON ft.league_id = fl.id WHERE ft.name = '" + sanitize(p.Team) + "' AND fl.name = '" + sanitize(p.League) + "' ));\n")
			if err != nil {
				log.Fatal("error writting string: ", err)
			}
		}
	}

	w.Flush()

}

func sanitize(s string) string {
	return strings.Replace(s, "'", "''", -1)
}
