# drill

Find files under the base path.

## Installation

```
go get github.com/hiroakis/drill
```

## Usage

```
drill.Down("BATH_PATH", DEPTH, IS_RESULT_ABS_PATH_OR_NOT, IS_RESULT_INCLUDING_DIRECTORY)
```

## Example

directory structure

```
.
├── dir1
│   ├── dir2
│   │   ├── dir3
│   │   │   └── file3
│   │   └── file2
│   └── file1
├── file
└── main.go

3 directories, 5 files
```

main.go

```
package main

import (
    "github.com/hiroakis/drill"
    "fmt"
)

func main() {
    files, err := drill.Down(".", 2, false, true)
     if err != nil {
        fmt.Println(err)
    }
    for _, v := range files {
        fmt.Println(v)
    }
```

Following is the result.

```
.
dir1
dir1/dir2
dir1/file1
file
main.go
```

## Use case

* check the number of files

```
files, err := drill.Down(".", 2, false, false)
if err != nil {
        fmt.Println(err)
}
fmt.Println(len(files))
```

* check all of the file size under base directory

```
files, err := drill.Down(".", 3, false, false)
if err != nil {
        fmt.Println(err)
}
var stat os.FileInfo
var info string
for _, v := range files {

        stat, _ = os.Stat(v)
        info = fmt.Sprintf("%d %s", stat.Size(), v)
        fmt.Println(info)
}
```

## License

MIT
