package output

import (
	"github.com/ffuf/ffuf/pkg/ffuf"
	"bufio"
	"os"

)

func writePlain(filename string, config *ffuf.Config, res []ffuf.Result) error {

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()


	writer := bufio.NewWriter(f)
	for _, resItem := range res {
		_, _ = writer.WriteString(resItem.Url+"\n",resItem.Url)
	}
	writer.Flush()

	return err
}
