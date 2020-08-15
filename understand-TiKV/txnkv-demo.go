package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/juju/errors"
	"github.com/pingcap/tidb/kv"
	"github.com/pingcap/tidb/store/tikv"
	"github.com/pingcap/parser/terror"

	goctx "golang.org/x/net/context"
)

type KV struct {
	K, V []byte
}

func (kv KV) String() string {
	return fmt.Sprintf("%s => %s (%v)", kv.K, kv.V, kv.V)
}

var (
	store  kv.Storage
	pdAddr = flag.String("pd", "172.18.0.11:2379", "pd address:172.18.0.11:2379")
)

// Init initializes information.
func initStore() {
	driver := tikv.Driver{}
	var err error
	store, err = driver.Open(fmt.Sprintf("tikv://%s", *pdAddr))
	terror.MustNil(err)
}

// key1 val1 key2 val2 ...
func puts(args ...[]byte) error {
	tx, err := store.Begin()
	if err != nil {
		return errors.Trace(err)
	}

	for i := 0; i < len(args); i += 2 {
		key, val := args[i], args[i+1]
		err := tx.Set(key, val)
		if err != nil {
			return errors.Trace(err)
		}
	}
	err = tx.Commit(goctx.Background())
	if err != nil {
		return errors.Trace(err)
	}

	return nil
}

func get(k []byte) (KV, error) {
	tx, err := store.Begin()
	if err != nil {
		return KV{}, errors.Trace(err)
	}
	v, err := tx.Get(goctx.Background(), k)
	if err != nil {
		return KV{}, errors.Trace(err)
	}
	return KV{K: k, V: v}, nil
}

func dels(keys ...[]byte) error {
	tx, err := store.Begin()
	if err != nil {
		return errors.Trace(err)
	}
	for _, key := range keys {
		err := tx.Delete(key)
		if err != nil {
			return errors.Trace(err)
		}
	}
	err = tx.Commit(goctx.Background())
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

func scan(keyPrefix []byte, limit int) ([]KV, error) {
	tx, err := store.Begin()
	if err != nil {
		return nil, errors.Trace(err)
	}
	it, err := tx.Iter(kv.Key(keyPrefix), nil)
	if err != nil {
		return nil, errors.Trace(err)
	}
	defer it.Close()
	var ret []KV
	for it.Valid() && limit > 0 {
		ret = append(ret, KV{K: it.Key()[:], V: it.Value()[:]})
		limit--
		it.Next()
	}
	return ret, nil
}

func main() {
	pdAddr := os.Getenv("PD_ADDR")
	if pdAddr != "" {
		os.Args = append(os.Args, "-pd", pdAddr)
	}
	flag.Parse()
	initStore()

	// set
	err := puts([]byte("key1"), []byte("value1"), []byte("key2"), []byte("value2"))
	terror.MustNil(err)

	// get
	kv, err := get([]byte("key1"))
	terror.MustNil(err)
	fmt.Println(kv)

	// scan
	ret, err := scan([]byte("key"), 10)
	for _, kv := range ret {
		fmt.Println(kv)
	}

	// delete
	err = dels([]byte("key1"), []byte("key2"))
	terror.MustNil(err)
}