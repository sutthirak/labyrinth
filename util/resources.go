package util

import (
	"os"
	"io/ioutil"
	"path/filepath"
	"github.com/revel/revel"
)

func GetResourceMaterialPath() (string , bool){
	dir , found := revel.Config.String("resources.dir")
	if found {
		return dir + "/material/" , true
	}
	revel.ERROR.Printf("Resources directory not found")
	return "" , false 
}

func GetResourceBabylonPath() (string , bool){
	dir , found := revel.Config.String("resources.dir")
	if found {
		return dir + "/babylon/" , true
	}
	return "" , false 
}

func GetResourceImagePath() (string , bool){
	dir , found := revel.Config.String("resources.dir")
	if found {
		return dir + "/image/" , true
	}
	return "" , false 
}

func GetResourceBabylon() []string {
	resources := make([]string, 0)
	dir , found := GetResourceBabylonPath()
	if found {
		files , _ := ioutil.ReadDir(dir)
		for _ , file := range files {
			if(filepath.Ext(file.Name()) == ".babylon"){
				resources = append(resources,file.Name())
			}
		}
	}

	return resources
}

func GetResourceImage() []string {
	resources := make([]string, 0)
	dir , found := GetResourceImagePath()
	if found {
		files , _ := ioutil.ReadDir(dir)
		for _ , file := range files {
			if(GetFileTypeSupportFromExtension(file) == "image" ){
				resources = append(resources,file.Name())
			}
		}
	}

	return resources
}

func GetResourceMaterial() []string {
	resources := make([]string, 0)
	dir , found := GetResourceMaterialPath()
	if found {
		files , _ := ioutil.ReadDir(dir)
		for _ , file := range files {
			if(GetFileTypeSupportFromExtension(file) == "image" ){
				resources = append(resources,file.Name())
			}
		}
	}

	return resources
}

func GetFileTypeSupportFromExtension(file os.FileInfo) string {
	if(filepath.Ext(file.Name()) == ".png") ||
					(filepath.Ext(file.Name()) == ".jpg") ||
					(filepath.Ext(file.Name()) == ".jpeg") ||
					(filepath.Ext(file.Name()) == ".gif"){

		return "image"
	}

	if(filepath.Ext(file.Name()) == ".babylon"){
		return "babylon"
	}

	return ""
}
