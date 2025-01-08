#include<iostream>
#include<string.h>
#include<stdio.h>
#include<cstdlib>
#include<stdlib.h>
#include<cstring>
#include<sstream>
#include<fstream>
#include "httplib.h"

using namespace std;

class Decores{
	private:
		int var;
	public:
		int function_banner(){
			system("cat banner.txt");
			return 0;
		}
};
class Text{
	private:
		int var;
	public:
		int help_function(){
			string var[] = {"===============", 
				"#CRACKBUSTER", "===============", "./crackbuster <url> <wordlists>"};
			size_t local = sizeof(var)/sizeof(string);
			for(int x=0; x<local;x++){
				cout<<var[x]<<endl;
			}
			return 0;
		}
};
class System{
	private:
		int var;
	public:
		int buster_function(string url, string directory){
			httplib::Client cli(url);
			auto res = cli.Get(directory);
			if(res && res->status != 404){
				cout<<url<<"/"<<directory<<" status: "<<res->status<<endl;
			}
			return 0;
		}
};


int main(int argc, char *argv[]){
	Decores ban;
	Decores* nab = &ban;
	int flag_a = nab->function_banner();
	if(argc<=1){
		cout<<"specify an argument, use -h for more information"<<endl;
	}
	if(argc==2){
		char help[] = "-h";
		int flag_diferent = strcmp(argv[1], help);
		if(flag_diferent==0){
			Text ctrl;
			Text* control = &ctrl;
			control->help_function();
		}else{
			cout<<"Invalid Argument"<<endl;
		}
	}
	if(argc==3){
		cout<<"{+} crackbuster initiating fuzzing attack"<<endl;
		//BUSTER FUNCTION
		System sys;
		System* ysy = &sys;
		//ysy->buster_function(argv[1], argv[2]);
		ifstream archivo(argv[2]);
		string linea;
		if(archivo.is_open()){
			while(getline(archivo, linea)){
				ysy->buster_function(argv[1], linea);
			}
		}

	}
	cout<<"CrackBuster Finish, attack Complete! "<<endl;
	return 0;
}
