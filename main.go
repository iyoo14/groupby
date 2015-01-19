package main

import (
    "fmt"
    "bufio"
    "os"
    "flag"
    "strings"
    "strconv"
    )

func main() {
    var f = flag.String("f", "", "file path.")
    var c = flag.Bool("c", false, "disp count for bool")
    var k = flag.String("k", "1", " target keys index. default 1")
    var v = flag.String("v", "2", " target values index. default 2")
    var d = flag.String("d", "\t", "demilter. default \t")
    flag.Parse()

    // argc
    if flag.NArg() > 0 {
        fmt.Printf("error:illegale args.\n")
        os.Exit(1)
    }

    // arg f
    if *f != "" {
        if strings.Index(*f, "-") == 0 {
            fmt.Printf("f is  %v\n", *f)
            os.Exit(1)
        }
    }
    var fp *os.File
    var err error

    if *f == "" {
        fp = os.Stdin
    } else {
        fp, err = os.Open(*f)
        if err != nil {
            panic(err)
        }
        defer fp.Close()
    }

    // arg k
    tkarray := strings.Split(*k, ",")
    var karray = make([]int, len(tkarray))
    for i, kv := range tkarray {
        ikv, _ := strconv.Atoi(kv)
        karray[i] = ikv
    }
    // arg v
    tvarray := strings.Split(*v, ",")
    var varray = make([]int, len(tvarray))
    for i, vv := range tvarray {
        ivv, _ := strconv.Atoi(vv)
        varray[i] = ivv
    }

    var inarray []string
    var cnt int = 0
    var kmap = make(map[int]string)
    var lkmap = make(map[int]string)
    var vmap = make(map[int]int)

    scanner := bufio.NewScanner(fp)
    for scanner.Scan() {
        val := scanner.Text()
        inarray = strings.SplitN(val, *d, 100)
        for _, k := range karray {
            kmap[k] = inarray[k-1]
        }
        mapcopy(lkmap, kmap)

        for _, k := range varray {
            v, _ := strconv.Atoi(inarray[k-1])
            vmap[k] = v
        }

        cnt++
        break
    }
    for scanner.Scan() {
        val := scanner.Text()
        inarray = strings.SplitN(val, *d, 100)
        for _, k := range karray {
            kmap[k] = inarray[k-1]
        }
        if comparemap(kmap, lkmap) {
            for _, k := range varray {
                v, _ := strconv.Atoi(inarray[k-1])
                vmap[k] += v
            }
            cnt++
        } else {
            disp(lkmap,vmap)
            fmt.Printf("\n")
            for _, k := range varray {
                v, _ := strconv.Atoi(inarray[k-1])
                vmap[k] = v
            }
            cnt = 0
        }
        mapcopy(lkmap, kmap)

    }
    disp(lkmap,vmap)
    if *c {
        fmt.Printf("\t%d", cnt)
    }
    fmt.Printf("\n")
    if err := scanner.Err(); err != nil {
        panic(err)
    }
}

func mapcopy(dst map[int]string, src map[int]string) {
    for k, v := range src {
        dst[k] = v
    }
}

func comparemap(amap map[int]string, bmap map[int]string) bool {
    for k, v := range amap {
        if v != bmap[k] {
            return false
        }
    }
    return true
}

func initmap(amap map[int]int) {
    for k, _ := range amap {
        amap[k] = 0
    }
}

func disp(kmap map[int]string, vmap map[int]int) {
    for _, v := range kmap {
        fmt.Printf("%s", v)
    }
    fmt.Printf("\t")
    for _, v := range vmap {
        fmt.Printf("%d", v)
    }
}
