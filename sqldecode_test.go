package mycode

import (
	"bufio"
	"encoding/csv"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func Test_3(t *testing.T) {
	docsv := func(idx string) {
		filepath := idx + ".out"
		file, err := os.OpenFile(filepath, os.O_RDWR, 0666)
		if err != nil {
			t.Logf("Open file error! %v", err)
			return
		}
		defer file.Close()

		stat, err := file.Stat()
		if err != nil {
			panic(err)
		}
		var size = stat.Size()
		t.Logf("file size=%v", size)

		res := make(map[string][]string, 739)
		buf := bufio.NewReader(file)
		for {
			line, err := buf.ReadString('\n')
			if err != nil {
				if err == io.EOF {
					t.Log("File read ok!")
					break
				} else {
					t.Logf("Read file error! %v", err)
					return
				}
			}

			ls := strings.Fields(line)
			if len(ls) != 2 {
				t.Errorf("line=%v", line)
				return
			}

			if _, ok := res[ls[1]]; !ok {
				res[ls[1]] = make([]string, 0, 100)
			}

			res[ls[1]] = append(res[ls[1]], ls[0])
		}

		pids := make([]string, 0, len(res))
		for parentID, _ := range res {
			pids = append(pids, parentID)
		}
		sort.Slice(pids, func(i, j int) bool {
			a, _ := strconv.Atoi(pids[i])
			b, _ := strconv.Atoi(pids[j])
			if a > b {
				return true
			}
			return false
		})

		f, err := os.Create("./csv/" + idx + ".csv")
		if err != nil {
			panic(err)
		}
		defer f.Close()

		f.WriteString("\xEF\xBB\xBF")
		w := csv.NewWriter(f)

		data := make([][]string, 0, len(pids))
		for _, pid := range pids {
			rec := make([]string, 0, 2)
			rec = append(rec, pid)
			ids := res[pid]
			sort.Slice(ids, func(i, j int) bool {
				a, _ := strconv.Atoi(ids[i])
				b, _ := strconv.Atoi(ids[j])
				if a > b {
					return true
				}
				return false
			})
			rec = append(rec, strings.Join(ids, " "))
			data = append(data, rec)
		}
		w.WriteAll(data)
		w.Flush()
	}
	idxs := []string{"33", "39", "41", "43", "44", "47", "48", "49", "51", "52"}
	for _, idx := range idxs {
		t.Log(idx)
		docsv(idx)
	}
}
