package main

import "C"

import (
	"fmt"
	"log"
	"os"

	config "./src/config"
	pawn_api "./src/pawn_api"
)

func main() {
	config.InitConfig()
	current_dir, _ := os.Getwd()
	src_info := pawn_api.SourceInfo{CurrentPath: current_dir, Name: "hello2", IncludeDir: current_dir + "/include"}
	var err_compile error
	var out_compiler string
	err_compile, out_compiler = pawn_api.GetStreamsOut(pawn_api.CompileSrc, &src_info)

	fmt.Println(out_compiler)
	if err_compile == nil {
		err_run := pawn_api.RunPwn(&src_info)
		if err_run != nil {
			log.Println(err_run)
		}
	}
}
