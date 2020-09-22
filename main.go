package main

import (
  "io/ioutil"
  "os"
  "path"

  "github.com/leaanthony/mewn"
  "github.com/wailsapp/wails"
)

type professions struct {
  firstProfession string;
  secondProfession string;
  cookingProfession string;
  cookingSkill int;
  archaeologySkill int;
}

type character struct {
  name string;
  level int;
  classID int;
  className string;
  race string;
  currentRole string;
  avatarURL string;
  itemlevel int;
  guild string;
}

func getDummy() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return " ", err
	}
	filename := path.Join(cwd, "dummyDatabase.json")
	result, err := ioutil.ReadFile(filename)
	return string(result), err
}

func main() {

  js := mewn.String("./frontend/dist/app.js")
  css := mewn.String("./frontend/dist/app.css")

  app := wails.CreateApp(&wails.AppConfig{
    Width:  1280,
    Height: 768,
    Title:  "WoWnomicon",
    JS:     js,
    CSS:    css,
    Colour: "#131313",
  })
  app.Bind(getDummy)
  app.Run()
}
