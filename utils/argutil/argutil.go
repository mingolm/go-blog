package argutil

import (
	"flag"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unsafe"
)

func Parse(v interface{}) {
	ParseSlice(v, os.Args[1:])
}

func ParseSlice(v interface{}, args []string) {
	ps := reflect.ValueOf(v).Elem()
	for i := 0; i < ps.Type().NumField(); i++ {
		field := ps.Type().Field(i)
		envTag := field.Tag.Get("env")
		if envTag != "" {
			envVal := strings.TrimSpace(os.Getenv(envTag))
			if envVal != "" {
				vField := reflect.ValueOf(v).Elem().FieldByName(field.Name)
				switch vField.Kind() {
				case reflect.String:
					vField.SetString(envVal)
				case reflect.Bool:
					b, _ := strconv.ParseBool(envVal)
					vField.SetBool(b)
				case reflect.Int:
					fallthrough
				case reflect.Int8:
					fallthrough
				case reflect.Int16:
					fallthrough
				case reflect.Int32:
					fallthrough
				case reflect.Int64:
					v, _ := strconv.ParseInt(envVal, 10, 0)
					vField.SetInt(v)
				case reflect.Uint:
					fallthrough
				case reflect.Uint8:
					fallthrough
				case reflect.Uint16:
					fallthrough
				case reflect.Uint32:
					fallthrough
				case reflect.Uint64:
					v, _ := strconv.ParseUint(envVal, 10, 0)
					vField.SetUint(v)
				}
			}
		}
		flagTag := field.Tag.Get("flag")
		if flagTag != "" {
			vField := reflect.ValueOf(v).Elem().FieldByName(field.Name)
			switch vField.Kind() {
			case reflect.String:
				flag.StringVar((*string)(unsafe.Pointer(vField.Addr().Pointer())), flagTag, vField.String(), field.Tag.Get("flagUsage"))
			case reflect.Bool:
				flag.BoolVar((*bool)(unsafe.Pointer(vField.Addr().Pointer())), flagTag, vField.Bool(), field.Tag.Get("flagUsage"))
			case reflect.Int:
				flag.IntVar((*int)(unsafe.Pointer(vField.Addr().Pointer())), flagTag, int(vField.Int()), field.Tag.Get("flagUsage"))
			case reflect.Int64:
				flag.Int64Var((*int64)(unsafe.Pointer(vField.Addr().Pointer())), flagTag, vField.Int(), field.Tag.Get("flagUsage"))
			case reflect.Uint:
				flag.UintVar((*uint)(unsafe.Pointer(vField.Addr().Pointer())), flagTag, uint(vField.Uint()), field.Tag.Get("flagUsage"))
			case reflect.Uint64:
				flag.Uint64Var((*uint64)(unsafe.Pointer(vField.Addr().Pointer())), flagTag, vField.Uint(), field.Tag.Get("flagUsage"))
			default:
				panic(fmt.Sprintf("kind of type %s not support for flag", vField.Kind()))
			}
		}
	}

	if err := flag.CommandLine.Parse(args); err != nil {
		panic(err)
	}
}
