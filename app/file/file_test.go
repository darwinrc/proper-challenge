package file

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestFile_Store(t *testing.T) {
	data, err := os.Open("../../test/1.jpg")
	defer data.Close()
	assert.Nil(t, err)
	assert.NotNil(t, data)

	f := &File{
		Name: "1.jpg",
		Url:  "https://i.chzbgr.com/full/9710998528/h05CBADD8/old-gods",
		Data: data,
	}
	err = f.Store("../../img/")
	assert.Nil(t, err)

	img, err := os.Open("../../img/1.jpg")
	defer img.Close()
	assert.Nil(t, err)
	assert.NotNil(t, img)

	assert.Contains(t, img.Name(), "1.jpg")

	stat, err := img.Stat()
	assert.Nil(t, err)
	assert.NotNil(t, stat)
	assert.Equal(t, int64(119096), stat.Size())
}
