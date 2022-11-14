package app

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"os"
	"proper-challenge/app/file"
	"testing"
)

func TestApp_GetImages(t *testing.T) {
	app := App{}

	files := app.GetImages(33, 3, 16)

	assert.NotNil(t, files)
	assert.Equal(t, 33, len(files))
	assert.Equal(t, "1.jpg", files[0].Name)
	assert.Equal(t, "33.jpg", files[32].Name)
}

func TestApp_StoreImages(t *testing.T) {
	app := App{}

	files := []*file.File{{
		Name: "1.jpg",
		Url:  "https://i.chzbgr.com/thumb800/18415365/h849F22ED/go-you-guys-walk-out-144-am-nov-10-2022-twitter-for-iphone-973-retweets-483-quote-tweets-423k-likes",
	},
		{
			Name: "2.jpg",
			Url:  "https://i.chzbgr.com/full/9712046336/h1F0003A5/incoming",
		},
		{
			Name: "3.jpg",
			Url:  "https://i.chzbgr.com/full/9710725120/h514D2447/get-yourself-a-new-cat-me-christmas-sweater-set-by-icanhascheezburger-just-in-time-for-the-holidays",
		},
		{
			Name: "4.jpg",
			Url:  "https://i.chzbgr.com/thumb800/18418181/hB3A65183/10-images-of-christmas-sweater",
		},
		{
			Name: "5.jpg",
			Url:  "https://i.chzbgr.com/thumb800/18417925/hF628886C/christmas-calendar-holidays-happy-holidays-18417925",
		},
	}

	app.StoreImages(files, 5)

	img, err := os.Open("img/1.jpg")
	defer img.Close()
	assert.Nil(t, err)
	assert.NotNil(t, img)
	assert.Contains(t, img.Name(), "1.jpg")

	stat, err := img.Stat()
	assert.Nil(t, err)
	assert.NotNil(t, stat)
	assert.Equal(t, int64(62910), stat.Size())

	img5, err := os.Open("../img/5.jpg")
	defer img.Close()
	assert.Nil(t, err)
	assert.NotNil(t, img5)
	assert.Contains(t, img5.Name(), "5.jpg")

	stat5, err := img5.Stat()
	assert.Nil(t, err)
	assert.NotNil(t, stat5)
	assert.Equal(t, int64(83143), stat5.Size())

	dir, err := ioutil.ReadDir("img")
	assert.Nil(t, err)
	assert.NotNil(t, dir)
	assert.Equal(t, 5, len(dir))

	if err := os.RemoveAll("img"); err != nil {
		t.Log(err)
	}
}
