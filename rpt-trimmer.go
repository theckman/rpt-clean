package main

import "os"
import "bufio"
import "fmt"
import "regexp"
import "flag"

var useExt, useSimple bool

// parse my command-line args during the init
func init() {
    flag.BoolVar(&useExt, "e", false, "parse the rpt fpr extended output")
    flag.BoolVar(&useSimple, "s", false, "parse the rpt for simple output")
    flag.Parse()
}

// function for error conditions
func die(msg string, code int) {
    fmt.Fprintln(os.Stderr, msg)
    os.Exit(1)
}

// main function, aka needful()
func main() {
    // define variable for mode
    var mode string

    // make sure we have reasonable command-line options
    if useExt && useSimple {
        die("unable to use extended and simple", 1)
    } else if !useExt && !useSimple {
        die("you need to tell rpt-trimmer what to do", 1)
    }

    // set the mode string based on the needful
    if useSimple {
        mode = "simple"
    } else  {
        mode = "ext"
    }

    // set up the scanner for reading from stdin
    scanner := bufio.NewScanner(os.Stdin)

    // all the regexeseses for matching what the line is
    killed, err_killed := regexp.Compile("has been killed by")
    wounded, err_wound := regexp.Compile("has been.*?by")
    died, err_died := regexp.Compile("has died at")
    bleeder, err_bleeder := regexp.Compile("has bled out")

    // make sure the regexes didn't fail, they shouldn't. but Go gets really
    // pissy if you don't do shit with variables you defined (err_* int this case)
    // oh well...
    if err_killed != nil || err_wound != nil || err_died != nil || err_bleeder != nil {
        die("error occurred with regex needful", 1)
    }

    // things seem sane, let's go!
    // define the string used at the end of each line. '."'
    const EOL string = ".\""

    // loop over stdin
    for scanner.Scan() {
        // current line
        line := scanner.Text()

        // if it's a killed line
        if killed.MatchString(line) {
            // define variables needed
            var reg *regexp.Regexp
            var err error

            // specify which regex to use based on mode
            if mode == "simple" {
                reg, err = regexp.Compile(" by .*? (.*?).*?$")
            } else {
                reg, err = regexp.Compile("\\. Near players.*$")
            }

            // again with this requirement
            if err != nil {
                die("something happened in the scanner loop compiling regex " + err.Error(), 1)
            }

            // replace the line and print
            r_line := reg.ReplaceAllLiteralString(line, EOL)
            fmt.Println(r_line)
        // or if it's a wounded line
        } else if wounded.MatchString(line) {
            if mode == "simple" {
                reg, err := regexp.Compile(" for .*? damage.*$")
                if err != nil {
                    die("something happened in the scanner loop compiling regex " + err.Error(), 1)
                }
                r_line := reg.ReplaceAllLiteralString(line, EOL)
                fmt.Println(r_line)
            }
        // or if it's a died line
        } else if died.MatchString(line) {
            var reg *regexp.Regexp
            var err error
            if mode == "simple" {
                reg, err = regexp.Compile(" at \\[.*$")
            } else {
                reg, err = regexp.Compile("\\. Near players.*$")
            }
            if err != nil {
                die("somethign happened in the scanner loop compiling regex " + err.Error(), 1)
            }
            r_line := reg.ReplaceAllLiteralString(line, EOL)
            fmt.Println(r_line)
        // or if it's a bleeder line
        } else if bleeder.MatchString(line) {
            var reg *regexp.Regexp
            var err error
            if mode == "simple" {
                reg, err = regexp.Compile(" at \\[.*$")
            } else {
                reg, err = regexp.Compile("\\. Near players.*$")
            }
            if err != nil {
                die("somethign happened in the scanner loop compiling regex " + err.Error(), 1)
            }
            r_line := reg.ReplaceAllLiteralString(line, EOL)
            fmt.Println(r_line)
        // otherwise just print the line
        } else {
            fmt.Println(line)
        }
    }

    if err := scanner.Err(); err != nil {
        die("error reading stdin:" + err.Error(), 1)
    }
}
