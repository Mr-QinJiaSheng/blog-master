package utils

import (
	"encoding/json"
	"fmt"
	models "go-blog/models/admin"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

const API = "https://api.niubi.plus/Music/Get.php?type=song&media=netease&id="

type Music struct {
	Name   string `json:"name"`
	Url    string `json:"url"`
	Cover  string `json:"cover"`
	SongId int    `json:"song_id"`
	Author string `json:"author"`
}

func MusicGet(id int) (music Music, newid int) {

	for {
		resp, err := http.Get(API + strconv.Itoa(id))
		if err != nil {
			id++
			fmt.Println(err.Error())
			//time.Sleep(2 * time.Second)
			continue
		}
		defer resp.Body.Close()
		s, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			id++
			fmt.Println(err.Error())
			//time.Sleep(2 * time.Second)
			continue
		}
		fmt.Println(string(s))
		err = json.Unmarshal(s, &music)
		if err != nil {
			id++
			fmt.Println(err.Error())
			//time.Sleep(2 * time.Second)
			continue
		}
		if music.Url == "" {
			id++
			//time.Sleep(2 * time.Second)
			continue
		}
		id++
		newid = id
		fmt.Printf("%d更新成功:%v\n", id, music)
		break
	}
	return
}

func MusicSave(id int) (nid int) {
	m, nid := MusicGet(id)
	models.AddMusic(&models.Music{
		Name:   m.Name,
		Cover:  m.Cover,
		SongId: m.SongId,
		Url:    m.Url,
		Author: m.Author,
	})
	return
}

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
