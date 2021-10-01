package grapher

import (
	"bytes"
	"io/ioutil"
	"os"

	"github.com/i582/CodeQuery/pkg/graph"
)

func AddGraphScalability(name string) error {
	data, err := ioutil.ReadFile(name + ".svg")
	if err != nil {
		return err
	}

	needStr := `xmlns:xlink="http://www.w3.org/1999/xlink">`
	startGraphData := bytes.Index(data, []byte(needStr))
	startGraphData += len(needStr)
	startSvg := bytes.Index(data, []byte("<svg ")) + 5
	startViewBox := bytes.Index(data, []byte(" viewBox"))
	startEndSvg := bytes.Index(data, []byte("</svg>"))
	var newData []byte
	newData = append(newData, data[0:startSvg]...)
	newData = append(newData, []byte("width=\"100%\" height=\"100%\"")...)
	newData = append(newData, data[startViewBox:startGraphData]...)
	newData = append(newData, []byte(graph.WebAdditionHeader)...)
	newData = append(newData, data[startGraphData:startEndSvg]...)
	newData = append(newData, []byte(graph.WebAdditionFooter)...)
	newData = append(newData, data[startEndSvg:]...)
	err = ioutil.WriteFile(name+".svg", newData, 0677)
	if err != nil {
		return err
	}
	return nil
}

func WriteGraph(name string, graphData string) error {
	graphFileName := name + ".gv"

	err := os.WriteFile(graphFileName, []byte(graphData), 0777)
	if err != nil {
		return err
	}

	dot := &Dot{
		Format:     Svg,
		InputFile:  graphFileName,
		OutputName: name + ".svg",
	}
	err = dot.Execute()
	if err != nil {
		return err
	}

	return nil
}
