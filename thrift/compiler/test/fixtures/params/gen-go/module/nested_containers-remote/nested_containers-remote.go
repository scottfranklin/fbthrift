// Autogenerated by Thrift Compiler (facebook)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING
// @generated

package main

import (
        "flag"
        "fmt"
        "math"
        "net"
        "net/url"
        "os"
        "strconv"
        "strings"
        "git.apache.org/thrift.git/lib/go/thrift"
        "module"
)

func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  void mapList( foo)")
  fmt.Fprintln(os.Stderr, "  void mapSet( foo)")
  fmt.Fprintln(os.Stderr, "  void listMap( foo)")
  fmt.Fprintln(os.Stderr, "  void listSet( foo)")
  fmt.Fprintln(os.Stderr, "  void turtles( foo)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  var parsedUrl url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Parse()
  
  if len(urlString) > 0 {
    parsedUrl, err := url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  client := module.NewNestedContainersClientFactory(trans, protocolFactory)
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "mapList":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "MapList requires 1 args")
      flag.Usage()
    }
    arg30 := flag.Arg(1)
    mbTrans31 := thrift.NewTMemoryBufferLen(len(arg30))
    defer mbTrans31.Close()
    _, err32 := mbTrans31.WriteString(arg30)
    if err32 != nil { 
      Usage()
      return
    }
    factory33 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt34 := factory33.GetProtocol(mbTrans31)
    containerStruct0 := module.NewNestedContainersMapListArgs()
    err35 := containerStruct0.ReadField1(jsProt34)
    if err35 != nil {
      Usage()
      return
    }
    argvalue0 := containerStruct0.Foo
    value0 := argvalue0
    fmt.Print(client.MapList(value0))
    fmt.Print("\n")
    break
  case "mapSet":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "MapSet requires 1 args")
      flag.Usage()
    }
    arg36 := flag.Arg(1)
    mbTrans37 := thrift.NewTMemoryBufferLen(len(arg36))
    defer mbTrans37.Close()
    _, err38 := mbTrans37.WriteString(arg36)
    if err38 != nil { 
      Usage()
      return
    }
    factory39 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt40 := factory39.GetProtocol(mbTrans37)
    containerStruct0 := module.NewNestedContainersMapSetArgs()
    err41 := containerStruct0.ReadField1(jsProt40)
    if err41 != nil {
      Usage()
      return
    }
    argvalue0 := containerStruct0.Foo
    value0 := argvalue0
    fmt.Print(client.MapSet(value0))
    fmt.Print("\n")
    break
  case "listMap":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "ListMap requires 1 args")
      flag.Usage()
    }
    arg42 := flag.Arg(1)
    mbTrans43 := thrift.NewTMemoryBufferLen(len(arg42))
    defer mbTrans43.Close()
    _, err44 := mbTrans43.WriteString(arg42)
    if err44 != nil { 
      Usage()
      return
    }
    factory45 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt46 := factory45.GetProtocol(mbTrans43)
    containerStruct0 := module.NewNestedContainersListMapArgs()
    err47 := containerStruct0.ReadField1(jsProt46)
    if err47 != nil {
      Usage()
      return
    }
    argvalue0 := containerStruct0.Foo
    value0 := argvalue0
    fmt.Print(client.ListMap(value0))
    fmt.Print("\n")
    break
  case "listSet":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "ListSet requires 1 args")
      flag.Usage()
    }
    arg48 := flag.Arg(1)
    mbTrans49 := thrift.NewTMemoryBufferLen(len(arg48))
    defer mbTrans49.Close()
    _, err50 := mbTrans49.WriteString(arg48)
    if err50 != nil { 
      Usage()
      return
    }
    factory51 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt52 := factory51.GetProtocol(mbTrans49)
    containerStruct0 := module.NewNestedContainersListSetArgs()
    err53 := containerStruct0.ReadField1(jsProt52)
    if err53 != nil {
      Usage()
      return
    }
    argvalue0 := containerStruct0.Foo
    value0 := argvalue0
    fmt.Print(client.ListSet(value0))
    fmt.Print("\n")
    break
  case "turtles":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Turtles requires 1 args")
      flag.Usage()
    }
    arg54 := flag.Arg(1)
    mbTrans55 := thrift.NewTMemoryBufferLen(len(arg54))
    defer mbTrans55.Close()
    _, err56 := mbTrans55.WriteString(arg54)
    if err56 != nil { 
      Usage()
      return
    }
    factory57 := thrift.NewTSimpleJSONProtocolFactory()
    jsProt58 := factory57.GetProtocol(mbTrans55)
    containerStruct0 := module.NewNestedContainersTurtlesArgs()
    err59 := containerStruct0.ReadField1(jsProt58)
    if err59 != nil {
      Usage()
      return
    }
    argvalue0 := containerStruct0.Foo
    value0 := argvalue0
    fmt.Print(client.Turtles(value0))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
