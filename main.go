package main

import (
	"fmt"
	"log"
	"os"

	pawn_api "./src/pawn_api"
)

func main() {
	current_dir, _ := os.Getwd()
	src_info := pawn_api.SourceInfo{CurrentPath: current_dir, Name: "hello2", IncludeDir: current_dir + "/include"}
	fmt.Println(src_info)
	err_compile := pawn_api.CompileSrc(&src_info)
	if err_compile == nil {
		err_run := pawn_api.RunPwn(&src_info)
		if err_run != nil {
			log.Println(err_run)
		}
	}
}
