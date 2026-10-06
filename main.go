package main

import (
	"errors"
	"fmt"
	"sort"
	"unicode/utf8"
)

var pos = []string{"директор", "зам. директора", "начальник цеха", "мастер", "рабочий"}

var (
    ErrByShortName error = errors.New("Too short name")
    ErrByUnexpectedPos error = errors.New("No exist position")
    ErrByNegativeNums error = errors.New("Negative salary or negative exp")
)

type CompanyInterface interface{
    AddWorkerInfo(name, position string, salary, experience uint) error
    SortWorkers() ([]string, error) 
}

type Worker struct {
    name string
    position string
    salary uint
    experience uint
}

type Company struct {
    workers []Worker
}

func (c *Company) AddWorkerInfo(name, position string, salary, experience uint) error {
    if utf8.RuneCountInString(name) < 2 {
        return ErrByShortName
    }

    flag := false
    for _, x := range pos {
        if x == position {
            flag = true
        }
    }
    if !flag {
        return ErrByUnexpectedPos
    }

    if salary <= 0 && experience <= 0 {
        return ErrByNegativeNums
    }
    c.workers = append(c.workers, Worker{name: name, position: position, salary: salary, experience: experience})
    return nil
}

func getInd(p string) (int, error) {
    for i, elemet := range pos {
        if elemet == p {
            return i, nil
        }
    }
    return 0, ErrByUnexpectedPos
}

func (c *Company) SortWorkers() ([]string, error) {
    res := []string{}
    sort.Slice(c.workers, func(i, j int) bool {
        if c.workers[i].salary * c.workers[i].experience != c.workers[j].salary * c.workers[j].experience {
            return c.workers[i].salary * c.workers[i].experience > c.workers[j].salary * c.workers[j].experience
        } else if c.workers[i].position != c.workers[j].position {
            ind1, _ := getInd(c.workers[i].position)
            ind2, _ := getInd(c.workers[j].position)
            return  ind1 < ind2
        } else {
            return false
        }
    })
    for _, elem := range c.workers {
        res = append(res, fmt.Sprintf("%s — %v — %s", elem.name, elem.experience * elem.salary, elem.position))
        
    }
    return res, nil
}
