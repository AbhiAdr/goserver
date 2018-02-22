package models

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"

	"go_mjolnir/structs"
	"io"
	"io/ioutil"
	"os"
)

var path = "/home/adr/gocode/src/go_mjolnir/license.txt"
var key = []byte("b9e2eb85b3af6e9a3f7877fa352e76ed")

func CreateLicenseFile() {

	var _, err = os.Stat(path)
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		checkErr(err)

		defer file.Close()
	}
	fmt.Println("==> done creating file", path)
	// writeLicenseFile()
}

// func writeLicenseFile() {
// 	// open file using READ & WRITE permission
// 	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
// 	checkErr(err)
// 	defer file.Close()

// 	lic_, err := encrypt([]byte("DemoVersion"))
// 	checkErr(err)
// 	date_ := time.Now().UTC().AddDate(1, 0, 0).Format("2006-01-02")
// 	exp, err := encrypt([]byte(date_))
// 	checkErr(err)

// 	demo_data := structs.LicenseJson{string(lic_), string(exp)}
// 	licenseJson, err := json.Marshal(demo_data)

// 	_, err = file.WriteString(string(licenseJson))
// 	checkErr(err)
// 	// save changes
// 	err = file.Sync()
// 	checkErr(err)

// 	fmt.Println("==> done writing to file")
// }

func readLicenseFile() {
	// re-open file
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	checkErr(err)
	defer file.Close()

	read_data, _ := ioutil.ReadAll(file)
	fmt.Println("==> done reading from file")

	fmt.Println(string(read_data))
	license_data := structs.LicenseJson{}
	err = json.Unmarshal([]byte(read_data), &license_data)
	checkErr(err)
	// fmt.Println(decrypt([]byte(license_data.Demo)))
	// fmt.Println(decrypt([]byte(license_data.Expiry)))
}

func deleteFile() {
	// delete file
	var err = os.Remove(path)
	checkErr(err)

	fmt.Println("==> done deleting file")
}

func Encrypt(text []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	ciphertext := make([]byte, aes.BlockSize+len(text))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	cfb := cipher.NewCFBEncrypter(block, iv)
	cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(text))
	encoded := base64.StdEncoding.EncodeToString(ciphertext)
	return []byte(encoded), nil
}

func Decrypt(encoded []byte) ([]byte, error) {
	data, err := base64.StdEncoding.DecodeString(string(encoded))
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	if len(data) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}
	iv := data[:aes.BlockSize]
	data = data[aes.BlockSize:]
	cfb := cipher.NewCFBDecrypter(block, iv)
	cfb.XORKeyStream(data, data)
	return data, nil
}
