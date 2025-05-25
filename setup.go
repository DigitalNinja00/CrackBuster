package main

import (
	"fmt"
	"os"
	"net/http"
	"bufio"
	"time"
	"strconv"
	"os/exec"
	"math/rand"
)
func main(){
	banner();
	args := os.Args
	count := len(args);
	if(count==2){
		if(args[1]=="-u"){
			fmt.Println("You must specify an argument for -u");
			return
		}
		if(args[1]=="-w"){
			fmt.Println("You must specify an argument for -w");
			return
		}
		if(args[1]=="-ig"){
			fmt.Println("You must specify an argument for -ig");
			return
		}

		if(args[1]=="--version" || args[1]=="-v"){
			fmt.Println("CRACKBUSTER VERSION: 2.0");
			return
		}
		if(args[1]=="--help" || args[1]=="-h"){
			lista := [10]string{"./crackbuster -u <url> -w <wordlist> -fc <response>",
			"--version or -v | see the version of the tool",
		    "--help or -h | display help information about the tool",
	        "-u | Specifies the target URL", 
		    "-w | It works to specify the wordlist to use.",
		    "-ig | It works to specify the wordlist to use."};
			for x:=0;x < 6; x++ {
				fmt.Println(lista[x]);
			}
			return
		}
		fmt.Println(args[1], "Invalid Option");
	}
	if(count==3){
		if(args[1]=="-u"){
			fmt.Println(".... -w <wordlist> -ig <response>");
			return
		}
		if(args[1]=="-w"){
			fmt.Println(".... -u <url> -ig <response>");
			return
		}
		if(args[1]=="-ig"){
			fmt.Println(".... -u <url> -w <wordlist>");
			return
		}
		fmt.Println(args[1], "Invalid Option");
	}
	if(count==4){
		if(args[1]=="-u" && args[3]=="-w"){
			fmt.Println("You must specify an argument for -w");
			return
		}
		if(args[1]=="-ig" && args[3]=="-w"){
			fmt.Println("You must specify an argument for -w");
			return
		}
		if(args[1]=="-w" && args[3]=="-u"){
			fmt.Println("You must specify an argument for -u");
			return
		}
		if(args[1]=="-ig" && args[3]=="-u"){
			fmt.Println("You must specify an argument for -u");
			return
		}
		if(args[1]=="-u" && args[3]=="-ig"){
			fmt.Println("You must specify an argument for -ig");
			return
		}
		if(args[1]=="-w" && args[3]=="-ig"){
			fmt.Println("You must specify an argument for -ig");
			return
		}
		if(args[1]==args[3] || args[3]==args[1]){
			fmt.Println("You cannot place the same argument more than once");
			return
		}
	}
	if(count==5){
		if((args[1]=="-u" && args[3]=="-w") || (args[1]=="-w" && args[3]=="-u")){
			fmt.Println("specifies the missing parameter -ig <response>");
			return
		}
		if((args[1]=="-u" && args[3]=="-ig") || args[1]=="-ig" && args[1]=="-u"){
			fmt.Println("specifies the missing parameter -w <wordlist>");
			return
		}
		if((args[1]=="-ig" && args[3]=="-w") || (args[1]=="-w" && args[3]=="-ig")){
			fmt.Println("specifies the missing parameter -u <url>");
			return
		}
		if((args[1]=="-ig" && args[3]=="-u") || (args[1]=="-u" && args[3]=="-ig")){
			fmt.Println("specifies the missing parameter -w <wordlist>");
			return
		}
		if(args[1]==args[3] || args[3]==args[1]){
			fmt.Println("You cannot place the same argument more than once");
			return
		}
	}
	if(count==6){
		if(args[1]=="-u" && args[3]=="-w" && args[5]=="-ig"){
			fmt.Println("You must specify an argument for -ig");
			return
		}
 
		if(args[1]=="u" && args[3]=="-ig" && args[5]=="-w"){
			fmt.Println("You must specify an argument for -w");
			return
		}
		if(args[1]=="-w" && args[3]=="-u" && args[5]=="-ig"){
			fmt.Println("You must specify an argument for -ig");
			return
		}
		if(args[1]=="-w" && args[3]=="-ig" && args[5]=="-u"){
			fmt.Println("You must specify an argument for -u");
			return;
		}
		if(args[1]=="-ig" && args[3]=="-u" && args[5]=="-w"){
			fmt.Println("You must specify an argument for -w");
			return;
		}
		if(args[1]=="-ig" && args[3]=="-w" && args[5]=="-u"){
			fmt.Println("You must specify an argument for -u");
			return;
		}
	}
	if(count==7){
		if(args[1]=="-u" && args[3]=="-w" && args[5]=="-ig"){
			str, err := strconv.Atoi(args[6]);
			if err != nil {
				fmt.Println(err);
			}
			indexar(args[2], args[4], str);
			fmt.Println("Attack ended with dictionary", args[4])
		}
		if(args[1]=="-u" && args[3]=="-ig" && args[5]=="-w"){
			str_, err := strconv.Atoi(args[4]);
            if err != nil {          
				fmt.Println(err);
			}
			indexar(args[2], args[6], str_);
			fmt.Println("Attack ended with dictionary", args[6])
		}
		if(args[1]=="-w" && args[3]=="-u" && args[5]=="-ig"){
			str_a, err := strconv.Atoi(args[6]);
            if err != nil {
				fmt.Println(err);
			}
			indexar(args[4], args[2], str_a);
			fmt.Println("Attack ended with dictionary", args[2])
		}
		if(args[1]=="-w" && args[3]=="-ig" && args[5]=="-u"){
			str_b, err := strconv.Atoi(args[4]);
            if err != nil {
				fmt.Println(err);
			}
			indexar(args[6], args[2], str_b);
			fmt.Println("Attack ended with dictionary", args[2])
		}
		if(args[1]=="-ig" && args[3]=="-u" && args[5]=="-w"){
			str_c, err := strconv.Atoi(args[2]);
            if err != nil {
				fmt.Println(err);
			}
			indexar(args[4], args[6], str_c);
			fmt.Println("Attack ended with dictionary", args[2])
		}
		if(args[1]=="-ig" && args[3]=="-w" && args[5]=="-u"){
			str_d, err := strconv.Atoi(args[2]);
            if err != nil {
				fmt.Println(err);
			}
			indexar(args[6], args[4], str_d);
			fmt.Println("Attack ended with dictionary", args[2])
		}
		
	}
}
func indexar(url string, file string, response int){
	red, err := os.Open(file);
	if err != nil {
		fmt.Println(err);
		return
	}
	defer red.Close();
	scan := bufio.NewScanner(red);
	for scan.Scan(){
		resp, err := http.Get(url+scan.Text());
		if err != nil {
			fmt.Println(err);
		}
		time.Sleep(1*time.Millisecond);
		var local int = resp.StatusCode;
		if local != response {
			fmt.Println("Response: ", resp.StatusCode, url+scan.Text());
		}
	}
}
func banner(){
	lista := [3]string{"ban1.txt", "ban2.txt", "ban3.txt"};
	num := rand.Intn(3);
	cmd := exec.Command("cat", "./banner/"+lista[num]);
	output, err := cmd.Output();
	if err != nil {
		fmt.Println(err);
	}
	fmt.Println("[*] CrackBuster Fuzzing Attack Initiated")
	fmt.Println(string(output));
}
