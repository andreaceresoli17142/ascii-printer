package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"math/rand"
	"time"
	"io/ioutil"
	"runtime"
	"path"
)

func RemoveIndex(s []string, index int) []string {
    return append(s[:index], s[index+1:]...)
}

func getCurrentPath() string {
    _, filename, _, _ := runtime.Caller(1)
    return path.Dir(filename)
}

func printFileToTerm(file_path string) error {
	content, err := ioutil.ReadFile( file_path )
  	if err != nil {
  		return err
	}
   fmt.Println(string(content))
	return nil
}

func main(){

	// get arguments
	argsN := len(os.Args[1:])

	if argsN > 2 {
		fmt.Println("erroe: too many args")
		return
	}

	if argsN == 0 {

		fmt.Println(`
	Help page:
		-h 	help page
		-l 	list all ascii-arts
		-ls 	list all ascii-arts and show them
		-r 	return random ascii-art
`)
		return
	}

	// get file path
	filePath := getCurrentPath()
	
	// get files in assets dir
	cmd := exec.Command("ls", filePath + "/assets/")
	output, err := cmd.CombinedOutput()
	
	if err != nil {
		fmt.Println(err)
		return 
	} 

	files := strings.Split( string(output), "\n")

	// remove empty string
	to_remove := -1
	for i, ascii := range files {
		
		if ascii == "" {
			to_remove = i
		}
	}

	if to_remove != -1 {
		files = RemoveIndex( files, to_remove )
	}

	// args switch
	switch (os.Args[1]){
		case "-h":
			fmt.Println(`
		Help page:
			-h 	help page
			-l 	list all ascii-arts
			-ls 	list all ascii-arts and show them
			-r		return random ascii-art
`)
		break

		case "-l":
		// list files
			for _, ascii := range files {
				fmt.Printf( "%s\n", ascii )
			}
		break
		
		case "-ls":
		// list files and print them
			for _, ascii := range files {
				fmt.Printf( "%s\n", ascii )
				printFileToTerm( filePath + "/assets/" + ascii )
	 		}
		break

		case "-r":
			// choose random file
			rand.Seed(time.Now().UnixNano())
			v := rand.Int() % len(files)
			err := printFileToTerm( filePath + "/assets/" + files[v] )
			
			if err != nil {
   			fmt.Println(err)
   			return
	 		}
		break
 	}
}
