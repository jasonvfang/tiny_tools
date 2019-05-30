package pack

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"sort"
	//"strconv"
	"strings"
	"io/ioutil"
	"time"
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
	log.Println("[1]Prepare:")
	log.Println("1.Please enter the build Region (US | EU | AU):")
	fmt.Scanln(&GlobalRelease.region)

	log.Println("2.Please enter R18 release(50 | 50.2 etc):")
	fmt.Scanln(&GlobalRelease.r18ReleaseNum)

	log.Println("3.Please enter cmbs release(50 etc):")
	fmt.Scanln(&GlobalRelease.cmbsReleaseNum)

	log.Println("4.Please enter handset release(37 etc):")
	fmt.Scanln(&GlobalRelease.handsetReleaseNum)

	log.Println("5.Is it for Debug(true | false):")
	fmt.Scanln(&GlobalRelease.IsForDebug)

	//To confirm
	if len(GlobalRelease.region) == 0 {
		log.Println("Region is not Specified, Exit !!")
		os.Exit(1)
	}

	RegionList := []string{"EU", "US", "AU"}
	GlobalRelease.region = strings.ToUpper(GlobalRelease.region)

	i := sort.Search(len(RegionList), func(i int) bool { return GlobalRelease.region <= RegionList[i] })

	if i >= len(RegionList) {
		log.Println("Invalid Region, We support AU/US/EU only, Exit !!")
		os.Exit(1)
	}

	if GlobalRelease.r18ReleaseNum <= 0 {
		log.Println("Release is not Specified, Exit !!")
		os.Exit(1)
	}

	GlobalRelease.r18ReleaseShort = int(GlobalRelease.r18ReleaseNum)

	log.Println("\n\n================Confirmation===========")
	log.Println("Region=", GlobalRelease.region)
	log.Println("R18Version=", GlobalRelease.r18ReleaseNum)
	log.Println("r18ReleaseShort=", GlobalRelease.r18ReleaseShort)
	log.Println("CmbsVersion=", GlobalRelease.cmbsReleaseNum)
	log.Println("HandsetVersion=", GlobalRelease.handsetReleaseNum)
	log.Println("isDebug=", GlobalRelease.IsForDebug)
	log.Println("======================\n")

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
	
	log.Println(cmdString)

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
	
	os.RemoveAll(PackTargetName)

	if GlobalRelease.region == "EU" {
		//otaDirPath += "EU/"
		if GlobalRelease.r18ReleaseNum == float32(GlobalRelease.r18ReleaseShort) {
			otaDirPath = fmt.Sprintf("%s/EU/%d", RootOtaDirNam, GlobalRelease.r18ReleaseShort)
		}else{
			otaDirPath = fmt.Sprintf("%s/EU/%.1f", RootOtaDirNam, GlobalRelease.r18ReleaseNum)
		}
		
	}else {
		if GlobalRelease.r18ReleaseNum == float32(GlobalRelease.r18ReleaseShort) {
			otaDirPath = fmt.Sprintf("%s/%d", RootOtaDirNam, GlobalRelease.r18ReleaseShort)
		}else{
			otaDirPath = fmt.Sprintf("%s/%.1f", RootOtaDirNam, GlobalRelease.r18ReleaseNum)
		}
	}

	log.Println("Dest OtaDirPath is", otaDirPath)

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
        if strings.HasSuffix(f.Name(), GlobalRelease.region) {
        	hsFileName = f.Name()
        	break
        }
    }
	
	return hsFileName
}


func GetCorrectR18FileName()string {
	var r18FileName string


	if GlobalRelease.IsForDebug == true {
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
	//cmbsPostString := "V00" + strconv.Itoa(GlobalRelease.r18ReleaseNum) + "." + GlobalRelease.region
	
	var shortCmbsReleaseNum = int(GlobalRelease.cmbsReleaseNum)
	var cmbsPostString string

	if float32(shortCmbsReleaseNum) == GlobalRelease.cmbsReleaseNum {
		cmbsPostString = fmt.Sprintf("V00%d.%s", shortCmbsReleaseNum, GlobalRelease.region)
	}else{
		cmbsPostString = fmt.Sprintf("V00%1f.%s", GlobalRelease.cmbsReleaseNum, GlobalRelease.region)
	}
	
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
        	tmpFileName = f.Name()
        	tmpFileType = "cmbs"
        	CmbsFileName = tmpFileName
        }else if strings.Contains(f.Name(), HandsetPrefixString){
        	tmpFileName = f.Name()
        	tmpFileType = "handset"
        	HsFileName = tmpFileName
        }else if strings.HasPrefix(f.Name(), BaseKernelPrefixString) {
        	tmpFileName = f.Name()
        	tmpFileType = "boot"
        	KernelFileName = tmpFileName
        }else if strings.HasPrefix(f.Name(), BaseRootfsPrefixString) {
        	tmpFileName = f.Name()
        	tmpFileType = "rootfs"
        	RootfsFileName = tmpFileName
        }else if strings.HasPrefix(f.Name(), ScriptName) {
        	tmpFileName = f.Name()
        	tmpFileType = "script"
        	ScriptFileName = tmpFileName
        }else {
        	continue
        }               

        if tmpFileType != "" && tmpFileName != "" {
        	
        	log.Println("Packing==>", tmpFileName)

        	outPackedFileName := DstOtaPkgPath + "/" + tmpFileName + "_packed"

        	cmdString := HelperToolPackImg + " " + PackTargetName + " " + tmpFileType + " " + DstOtaPkgPath + "/" + tmpFileName + " " + outPackedFileName	
        	//log.Println(cmdString)
			succ := ExecuteSystemCommand(cmdString)
			if succ != true {
				log.Println(cmdString, "Failed")
				return false
			}        	

			cmdString = HelperToolOpenSSL + " " + "enc -e -aes-128-cbc -kfile " + HelperAESKeyFile + " -in " + outPackedFileName + " -out " + DstOtaPkgPath + "/" + tmpFileName
			//log.Println(cmdString)
			succ = ExecuteSystemCommand(cmdString)
			if succ != true {
				log.Println(cmdString, "Failed")
				return false
			}   

			cmdString = "rm -rf " + outPackedFileName
			//log.Println(cmdString)
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
	//log.Println(cmdString)
	succ := ExecuteSystemCommand(cmdString)
	if succ != true {
		log.Println(cmdString, "Failed")
		return false
	}   		
	
	cmdString = "cp -fr " + ReleaseLinkplaySdkVerFile + " " + DstOtaPkgPath
	//log.Println(cmdString)
	succ = ExecuteSystemCommand(cmdString)
	if succ != true {
		log.Println(cmdString, "Failed")
		return false
	}  

	return true
}

func check(e error) {
    if e != nil {
        panic(e)
    }
}


func generateProductXML(pkgDirPrefix string, urlPrefixString string) bool {
	CurrentLocalTime := time.Now()

	//log.Println("Current local time is", CurrentLocalTime)
	//Year := CurrentLocalTime.Year()
	//Mon := CurrentLocalTime.Month()
	//Day := CurrentLocalTime.Day()
	
	Year, Mon, Day := CurrentLocalTime.Date()

	//TimeTagString := strconv.Itoa(Year) + strconv.Itoa(int(Mon)) + strconv.Itoa(Day)
	TimeTagString := fmt.Sprintf("%04d%02d%02d", Year, int(Mon), Day)
	log.Println("Current local time is", TimeTagString)

	productXMLFile := DstOtaPkgPath + "/" + PackageProductionXmlFile

    f, err := os.Create(productXMLFile)
    check(err)
 	
 	defer func (){
 		f.Sync()
 		f.Close()
	}()

 	var writeBuffer string 

 	writeBuffer = "<?xml version=\"1.0\" encoding=\"ISO-8859-1\" ?>\n"
 	f.WriteString(writeBuffer)

 	writeBuffer = "<product>\n"
	f.WriteString(writeBuffer)

	writeBuffer = fmt.Sprintf("<major-version>%s</major-version>\n", TimeTagString)
	f.WriteString(writeBuffer)

	if GlobalRelease.r18ReleaseNum == float32(GlobalRelease.r18ReleaseShort) {
		writeBuffer = fmt.Sprintf("<md5-url>%s/%s/%d/md5.txt</md5-url>\n", urlPrefixString, pkgDirPrefix, GlobalRelease.r18ReleaseShort)
		f.WriteString(writeBuffer)

		writeBuffer = fmt.Sprintf("<ver-url>%s/%s/%d/version.txt</ver-url>\n", urlPrefixString, pkgDirPrefix, GlobalRelease.r18ReleaseShort)
		f.WriteString(writeBuffer)

		//hs
		writeBuffer = fmt.Sprintf("<image-handset>%s/%s/%d/%s</image-handset>\n", urlPrefixString, pkgDirPrefix, GlobalRelease.r18ReleaseShort, HsFileName)
		f.WriteString(writeBuffer)

		//cmbs
		writeBuffer = fmt.Sprintf("<image-cmbs>%s/%s/%d/%s</image-cmbs>\n", urlPrefixString, pkgDirPrefix, GlobalRelease.r18ReleaseShort, CmbsFileName)
		f.WriteString(writeBuffer)

		//kernel
		writeBuffer = fmt.Sprintf("<image-kernel>%s/%s/%d/%s</image-kernel>\n", urlPrefixString, pkgDirPrefix, GlobalRelease.r18ReleaseShort, KernelFileName)
		f.WriteString(writeBuffer)

		//rootfs
		writeBuffer = fmt.Sprintf("<image-rootfs>%s/%s/%d/%s</image-rootfs>\n", urlPrefixString, pkgDirPrefix, GlobalRelease.r18ReleaseShort, RootfsFileName)
		f.WriteString(writeBuffer)

		//script 
		writeBuffer = fmt.Sprintf("<image-script>%s/%s/%d/%s</image-script>\n", urlPrefixString, pkgDirPrefix, GlobalRelease.r18ReleaseShort, ScriptFileName)
		f.WriteString(writeBuffer)

	}else{
		writeBuffer = fmt.Sprintf("<md5-url>%s/%s/%0.1f/md5.txt</md5-url>\n", urlPrefixString, pkgDirPrefix, GlobalRelease.r18ReleaseNum)
		f.WriteString(writeBuffer)

		writeBuffer = fmt.Sprintf("<ver-url>%s/%s/%0.1f/version.txt</ver-url>\n", urlPrefixString, pkgDirPrefix, GlobalRelease.r18ReleaseNum)
		f.WriteString(writeBuffer)

		//hs
		writeBuffer = fmt.Sprintf("<image-handset>%s/%s/%0.1f/%s</image-handset>\n", urlPrefixString, pkgDirPrefix, GlobalRelease.r18ReleaseNum, HsFileName)
		f.WriteString(writeBuffer)

		//cmbs
		writeBuffer = fmt.Sprintf("<image-cmbs>%s/%s/%0.1f/%s</image-cmbs>\n", urlPrefixString, pkgDirPrefix, GlobalRelease.r18ReleaseNum, CmbsFileName)
		f.WriteString(writeBuffer)

		//kernel
		writeBuffer = fmt.Sprintf("<image-kernel>%s/%s/%0.1f/%s</image-kernel>\n", urlPrefixString, pkgDirPrefix, GlobalRelease.r18ReleaseNum, KernelFileName)
		f.WriteString(writeBuffer)

		//rootfs
		writeBuffer = fmt.Sprintf("<image-rootfs>%s/%s/%0.1f/%s</image-rootfs>\n", urlPrefixString, pkgDirPrefix, GlobalRelease.r18ReleaseNum, RootfsFileName)
		f.WriteString(writeBuffer)

		//script 
		writeBuffer = fmt.Sprintf("<image-script>%s/%s/%0.1f/%s</image-script>\n", urlPrefixString, pkgDirPrefix, GlobalRelease.r18ReleaseNum, ScriptFileName)
		f.WriteString(writeBuffer)
	}
	
	f.WriteString("</product>\n")

	return true
}


func generateProductListXML(pkgDirPrefix string, urlPrefixString string) bool {
	var writeBuffer string 

	productListXMLFile := DstOtaPkgPath + "/../" + PackageProductListXmlFile

    f, err := os.Create(productListXMLFile)
    check(err)

 	defer func (){
 		f.Sync()
 		f.Close()
	}()


	writeBuffer = "<?xml version=\"1.0\" encoding=\"ISO-8859-1\" ?>\n"
	f.WriteString(writeBuffer)

	f.WriteString("<productList>\n")

	f.WriteString("<product>\n")
	
	writeBuffer = fmt.Sprintf("<productid>%s</productid>\n", SgwProductName);
	f.WriteString(writeBuffer)

	writeBuffer = fmt.Sprintf("<hardwareversion>%s</hardwareversion>\n", SgwHardwareName);
	f.WriteString(writeBuffer)

	writeBuffer = fmt.Sprintf("<UUID>%s</UUID>\n", SgwProjectUUID);
	f.WriteString(writeBuffer)

	if GlobalRelease.r18ReleaseNum == float32(GlobalRelease.r18ReleaseShort) {
		writeBuffer = fmt.Sprintf("<major-url>%s/%s/%d/%s</major-url>\n", urlPrefixString, pkgDirPrefix, GlobalRelease.r18ReleaseShort, PackageProductionXmlFile);
	}else{
		writeBuffer = fmt.Sprintf("<major-url>%s/%s/%0.1f/%s</major-url>\n", urlPrefixString, pkgDirPrefix, GlobalRelease.r18ReleaseNum, PackageProductionXmlFile);
	}

	f.WriteString(writeBuffer)

	f.WriteString("</product>\n")

	f.WriteString("</productList>\n")

	return true
}


func GenerateProductXML() bool {
 	var pkgDirPrefix, urlPrefixString string

 	if GlobalRelease.region == "EU" {
 		pkgDirPrefix = PackTargetName + "/" + GlobalRelease.region	
 	}else{
 		pkgDirPrefix = PackTargetName
 	}

	if GlobalRelease.IsForDebug {
		urlPrefixString = URL_PREFIX_DEBUG
	}else{
		urlPrefixString = URL_PREFIX
	}

	generateProductXML(pkgDirPrefix, urlPrefixString)
	generateProductListXML(pkgDirPrefix, urlPrefixString)

	return true
}


func DoImagePackRsaEnc() bool {
	var cmdString string 
	cmdString = fmt.Sprintf("%s -encrypt -in %s/%s -inkey %s -pubin -out %s/product.xml", HelperToolRsaFile, DstOtaPkgPath, PackageProductionXmlFile, HelperRsaPublicKeyFile, DstOtaPkgPath)
	
	//log.Println(cmdString)
	succ := ExecuteSystemCommand(cmdString)
	if succ != true {
		log.Println(cmdString, "Failed")
		return false
	}  

	cmdString = fmt.Sprintf("%s -encrypt -in %s/../%s -inkey %s -pubin -out %s/../products.xml", HelperToolRsaFile, DstOtaPkgPath, PackageProductListXmlFile, HelperRsaPublicKeyFile, DstOtaPkgPath)
	//log.Println(cmdString)
	succ = ExecuteSystemCommand(cmdString)
	if succ != true {
		log.Println(cmdString, "Failed")
		return false
	}  

	return true
}


func DoFinallyFileZip() bool {
	
	var finallyOutputZipFileName string 

	if GlobalRelease.IsForDebug {
		if GlobalRelease.r18ReleaseNum == float32(GlobalRelease.r18ReleaseShort){
			finallyOutputZipFileName = fmt.Sprintf("%s/%s_ota_v00%d_debug_%s.zip", ConfigOutputOtaPackageDir, PackTargetName, GlobalRelease.r18ReleaseShort, GlobalRelease.region)	
		}else{
			finallyOutputZipFileName = fmt.Sprintf("%s/%s_ota_v00%0.1f_debug_%s.zip", ConfigOutputOtaPackageDir, PackTargetName, GlobalRelease.r18ReleaseNum, GlobalRelease.region)		
		}
	}else{
		finallyOutputZipFileName = fmt.Sprintf("%s/%s_ota_v00%0.1f_%s.zip", ConfigOutputOtaPackageDir, PackTargetName, GlobalRelease.r18ReleaseNum, GlobalRelease.region)
	}
	
	log.Println("Zip file ", finallyOutputZipFileName)

	cmdString := fmt.Sprintf("rm -fr %s ; zip -r %s %s/*", finallyOutputZipFileName, finallyOutputZipFileName, PackTargetName)
	//log.Println(cmdString)
	succ := ExecuteSystemCommand(cmdString)
	if succ != true {
		log.Println(cmdString, "Failed")
		return false
	}  

	return true
}


func DoPack() {
	log.Println("Doing task")
	var succ bool 

	CreateDirectoryPerRegion()

	log.Println("[2]COPY Files:")

	if 0 != CopyFiles() {
		log.Fatal("Copy files Failed");
	}
	
	log.Println("[3]Pack Files:")

	GenerateVersionFile()

	succ = DoPreImagePack()
	if succ != true {
		log.Fatal("DoPreImagePack Failed");
	}

	succ = GenerateProductXML()
	if succ != true {
		log.Fatal("GenerateProductXML Failed");
	}

	succ = DoImagePackRsaEnc()
	if succ != true {
		log.Fatal("DoImagePackRsaEnc Failed");
	}
	
	log.Println("[4]ZIP Files:")

	succ = DoFinallyFileZip()
	if succ != true {
		log.Fatal("DoPostImagePack Failed");
	}	


	log.Println("Great, Successfully to pack OTA image !!")
}


func CleanUp() {
	log.Println("Clean UP")
}

