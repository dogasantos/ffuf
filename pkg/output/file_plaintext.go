package output

import (
	"github.com/ffuf/ffuf/pkg/ffuf"
	"bufio"
	"os"


)
/*

ffuf.Result{
		Input:            inputs,
		Position:         resp.Request.Position,
		StatusCode:       resp.StatusCode,
		ContentLength:    resp.ContentLength,
		ContentWords:     resp.ContentWords,
		ContentLines:     resp.ContentLines,
		ContentType:      resp.ContentType,
		RedirectLocation: resp.GetRedirectLocation(false),
		Url:              resp.Request.Url,
		Duration:         resp.Time,
		ResultFile:       resp.ResultFile,
		Host:             resp.Request.Host,
	}
*/
func writePlain(filename string, config *ffuf.Config, res []ffuf.Result) error {

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()


	writer := bufio.NewWriter(f)
	for _, resItem := range res {
		_, _ = writer.WriteString(resItem.Request.url + "\n")
	}
	writer.Flush()
	

	return err
}
