/*
Copyright (C) 2018 Jeroen Tietema

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// A simple Go CLI application that turns of autosuspend on USB input devices
// useful after running `powertop --auto-tune`
package main

import (
	"io/ioutil"
	"log"
	"strings"
)

var usbDevicesDir = "/sys/bus/usb/devices/"

func main() {
	dirInfo, err := ioutil.ReadDir(usbDevicesDir)
	if err != nil {
		log.Fatal("Can't read dir: ", usbDevicesDir, err)
	}
	for _, dir := range dirInfo {
		// fmt.Println(dir.Name())
		fileContent, err := ioutil.ReadFile(usbDevicesDir + dir.Name() + "/product")
		if err == nil {
			product := string(fileContent)
			// fmt.Print("Product: ", product)
			if strings.Contains(product, "USB Keyboard") || strings.Contains(product, "USB Receiver") {
				err := ioutil.WriteFile(usbDevicesDir+dir.Name()+"/power/control", []byte("on\n"), 0644)
				if err != nil {
					log.Fatal("Can't write file: ", usbDevicesDir, dir.Name(), "/power/control ", err)
				}
				log.Println("Set", strings.Trim(product, "\n"), "to \"on\"")
			}
		}

	}
}
