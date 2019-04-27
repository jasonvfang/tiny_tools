package pack

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"io/ioutil"
)


func createDirectory(dirPath string) error {
	os.RemoveAll(dirPath)
	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		panic(err)
	}

	return err
}


func Prepare() {
	log.Println("1.Please enter the build Region (US | EU | AU):")
	fmt.Scanln(&Region)

	log.Println("2.Please enter the build release(45 etc):")
	fmt.Scanln(&ReleaseVersion)

	log.Println("3.Is it for Debug(true | false):")
	fmt.Scanln(&IsForDebug)

	//To confirm
	if len(Region) == 0 {
		log.Println("Region is not Specified, Exit !!")
		os.Exit(1)
	}

	RegionList := []string{"EU", "US", "AU"}
	Region = strings.ToUpper(Region)

	i := sort.Search(len(RegionList), func(i int) bool { return Region <= RegionList[i] })

	if i >= len(RegionList) {
		log.Println("Invalid Region, We support AU/US/EU only, Exit !!")
		os.Exit(1)
	}

	if ReleaseVersion <= 0 {
		log.Println("Release is not Specified, Exit !!")
		os.Exit(1)
	}

	log.Println("\n\n================Confirmation===========")
	log.Println("Region=", Region)
	log.Println("Version=", ReleaseVersion)
	log.Println("isDebug=", IsForDebug)
	log.Println("\n=======================================\n")

}


func GetCurrentDirectory() (string,error){
	dir, err := os.Getwd()

	if err != nil {
		log.Println("GetCurrentDirectory error")
	}

	return dir, err
}


func ExecuteSystemCommand(cmdString string) bool {

	if (cmdString == "" || len(cmdString) <= 0) {
		log.Println("Execute cmd error, Cmd is nil")
		return false
	}
	
	shCmdString := "/bin/sh"
	
	shOptionList := []string{"-c"}

	//append cmd
	shOptionList = append(shOptionList,cmdString)

	//Create command but not started yet
	cmd := exec.Command(shCmdString, shOptionList...)
  	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
    if err != nil {
        log.Println("Execute cmd err:", err)
        return false
    }

    return true
}


func CreateDirectoryPerRegion(){
	var otaDirPath string

	otaDirPath = RootOtaDirNam + "/"
	
	os.RemoveAll(otaDirPath)

	if Region == "EU" {
		otaDirPath += "EU/"
	}

	otaDirPath += strconv.Itoa(ReleaseVersion)

	log.Println("otaDirPath is", otaDirPath)

	err := createDirectory(otaDirPath)

	if err != nil {
		log.Println("Create directory fatal error !")
		os.Exit(1)
	}

	DstOtaPkgPath = otaDirPath

	err = createDirectory(ConfigOutputOtaPackageDir)
	if err != nil {
		log.Println("Create directory fatal error !")
		os.Exit(1)
	}
}


func GetFileListFromDir(dirName string)([]os.FileInfo, error){
	files, err := ioutil.ReadDir(dirName)
    if err != nil {
        log.Println(err)
        return []os.FileInfo(nil), err
    }	
/*
    for _, f := range files {
        log.Println(f.Name())
    }
*/

	return files, err
}


func GetCorrectHsFileName()string {
	var hsFileName string

	files, err:= GetFileListFromDir(ConfigHandsetImgDir)

	if err != nil {
		log.Println("get dirlist failed!")
		return ""
	}

    for _, f := range files {
        //log.Println(f.Name())
        if strings.HasSuffix(f.Name(), Region) {
        	hsFileName = f.Name()
        	break
        }
    }
	
	return hsFileName
}


func GetCorrectR18FileName()string {
	var r18FileName string


	if IsForDebug == true {
		r18FileName = (ConfigDebugR18ImgDir)
	}else{
		r18FileName = (ConfigReleaseR18ImgDir)
	}

	return r18FileName
}


func CopyFiles() int{
	var CpCmd string

	CpCmd = "cp -rf "
	
	CurrentDir, err := GetCurrentDirectory()
	if err != nil {
		return -1
	}

	log.Println("Current direcotry:", CurrentDir)

	//1.cmbs
	cmbsPostString := "V00" + strconv.Itoa(ReleaseVersion) + "." + Region
	cmbsFilePath := ConfigCmbsImgDir + CmbsPrefixString + cmbsPostString + "/*"
	
	cmdString := CpCmd + cmbsFilePath + " " + DstOtaPkgPath

	succ := ExecuteSystemCommand(cmdString)
	if succ != true {
		log.Println("Copy Cmbs failed ! ==>", cmbsFilePath)
		return -1
	}

	//2.hs
	hsName := GetCorrectHsFileName()
	if (hsName == ""){
		log.Println("Copy hs failed, hs name is nil !")
		return -1	
	}

	hsFilePath := ConfigHandsetImgDir + hsName + "/*"
	cmdString = CpCmd + hsFilePath + " " + DstOtaPkgPath

	succ = ExecuteSystemCommand(cmdString)
	if succ != true {
		log.Println("Copy hs failed ! ==>", hsName)
		return -1
	}

	//3.R18
	r18FileName := GetCorrectR18FileName()
	if (r18FileName == ""){
		log.Println("Copy r18 failed, name is nil !")
		return -1	
	}

	cmdString = CpCmd + r18FileName + "/* " + DstOtaPkgPath

	succ = ExecuteSystemCommand(cmdString)
	if succ != true {
		log.Println("Copy r18 failed ! ==>", r18FileName)
		return -1
	}

	return 0
}


func GenerateVersionFile() {
	cmdString := "cp -rf " + HelperToolDirPath + GenVersionScript + " " +DstOtaPkgPath
	
	log.Println(cmdString)

	succ := ExecuteSystemCommand(cmdString)
	if succ != true {
		log.Println(cmdString, "Failed")
		return
	}

	cmdString = "cd " + DstOtaPkgPath + ";" + "./" + GenVersionScript + ";" + "rm -rf " + GenVersionScript
	
	log.Println(cmdString)

	succ = ExecuteSystemCommand(cmdString)
	if succ != true {
		log.Println(cmdString, "Failed")
	}
}


func DoPreImagePack() bool {
	var tmpFileName string 
	var tmpFileType string 

	files, err:= GetFileListFromDir(DstOtaPkgPath)

	if err != nil {
		log.Println("get dirlist failed!")
		return false
	}

    for _, f := range files {
    	
        if strings.HasPrefix(f.Name(), CmbsPrefixString) {
        	log.Println("Pre-Pack-->",f.Name())
        	tmpFileName = f.Name()
        	tmpFileType = "cmbs"
        }else if strings.Contains(f.Name(), HandsetPrefixString){
        	log.Println("Pre-Pack-->",f.Name())
        	tmpFileName = f.Name()
        	tmpFileType = "handset"
        }else if strings.HasPrefix(f.Name(), BaseKernelPrefixString) {
        	log.Println("Pre-Pack-->",f.Name())
        	tmpFileName = f.Name()
        	tmpFileType = "boot"
        }else if strings.HasPrefix(f.Name(), BaseRootfsPrefixString) {
        	log.Println("Pre-Pack-->",f.Name())
        	tmpFileName = f.Name()
        	tmpFileType = "rootfs"
        }else if strings.HasPrefix(f.Name(), ScriptName) {
        	log.Println("Pre-Pack-->",f.Name())
        	tmpFileName = f.Name()
        	tmpFileType = "script"
        }else {
        	continue
        }               

        if tmpFileType != "" && tmpFileName != "" {
        	
        	outPackedFileName := DstOtaPkgPath + "/" + tmpFileName + "_packed"

        	cmdString := HelperToolPackImg + " " + PackTargetName + " " + tmpFileType + " " + DstOtaPkgPath + "/" + tmpFileName + " " + outPackedFileName	
        	log.Println(cmdString)
			succ := ExecuteSystemCommand(cmdString)
			if succ != true {
				log.Println(cmdString, "Failed")
				return false
			}        	

			cmdString = HelperToolOpenSSL + " " + "enc -e -aes-128-cbc -kfile " + HelperAESKeyFile + " -in " + outPackedFileName + " -out " + DstOtaPkgPath + "/" + tmpFileName
			log.Println(cmdString)
			succ = ExecuteSystemCommand(cmdString)
			if succ != true {
				log.Println(cmdString, "Failed")
				return false
			}   

			cmdString = "rm -rf " + outPackedFileName
			log.Println(cmdString)
			succ = ExecuteSystemCommand(cmdString)
			if succ != true {
				log.Println(cmdString, "Failed")
				return false
			}   			
        }	

        //reset strings
        tmpFileName = ""
        tmpFileType = ""
        fmt.Println();
    }

    //Generate MD5 
    cmdString := "cd " + DstOtaPkgPath + ";" + "md5sum * > md5.txt"
	log.Println(cmdString)
	succ := ExecuteSystemCommand(cmdString)
	if succ != true {
		log.Println(cmdString, "Failed")
		return false
	}   		
	
	cmdString = "cp -fr " + ReleaseLinkplaySdkVerFile + " " + DstOtaPkgPath
	log.Println(cmdString)
	succ = ExecuteSystemCommand(cmdString)
	if succ != true {
		log.Println(cmdString, "Failed")
		return false
	}  

	return true
}


func GenerateProductXML(){

}


func DoPostImagePack(){

}


func DoFinallyFileZip(){

}


func DoPack() {
	log.Println("Doing task")

	CreateDirectoryPerRegion()

	CopyFiles()
	
	GenerateVersionFile()

	DoPreImagePack()

	GenerateProductXML()
	
	DoPostImagePack()

	DoFinallyFileZip()
}


func CleanUp() {
	log.Println("Clean UP")
}

