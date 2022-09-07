package main

import (
  "flag"
  "fmt"
  "os"
  "log"
  "path/filepath"
)

func main() {
  dir, err := os.Getwd()
  if err != nil {
    log.Fatal(err)
  }

  var path string
  flag.StringVar(&path, "path", dir, "root path")
  flag.Parse()

  // fmt.Println(dir)
  // get files by os.ReadDir
  fmt.Println("---get files from os.ReadDir---")
  files, err := OSReadDir(path)
  if err != nil {
    log.Fatal(err)
    panic(err)
  }

  for _, f := range files {
    fmt.Println(f)
  }

  // get files by filepath
  fmt.Println("---get files from filepath---")
  files, err = FilePathWalkDir(path)
  if err != nil {
    log.Fatal(err)
    panic(err)
  }
  for _, f := range files {
    fmt.Println(f)
  }

  // get files by os.File.Readdir
  fmt.Println("---get files from os.File.Readdir---")
  files, err = FromOSReadDir(path)
  if err != nil {
    log.Fatal(err)
    panic(err)
  }

  for _, f := range files {
    fmt.Println(f)
  }
}

func FromOSReadDir(root string) ([]string, error) {
  var files []string
  f, err := os.Open(root)
  if err != nil {
    return files, err
  }

  fileInfo, err := f.Readdir(-1)
  f.Close()
  if err != nil {
    return files, err
  }

  for _, file := range fileInfo {
    files = append(files, file.Name())
  }

  return files, nil
}

func OSReadDir(root string) ([]string, error) {
  var files []string
  filesInfo, err := os.ReadDir(root)

  if err != nil {
    return files, err
  }

  for _, f := range filesInfo {
    files = append(files, f.Name())
  }

  return files, nil
}

func FilePathWalkDir(root string)([]string, error) {
  var files []string
  err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
    if !info.IsDir() {
      files = append(files, path)
    }

    return nil
  })

  return files, err
}
