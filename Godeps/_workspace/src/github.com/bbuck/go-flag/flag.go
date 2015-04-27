package flag

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type flag struct {
	long, short string
	desc        string
	value       interface{}
}

type Value interface {
	Set(val interface{}) error
}

var flagDefs []flag = make([]flag, 0)
var finalArgs []string = make([]string, 0)

func keys(name string) (string, string, error) {
	parts := strings.Split(name, ",")
	for i, part := range parts {
		parts[i] = strings.Trim(part, " ")
	}
	if len(parts) == 0 {
		return "", "", errors.New("Invalid name given for flag!")
	}
	var long, short string
	long = parts[0]
	if len(parts) >= 2 {
		short = parts[1]
	} else {
		short = ""
	}

	return long, short, nil
}

func String(name string, def string, desc string) (*string, error) {
	long, short, err := keys(name)
	if err != nil {
		return nil, err
	}
	val := new(string)
	*val = def
	f := flag{
		long:  long,
		short: short,
		desc:  desc,
		value: val}
	flagDefs = append(flagDefs, f)

	return val, nil
}

func Int(name string, def int64, desc string) (*int64, error) {
	long, short, err := keys(name)
	if err != nil {
		return nil, err
	}
	val := new(int64)
	*val = def
	f := flag{
		long:  long,
		short: short,
		desc:  desc,
		value: val}
	flagDefs = append(flagDefs, f)

	return val, nil
}

func Float(name string, def float64, desc string) (*float64, error) {
	long, short, err := keys(name)
	if err != nil {
		return nil, err
	}
	val := new(float64)
	*val = def
	f := flag{
		long:  long,
		short: short,
		desc:  desc,
		value: val}
	flagDefs = append(flagDefs, f)

	return val, nil
}

func Bool(name string, def bool, desc string) (*bool, error) {
	long, short, err := keys(name)
	if err != nil {
		return nil, err
	}
	val := new(bool)
	*val = def
	f := flag{
		long:  long,
		short: short,
		desc:  desc,
		value: val}
	flagDefs = append(flagDefs, f)

	return val, nil
}

func Var(val interface{}, name string, desc string) error {
	long, short, err := keys(name)
	if err != nil {
		return err
	}
	f := flag{
		long:  long,
		short: short,
		desc:  desc,
		value: val}
	flagDefs = append(flagDefs, f)

	return nil
}

func Args() []string {
	return finalArgs
}

func assignFlagVal(flagVal flag, val string) error {
	switch t := flagVal.value.(type) {
	case (*int64):
		argVal, err := strconv.ParseInt(val, 10, 64)
		if err != nil {
			return errors.New(fmt.Sprintf("Invalid integer value specified for flag \"%s\"", flagVal.long))
		}
		*t = argVal
	case (*float64):
		argVal, err := strconv.ParseFloat(val, 64)
		if err != nil {
			return errors.New(fmt.Sprintf("Invalid float value specified for flag \"%s\"", flagVal.long))
		}
		*t = argVal
	case (*string):
		*t = val
	default:
		if fval, ok := flagVal.value.(Value); ok {
			err := fval.Set(val)
			if err != nil {
				return err
			}
		} else {
			return errors.New(fmt.Sprintf("The flag variable for \"%s\" does not conform to the flag.Value interface!", flagVal.long))
		}
	}

	return nil
}

func Parse() error {
	args := os.Args[1:]
	longToFlag := make(map[string]flag)
	shortToFlag := make(map[string]flag)
	for _, f := range flagDefs {
		longToFlag[f.long] = f
		shortToFlag[f.short] = f
	}
	done := false
	for i := 0; i < len(args); i++ {
		arg := args[i]
		if len(arg) >= 2 {
			if arg == "--" {
				done = true
			}
			if done {
				finalArgs = append(finalArgs, arg)
			} else if arg[0:2] == "--" {
				if flagVal, ok := longToFlag[arg[2:]]; ok {
					if b, ok := flagVal.value.(*bool); ok {
						*b = true
					} else {
						i += 1
						if err := assignFlagVal(flagVal, args[i]); err != nil {
							return err
						}
					}
				} else {
					return errors.New(fmt.Sprintf("Unknown flag passed in \"%s\"", arg))
				}
			} else if arg[0] == '-' {
				shortFlags := strings.Split(arg[1:], "")
				for j := 0; j < (len(shortFlags) - 1); j++ {
					if flagVal, ok := shortToFlag[shortFlags[j]]; ok {
						if b, ok := flagVal.value.(*bool); ok {
							*b = true
						} else {
							return errors.New(fmt.Sprintf("Flag (\"%s\") specified but requires a value.", shortFlags[j]))
						}
					} else {
						return errors.New(fmt.Sprintf("Unknown flag passed in \"%s\"", arg))
					}
				}
				final := shortFlags[len(shortFlags)-1]
				if flagVal, ok := shortToFlag[final]; ok {
					if b, ok := flagVal.value.(*bool); ok {
						*b = true
					} else {
						i += 1
						if err := assignFlagVal(flagVal, args[i]); err != nil {
							return err
						}
					}
				} else {
					return errors.New(fmt.Sprintf("Unknown flag passed in \"%s\"", final))
				}
			} else {
				finalArgs = append(finalArgs, arg)
			}
		} else {
			finalArgs = append(finalArgs, arg)
		}
	}

	return nil
}
