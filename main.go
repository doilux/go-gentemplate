package main

import (
	"flag"
	"go-gentemplate/gen"
	"io"
	"os"
	"strings"
)

const generatorName = "gentemplate"

var (
	// 処理対象の型（カンマ区切り）
	targetTypes = flag.String("types", "", "")
	// ソースファイル
	src = flag.String("src", "", "source file")
	// 出力ファイル
	dist = flag.String("dist", "", "distination file")
	// 出力ファイル
	debugMode = flag.Bool("debug", false, "debug mode")
)

func main() {
	flag.Usage()
	flag.Parse()
	// 引数チェック
	checkArgs(*src, *dist, *targetTypes)


	// コード生成
	spec := generateFileSpec(*targetTypes, generatorName, *debugMode)

	// ファイルの書き込み
	writeToFile(spec, *dist)

}

func checkArgs(src string, dist string, targetTypes string) {
	if src == "" || dist == "" || targetTypes == "" {
		panic("option must not be empty")
	}

}

// generateFileSpec はファイルの中身を生成する
func generateFileSpec(targetStructs string, generatorName string, debug bool) []byte {
	spec, err := gen.Generate(*src, gen.TargetTypes(strings.Split(targetStructs, ",")), generatorName, debug)
	if err != nil {
		panic(err)
	}
	return spec
}

// writeToFile はdistFileを生成する
func writeToFile(spec []byte, distFile string) {
	// ファイルに書き込む
	var writer io.Writer
	writer, err := os.Create(*dist)
	if err != nil {
		panic(err)
	}
	if closer, ok := writer.(io.Closer); ok {
		defer closer.Close()
	}
	if _, err := writer.Write(spec); err != nil {
		panic(err)
	}
}


