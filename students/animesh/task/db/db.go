package db

import (
	"encoding/binary"
	"errors"
	"fmt"
	"time"

	"github.com/boltdb/bolt"
)

const (
	TasksTbl = "TASKTBL"
)

var db *bolt.DB

func InitDb(dbPath string) error {
	var err error
	db, err = bolt.Open(dbPath, 0600, &bolt.Options{Timeout: 1 * time.Second, ReadOnly: false})
	if err != nil {
		return err
	}
	return createTable(TasksTbl)
}

func CloseDb() {
	db.Close()
}

func createTable(tablename string) error {
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(tablename))
		if err != nil {
			return fmt.Errorf("create bucket: %s", err.Error())
		}
		return nil
	})
}

func AddTask(t string) (int, error) {
	var id int
	e := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TasksTbl))
		uid, _ := b.NextSequence()
		id = int(uid)
		t := Task{Id: id, Desc: t, CreateTS: time.Now(), DoneTS: time.Now(), Status: 0, Critic: 0, Urge: 0, Effor: 0}
		bs, err := t.toByte()
		if err != nil {
			return err
		}
		return b.Put(itob(id), bs)
	})
	if e != nil {
		return -1, e
	}
	return id, nil
}

func ListAll() []Task {
	var t []Task
	db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TasksTbl))
		b.ForEach(func(k, v []byte) error {
			task, err := Decode(v)
			if err != nil {
				return err
			}
			t = append(t, task)
			return nil
		})
		return nil
	})
	return t
}

func DoneTask(id int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TasksTbl))
		v := b.Get(itob(id))
		if v == nil {
			return errors.New("Warning : no task with id " + string(id))
		}
		t, err := Decode(v)
		if err != nil {
			return err
		}
		(&t).Done()
		bytes, err := t.toByte()
		if err != nil {
			return err
		}
		return b.Put(itob(id), bytes)
	})
}

func WaiveTask(id int, reclaim bool) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TasksTbl))
		v := b.Get(itob(id))
		if v == nil {
			return errors.New("Warning : no task with id " + string(id))
		}
		t, err := Decode(v)
		if err != nil {
			return err
		}
		if reclaim {
			(&t).Todo()
		} else {
			(&t).Waive()
		}
		bytes, err := t.toByte()
		if err != nil {
			return err
		}
		return b.Put(itob(id), bytes)
	})
}

func UpdateTask(id int, desc string) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TasksTbl))
		v := b.Get(itob(id))
		if v == nil {
			return errors.New("Warning : no task with id " + string(id))
		}
		t, err := Decode(v)
		if err != nil {
			return err
		}
		(&t).UpdateDesc(desc)
		bytes, err := t.toByte()
		if err != nil {
			return err
		}
		return b.Put(itob(id), bytes)
	})
}

func updateInt(id, value, cas int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TasksTbl))
		v := b.Get(itob(id))
		if v == nil {
			return errors.New("Warning : no task with id " + string(id))
		}
		t, err := Decode(v)
		if err != nil {
			return err
		}
		switch cas {
		case 0:
			err = (&t).UpdateEffort(value)
		case 1:
			err = (&t).UpdateUrgency(value)
		case 2:
			err = (&t).UpdateCriticality(value)
		}
		if err != nil {
			return err
		}
		bytes, err := t.toByte()
		if err != nil {
			return err
		}
		return b.Put(itob(id), bytes)
	})
}

func UpdateEffort(id int, effort int) error {
	return updateInt(id, effort, 0)
}

func UpdateUrgency(id int, urgency int) error {
	return updateInt(id, urgency, 1)
}

func UpdateCriticality(id int, criticality int) error {
	return updateInt(id, criticality, 2)
}

func FindTask(id int) (Task, error) {
	var t Task
	var err error
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TasksTbl))
		bytes := b.Get(itob(id))
		if bytes == nil {
			return errors.New("Warning : no task with id " + string(id))
		}
		t, err = Decode(bytes)
		return err
	})
	return t, err
}

func DeleteTask(id int) error {
	return db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(TasksTbl))
		c := b.Cursor()
		k, _ := c.Seek(itob(id))
		if k == nil {
			return errors.New("Warning : no task with id " + string(id))
		}
		c.Delete()
		return nil
	})
}

func itob(v int) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func btoi(b []byte) int {
	return int(binary.BigEndian.Uint64(b))
}
