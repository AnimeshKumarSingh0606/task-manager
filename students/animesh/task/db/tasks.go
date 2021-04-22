package db

import (
	"bytes"
	"encoding/gob"
	"errors"
	"time"
)

type Task struct {
	Id int       
	Desc string    
	CreateTS time.Time 
	DoneTS   time.Time 
	Status int       
	Critic int       
	Urge int       
	Effor int       
}

func (t Task) State() string {
	switch t.Status {
	default:
		return "TODO"
	case 1:
		return "DOING"
	case 2:
		return "DONE"
	case 3:
		return "WAIVE"
	}
}

func (t Task) Effort() string {
	switch t.Effor {
	default:
		return "EASY"
	case 1:
		return "COMPLEX"
	case 2:
		return "COMPLIC"
	case 3:
		return "HARD"
	}
}

func (t Task) Urgency() string {
	switch t.Urge {
	default:
		return "INTIME"
	case 1:
		return "WAITED"
	case 2:
		return "URGENT"
	}
}

func (t Task) Criticality() string {
	switch t.Critic {
	default:
		return "WISHED"
	case 1:
		return "WANTED"
	case 2:
		return "NEEDED"
	}
}

func (t *Task) UpdateDesc(s string) {
	t.Desc = s
}

func (t *Task) UpdateCriticality(i int) error {
	if i >= 0 && i <= 2 {
		t.Critic = i
		return nil
	}
	return errors.New("Bad criticality value (must be in [0, 2])")
}

func (t *Task) UpdateUrgency(i int) error {
	if i >= 0 && i <= 2 {
		t.Urge = i
		return nil
	}
	return errors.New("Bad urgency value (must be in [0, 2])")
}

func (t *Task) UpdateEffort(i int) error {
	if i >= 0 && i <= 3 {
		t.Effor = i
		return nil
	}
	return errors.New("Bad effort value (must be in [0, 3])")
}

func (t *Task) Todo() {
	t.Status = 0
	t.DoneTS = time.Now()
}

func (t *Task) Done() {
	t.Status = 2
	t.DoneTS = time.Now()
}

func (t *Task) Waive() {
	t.Status = 3
	t.DoneTS = time.Now()
}

func (t Task) IsTodo() bool {
	return t.Status == 0
}

func (t Task) IsDone() bool {
	return t.Status == 2
}

func (t Task) IsWaived() bool {
	return t.Status == 3
}

func (t Task) toByte() ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(t)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func Decode(b []byte) (Task, error) {
	pbuf := bytes.NewBuffer(b)
	dec := gob.NewDecoder(pbuf)
	var t Task
	err := dec.Decode(&t)
	if err != nil {
		return t, err
	}
	return t, nil
}
