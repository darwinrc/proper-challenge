package web

import (
	"github.com/PuerkitoBio/goquery"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"strings"
	"testing"
)

func TestWeb_GetImages(t *testing.T) {
	data, err := ioutil.ReadFile("../../test/test_doc.html")
	assert.Nil(t, err)
	assert.NotNil(t, data)

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(data)))
	assert.Nil(t, err)
	assert.NotNil(t, doc)

	w := &Web{
		Url: "https://icanhas.cheezburger.com/",
		Doc: doc,
	}

	images := w.GetImages(".mu-content-card", 10, 1, 16)
	assert.NotNil(t, images)

	assert.Equal(t, 10, len(images))
	assert.Equal(t, "1.jpg", images[0].Name)
	assert.Equal(t, "https://i.chzbgr.com/thumb800/1823751/h217C9FDC/found-blind-kitten-is-nursed-back-to-full-recovery-can-even-see-again", images[0].Url)
	assert.Equal(t, "10.jpg", images[9].Name)
	assert.Equal(t, "https://i.chzbgr.com/thumb800/18365957/h71539753/cat-this-cat-at-our-work-tracked-a-mouse-on-the-security-cameras-and-then-went-and-caught-it-res", images[9].Url)
}
