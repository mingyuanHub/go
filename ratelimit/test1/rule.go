package main

import (
	"fmt"
	"time"
)

type Rule struct {
	rules []*singleRule
}


func NewRule() *Rule {
	return &Rule{
		rules : []*singleRule{},
	}
}

func (this *Rule) AddSingleRule(defaultExpiration time.Duration, numberOfAllowedAccesses int) {
	this.rules = append(this.rules, newSingleRule(defaultExpiration, numberOfAllowedAccesses))
}

func (this *Rule) AllowVisit(key interface{}) bool {
	if len(this.rules) == 0 {
		fmt.Println("rule is empty，please add rule by AddSingleRule")
		return true
	}
	for _, rule := range this.rules {
		if !rule.allowVisit(key) {
			return false
		}
	}
	return true
}
