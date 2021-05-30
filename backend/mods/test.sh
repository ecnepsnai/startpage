#!/bin/bash

for MODULE in $(find . -maxdepth 1 -type d | grep './' | sed 's@./@@'); do
    printf "${MODULE} "
    rm -f mods.go
    touch mods.go
    echo "package mods" >> mods.go
    echo "import (" >> mods.go
    echo "\"fmt\"" >> mods.go
    echo "\"github.com/ecnepsnai/startpage/mods/${MODULE}\"" >> mods.go
    echo ")" >> mods.go
    echo "func Foo() {" >> mods.go
    echo "i, err := ${MODULE}.Setup(\"\", &${MODULE}.Options{})" >> mods.go
    echo "fmt.Println(err.Error())" >> mods.go
    echo "fmt.Println(i.Refresh().Error())" >> mods.go
    echo "fmt.Printf(\"%v\", i.Get())" >> mods.go
    echo "i.Teardown()" >> mods.go
    echo "}" >> mods.go
    go build > /dev/null 2>&1
    if [ $? == 0 ]; then
        rm -f mods.go
        printf "\033[0;32mPASS\033[0m\n"
    else
        printf "\033[0;31mFAIL\033[0m\n"
        go build
        rm -f mods.go
        exit 1
    fi
done
