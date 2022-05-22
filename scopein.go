package main

import (
	"fmt"
	"strings"
	"bufio"
	"flag"
	"net/url"
	"regexp"
	"io/ioutil"
	"os"
	"log"
	"sync"
)

func init() {
	flag.Usage = func() {
	help := []string{
			"",
			"[buffer] | scopein [flags]",
			"+====================================================================================+",
			"-s <scope>		In Scope targets separated by '|'",
			"-b <outscope>		Out of scope targets separated by '|'",
			"-f <file>		File with In Scope urls",
			"-bf <file>		File with Out of Scope urls",
			"-h,			Show This Help Message",
			"",
			"+====================================================================================+",
			"",
	}

	fmt.Println(`
	  A
         /!\
        / ! \
 /\     )___(
(  '.____(_)_________
|           __..--""
(       _.-|
 \    ,' | |
  \  /   | |
   \(    | |
    '    | |
         | |
	`)
	fmt.Fprintf(os.Stderr, strings.Join(help, "\n"))

}

}

func main(){

	log.SetOutput(ioutil.Discard)
	
	var single string
	flag.StringVar(&single, "s", "", "")

	var scopeFile string
	flag.StringVar(&scopeFile,"f", "","")

	var outscope string
	flag.StringVar(&outscope, "b", "", "")

	var outscopeFile string
	flag.StringVar(&outscopeFile, "bf", "", "")

	flag.Parse()

	visto := make(map[string]bool)
	std := bufio.NewScanner(os.Stdin)
	targets := make(chan string)

	var wg sync.WaitGroup
	for i:=0;i<50;i++ {
			wg.Add(1)
			go func() {
					defer wg.Done()
					for v := range targets{
						if !strings.HasPrefix(v, "http"){
							v = "https://" + v
						}

						u, err := url.Parse(v)
						if err != nil{
							continue
						}
						
						scopeurl := scopein(u, single, scopeFile, outscope, outscopeFile)
						if scopeurl != "not"{
							fmt.Println(scopeurl)
						}

						}
					
				

			}()
	}

	for std.Scan() {
		var line string = std.Text()
		if visto[line] != true{
			targets <- line
		}
		visto[line] = true

	}
	close(targets)
	wg.Wait()

}

func scopein(v *url.URL, single string, scopeFile string, outscope string, outscopeFile string) string{
	//fmt.Println(v.Host)
	host := v.Host
	
	if strings.Contains(host, ":"){
		host = strings.Split(host, ":")[0]
		
	}
	

	

	// SPLIT REGEX WITH |, -s "xx|xxx|xxx"
	if single != ""{
		if strings.HasPrefix(single, "*."){
			single = strings.Replace(single, "*.", "", 1)
		}
		

		re, err := regexp.Compile(`.*\.?` + single + `\/?.*`)
		if err != nil{
			log.Println("Failed to compile regex!")
		}
	
		for _, value := range re.FindAllString(host, -1){
			log.Println(value)
			if v.String() == "https:"{
				return "not"
			}
			return v.String()
			
		}
		
		//(http:\/\/|https:\/\/)?(\.*\.[a-z0-9.-])?(www.)?(scope.com)\/?.*
		

	}else if scopeFile != ""{
		f, err := os.ReadFile(scopeFile)

		if err != nil{
			log.Println("Error opening file!")
		}
		arr := string(f)
		var array string
		
		str := strings.Split(arr, "\n")
		for _, p := range str{
			
			if !strings.HasPrefix(p, " "){
				array = array + p + "|"
			}
		}

		textstr := string(array[0:len(array)-2])

		if strings.HasPrefix(textstr, "*."){
			textstr = strings.Replace(textstr, "*.", "", -1)
		}
			
	
			
		re, err := regexp.Compile(`.*\.?` + textstr + `\/?.*`)
		if err != nil{
			log.Println("Failed to compile regex!")
		}
	
		for _, value := range re.FindAllString(host, -1){
			log.Println(value)
			if v.String() == "https:"{
				return "not"
			}
			return v.String()
			}
			
			//return "not"
			
		
		
	

	}else if outscope != ""{
		if strings.HasPrefix(outscope, "*."){
			outscope = strings.Replace(outscope, "*.", "", 1)
		}
		

		//fmt.Println(single)
		re, err := regexp.Compile(`.*\.?` + outscope + `\/?.*`)
		if err != nil{
			log.Println("Failed to compile regex!")
		}
		
		search := re.FindAllString(host, -1)
		if (search == nil) == true{
			if v.String() == "https:"{
				return "not"
			}
			return v.String()
		}else{
			return "not"
		}
		
		
		
	}else if outscopeFile != ""{
		f, err := os.ReadFile(outscopeFile)

		if err != nil{
			log.Println("Error opening file!")
		}
		arr := string(f)
		var array string
		
		str := strings.Split(arr, "\n")
		for _, p := range str{
			
			if !strings.HasPrefix(p, " "){
				array = array + p + "|"
			}
		}

		outs := string(array[0:len(array)-2])

		if strings.HasPrefix(outs, "*."){
			outs = strings.Replace(outs, "*.", "", -1)
		}
			

			
		re, err := regexp.Compile(`.*\.?` + outs + `\/?.*`)
		if err != nil{
			log.Println("Failed to compile regex!")
			
		}
		
		search := re.FindAllString(host, -1)
		if (search == nil) == true{
			if v.String() == "https:"{
				return "not"
			}
			return v.String()
		}else{
			return "not"
		}
		
	}else{
		// read ~/.config/scopein/scope.conf
		return "not"
	}
	return "not"
	
}

