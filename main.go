package main

import "C"
import (
	"fmt"
	"github.com/ledgerwatch/lmdb-go/lmdb"
	_ "github.com/ledgerwatch/lmdb-go/lmdb"
)

func main() {

	env, err := lmdb.NewEnv()
	if err != nil {

	}
	env.Open("error", lmdb.WriteMap, 0664)

	//var database lmdb.DBI
	err = env.Update(func(txn *lmdb.Txn) (err error) {
		db, err := txn.OpenRoot(0)
		if err != nil {
			return err
		}
		return txn.Put(db, []byte("key2"), []byte("value"), 0)
	})

	env.View(func(txn *lmdb.Txn) (err error) {
		db, err := txn.OpenRoot(0)
		cur, _ := txn.OpenCursor(db)
		if err != nil {
			return err
		}
		defer cur.Close()

		for {
			k, first, err := cur.Get(nil, nil, lmdb.Next)
			fmt.Println("key", string(k), "first", string(first),
				"err", err)
			if lmdb.IsNotFound(err) {
				fmt.Println("not found")
				return nil
			}
			if err != nil {
				fmt.Println("not found")
				return err
			}

			//stride := len(first)

			//for {
			//	_, v, err := cur.Get(nil, nil, lmdb.NextMultiple)
			//	if lmdb.IsNotFound(err) {
			//		break
			//	}
			//	if err != nil {
			//		return err
			//	}
			//
			//	multi := lmdb.WrapMulti(v, stride)
			//	for i := 0; i < multi.Len(); i++ {
			//		fmt.Printf("%s %s\n", k, multi.Val(i))
			//	}
			//}
		}
	})

	//env.View(func(txn *lmdb.Txn) error {
	//	db, err := txn.OpenRoot(0)
	//	value, err := txn.Get(db, []byte("key1"))
	//	fmt.Println("value", value)
	//	return err
	//})

	stat, err := env.Stat()
	info, _ := env.Info()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(stat.Entries)
	fmt.Println(info.MapSize)

}

//
//func ExampleCursor_Get_dupFixed() {
//	env.View(func(txn *lmdb.Txn) (err error) {
//		cur, err := txn.OpenCursor(dbi)
//		if err != nil {
//			return err
//		}
//		defer cur.Close()
//
//		for {
//			k, first, err := cur.Get(nil, nil, lmdb.NextNoDup)
//			if lmdb.IsNotFound(err) {
//				return nil
//			}
//			if err != nil {
//				return err
//			}
//
//			stride := len(first)
//
//			for {
//				_, v, err := cur.Get(nil, nil, lmdb.NextMultiple)
//				if lmdb.IsNotFound(err) {
//					break
//				}
//				if err != nil {
//					return err
//				}
//
//				multi := lmdb.WrapMulti(v, stride)
//				for i := 0; i < multi.Len(); i++ {
//					fmt.Printf("%s %s\n", k, multi.Val(i))
//				}
//			}
//		}
//	})
//}
