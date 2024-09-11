package main

import (
    "bufio"
    "fmt"
    "io/ioutil"
    "log"
    "os"
    "path/filepath"
    "strconv"
    "strings"
)


func binToHex(bin string) string {
    remainder := len(bin) % 4
    if remainder != 0 {
        padding := 4 - remainder
        bin += strings.Repeat("0", padding)
    }

    var hexStr string
    for i := 0; i < len(bin); i += 4 {
        nibble := bin[i : i+4]
        val, err := strconv.ParseUint(nibble, 2, 4)
        if err != nil {
            log.Fatal(err)
        }
        hexStr += fmt.Sprintf("%X", val)
    }
    return hexStr
}


func hexToBin(hex string) string {
    var binStr string
    for i := 0; i < len(hex); i++ {
        hexChar := hex[i : i+1]
        val, err := strconv.ParseUint(hexChar, 16, 4)
        if err != nil {
            log.Fatal(err)
        }
        binStr += fmt.Sprintf("%04b", val)
    }
    return binStr
}


func convertInToX(filePath string) {
    file, err := os.Open(filePath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var output []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, ":")
        matrixInfo := parts[0]
        binaryData := parts[1]

     
        hexData := binToHex(binaryData)

        
        output = append(output, fmt.Sprintf("%s:%s", matrixInfo, hexData))
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

   
    newFilePath := filePath + ".x"
    ioutil.WriteFile(newFilePath, []byte(strings.Join(output, "\n")+"\n"), 0644)

    
    os.Remove(filePath)
}


func convertXToIn(filePath string) {
    file, err := os.Open(filePath)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    var output []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := scanner.Text()
        parts := strings.Split(line, ":")
        matrixInfo := parts[0]
        hexData := parts[1]

        
        binData := hexToBin(hexData)

       
        output = append(output, fmt.Sprintf("%s:%s", matrixInfo, binData))
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

   
    newFilePath := strings.TrimSuffix(filePath, ".x")
    ioutil.WriteFile(newFilePath, []byte(strings.Join(output, "\n")+"\n"), 0644)

   
    os.Remove(filePath)
}


func processPath(path string) {
    info, err := os.Stat(path)
    if err != nil {
        log.Fatal(err)
    }

    if info.IsDir() {
        files, err := ioutil.ReadDir(path)
        if err != nil {
            log.Fatal(err)
        }
        for _, file := range files {
            fileName := file.Name()
            if strings.HasSuffix(fileName, ".in") && !strings.HasSuffix(fileName, ".in.x") {
                convertInToX(filepath.Join(path, fileName))
            } else if strings.HasSuffix(fileName, ".in.x") {
                convertXToIn(filepath.Join(path, fileName))
            }
        }
    } else {
        if strings.HasSuffix(path, ".in") && !strings.HasSuffix(path, ".in.x") {
            convertInToX(path)
        } else if strings.HasSuffix(path, ".in.x") {
            convertXToIn(path)
        } else {
            fmt.Println("Invalid file format.")
        }
    }
}

func main() {
    if len(os.Args) != 2 {
        fmt.Println("Usage: go run convert.go <file-name>.in or <file-name>.in.x or /path/to/dir")
        return
    }

    path := os.Args[1]
    processPath(path)
}
