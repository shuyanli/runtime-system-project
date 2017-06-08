package stats

import (
	"flag"
	"fmt"
	//"regexp"
	//"strings"
	"time"
	"bytes"
	"sync" //for sync.Pool
)

var printStats = flag.Bool("printStats", false, "Print stats to console")


//using bufferpool to get rid of the allocation
var BufferPool = sync.Pool{
       New: func () interface{} {
              return &bytes.Buffer{}
       },
}

// IncCounter increments a counter.
func IncCounter(name string, tags map[string]string, value int64) {
	name = addTagsToName(name, tags)
	if *printStats {
		fmt.Printf("IncCounter: %v = %v\n", name, value)
	}
}

// UpdateGauge updates a gauge.
func UpdateGauge(name string, tags map[string]string, value int64) {
	name = addTagsToName(name, tags)
	if *printStats {
		fmt.Printf("UpdateGauge: %v = %v\n", name, value)
	}
}

// RecordTimer records a timer.
func RecordTimer(name string, tags map[string]string, d time.Duration) {
	name = addTagsToName(name, tags)
	if *printStats {
		fmt.Printf("RecordTimer: %v = %v\n", name, d)
	}
}

func addTagsToName(name string, tags map[string]string) string {
	// The format we want is: host.endpoint.os.browser
	// if there's no host tag, then we don't use it.

	//var keyOrder []string

	keyOrder := make([]string,0,4)
	if _, ok := tags["host"]; ok {
		keyOrder = append(keyOrder, "host")
	}
	keyOrder = append(keyOrder, "endpoint", "os", "browser")

	//instead of creatinga buffer directly, we get that from pool
	buff := BufferPool.Get().(*bytes.Buffer)
	buff.Reset()
	defer BufferPool.Put(buff) //after finish using the pool, we put it back into the pool

	//buff := &bytes.Buffer{}
	//parts := []string{name}  //first modify
	// parts := make([]string,0,5)
	// parts = append(parts, name)
	buff.WriteString(name)
	for _, k := range keyOrder {
		//add the '.' here instead of using te join function
		buff.WriteByte('.')
		v, ok := tags[k]
		if !ok || v == "" {
			buff.WriteString("no-")
			buff.WriteString(k)
			//parts = append(parts, "no-"+k)
			continue
		}

		//buff.WriteString(clean(v))
		clean(buff, v)
		//parts = append(parts, clean(v))
	}

	//return strings.Join(parts, ".")
	return buff.String()
}



// clean takes a string that may contain special characters, and replaces these
// characters with a '-'.

// ============== orirignal clean function that call replaceAllString in library 

// var specialChars = regexp.MustCompile(`[{}/\\:\s.]`)
// func clean(value string) string {
// 	return specialChars.ReplaceAllString(value, "-")
// }

//faster clean function
 
//second implementation:
//write out to the buffer directly instead of creating intermediate object
func clean(buff *bytes.Buffer, value string) {

	//res := make([]byte, len(value )) 
	for i := 0; i < len(value); i++ {
		switch c := value[i]; c {
		case '{', '}', '/', '\\', ':', ' ', '\t', '.':
			buff.WriteByte('-')
			//res[i] = '-'
		default:
			buff.WriteByte(c)
			//res[i] = c
		}
	}

	//return string(res)
}
